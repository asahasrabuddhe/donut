// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/internal/ast"
	"go.ajitem.com/donut/internal/lexer"
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
