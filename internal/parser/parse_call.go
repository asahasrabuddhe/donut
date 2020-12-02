// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/token"
)

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	callExpression := &ast.CallExpression{Token: p.currentToken, Function: function}
	callExpression.Arguments = p.parseCallArguments()

	return callExpression
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var arguments []ast.Expression

	if p.peekTokenIs(token.RightParenthesis) {
		p.nextToken()
		return arguments
	}

	p.nextToken()
	arguments = append(arguments, p.parseExpression(Lowest))

	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()
		arguments = append(arguments, p.parseExpression(Lowest))
	}

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	return arguments
}
