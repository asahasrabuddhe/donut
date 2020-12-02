// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/token"
)

func (p *Parser) parseBoolean() ast.Expression {
	return &ast.BooleanLiteral{Token: p.currentToken, Value: p.currentTokenIs(token.True)}
}
