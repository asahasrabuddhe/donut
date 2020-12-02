// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package ast

import (
	"fmt"
	"go.ajitem.com/donut/token"
	"strings"
)

// return <expression>;

type ReturnStatement struct {
	Token token.Token
	Value Expression
}

func (r *ReturnStatement) statementNode() {}

func (r *ReturnStatement) TokenLiteral() string {
	return r.Token.Literal
}

func (r *ReturnStatement) String() string {
	var out strings.Builder

	out.WriteString(fmt.Sprintf("%s ", r.TokenLiteral()))

	if r.Value != nil {
		out.WriteString(r.Value.String())
	}

	out.WriteString(";")

	return out.String()
}
