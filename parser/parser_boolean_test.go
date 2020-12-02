// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"testing"
)

func TestBooleanLiteralExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{
			input:    "true;",
			expected: true,
		},
		{
			input:    "false;",
			expected: false,
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)

		p := NewParser(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program does not have enough statements, got=%d", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("program.Statements[0] is not ast.ExpressionStatement. got=%T", program.Statements[0])
		}

		testLiteralExpression(t, stmt.Expression, tt.expected)
	}
}
