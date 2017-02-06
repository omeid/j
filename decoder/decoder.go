package decoder

import (
	"errors"
	"fmt"

	"github.com/omeid/j"
	"github.com/omeid/j/mutable"
)

// ParseError holds an error and offset, the error maybe a
// a SyntaxError
type ParseError struct {
	err    error
	offset int
}

func (pe ParseError) Error() string {
	return fmt.Sprintf("<%d> %v", pe.offset, pe.err)
}

// Decode loads your input int a j.Value object.
func Decode(input []byte) (j.Value, error) {

	err := Valid(input)
	if err != nil {
		return nil, err
	}

	var scan Scanner
	read := &reader{src: input, index: 0}

	scan.Reset()

	for scan.state == stateBeginJSON {
		_, err := step(&scan, read)
		if err != nil {
			return nil, err
		}
	}

	switch scan.state {
	case stateBeginArray:
		return nextArray(&scan, read)
	case stateBeginObject:
		return parseObject(&scan, read)
	default:
		return nil, ParseError{errors.New("parser-error: unexpected state"), read.index}
	}
}

func parseObject(scan *Scanner, read *reader) (j.Value, error) {
	return mutable.NewObject().Value(), nil
}

func nextValue(scan *Scanner, read *reader) (j.Value, error) {

	switch scan.state {
	case stateBeginObject:
	case stateBeginArray:
		return nextArray(scan, read)
	case stateBeginMember:
	case stateBeginNumber:
	case stateBeginString:
		return nextString(scan, read)
	default:
		ss := fmt.Sprintf("Whop the dope. No handler for state %s", scan.state)

		panic(ss)
	}

	return nil, nil
}

func nextArray(scan *Scanner, read *reader) (j.Value, error) {
	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginArray {
		panic("called parseArray while not in stateBeginArray")
	}

	var err error
	for scan.state == stateBeginArray {
		_, err = step(scan, read)
		if err != nil {
			return nil, err
		}
	}

	arr := mutable.NewArray()
	if scan.state == stateEndArray {
		return arr.Value(), nil
	}

	// depth := len(scan.stack)

	for {
		// for len(scan.stack) > depth || scan.state != StateEndArray {
		v, err := nextValue(scan, read)
		if err != nil {
			return nil, err
		}

		arr.Add(v)

		if scan.state == stateEndArray {
			break
		}

		for scan.state == stateBeginArrayValue {
			_, err = step(scan, read)
			if err != nil {
				return nil, err
			}
		}
	}
	return arr.Value(), nil
}

func nextString(scan *Scanner, read *reader) (j.Value, error) {

	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginString {
		panic("called parseArray while not in stateBeginArray")
	}

	var err error
	var s []byte
	var b byte

	b, err = step(scan, read) // step over "
	if err != nil {
		return nil, err
	}

	s = append(s, b)

	for {
		b, err = read.Next()
		if err != nil {
			return nil, err
		}
		scan.Step(b)
		if scan.Error() != nil {
			return nil, err
		}

		if scan.state == stateInString {
			s = append(s, b)
		} else {
			break
		}
	}

	if scan.state != stateEndString {
		panic("Expected end of string")
	}

	_, err = step(scan, read) // step over "
	if err != nil {
		return nil, err
	}

	return mutable.NewString(s).Value(), nil
}
