package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/token"
)

func (p *Parser) parseFunctionLiteral() ast.Expression {
	functionLiteral := &ast.FunctionLiteral{Token: p.currentToken}

	if !p.expectPeek(token.LPAREN) {
		return nil
	}

	functionLiteral.Parameters = p.parseFunctionParameters()

	if !p.expectPeek(token.LBRACE) {
		return nil
	}

	functionLiteral.Body = p.parseBlockStatement()

	return functionLiteral
}

func (p *Parser) parseFunctionParameters() []*ast.Identifier {
	var identifiers []*ast.Identifier

	if p.peekTokenIs(token.RPAREN) {
		p.nextToken()
		return identifiers
	}

	p.nextToken()

	identifier := &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
	identifiers = append(identifiers, identifier)

	for p.peekTokenIs(token.COMMA) {
		p.nextToken()
		p.nextToken()

		identifier = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}
		identifiers = append(identifiers, identifier)
	}

	if !p.expectPeek(token.RPAREN) {
		return nil
	}

	return identifiers
}
