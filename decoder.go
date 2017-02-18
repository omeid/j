package j

import (
	"errors"
	"fmt"
	// "fmt"
	"io"

	"github.com/omeid/j/strings"
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

type decoder struct {
	scan *Scanner
	read *reader

	b   byte
	err error
}

func (d *decoder) step() {
	d.b, d.err = d.read.Next()
	if d.err != nil {
		return
	}

	d.scan.Step(d.b)
	d.err = d.scan.Error()
}

// Decode loads your input int a Value object.
func Decode(input []byte) (Value, error) {

	var scan Scanner
	// fmt.Printf("\n\nvalid:\n")
	err := valid(input, &scan)
	if err != nil {
		// fmt.Printf("\n\n  !!! Failed At Valid !!! \n\n")
		return nil, err
	}
	scan.Reset()

	read := &reader{src: input, index: 0}

	d := &decoder{
		scan: &scan,
		read: read,
	}

	// fmt.Printf("\n\ndecode:\n")
	for {
		if scan.state != stateBeginJSON {
			break
		}
		d.step()
		if d.err != nil {
			return nil, err
		}
	}

	var v Value
	switch d.scan.state {
	case stateBeginArray:
		// fmt.Println("array document")
		v = nextArray(d)
	case stateBeginObject:
		// fmt.Println("object document")
		v = nextObject(d)
	default:
		return nil, ParseError{errors.New("parser-error: unexpected state"), read.index}
	}

	if d.err != nil && d.err != io.EOF {
		return nil, d.err
	}

	return v, nil
}

func nextValue(d *decoder) Value {

	switch d.scan.state {
	case stateBeginObject:
		// fmt.Println("::next object")
		return nextObject(d)
	case stateBeginArray:
		// fmt.Println("::next array")
		return nextArray(d)
	case stateBeginNumber:
		// fmt.Println("::next number")
		return nextNumber(d)
	case stateBeginString:
		// fmt.Println("::next string")
		return nextString(d)
	case stateBeginNull:
		// fmt.Println("::next null")
		return nextNull(d)
	case stateInFalse:
		// fmt.Println("::next false")
		return nextFalse(d)
	case stateInTrue:
		// fmt.Println("::next true")
		return nextTrue(d)

	default:
		ss := fmt.Sprintf("Whop the dup. No handler for state %s %s", d.scan.state, d.scan.pos.string())

		panic(ss)
	}

}

func nextObject(d *decoder) Value {
	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateBeginObject {
		panic("called nextObject while not in stateBeginObject")
	}

	for {
		d.step()
		if d.err != nil {
			return nil
		}
		if d.scan.state != stateBeginObject {
			break
		}
	}

	if d.scan.state == stateEndObject {
		d.step()
		if d.err != nil {
			return nil
		}
		return NewObject(nil)
	}

	// depth := len(scan.stack)

	members := make([]Member, 0, 40) //Good enough?

	for {

		for d.scan.state == stateInObject {
			d.step()
			if d.err != nil {
				return nil
			}
		}

		if d.scan.state == stateEndObject {
			d.step()
			if d.err != nil && d.err != io.EOF {
				return nil
			}
			// fmt.Printf("breaking %v %c\n", d.scan.state, scan.last)
			break
		}
		// fmt.Printf("nextObject didn't break %v %c\n", d.scan.state, scan.last)

		// fmt.Printf("yield to nextkv  %v %c\n", d.scan.state, scan.last)
		m := nextkv(d)
		if d.err != nil {
			return nil
		}

		members = append(members, m)
	}

	return NewObject(members)
}

func nextkv(d *decoder) Member {

	for d.scan.state == stateBeginMember {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	if d.scan.state != stateBeginMemberKey {
		// fmt.Printf("Position: %v\n", scan.pos.string())
		panic("expected state begin member key")
	}

	key := make([]byte, 0, 30)

	for {

		d.step()
		if d.err != nil {
			return nil
		}

		if d.scan.state != stateInString {
			break
		}

		key = append(key, d.b)
	}

	if d.scan.state != stateEndString {
		// // fmt.Println(d.scan.state)
		panic("expected end of string for member key")
	}

	d.step() // step over "
	if d.err != nil {
		return nil
	}

	// any possible whitespace.
	for d.scan.state == stateEndMemberKey {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	for d.scan.state == stateInMemberValue {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	value := nextValue(d)
	if d.err != nil {
		return nil
	}

	//TODO: escape the key!
	m := NewMember("", string(key), value)

	// fmt.Printf("returning a member\n")
	return m
}

func nextArray(d *decoder) Value {
	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateBeginArray {
		panic("called parseArray while not in stateBeginArray")
	}

	sd := len(d.scan.stack)

	for {
		d.step()
		if d.err != nil {
			return nil
		}
		if d.scan.state != stateBeginArray || len(d.scan.stack) != sd {
			break
		}
	}

	if d.scan.state == stateEndArray {
		d.step()
		if d.err != nil && d.err != io.EOF {
			return nil
		}
		return NewArray(nil)
	}

	values := []Value{}
	for {
		v := nextValue(d)
		if d.err != nil && d.err != io.EOF {
			return nil
		}

		values = append(values, v)

		for d.scan.state == stateBeginArrayValue || d.scan.state == stateInArray {
			d.step()
			if d.err != nil && d.err != io.EOF {
				return nil
			}
		}

		if d.scan.state == stateEndArray && len(d.scan.stack) == sd {
			d.step()
			if d.err != nil && d.err != io.EOF {
				return nil
			}
			break
		}

		if d.scan.state == stateEndArray && len(d.scan.stack) == sd {
			d.step()
			if d.err != nil && d.err != io.EOF {
				return nil
			}
			break
		}
	}
	return NewArray(values)
}

func nextString(d *decoder) Value {

	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateBeginString {
		panic("called nextString while not in stateBeginString")
	}

	var s []byte
	d.step() // step over "
	if d.err != nil {
		return nil
	}

	if d.scan.state == stateEndString {
		d.step()
		if d.err != nil {
			return nil
		}
		return NewString("")
	}

	s = append(s, d.b)

	for {

		d.step()

		if d.err != nil {
			fmt.Printf("errxr: %v\n", d.err)
			return nil
		}

		// fmt.Printf("state %v\n", d.scan.state)
		if d.scan.state == stateInString ||
			(d.scan.state >= stateInStringEscape &&
				d.scan.state <= stateInStringEscapeUxxxx) {
			s = append(s, d.b)
		} else {
			// fmt.Printf("breaking %v\n", d.scan.state)
			break
		}
	}

	if d.scan.state != stateEndString {
		panic("Unexpected end of string")
	}

	d.step() // step over "
	if d.err != nil {
		return nil
	}

	str := strings.Decode(s)
	return NewString(str)
}

func nextNumber(d *decoder) Value {

	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateBeginNumber {
		panic("called nextNumber while not in stateBeginNumber")
	}

	var s []byte
	// the value for stateBeginNumber is part of the number value.
	d.b = d.scan.last

	for {

		s = append(s, d.b)

		d.step()
		if d.err != nil {
			return nil
		}

		if d.scan.state >= stateInNumber && d.scan.state <= stateInNumberExp {
			continue
		}

		break
	}

	if d.scan.state == stateEndNumber {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	// fmt.Println("leaving nextNumber")
	return NewNumber(s)
}

func nextNull(d *decoder) Value {

	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateBeginNull {
		panic("called nextNull while not in stateBeginNull")
	}

	for {

		d.step()
		if d.err != nil {
			return nil
		}

		if d.scan.state != stateInNull {
			break
		}

	}

	if d.scan.state != stateEndNull {
		panic("Unexpected end of null")
	}

	//d.step over last `l`.
	d.step()
	if d.err != nil {
		return nil
	}

	return NewNull()
}

func nextFalse(d *decoder) Value {

	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateInFalse {
		panic("called nextFalse while not in stateBeginFalse")
	}

	for d.scan.state == stateInFalse {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	if d.scan.state != stateEndBool {
		panic("Unexpected end of false")
	}

	//d.step over last `stateEndBool`.
	d.step()
	if d.err != nil {
		return nil
	}

	return NewBool(false)
}

func nextTrue(d *decoder) Value {

	//scanner is in d.scan.stateBeginArray
	if d.scan.state != stateInTrue {
		panic("called nextTrue while not in stateBeginFalse")
	}

	for d.scan.state == stateInTrue {
		d.step()
		if d.err != nil {
			return nil
		}
	}

	if d.scan.state != stateEndBool {
		panic("Unexpected end of false")
	}

	d.step()
	if d.err != nil {
		return nil
	}

	return NewBool(true)
}
