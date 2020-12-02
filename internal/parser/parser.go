// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/lexer"
	"go.ajitem.com/donut/internal/token"
)

type (
	prefixParser  func() ast.Expression
	infixParser   func(ast.Expression) ast.Expression
	postfixParser func(ast.Expression) ast.Expression
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

	prefixParsers  map[token.Type]prefixParser
	infixParsers   map[token.Type]infixParser
	postfixParsers map[token.Type]postfixParser
}

func NewParser(l *lexer.Lexer) *Parser {
	p := &Parser{
		l:      l,
		errors: []error{},
	}

	p.prefixParsers = make(map[token.Type]prefixParser)

	p.registerPrefixParser(token.Identifier, p.parseIdentifier)
	p.registerPrefixParser(token.Integer, p.parseIntegerLiteral)
	p.registerPrefixParser(token.Float, p.parseFloatLiteral)

	p.registerPrefixParser(token.Bang, p.parsePrefixExpression)
	p.registerPrefixParser(token.Subtract, p.parsePrefixExpression)
	p.registerPrefixParser(token.Increment, p.parsePrefixExpression)
	p.registerPrefixParser(token.Decrement, p.parsePrefixExpression)

	p.registerPrefixParser(token.True, p.parseBoolean)
	p.registerPrefixParser(token.False, p.parseBoolean)

	p.registerPrefixParser(token.LeftParenthesis, p.parseGroupedExpression)
	p.registerPrefixParser(token.If, p.parseIfExpression)
	p.registerPrefixParser(token.Function, p.parseFunctionLiteral)

	p.infixParsers = make(map[token.Type]infixParser)

	p.registerInfixParser(token.Add, p.parseInfixExpression)
	p.registerInfixParser(token.AddAssign, p.parseInfixExpression)
	p.registerInfixParser(token.Subtract, p.parseInfixExpression)
	p.registerInfixParser(token.Multiply, p.parseInfixExpression)
	p.registerInfixParser(token.Divide, p.parseInfixExpression)
	p.registerInfixParser(token.Equals, p.parseInfixExpression)
	p.registerInfixParser(token.NotEquals, p.parseInfixExpression)
	p.registerInfixParser(token.LessThan, p.parseInfixExpression)
	p.registerInfixParser(token.LessThanOrEquals, p.parseInfixExpression)
	p.registerInfixParser(token.GreaterThan, p.parseInfixExpression)
	p.registerInfixParser(token.GreaterThanOrEquals, p.parseInfixExpression)
	p.registerInfixParser(token.LeftParenthesis, p.parseCallExpression)

	p.postfixParsers = make(map[token.Type]postfixParser)

	p.registerPostfixParser(token.Increment, p.parsePostfixExpression)
	p.registerPostfixParser(token.Decrement, p.parsePostfixExpression)

	// Reading the next two tokens ensures that both currentToken and peekTokens are set
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = ast.Statements{}

	for !p.currentTokenIs(token.Eof) {
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
