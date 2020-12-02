// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/token"
)

func (p *Parser) parseBlockStatement() *ast.BlockStatement {
	block := &ast.BlockStatement{Token: p.currentToken}
	block.Statements = make(ast.Statements, 0)

	p.nextToken()

	if !p.currentTokenIs(token.RightBrace) && !p.currentTokenIs(token.Eof) {
		stmt := p.parseStatement()
		if stmt != nil {
			block.Statements = append(block.Statements, stmt)
		}

		p.nextToken()
	}

	return block
}
