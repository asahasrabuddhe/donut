// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go.ajitem.com/donut/ast"
	"strconv"
	"strings"
)

func (p *Parser) parseIntegerLiteral() ast.Expression {
	value, err := strconv.ParseInt(p.currentToken.Literal, 0, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("could not parse %q as integer", p.currentToken.Literal))
		return nil
	}

	return &ast.IntegerLiteral{Token: p.currentToken, Value: value}
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	if strings.HasSuffix(p.currentToken.Literal, ".") {
		p.errors = append(p.errors, fmt.Errorf("could not parse %q as float", p.currentToken.Literal))
		return nil
	}

	value, err := strconv.ParseFloat(p.currentToken.Literal, 64)
	if err != nil {
		p.errors = append(p.errors, fmt.Errorf("could not parse %q as float", p.currentToken.Literal))
		return nil
	}

	return &ast.FloatLiteral{Token: p.currentToken, Value: value}
}
