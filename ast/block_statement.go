// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package ast

import (
	"bytes"
	"go.ajitem.com/donut/token"
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
	var out bytes.Buffer

	for _, statement := range b.Statements {
		out.WriteString(statement.String())
	}

	return out.String()
}
