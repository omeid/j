package j

import
// "fmt"

"strconv"

// Valid check if the provide data is correct json document.
func Valid(src []byte) error {
	var scan scanner
	return valid(src, &scan)
}

// Valid check if the provide data is correct json document.
func valid(src []byte, scan *scanner) error {
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
	return se.Message
}

//go:generate stringer -type=state

// state is the current state of the scanners state machine.
type state int

// The list of possible status the scanner can be in.
const (
	stateBeginJSON state = iota // At the beginning there was JSON.
	stateEndJSON
	// stateWhitespace
	stateBeginObject // { member or end.
	stateInObject    // we are in an object
	stateEndObject   // }

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
	stateEndNumber       // end of a number

	stateBeginNull // At n.
	stateInNull    // nu, ul,
	stateEndNull   // ll?

	stateInFalse // false
	stateInTrue  // true
	stateEndBool

	stateError // Something is wrong, check s.err
)

var steps = []func(*scanner, byte){
	stepBeginJSON,
	stepEndJSON,
	stepBeginObject,
	stepInObject,
	stepEnd, // stateEndObject,

	stepBeginMember,
	stepInMember,
	stepBeginMemberKey,
	stepEndMemberKey,
	stepInMemberValue,

	stepBeginArray,
	stepInArray,
	stepBeginArrayValue,
	stepEnd, // stateEndArray,

	stepBeginString,
	stepInString,
	stepEnd, //stateEndString

	stepInStringEscape,

	stepInStringEscapeU,
	stepInStringEscapeUx,
	stepInStringEscapeUxx,
	stepInStringEscapeUxxx,
	stepInStringEscapeUxxxx,

	stepBeginNumber,
	stepInNumber,
	stepBeginNumberFrac,
	stepInNumberFrac,
	stepBeginNumberExp,
	stepInNumberExp,
	stepEnd, // stateEndNumber

	stepBeginNull,
	stepInNull,
	stepEnd, // stateEndNull,

	stepInFalse,
	stepInTrue,
	stepEnd, // stateEndBool,

	stepError,
}

// var steps = map[state]func(*scanner, byte){
// 	stateBeginJSON: stepBeginJSON,
// 	stateEndJSON:   stepEndJSON,
//
// 	stateBeginObject:    stepBeginObject,
// 	stateEndObject:      stepEnd,
// 	stateInObject:       stepInObject,
// 	stateBeginMember:    stepBeginMember,
// 	stateBeginMemberKey: stepBeginMemberKey,
// 	stateEndMemberKey:   stepEndMemberKey,
// 	stateInMemberValue:  stepInMemberValue,
// 	stateInMember:       stepInMember,
//
// 	stateBeginArray:      stepBeginArray,
// 	stateInArray:         stepInArray,
// 	stateBeginArrayValue: stepBeginArrayValue,
// 	stateEndArray:        stepEnd,
//
// 	stateBeginString: stepInString,
// 	stateInString:    stepInString,
// 	stateEndString:   stepEnd,
//
// 	stateInStringEscape:      stepInStringEscape,
// 	stateInStringEscapeU:     stepInStringEscapeU,
// 	stateInStringEscapeUx:    stepInStringEscapeUx,
// 	stateInStringEscapeUxx:   stepInStringEscapeUxx,
// 	stateInStringEscapeUxxx:  stepInStringEscapeUxxx,
// 	stateInStringEscapeUxxxx: stepInStringEscapeUxxxx,
//
// 	stateBeginNull: stepBeginNull,
// 	stateInNull:    stepInNull,
// 	stateEndNull:   stepEnd,
// 	stateInFalse:   stepInFalse,
// 	stateEndBool:   stepEnd,
// 	stateInTrue:    stepInTrue,
//
// 	stateBeginNumber:     stepBeginNumber,
// 	stateInNumber:        stepInNumber,
// 	stateBeginNumberFrac: stepBeginNumberFrac,
// 	stateInNumberFrac:    stepInNumberFrac,
// 	stateBeginNumberExp:  stepBeginNumberExp,
// 	stateInNumberExp:     stepInNumberExp,
// 	stateEndNumber:       stepEnd,
// }

// Scanner is a JSON Scanner.
type scanner struct {
	state state
	pos   Position
	last  byte //The last token we had.
	err   *SyntaxError
	stack []state

	step func(*scanner, byte)

	yield bool
	// the stack of state
}

func (s *scanner) pushstate() {
	// fmt.Printf("push %v\n", s.state)
	s.stack = append(s.stack, s.state)
}

func (s *scanner) popstate() {
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
func (s *scanner) Reset() {
	s.pos.reset()
	s.err = nil
	s.yield = false
	s.stack = s.stack[0:0]
	// since a whitespace is valid at the start we use
	// a space instead of 0
	s.last = ' '
	s.state = stateBeginJSON
}

func (s *scanner) Error() error {
	if s.err != nil {
		return s.err
	}
	return nil
}

// Step is moving the Scanner state machine one step ahead
// using the byte provided as the next charachter.
func (s *scanner) Step(c byte) {

	if s.err != nil {
		return
	}

	s.pos.advance(c)

	s.step = steps[s.state]
	s.step(s, c)
	// fmt.Printf(" [%t] %v -> %v %v\n", s.yield, quoteChar(c), s.stack, s.state)
	for s.yield {
		s.yield = false
		s.step = steps[s.state]
		s.step(s, c)
		// fmt.Printf("![%t] %v -> %v %v\n", s.yield, quoteChar(c), s.stack, s.state)
	}
	s.last = c
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) {
	s.state = stateError
	s.err = &SyntaxError{
		Message:  "invalid character " + quoteChar(c) + " " + context,
		Position: s.pos,
	}
}

func stepError(s *scanner, c byte) {
	// Do nothing.
	return
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

	s.error(c, "expected whitespace or nothing")
}

func stepBeginObject(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	switch c {
	case '}':
		s.state = stateEndObject
	case '"':
		s.state = stateInObject
		s.pushstate()
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
		s.state = stateEndObject
	case ',':
		s.pushstate()
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
	s.pushstate()
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
	helperBeginValue(s, c)
}

func stepInMember(s *scanner, c byte) {
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

func stepBeginArray(s *scanner, c byte) {
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

func stepInArray(s *scanner, c byte) {

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

func stepBeginArrayValue(s *scanner, c byte) {
	// ignore white space nothing changes
	if c <= ' ' && isSpace(c) {
		return
	}

	// s.state = stateInArray
	// s.pushstate()
	helperBeginValue(s, c)
}

func helperBeginValue(s *scanner, c byte) {
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
		s.state = stateBeginNull
	case '-':
		s.state = stateBeginNumber
	case '0':
		s.state = stateBeginNumber
	default:
		s.error(c, "expected a value type")
	}
}

func stepBeginString(s *scanner, c byte) {
	s.state = stateInString
	stepInString(s, c)
}

func stepInString(s *scanner, c byte) {

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

// basic escaping, \uXXXX case is it's own state.
func stepInStringEscape(s *scanner, c byte) {

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
		s.popstate()
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
		s.state = stateEndBool
	}
}

func stepInTrue(s *scanner, c byte) {

	if (s.last == 't' && c != 'r') ||
		(s.last == 'r' && c != 'u') ||
		(s.last == 'u' && c != 'e') {
		s.error(c, "Unexpected value")
	} else if s.last == 'u' && c == 'e' {
		s.state = stateEndBool
	}
}

func stepBeginNull(s *scanner, c byte) {

	if s.last == 'n' && c != 'u' {
		s.error(c, "Unexpected value")
	} else {
		s.state = stateInNull
	}
}

func stepInNull(s *scanner, c byte) {

	if (s.last == 'n' && c != 'u') ||
		(s.last == 'u' && c != 'l') ||
		(s.last == 'l' && c != 'l') {
		s.error(c, "Unexpected value")
	} else if s.last == 'l' && c == 'l' {
		s.state = stateEndNull
	}
}

func stepBeginNumber(s *scanner, c byte) {

	// TODO: cleanup this!
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
			//Possibly just `0`
			s.state = stateEndNumber
			s.yield = true
		}
		return
	}

	if s.last == '-' {
		if '0' <= c && c <= '9' {
			s.state = stateInNumber
		} else {
			s.error(c, "expected digit")
		}
		return

	}

	if '1' <= s.last && s.last <= '9' {
		// after 1-9
		if '0' <= c && c <= '9' {
			s.state = stateInNumber
		} else {
			switch c {
			case 'e', 'E':
				s.state = stateBeginNumberExp
			case '.':
				s.state = stateBeginNumberFrac
			default:
				//Possibly just `0`
				s.state = stateEndNumber
				s.yield = true
			}
		}

		return
	}

	s.error(c, "scanner-error: In Begin Number state not after 0 or -")
}

func stepInNumber(s *scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	if c <= ' ' && isSpace(c) {
		s.state = stateEndNumber
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	case '.':
		s.state = stateBeginNumberFrac
	case ',', ']', '}':
		s.state = stateEndNumber
		s.yield = true
	default:
		s.error(c, "in number")
	}
}

func stepBeginNumberFrac(s *scanner, c byte) {
	// we have seen a '.' in a number
	if '0' <= c && c <= '9' {
		s.state = stateInNumberFrac
		return
	}

	s.error(c, "expected 0-9")
}

func stepInNumberFrac(s *scanner, c byte) {

	if '0' <= c && c <= '9' {
		return
	}

	switch c {
	case 'e', 'E':
		s.state = stateBeginNumberExp
	default:
		// Not digit, not eE, not must be end of number? yeild.
		s.state = stateEndNumber
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

	s.yield = true
	s.state = stateEndNumber
}

func stepEnd(s *scanner, c byte) {
	// fmt.Printf("%v:\n", quoteChar(c))
	// fmt.Printf("was: %v %v\n", s.stack, s.state)
	s.popstate() //
	if s.state == stateInArray && s.last == ',' {
		s.pushstate()
		s.state = stateBeginArrayValue
	}
	// fmt.Printf("is:  %v %v\n", s.stack, s.state)
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
