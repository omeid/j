package scanner

import "fmt"

// Position is a location in the json document, used for errors.
type Position struct {
	Line   int
	Column int
	Offset int
}

func (p *Position) advance(c byte) {
	p.Offset++

	if c == '\r' || c == '\n' {
		p.Line++
		p.Column = 1
	} else {
		p.Column++
	}

	if p.Line == 0 {
		p.Line = 1
	}
}

func (p *Position) reset() {
	p.Line = 0
	p.Column = 0
	p.Offset = 0
}

func (p *Position) string() string {
	return fmt.Sprintf(
		"%d:%d (%d) ",
		p.Line,
		p.Column,
		p.Offset,
	)
}
