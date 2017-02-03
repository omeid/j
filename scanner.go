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

	stateBeginMember // seen "," in stateInObject
	stateInMember    // end of object or , + value
	stateInMemberKey
	stateEndMemberKey  // expect :
	stateInMemberValue // after :, expect a value type.

	stateBeginArray // value or end.
	stateInArray    // after value in array, expect Seperator Followed by Value or ].
	stateArrayValue // after , in a value, expect value type.

	// stateEndObjectValue   // object valued finished, expect key or end of object.

	stateBeginValue // find a value type.

	stateInString // string value.
	stateInStringEscape

	stateInFalse // false
	stateInTrue  // true
	stateInNull  // null

	stateError // Something is wrong, check s.err
)

var steps = map[State]func(*scanner, byte){
	stateBeginJSON: stepBeginJSON,
	stateEndJSON:   stepEndJSON,

	stateBeginObject:   stepBeginObject,
	stateInObject:      stepInObject,
	stateBeginMember:   stepBeginMember,
	stateInMemberKey:   stepInMemberKey,
	stateEndMemberKey:  stepEndMemberKey,
	stateInMemberValue: stepInMemberValue,
	stateInMember:      stepInMember,

	stateBeginArray: stepBeginArray,
	stateInArray:    stepInArray,
	stateArrayValue: stepArrayValue,

	stateInString:       stepInString,
	stateInStringEscape: stepInStringEscape,
	stateInFalse:        stepInFalse,
	stateInTrue:         stepInTrue,
	stateInNull:         stepInNull,
}

type scanner struct {
	state State
	pos   Position
	last  byte //The last token we had.
	err   *SyntaxError
	stack []State
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

	// if there is an error, do not step.
	if s.err != nil {
		return
	}

	defer func() {
		//TODO: REMOVE POST DEBUG.
		// fmt.Printf("%c -> %v %s\n", c, s.stack, s.state)
		s.last = c
	}()

	s.pos.advance(c)
	// steps[s.state](s, c)

	step, ok := steps[s.state]
	if !ok {
		// fmt.Printf("Missing handler for state %s", s.state)
		panic("Whoops")
	}
	step(s, c)
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
		s.state = stateInMemberKey
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
		s.state = stateInMemberKey
	} else {
		s.error(c, "expected string.")
	}
}

func stepInMemberKey(s *scanner, c byte) {
	if c == '"' && s.last != '\\' {
		s.state = stateEndMemberKey // find colon ':'
	}
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
	// ignore white space nothing changes
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

	if s.last != '\\' {
		// we shouldn't be here!
		s.error(c, "parser-error: in escape-state without Escape.")
		return
	}

	switch c {
	case '"', '\\', '/', 'b', 'f', 'n', 'r', 't':
		s.popState() // end of escape.
	case 'u':
		s.pushState()
		// s.state = // HexLiteral.
	default:
		s.error(c, "Expected a qoutation mark, reverse solidus, or a control character")
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
