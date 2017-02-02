package json

import "strconv"

// Valid check if the provide data is correct json document.
func Valid(src []byte) error {
	var scan scanner

	scan.reset()
	for _, c := range src {
		if scan.step(c) == stateError {
			return scan.err
		}
	}

	// scan.eof() // check for the end.
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

// State is the current state of the scanner state machine.
type State int

// Tokens as per the JSON spec.
const (
	stateBeginJSON State = iota // At the beginning there was JSON.
	// stateWhitespace
	stateInObject  // we are in an object
	stateEndObject // object properly ended
	stateInArray   // we are in an array

	stateInObjectKey      // we are in an object key
	stateEndObjectKey     // object key finished, expect value.
	stateBeginObjectValue //process object value
	stateEndObjectValue   // object valued finished, expect key or end of object.

	stateBeginValue // find a value type.

	stateError // Something is wrong, check s.err
)

// parseTarget is the type of object being parsed.
// type parseTarget int

// Type of objects we are parsing, kept in the scanner stack.
// const (
// 	memberKey parseTarget = iota
// 	memberValue
// 	arrayvalue
// )

type scanner struct {
	state State
	pos   Position
	last  byte //The last token we had.
	err   *SyntaxError
	stack []State
	// the stack of state
}

func (s scanner) reset() {
	s.pos.reset()
	s.err = nil
	// since a whitespace is valid at the start we use
	// a space instead of 0
	s.last = ' '
	s.state = stateBeginJSON
}

// error records an error and switches to the error state.
func (s *scanner) error(c byte, context string) State {
	s.state = stateError
	s.err = &SyntaxError{
		Message:  "invalid character " + quoteChar(c) + " " + context,
		Position: s.pos,
	}
	return s.state
}

func (s *scanner) step(c byte) State {

	if s.err != nil {
		return stateError
	}

	defer func() {
		s.last = c
	}()

	s.pos.advance(c)

	switch s.state {

	case stateBeginJSON:
		return stepBeginJSON(s, c)
	case stateInObject:
		return stepInObject(s, c)
	case stateInObjectKey:
		return stepInObjectKey(s, c)
	case stateEndObjectKey:
		return stepBeginObjectValue(s, c)
	case stateBeginValue:
		return stepBeginValue(s, c)
		// case stateInArray:
		// return stepInArray(s, c)
	}
	// stateEndObject

	return s.state
}

func stepBeginJSON(s *scanner, c byte) State {
	if c <= ' ' && isSpace(c) {
		return s.state // nothing changes
	}

	if c == '{' {
		s.state = stateInObject
		// s.pushParseState(parseObjectKey)
		return s.state
	}

	if c == '[' {
		s.state = stateInArray
		// s.pushParseState(parseArrayValue)
		return s.state
	}

	return s.error(c, "expected beginning of json ('{' or '[')")
}

func stepInObject(s *scanner, c byte) State {
	if c <= ' ' && isSpace(c) {
		return s.state // nothing changes
	}

	if c == '}' {
		s.state = stateEndObject
		// s.pushParseState(parseObjectKey)
		return s.state
	}

	if c == '"' {
		s.state = stateInObjectKey
		// s.pushParseState(parseArrayValue)
		return s.state
	}

	return s.error(c, "expected a pair or end of object")
}

func stepInObjectKey(s *scanner, c byte) State {
	if c == '"' && s.last != '\\' {
		s.state = stateEndObjectKey
	}
	return s.state
}

func stepBeginObjectValue(s *scanner, c byte) State {

	if c <= ' ' && isSpace(c) {
		return s.state // nothing changes
	}

	if c != ':' {
		return s.error(c, "expected a colon")
	}

	s.state = stateBeginValue

	return s.state
}

func stepBeginValue(s *scanner, c byte) State {
	if c <= ' ' && isSpace(c) {
		return s.state // nothing changes
	}

	if c == '{' {
		s.state = stateInObject
		// s.pushParseState(parseObjectKey)
		return s.state
	}

	if c == '[' {
		s.state = stateInArray
		// s.pushParseState(parseArrayValue)
		return s.state
	}

	return s.error(c, "expected a value type")
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
