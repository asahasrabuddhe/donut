// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package ast

import (
	"go.ajitem.com/donut/token"
	"strings"
)

// <expression> <postfix operator>;

type PostfixExpression struct {
	Left     Expression
	Token    token.Token
	Operator string
}

func (p PostfixExpression) expressionNode() {}

func (p PostfixExpression) TokenLiteral() string {
	return p.Token.Literal
}

func (p PostfixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(p.Left.String())
	out.WriteString(p.Operator)
	out.WriteString(")")

	return out.String()
}
