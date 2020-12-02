// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
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
