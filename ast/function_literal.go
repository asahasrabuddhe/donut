package ast

import (
	"go.ajitem.com/donut/token"
	"strings"
)

// func <parameters> <block statement>

type FunctionLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (f *FunctionLiteral) expressionNode() {}

func (f *FunctionLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FunctionLiteral) String() string {
	var out strings.Builder

	var params []string

	for _, parameter := range f.Parameters {
		params = append(params, parameter.String())
	}

	out.WriteString(f.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(")")
	out.WriteString(f.Body.String())

	return out.String()
}
