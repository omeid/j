package j

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

}

func (p *Position) reset() {
	p.Line = 1
	p.Column = 0
	p.Offset = 0
}
