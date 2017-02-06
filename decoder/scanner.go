package decoder

import (
	"fmt"
	"strconv"
)

// Valid check if the provide data is correct json document.
func Valid(src []byte) error {
	var scan Scanner
	return valid(src, &scan)
}

// Valid check if the provide data is correct json document.
func valid(src []byte, scan *Scanner) error {
	scan.Reset()
	for _, c := range src {
		scan.Step(c)
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

// SyntaxError is JSON syntax error with a message and int's position.
type SyntaxError struct {
	Message  string
	Position Position
}

func (se SyntaxError) Error() string {
	return se.Position.string() + se.Message
}

//go:generate stringer -type=state

// state is the current state of the scanners state machine.
type state int

// The list of possible status the scanner can be in.
const (
	stateBeginJSON state = iota // At the beginning there was JSON.
	stateEndJSON
	// stateWhitespace
	stateBeginObject // member or end.
	stateInObject    // we are in an object

	stateBeginMember    // seen "," in stateInObject
	stateInMember       // end of object or , + value
	stateBeginMemberKey // " then jump into a string.
	stateEndMemberKey   // expect :
	stateInMemberValue  // after :, expect a value type.

	stateBeginArray      // value or end.
	stateInArray         // after value in array, expect Seperator Followed by Value or ].
	stateBeginArrayValue // after , in a value, expect value type.
	stateEndArray        // ]

	// stateEndObjectValue   // object valued finished, expect key or end of object.

	stateBeginString // after "
	stateInString    // string value.
	stateEndString   // " string end.

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
	stateEndBool
	stateInNull // null

	stateError // Something is wrong, check s.err
)

var steps = map[state]func(*Scanner, byte){
	stateBeginJSON: stepBeginJSON,
	stateEndJSON:   stepEndJSON,

	stateBeginObject:    stepBeginObject,
	stateInObject:       stepInObject,
	stateBeginMember:    stepBeginMember,
	stateBeginMemberKey: stepBeginMemberKey,
	stateEndMemberKey:   stepEndMemberKey,
	stateInMemberValue:  stepInMemberValue,
	stateInMember:       stepInMember,

	stateBeginArray:      stepBeginArray,
	stateInArray:         stepInArray,
	stateBeginArrayValue: stepBeginArrayValue,
	stateEndArray:        stepEndArray,

	stateBeginString: stepInString,
	stateInString:    stepInString,
	stateEndString:   stepEndString,

	stateInStringEscape:      stepInStringEscape,
	stateInStringEscapeU:     stepInStringEscapeU,
	stateInStringEscapeUx:    stepInStringEscapeUx,
	stateInStringEscapeUxx:   stepInStringEscapeUxx,
	stateInStringEscapeUxxx:  stepInStringEscapeUxxx,
	stateInStringEscapeUxxxx: stepInStringEscapeUxxxx,

	stateInFalse: stepInFalse,
	stateEndBool: stepEnd,
	stateInTrue:  stepInTrue,
	stateInNull:  stepInNull,

	stateBeginNumber:     stepBeginNumber,
	stateInNumber:        stepInNumber,
	stateBeginNumberFrac: stepBeginNumberFrac,
	stateInNumberFrac:    stepInNumberFrac,
	stateBeginNumberExp:  stepBeginNumberExp,
	stateInNumberExp:     stepInNumberExp,
}

// Scanner is a JSON Scanner.
type Scanner struct {
	state state
	pos   Position
	last  byte //The last token we had.
	err   *SyntaxError
	stack []state

	yield bool
	// the stack of state
}

func (s *Scanner) pushstate() {
	s.stack = append(s.stack, s.state)
}

func (s *Scanner) popstate() {
	l := len(s.stack)

	if l == 0 {
		s.state = stateEndJSON
	} else {
		s.state, s.stack = s.stack[l-1], s.stack[:l-1]
	}

}

func (s *Scanner) eof() {
	if s.err != nil {
		return
	}

	s.Step(' ') // we should be okay with whitespace at this point.

	if s.err != nil {
		return
	}

	if s.state != stateEndJSON {
		s.err = &SyntaxError{
			//TODO: Add more context about what state we were.
			Message:  "Unexpected end of json",
			Position: s.pos,
		}
	}

}

// Reset sets the scanner to correct starting state.
func (s *Scanner) Reset() {
	s.pos.reset()
	s.err = nil
	s.yield = false
	// since a whitespace is valid at the start we use
	// a space instead of 0
	s.last = ' '
	s.state = stateBeginJSON
}

func (s *Scanner) Error() error {
	if s.err != nil {
		return s.err
	}
	return nil
}

// Step is moving the Scanner state machine one step ahead
// using the byte provided as the next charachter.
func (s *Scanner) Step(c byte) {

	if s.err != nil {
		return
	}

	s.pos.advance(c)

	step := steps[s.state]
	//fmt.Printf(":::%c: %v\n", c, s.state)
	step(s, c)
	for s.yield {
		s.yield = false
		step = steps[s.state]
		step(s, c)
		//fmt.Printf(":.:%c: %v\n", c, s.state)
	}
	fmt.Printf(":::%c: %v\n", c, s.state)
	s.last = c
}

// error records an error and switches to the error state.
func (s *Scanner) error(c byte, context string) {
	s.state = stateError
	s.err = &SyntaxError{
		Message:  "invalid character " + quoteChar(c) + " " + context,
		Position: s.pos,
	}
}

func stepBeginJSON(s *Scanner, c byte) {
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

func stepEndJSON(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.error(c, "unexpected end of json")
}

func stepBeginObject(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {
	case '}':
		s.popstate() //
	case '"':
		s.state = stateInObject
		s.pushstate()
		// s.state = stateInMemberKey
		s.state = stateBeginMemberKey
	default:
		s.error(c, "expected a pair or end of object")
	}

}
func stepInObject(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case '}':
		s.popstate()
	case ',':
		s.pushstate()
		s.state = stateBeginMember
	default:
		s.error(c, "expected a pair or end of object")
	}

}

func stepBeginMember(s *Scanner, c byte) {
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

func stepBeginMemberKey(s *Scanner, c byte) {
	s.state = stateEndMemberKey
	s.pushstate()
	s.state = stateInString
}

func stepEndMemberKey(s *Scanner, c byte) {
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

func stepInMemberValue(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.state = stateInMember
	helperBeginValue(s, c)
}

func stepInMember(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case '}':
		s.popstate() // end of object.
	case ',':
		s.state = stateBeginMember
	default:
		s.error(c, "expected } or a value seperator")
	}
}

func stepBeginArray(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	// end or value
	if c == ']' {
		s.state = stateEndArray
		return
	}

	s.state = stateInArray
	s.pushstate()
	helperBeginValue(s, c)
}

func stepInArray(s *Scanner, c byte) {

	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {

	case ']':
		s.state = stateEndArray
	case ',':
		s.pushstate()
		s.state = stateBeginArrayValue
	default:
		s.error(c, "expected ] or a value seperator")
	}
}

func stepEndArray(s *Scanner, c byte) {
	s.yield = true
	s.popstate() // end of string.
}

func stepBeginArrayValue(s *Scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	s.state = stateInArray
	helperBeginValue(s, c)
}

func helperBeginValue(s *Scanner, c byte) {
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
		s.state = stateBeginString
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

func stepInString(s *Scanner, c byte) {

	if c < 0x20 {
		s.error(c, "in string literal")
		return
	}

	switch c {
	case '\\':
		s.pushstate()
		s.state = stateInStringEscape
	case '"':
		s.state = stateEndString
	default:
		s.state = stateInString
	}
}

func stepEndString(s *Scanner, c byte) {
	s.yield = true
	s.popstate() // end of string.
}

// basic escaping, \uXXXX case is it's own state.
func stepInStringEscape(s *Scanner, c byte) {

	switch c {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		s.popstate() // end of escape.
	case 'u':
		s.state = stateInStringEscapeU
	default:
		s.error(c, "Expected a qoutation mark, reverse solidus, or a control character")
	}

}

const errorExpectedHexDigit = "Expected a Hexadecimal digit."

func stepInStringEscapeU(s *Scanner, c byte) {
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

func stepInStringEscapeUx(s *Scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxx(s *Scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxxx(s *Scanner, c byte) {
	if isHexDig(c) {
		s.state = stateInStringEscapeUxxxx
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}

func stepInStringEscapeUxxxx(s *Scanner, c byte) {
	if isHexDig(c) {
		s.popstate()
	} else {
		s.error(c, errorExpectedHexDigit)
	}
}
func stepInFalse(s *Scanner, c byte) {

	if (s.last == 'f' && c != 'a') ||
		(s.last == 'a' && c != 'l') ||
		(s.last == 'l' && c != 's') ||
		(s.last == 's' && c != 'e') {
		s.error(c, "Unexpected value")
	} else if s.last == 's' && c == 'e' {
		s.state = stateEndBool
	}
}

func stepInTrue(s *Scanner, c byte) {

	if (s.last == 't' && c != 'r') ||
		(s.last == 'r' && c != 'u') ||
		(s.last == 'u' && c != 'e') {
		s.error(c, "Unexpected value")
	} else if s.last == 'u' && c == 'e' {
		s.state = stateEndBool
	}
}

func stepInNull(s *Scanner, c byte) {

	if (s.last == 'n' && c != 'u') ||
		(s.last == 'u' && c != 'l') ||
		(s.last == 'l' && c != 'l') {
		s.error(c, "Unexpected value")
	} else if s.last == 'l' && c == 'l' {
		s.popstate() //nullols?
	}
}

func stepBeginNumber(s *Scanner, c byte) {

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
			s.popstate()
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

func stepInNumber(s *Scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	case '.':
		s.state = stateBeginNumberFrac
	default:
		s.popstate() // Not digit, not eE, not ., must be end of number? yeild.
		s.yield = true
	}
}

func stepBeginNumberFrac(s *Scanner, c byte) {
	// we have seen a '.' in a number
	if '0' <= c && c <= '9' {
		s.state = stateInNumberFrac
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateInNumberExp
	default:
		s.popstate() // Not digit, not eE, not must be end of number? yeild.
		s.yield = true
	}
}

func stepInNumberFrac(s *Scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	default:
		s.popstate() // Not digit, not eE, not must be end of number? yeild.
		s.yield = true
	}
}

func stepBeginNumberExp(s *Scanner, c byte) {

	if c == '-' || c == '+' || '1' <= c && c <= '9' {
		s.state = stateInNumberExp
		return
	}

	s.error(c, "Expected Minus, Plus or non zero digit.")
}

func stepInNumberExp(s *Scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	// We know these values are never valid.
	if c == 'e' || c == 'E' || c == '.' {
		s.error(c, "Expected digit")
		return
	}

	s.popstate() // Not digit, must be end of number? yeild.
	s.yield = true
}

func stepEnd(s *Scanner, c byte) {
	s.popstate() //
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