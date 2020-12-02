// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package ast

import (
	"fmt"
	"go.ajitem.com/donut/internal/token"
	"strings"
)

// <expression> <infix operator> <expression>;

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (i *InfixExpression) expressionNode() {}

func (i *InfixExpression) TokenLiteral() string {
	return i.Token.Literal
}

func (i *InfixExpression) String() string {
	var out strings.Builder

	out.WriteString("(")
	out.WriteString(i.Left.String())
	out.WriteString(fmt.Sprintf(" %s ", i.Operator))
	out.WriteString(i.Right.String())
	out.WriteString(")")

	return out.String()
}
