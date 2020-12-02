package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/token"
)

func (p *Parser) parseCallExpression(function ast.Expression) ast.Expression {
	callExpression := &ast.CallExpression{Token: p.currentToken, Function: function}
	callExpression.Arguments = p.parseCallArguments()

	return callExpression
}

func (p *Parser) parseCallArguments() []ast.Expression {
	var arguments []ast.Expression

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return arguments
	}

	p.nextToken()
	arguments = append(arguments, p.parseExpression(LOWEST))

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()
		arguments = append(arguments, p.parseExpression(LOWEST))
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return arguments
}
