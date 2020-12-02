// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package ast

import (
	"go.ajitem.com/donut/internal/token"
	"strings"
)

type BlockStatement struct {
	Token      token.Token
	Statements Statements
}

func (b *BlockStatement) statementNode() {}

func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) String() string {
	var out strings.Builder

	for _, statement := range b.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}
