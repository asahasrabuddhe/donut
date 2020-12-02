// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import "go.ajitem.com/donut/token"

var precedences = map[token.Type]int{
	token.Equals:              Equals,
	token.NotEquals:           Equals,
	token.LessThan:            LessOrGreater,
	token.LessThanOrEquals:    LessOrGreater,
	token.GreaterThan:         LessOrGreater,
	token.GreaterThanOrEquals: LessOrGreater,
	token.Add:                 Sum,
	token.AddAssign:           Sum,
	token.Subtract:            Sum,
	token.SubtractAssign:      Sum,
	token.Divide:              Product,
	token.DivideAssign:        Product,
	token.Multiply:            Product,
	token.MultiplyAssign:      Product,
	token.LeftParenthesis:     Call,
}

func (p *Parser) peekTokenPrecedence() int {
	if p, ok := precedences[p.peekToken.Type]; ok {
		return p
	}

	return Lowest
}

func (p *Parser) currentTokenPrecedence() int {
	if p, ok := precedences[p.currentToken.Type]; ok {
		return p
	}

	return Lowest
}
