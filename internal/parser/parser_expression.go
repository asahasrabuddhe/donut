// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/token"
)

const (
	_ int = iota
	Lowest
	Equals
	LessOrGreater
	Sum
	Product
	Prefix
	Call
)

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	stmt := &ast.ExpressionStatement{Token: p.currentToken}

	stmt.Expression = p.parseExpression(Lowest)

	if p.peekTokenIs(token.Semicolon) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParsers[p.currentToken.Type]

	if prefix == nil {
		p.noPrefixParseFunctionError(p.currentToken.Type)
		return nil
	}

	leftExp := prefix()

	postfix := p.postfixParsers[p.peekToken.Type]
	if postfix != nil {
		return postfix(leftExp)
	}

	for !p.peekTokenIs(token.Semicolon) && precedence < p.peekTokenPrecedence() {
		infix := p.infixParsers[p.peekToken.Type]
		if infix == nil {
			return leftExp
		}

		p.nextToken()

		leftExp = infix(leftExp)
	}

	return leftExp
}
