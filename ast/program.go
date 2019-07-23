package ast

import "bytes"

type Program struct {
	Statements Statements
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		_, _ = out.WriteString(s.String())
	}

	return out.String()
}
