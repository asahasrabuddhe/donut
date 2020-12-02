// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"testing"
)

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  interface{}
		operator   string
		rightValue interface{}
	}{
		{input: "5 + 5", leftValue: 5, operator: "+", rightValue: 5},
		{input: "5 - 5", leftValue: 5, operator: "-", rightValue: 5},
		{input: "6.5 * 5.5", leftValue: 6.5, operator: "*", rightValue: 5.5},
		{input: "5 / 5", leftValue: 5, operator: "/", rightValue: 5},
		{input: "5 < 6.5", leftValue: 5, operator: "<", rightValue: 6.5},
		{input: "5.5 > 5", leftValue: 5.5, operator: ">", rightValue: 5},
		{input: "5 <= 5", leftValue: 5, operator: "<=", rightValue: 5},
		{input: "5 >= 5", leftValue: 5, operator: ">=", rightValue: 5},
		{input: "5 == 5", leftValue: 5, operator: "==", rightValue: 5},
		{input: "7.15 != 5.32", leftValue: 7.15, operator: "!=", rightValue: 5.32},
		{input: "true == true", leftValue: true, operator: "==", rightValue: true},
		{input: "true != false", leftValue: true, operator: "!=", rightValue: false},
		{input: "false == false", leftValue: false, operator: "==", rightValue: false},
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

		testInfixExpression(t, stmt.Expression, tt.leftValue, tt.operator, tt.rightValue)
	}
}

func testInfixExpression(t *testing.T, expression ast.Expression, left interface{}, operator string, right interface{}) bool {
	exp, ok := expression.(*ast.InfixExpression)
	if !ok {
		t.Fatalf("exp is not ast.InfixExpression. got=%T", expression)
	}

	if !testLiteralExpression(t, exp.Left, left) {
		return false
	}

	if exp.Operator != operator {
		t.Fatalf("exp.Operator is not '%s'. got=%s", operator, exp.Operator)
		return false
	}

	if !testLiteralExpression(t, exp.Right, right) {
		return false
	}

	return true
}
