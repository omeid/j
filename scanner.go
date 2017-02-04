package json

import "strconv"

// Valid check if the provide data is correct json document.
func Valid(src []byte) error {
	var scan scanner

	scan.reset()
	for _, c := range src {
		scan.step(c)
		if scan.state == stateError {
			return scan.err
		}
	}

	scan.eof() // check for the end.

	if scan.err != nil {
		return scan.err
	}

	return nil
}

// SyntaxError is an error and it's position.
type SyntaxError struct {
	Message  string
	Position Position
}

func (se SyntaxError) Error() string {
	return se.Position.string() + se.Message
}

//go:generate stringer -type=State

// State is the current state of the scanner state machine.
type State int

// Tokens as per the JSON spec.
const (
	stateBeginJSON State = iota // At the beginning there was JSON.
	stateEndJSON
	// stateWhitespace
	stateBeginObject // member or end.
	stateInObject    // we are in an object

	stateBeginMember    // seen "," in stateInObject
	stateInMember       // end of object or , + value
	stateBeginMemberKey // " then jump into a string.
	stateEndMemberKey   // expect :
	stateInMemberValue  // after :, expect a value type.

	stateBeginArray // value or end.
	stateInArray    // after value in array, expect Seperator Followed by Value or ].
	stateArrayValue // after , in a value, expect value type.

	// stateEndObjectValue   // object valued finished, expect key or end of object.

	stateBeginValue // find a value type.

	stateInString // string value.

	stateInStringEscape

	stateInStringEscapeU
	stateInStringEscapeUx
	stateInStringEscapeUxx
	stateInStringEscapeUxxx
	stateInStringEscapeUxxxx

	stateBeginNumber     // we have seen - or a digit.
	stateInNumber        // just digits, nothing special.
	stateBeginNumberFrac // we have seen a . after 0 or a digit
	stateInNumberFrac    // we have seen a digit (0-9) after . eE now allowed.
	stateBeginNumberExp  // we have seen e or E
	stateInNumberExp     // we are past (e|E)(+|-|1-9).

	stateInFalse // false
	stateInTrue  // true
	stateInNull  // null

	stateError // Something is wrong, check s.err
)

var steps = map[State]func(*scanner, byte){
	stateBeginJSON: stepBeginJSON,
	stateEndJSON:   stepEndJSON,

	stateBeginObject:    stepBeginObject,
	stateInObject:       stepInObject,
	stateBeginMember:    stepBeginMember,
	stateBeginMemberKey: stepBeginMemberKey,
	stateEndMemberKey:   stepEndMemberKey,
	stateInMemberValue:  stepInMemberValue,
	stateInMember:       stepInMember,

	stateBeginArray: stepBeginArray,
	stateInArray:    stepInArray,
	stateArrayValue: stepArrayValue,

	stateInString:            stepInString,
	stateInStringEscape:      stepInStringEscape,
	stateInStringEscapeU:     stepInStringEscapeU,
	stateInStringEscapeUx:    stepInStringEscapeUx,
	stateInStringEscapeUxx:   stepInStringEscapeUxx,
	stateInStringEscapeUxxx:  stepInStringEscapeUxxx,
	stateInStringEscapeUxxxx: stepInStringEscapeUxxxx,

	stateInFalse: stepInFalse,
	stateInTrue:  stepInTrue,
	stateInNull:  stepInNull,

	stateBeginNumber:     stepBeginNumber,
	stateInNumber:        stepInNumber,
	stateBeginNumberFrac: stepBeginNumberFrac,
	stateInNumberFrac:    stepInNumberFrac,
	stateBeginNumberExp:  stepBeginNumberExp,
	stateInNumberExp:     stepInNumberExp,
}

type scanner struct {
	state State
	pos   Position
	last  byte //The last token we had.
	err   *SyntaxError
	stack []State

	yield bool
	// the stack of state
}

func (s *scanner) pushState() {
	s.stack = append(s.stack, s.state)
}

func (s *scanner) popState() {
	l := len(s.stack)

	if l == 0 {
		s.state = stateEndJSON
	} else {
		s.state, s.stack = s.stack[l-1], s.stack[:l-1]
	}

}

func (s *scanner) eof() {
	if s.err != nil {
		return
	}

	s.step(' ') // we should be okay with whitespace at this point.

	if s.err != nil {
		return
	}

	if s.state != stateEndJSON {
		s.err = &SyntaxError{
			Message:  "Unexpected end of json", // + stateEndJSON.String(),
			Position: s.pos,
		}
	}

}

func (s *scanner) reset() {
	///TODO: REMOVE POST DEBUG
	s.pos.reset()
	s.err = nil
	// is used to pop the stack and "yield".
	s.yield = false
	// since a whitespace is valid at the start we use
	// a space instead of 0
	s.last = ' '
	s.state = stateBeginJSON
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) {
	s.state = stateError
	s.err = &SyntaxError{
		Message:  "invalid character " + quoteChar(c) + " " + context,
		Position: s.pos,
	}
}

func (s *scanner) step(c byte) {

	if s.err != nil {
		return
	}

	s.pos.advance(c)

	step := steps[s.state]
	step(s, c)
	for s.yield {
		s.yield = false
		step = steps[s.state]
		step(s, c)
	}
	s.last = c
}

func stepBeginJSON(s *scanner, c byte) {
	// ignore whitespaces
	if c <= ' ' && isSpace(c) {
		return // nothing changes.
	}

	switch c {

	case '{':
		s.state = stateBeginObject
	case '[':
		s.state = stateBeginArray
	default:
		s.error(c, "expected beginning of json ('{' or '[')")
	}
}

func stepEndJSON(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.error(0, "unexpected end of json")
}

func stepBeginObject(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {
	case '}':
		s.popState() //
	case '"':
		s.state = stateInObject
		s.pushState()
		// s.state = stateInMemberKey
		s.state = stateBeginMemberKey
	default:
		s.error(c, "expected a pair or end of object")
	}

}
func stepInObject(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case '}':
		s.popState()
	case ',':
		s.pushState()
		s.state = stateBeginMember
	default:
		s.error(c, "expected a pair or end of object")
	}

}

func stepBeginMember(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	if c == '"' {
		// s.state = stateInMemberKey
		s.state = stateBeginMemberKey
	} else {
		s.error(c, "expected string.")
	}
}

func stepBeginMemberKey(s *scanner, c byte) {
	s.state = stateEndMemberKey
	s.pushState()
	s.state = stateInString
}

func stepEndMemberKey(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	if c != ':' {
		s.error(c, "expected a colon")
	} else {
		s.state = stateInMemberValue
	}
}

func stepInMemberValue(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.state = stateInMember
	stepBeginValue(s, c)
}

func stepInMember(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case '}':
		s.popState() // end of object.
	case ',':
		s.state = stateBeginMember
	default:
		s.error(c, "expected } or a value seperator")
	}
}

func stepBeginArray(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	// end or value
	if c == ']' {
		s.popState()
		return
	}

	s.state = stateInArray
	s.pushState()
	stepBeginValue(s, c)
}

func stepInArray(s *scanner, c byte) {

	// if s.last == ',' || s.last == ']' {
	// 	c = s.last
	// }

	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case ']':
		s.popState() // array end.
	case ',':
		s.pushState()
		s.state = stateArrayValue
	default:
		s.error(c, "expected ] or a value seperator")
	}
}

func stepArrayValue(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.state = stateInArray
	stepBeginValue(s, c)
}

func stepBeginValue(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	if '1' <= c && c <= '9' {
		s.state = stateBeginNumber
		return
	}

	switch c {
	case '{':
		s.state = stateBeginObject
	case '[':
		s.state = stateBeginArray
	case '"':
		s.state = stateInString
	case 'f':
		s.state = stateInFalse
	case 't':
		s.state = stateInTrue
	case 'n':
		s.state = stateInNull
	case '-':
		s.state = stateBeginNumber
	case '0':
		s.state = stateBeginNumber
	default:
		s.error(c, "expected a value type")
	}
}

func stepInString(s *scanner, c byte) {

	switch c {
	case '\\':
		s.pushState()
		s.state = stateInStringEscape
	case '"':
		s.popState() // end of string.
	}
}

// basic escaping, \uXXXX case is it's own state.
func stepInStringEscape(s *scanner, c byte) {

	switch c {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		s.popState() // end of escape.
	case 'u':
		s.state = stateInStringEscapeU
	default:
		s.error(c, "Expected a qoutation mark, reverse solidus, or a control character")
	}

}

const errorExpectedHexDigit = "Expected a Hexadecimal digit."

func stepInStringEscapeU(s *scanner, c byte) {
	if s.last != 'u' {
		// we shouldn't be here!
		s.error(c, "parser-error: in escape-U state not after 'u'.")
		return
	}

	if isHexDig(c) {
		s.state = stateInStringEscapeUx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUx(s *scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxx(s *scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxxx(s *scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxxxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxxxx(s *scanner, c byte) {
	if isHexDig(c) {
		s.popState()
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}
func stepInFalse(s *scanner, c byte) {

	if (s.last == 'f' && c != 'a') ||
		(s.last == 'a' && c != 'l') ||
		(s.last == 'l' && c != 's') ||
		(s.last == 's' && c != 'e') {
		s.error(c, "Unexpected value")
	} else if s.last == 's' && c == 'e' {
		s.popState() //falsey?
	}
}

func stepInTrue(s *scanner, c byte) {

	if (s.last == 't' && c != 'r') ||
		(s.last == 'r' && c != 'u') ||
		(s.last == 'u' && c != 'e') {
		s.error(c, "Unexpected value")
	} else if s.last == 'u' && c == 'e' {
		s.popState() //trues?
	}
}

func stepInNull(s *scanner, c byte) {

	if (s.last == 'n' && c != 'u') ||
		(s.last == 'u' && c != 'l') ||
		(s.last == 'l' && c != 'l') {
		s.error(c, "Unexpected value")
	} else if s.last == 'l' && c == 'l' {
		s.popState() //nullols?
	}
}

func stepBeginNumber(s *scanner, c byte) {

	if s.last == '0' {
		// after 0
		if '0' <= c && c <= '9' {
			s.error(c, "expected  e|E, decimal point after 0, or exactly 0")
			return
		}

		switch c {
		case 'e', 'E':
			s.state = stateBeginNumberExp
		case '.':
			s.state = stateBeginNumberFrac
		default:
			// possibly the value '0', yeild.
			s.popState()
			s.yield = true
		}
	} else if s.last == '-' {
		if '0' <= c && c <= '9' {
			s.state = stateInNumber
		} else {
			s.error(c, "expected digit")
		}
	} else if '1' <= s.last && s.last <= '9' {
		// after 1-9
		s.state = stateInNumber
		s.yield = true //We are in number.
	} else {
		// why are we here?
		s.error(c, "scanner-error: In Begin Number state not after 0 or -")
	}

}

func stepInNumber(s *scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	case '.':
		s.state = stateBeginNumberFrac
	default:
		s.popState() // Not digit, not eE, not ., must be end of number? yeild.
		s.yield = true
	}
}

func stepBeginNumberFrac(s *scanner, c byte) {
	// we have seen a '.' in a number
	if '0' <= c && c <= '9' {
		s.state = stateInNumberFrac
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateInNumberExp
	default:
		s.popState() // Not digit, not eE, not must be end of number? yeild.
		s.yield = true
	}
}

func stepInNumberFrac(s *scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	default:
		s.popState() // Not digit, not eE, not must be end of number? yeild.
		s.yield = true
	}
}

func stepBeginNumberExp(s *scanner, c byte) {

	if c == '-' || c == '+' || '1' <= c && c <= '9' {
		s.state = stateInNumberExp
		return
	}

	s.error(c, "Expected Minus, Plus or non zero digit.")
}

func stepInNumberExp(s *scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	// We know these values are never valid.
	if c == 'e' || c == 'E' || c == '.' {
		s.error(c, "Expected digit")
		return
	}

	s.popState() // Not digit, must be end of number? yeild.
	s.yield = true
}

func isHexDig(c byte) bool {
	return '0' <= c && c <= '9' || 'a' <= c && c <= 'f' || 'A' <= c && c <= 'F'
}

func isSpace(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

// quoteChar formats c as a quoted character literal
func quoteChar(c byte) string {
	// special cases - different from quoted strings
	if c == '\'' {
		return `'\''`
	}
	if c == '"' {
		return `'"'`
	}

	// use quoted string with different quotation marks
	s := strconv.Quote(string(c))
	return "'" + s[1:len(s)-1] + "'"
}
