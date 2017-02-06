package decoder

import "io"

type reader struct {
	src   []byte
	index int // current reading index
}

func (r *reader) Next() (byte, error) {

	if r.index >= len(r.src) {
		return 0, io.EOF
	}

	b := r.src[r.index]

	r.index++
	return b, nil
}

func step(scan *Scanner, read *reader) (byte, error) {
	var b byte
	var err error
	b, err = read.Next()
	if err != nil {
		return b, err
	}
	scan.Step(b)
	if scan.Error() != nil {
		return b, err
	}
	return b, nil
}
