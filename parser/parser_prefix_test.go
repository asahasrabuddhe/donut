// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"go.ajitem.com/donut/token"
	"testing"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input           string
		operator        string
		integerValue    int64
		floatValue      float64
		identifierValue string
		tokenType       string
	}{
		{input: "!5;", operator: "!", integerValue: 5, tokenType: token.INT},
		{input: "-15.51;", operator: "-", floatValue: 15.51, tokenType: token.FLOAT},
		{input: "--x;", operator: "--", identifierValue: "x", tokenType: token.IDENT},
		{input: "++y;", operator: "++", identifierValue: "y", tokenType: token.IDENT},
	}

	for _, tt := range prefixTests {
		l := lexer.NewLexer(tt.input)

		p := NewParser(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not have enough statements. got=%d",
				len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("stmt is not ast.PrefixExpression. got=%T", stmt.Expression)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}

		switch tt.tokenType {
		case token.INT:
			if !testIntegerLiteral(t, exp.Right, tt.integerValue) {
				return
			}
		case token.FLOAT:
			if !testFloatLiteral(t, exp.Right, tt.floatValue) {
				return
			}
		case token.IDENT:
			if !testIdentifier(t, exp.Right, tt.identifierValue) {
				return
			}
		}
	}
}

func testIntegerLiteral(t *testing.T, il ast.Expression, value int64) bool {
	integer, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", value, integer.Value)
		return false
	}
	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integer.TokenLiteral not %d. got=%s", value,
			integer.TokenLiteral())
		return false
	}

	return true
}

func testFloatLiteral(t *testing.T, il ast.Expression, value float64) bool {
	float, ok := il.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	if float.TokenLiteral() != fmt.Sprintf("%.2f", value) {
		t.Errorf("float.TokenLiteral not %f. got=%s", value,
			float.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, il ast.Expression, value string) bool {
	identifier, ok := il.(*ast.Identifier)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", il)
		return false
	}
	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got=%s", value, identifier.Value)
		return false
	}
	if identifier.TokenLiteral() != fmt.Sprintf("%s", value) {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", value,
			identifier.TokenLiteral())
		return false
	}

	return true
}
