// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"go.ajitem.com/donut/token"
)

type (
	prefixParser func() ast.Expression
	infixParser  func(ast.Expression) ast.Expression
)

// The Parser has three fields:
// l - a pointer to an instance of the lexer to get the next token in the input
// currentToken, peekToken - point to the current and the next token in the input. Analogus to position and readPosition
// in the lexer.
type Parser struct {
	l *lexer.Lexer

	errors []error

	currentToken token.Token
	peekToken    token.Token

	prefixParsers map[token.Type]prefixParser
	infixParsers  map[token.Type]infixParser
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	p.prefixParsers = make(map[token.Type]prefixParser)

	p.registerPrefixParser(token.IDENT, p.parseIdentifier)
	p.registerPrefixParser(token.INT, p.parseIntegerLiteral)
	p.registerPrefixParser(token.FLOAT, p.parseFloatLiteral)

	p.registerPrefixParser(token.BANG, p.parsePrefixExpression)
	p.registerPrefixParser(token.SUB, p.parsePrefixExpression)
	p.registerPrefixParser(token.INCR, p.parsePrefixExpression)
	p.registerPrefixParser(token.DECR, p.parsePrefixExpression)

	// Reading the next two tokens ensures that both currentToken and peekTokens are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = ast.Statements{}

	for !p.currentTokenIs(token.EOF) {
		stmt := p.parseStatement()

		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.nextToken()
	}

	return program
}

func (p *Parser) Errors() []error {
	return p.errors
}
