// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"go.ajitem.com/donut/token"
	"testing"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input             string
		leftIntegerValue  int64
		leftFloatValue    float64
		leftTokenType     token.Type
		operator          string
		rightIntegerValue int64
		rightFloatValue   float64
		rightTokenType    token.Type
	}{
		{input: "5 + 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "+", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "5 - 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "-", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "6.5 * 5.5", leftFloatValue: 6.5, leftTokenType: token.FLOAT, operator: "*", rightFloatValue: 5.5, rightTokenType: token.FLOAT},
		{input: "5 / 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "/", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "5 < 6.5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "<", rightFloatValue: 6.5, rightTokenType: token.FLOAT},
		{input: "5.5 > 5", leftFloatValue: 5.5, leftTokenType: token.FLOAT, operator: ">", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "5 <= 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "<=", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "5 >= 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: ">=", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "5 == 5", leftIntegerValue: 5, leftTokenType: token.INT, operator: "==", rightIntegerValue: 5, rightTokenType: token.INT},
		{input: "7.15 != 5.32", leftFloatValue: 7.15, leftTokenType: token.FLOAT, operator: "!=", rightFloatValue: 5.32, rightTokenType: token.FLOAT},
	}

	for _, tt := range infixTests {
		l := lexer.NewLexer(tt.input)

		p := NewParser(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}
		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T",
				program.Statements[0])
		}
		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("exp is not ast.InfixExpression. got=%T", stmt.Expression)
		}
		switch tt.leftTokenType {
		case token.INT:
			if !testIntegerLiteral(t, exp.Left, tt.leftIntegerValue) {
				return
			}
		case token.FLOAT:
			if !testFloatLiteral(t, exp.Left, tt.leftFloatValue) {
				return
			}
		}
		if exp.Operator != tt.operator {
			t.Fatalf("exp.Operator is not '%s'. got=%s",
				tt.operator, exp.Operator)
		}
		switch tt.rightTokenType {
		case token.INT:
			if !testIntegerLiteral(t, exp.Right, tt.rightIntegerValue) {
				return
			}
		case token.FLOAT:
			if !testFloatLiteral(t, exp.Right, tt.rightFloatValue) {
				return
			}
		}
	}
}
