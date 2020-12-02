// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package ast

import (
	"fmt"
	"go.ajitem.com/donut/token"
	"strings"
)

// let <identifier> = <expression>;

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) TokenLiteral() string {
	return l.Token.Literal
}

func (l *LetStatement) String() string {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("%s %s = ", l.TokenLiteral(), l.Name.String()))

	if l.Value != nil {
		out.WriteString(l.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
