// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package ast

import "go.ajitem.com/donut/token"

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (e *ExpressionStatement) statementNode() {}

func (e *ExpressionStatement) TokenLiteral() string {
	return e.Token.Literal
}

func (e *ExpressionStatement) String() string {
	if e.Expression != nil {
		return e.Expression.String()
	}

	return ""
}
