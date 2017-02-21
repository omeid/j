package j

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
