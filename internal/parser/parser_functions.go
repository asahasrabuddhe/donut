// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/token"
)

func (p *Parser) parseStatement() ast.Statement {
	switch p.currentToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) currentTokenIs(t token.Type) bool {
	return p.currentToken.Type == t
}

func (p *Parser) peekTokenIs(t token.Type) bool {
	return p.peekToken.Type == t
}

func (p *Parser) expectPeek(t token.Type) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		p.peekError(t)
		return false
	}
}

func (p *Parser) peekError(t token.Type) {
	err := fmt.Errorf("expected next to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, err)
}

func (p *Parser) registerPrefixParser(t token.Type, fn prefixParser) {
	p.prefixParsers[t] = fn
}

func (p *Parser) registerInfixParser(t token.Type, fn infixParser) {
	p.infixParsers[t] = fn
}

func (p *Parser) registerPostfixParser(t token.Type, fn postfixParser) {
	p.postfixParsers[t] = fn
}

func (p *Parser) noPrefixParseFunctionError(t token.Type) {
	p.errors = append(p.errors, fmt.Errorf("no prefix parse function for %s found", t))
}
