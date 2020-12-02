// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	functionLiteral := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeek(token.LeftParenthesis) {
		return nil
	}

	functionLiteral.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LeftBrace) {
		return nil
	}

	functionLiteral.Body = p.parseBlockStatement()

	return functionLiteral
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var identifiers []*ast.Identifier

	if p.peekTokenIs(token.RightParenthesis) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	identifier := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	identifiers = append(identifiers, identifier)

	for p.peekTokenIs(token.Comma) {
		p.nextToken()
		p.nextToken()

		identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		identifiers = append(identifiers, identifier)
	}

	if !p.expectPeek(token.RightParenthesis) {
		return nil
	}

	return identifiers
}
