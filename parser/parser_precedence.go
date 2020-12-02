// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import "go.ajitem.com/donut/token"

var precedences = map[token.Type]int{
	token.EQ:         EQUALS,
	token.NOTEQ:      EQUALS,
	token.LT:         LESSGREATER,
	token.LTE:        LESSGREATER,
	token.GT:         LESSGREATER,
	token.GTE:        LESSGREATER,
	token.ADD:        SUM,
	token.ADD_ASSIGN: SUM,
	token.SUB:        SUM,
	token.SUB_ASSIGN: SUM,
	token.DIV:        PRODUCT,
	token.DIV_ASSIGN: PRODUCT,
	token.MUL:        PRODUCT,
	token.MUL_ASSIGN: PRODUCT,
}

func (p *Parser) peekTokenPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return LOWEST
}

func (p *Parser) currentTokenPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return LOWEST
}
