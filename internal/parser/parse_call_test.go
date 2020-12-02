// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/lexer"
	"testing"
)

func TestCallExpressionParsing(t *testing.T) {
	input := `add(1, 2 * 3, 4 + 5)`

	l := lexer.NewLexer(input)

	p := NewParser(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements does not contain 1 statement. got=%d\n", len(program.Statements))
	}

	expressionStatement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("expressionStatement is not ast.ExpressionStatement. got=%T\n", program.Statements[0])
	}

	callExpression, ok := expressionStatement.Expression.(*ast.CallExpression)
	if !ok {
		t.Fatalf("expressionStatement.Expression is not ast.CallExpression. got=%T\n", expressionStatement.Expression)
	}

	if !testIdentifier(t, callExpression.Function, "add") {
		return
	}

	if len(callExpression.Arguments) != 3 {
		t.Fatalf("call expression did not receive 3 arguments. got=%d\n", len(callExpression.Arguments))
	}

	testLiteralExpression(t, callExpression.Arguments[0], 1)
	testInfixExpression(t, callExpression.Arguments[1], 2, "*", 3)
	testInfixExpression(t, callExpression.Arguments[2], 4, "+", 5)
}
