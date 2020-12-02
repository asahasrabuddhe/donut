// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/token"
)

func (p *Parser) parseIfExpression() ast.Expression {
	expression := &ast.IfExpression{Token: p.currentToken}

	if !p.expectPeek(token.LeftParenthesis) {
		return nil
	}

	p.nextToken()
	expression.Condition = p.parseExpression(Lowest)

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	expression.Consequence = p.parseBlockStatement()

	if p.peekTokenIs(token.Else) {
		p.nextToken()

		if !p.expectPeek(token.LeftBrace) {
			return nil
		}

		expression.Alternative = p.parseBlockStatement()
	}

	return expression
}
