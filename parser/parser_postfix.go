// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
)

func (p *Parser) parsePostfixExpression(left ast.Expression) ast.Expression {
	expression := &ast.PostfixExpression{
		Token:    p.peekToken,
		Operator: p.peekToken.Literal,
	}

	p.nextToken()

	expression.Left = left

	return expression
}
