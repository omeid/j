package codec

import (
	"errors"
	"fmt"
	// "fmt"
	"io"

	"github.com/omeid/j"
	"github.com/omeid/j/codec/strings"
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

	var scan Scanner
	// fmt.Printf("\n\nvalid:\n")
	err := valid(input, &scan)
	if err != nil {
		// fmt.Printf("\n\n  !!! Failed At Valid !!! \n\n")
		return nil, err
	}

	read := &reader{src: input, index: 0}

	scan.Reset()

	// fmt.Printf("\n\ndecode:\n")
	for {
		if scan.state != stateBeginJSON {
			break
		}
		_, err := step(&scan, read)
		if err != nil {
			return nil, err
		}
	}

	switch scan.state {
	case stateBeginArray:
		// fmt.Println("array document")
		return nextArray(&scan, read)
	case stateBeginObject:
		// fmt.Println("object document")
		return nextObject(&scan, read)
	default:
		return nil, ParseError{errors.New("parser-error: unexpected state"), read.index}
	}
}

func nextValue(scan *Scanner, read *reader) (j.Value, error) {

	switch scan.state {
	case stateBeginObject:
		// fmt.Println("::next object")
		return nextObject(scan, read)
	case stateBeginArray:
		// fmt.Println("::next array")
		return nextArray(scan, read)
	case stateBeginNumber:
		// fmt.Println("::next number")
		return nextNumber(scan, read)
	case stateBeginString:
		// fmt.Println("::next string")
		return nextString(scan, read)

	default:
		ss := fmt.Sprintf("Whop the dup. No handler for state %s", scan.state)

		panic(ss)
	}

}

func nextObject(scan *Scanner, read *reader) (j.Value, error) {
	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginObject {
		panic("called nextObject while not in stateBeginObject")
	}

	var err error
	for {
		_, err = step(scan, read)
		if err != nil {
			return nil, err
		}
		if scan.state != stateBeginObject {
			break
		}
	}

	obj := mutable.NewObject()
	if scan.state == stateEndObject {
		_, err = step(scan, read)
		if err != nil {
			return nil, err
		}
		return obj.Value(), nil
	}

	// depth := len(scan.stack)

	for {
		// for len(scan.stack) > depth || scan.state != StateEndArray {
		if scan.state == stateEndObject || scan.state == stateEndJSON {
			break
		}

		m, err := nextkv(scan, read)
		if err != nil {
			return nil, err
		}

		obj.Add(m)

	}

	return obj.Value(), nil
}

func nextkv(scan *Scanner, read *reader) (j.Member, error) {

	for scan.state == stateBeginMember {
		_, err := step(scan, read) //step over "
		if err != nil {
			return nil, err
		}
	}

	if scan.state != stateBeginMemberKey {
		panic("expected state begin member key")
	}

	var key []byte

	for {

		b, err := step(scan, read)
		if err != nil {
			return nil, err
		}

		if scan.state != stateInString {
			break
		}

		key = append(key, b)
	}

	if scan.state != stateEndString {
		// // fmt.Println(scan.state)
		panic("expected end of string for member key")
	}

	_, err := step(scan, read) //step over "
	if err != nil {
		return nil, err
	}

	for scan.state == stateEndMemberKey {
		_, err := step(scan, read) //step over "
		if err != nil {
			return nil, err
		}
	}

	for scan.state == stateInMemberValue {
		_, err := step(scan, read) //step over "
		if err != nil {
			return nil, err
		}
	}

	value, err := nextValue(scan, read)
	if err != nil {
		return nil, err
	}

	//TODO: escape the key!
	m := mutable.NewMember("", string(key), value)

	return m, nil
}

func nextArray(scan *Scanner, read *reader) (j.Value, error) {
	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginArray {
		panic("called parseArray while not in stateBeginArray")
	}

	sd := len(scan.stack)
	var err error

	for {
		_, err = step(scan, read)
		if err != nil {
			return nil, err
		}
		if scan.state != stateBeginArray || len(scan.stack) != sd {
			break
		}
	}

	arr := mutable.NewArray()
	if scan.state == stateEndArray {
		_, err = step(scan, read)
		if err != nil && err != io.EOF {
			return nil, err
		}
		return arr.Value(), nil
	}

	for {
		v, err := nextValue(scan, read)
		if err != nil {
			// fmt.Printf("errr %v\n", err)
			return nil, err
		}

		arr.Add(v)

		if scan.state == stateEndArray && len(scan.stack) == sd {
			_, err = step(scan, read)
			if err != nil && err != io.EOF {
				return nil, err
			}
			break
		}

		for scan.state == stateBeginArrayValue || scan.state == stateInArray || scan.state == stateEndNumber {
			_, err = step(scan, read)
			if err != nil && err != io.EOF {
				return nil, err
			}
		}

		if scan.state == stateEndArray && len(scan.stack) == sd {
			_, err = step(scan, read)
			if err != nil && err != io.EOF {
				return nil, err
			}
			break
		}
	}
	return arr.Value(), nil
}

func nextString(scan *Scanner, read *reader) (j.Value, error) {

	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginString {
		panic("called nextString while not in stateBeginString")
	}

	var err error
	var s []byte
	var b byte

	b, err = step(scan, read) // step over "
	if err != nil {
		return nil, err
	}

	if scan.state == stateEndString {
		_, err = step(scan, read)
		if err != nil {
			return nil, err
		}
		//TODO: escape string
		str := strings.Decode(s)
		return mutable.NewString(str).Value(), nil
	}

	s = append(s, b)

	for {

		b, err = step(scan, read)

		if err != nil {
			// fmt.Printf("errxr: %v\n", err)
			return nil, err
		}

		// fmt.Printf("state %v\n", scan.state)
		if scan.state == stateInString ||
			(scan.state >= stateInStringEscape &&
				scan.state <= stateInStringEscapeUxxxx) {
			s = append(s, b)
		} else {
			// fmt.Printf("breaking %v\n", scan.state)
			break
		}
	}

	if scan.state != stateEndString {
		panic("Unexpected end of string")
	}

	_, err = step(scan, read) // step over "
	if err != nil {
		return nil, err
	}

	str := strings.Decode(s)
	return mutable.NewString(str).Value(), nil
}

func nextNumber(scan *Scanner, read *reader) (j.Value, error) {

	//scanner is in scan.stateBeginArray
	if scan.state != stateBeginNumber {
		panic("called nextNumber while not in stateBeginNumber")
	}

	var err error
	var s []byte
	var b byte

	// the value for stateBeginNumber is part of the number value.
	b = scan.last

	for {

		s = append(s, b)

		b, err = step(scan, read)
		if err != nil {
			return nil, err
		}

		if scan.state >= stateInNumber && scan.state <= stateInNumberExp {
			continue
		}

		break
	}

	return mutable.NewNumber(s).Value(), nil
}
