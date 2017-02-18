package j

import (
	"bytes"
	"errors"
	"io"

	"github.com/omeid/j/strings"
)

var (
	// ErrInvalidType means your Value is of InvalidType.
	ErrInvalidType = errors.New("Invalid Value")
	// ErrNotDocument means you have provided a non-array or non-object type where
	// a valid JSON Document type is expected.
	ErrNotDocument = errors.New("Not a JSON Document. Expected Array or Object Type")
)

type writer interface {
	io.Writer
	io.ByteWriter
}

// Encode creates a json document the provided j.Value
func Encode(value Value) ([]byte, error) {

	var out bytes.Buffer
	var err error

	switch value.Type() {
	case InvalidType:
		err = ErrInvalidType
	case ObjectType:
		err = writeObject(&out, value)
	case ArrayType:
		err = writeArray(&out, value)
	default:
		err = ErrNotDocument
	}

	if err != nil {
		return nil, err
	}

	return out.Bytes(), nil
}

func writeValue(out writer, v Value) error {
	var err error

	switch v.Type() {
	case InvalidType:
		err = ErrInvalidType
	case ObjectType:
		err = writeObject(out, v)
	case ArrayType:
		err = writeArray(out, v)
	case BoolType:
		err = writeBool(out, v)
	case NumberType:
		err = writeNumber(out, v)
	case StringType:
		err = writeString(out, v)
	case NullType:
		err = writeNull(out)
	}

	return err
}

func writeObject(out writer, v Value) error {

	out.WriteByte('{')

	ms := v.Members()
	l := len(ms) - 1

	var err error
	for i, m := range ms {

		err = writeMember(out, m)
		if err != nil {
			return err
		}

		if i < l {
			out.WriteByte(',')
		}
	}

	out.WriteByte('}')

	return nil
}

func writeArray(out writer, v Value) error {

	var err error

	err = out.WriteByte('[')
	if err != nil {
		return err
	}

	vs := v.Values()
	l := len(vs) - 1

	for i, v := range vs {

		if v == nil {
			// fmt.Printf("value %v\nvs: %v\n", v, vs)
			panic("Nil value!")
		}

		err = writeValue(out, v)
		if err != nil {
			return err
		}

		if i < l {
			err = out.WriteByte(',')
			if err != nil {
				return err
			}
		}
	}

	err = out.WriteByte(']')

	return err
}

func writeMember(out writer, m Member) error {

	err := out.WriteByte('"')
	if err != nil {
		return err
	}
	_, err = out.Write([]byte(m.Name()))
	if err != nil {
		return err
	}

	err = out.WriteByte('"')
	if err != nil {
		return err
	}

	err = out.WriteByte(':')
	if err != nil {
		return err
	}

	// fmt.Printf("will write value for %v\n", m.Name())
	return writeValue(out, m.Value())
}

func writeBool(out writer, v Value) error {

	var err error
	if v.Bool() {
		_, err = out.Write([]byte(`true`))
	} else {
		_, err = out.Write([]byte(`false`))
	}

	return err
}

func writeNull(out writer) error {
	_, err := out.Write([]byte(`null`))
	return err
}

func writeString(out writer, v Value) error {
	err := out.WriteByte('"')
	if err != nil {
		return err
	}

	sb := strings.Encode(string(v.String()))
	_, err = out.Write(sb)
	if err != nil {
		return err
	}
	err = out.WriteByte('"')
	if err != nil {
		return err
	}
	return err
}

func writeNumber(out writer, v Value) error {
	_, err := out.Write(v.Raw())
	return err
}
