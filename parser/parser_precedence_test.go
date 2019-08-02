// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package parser

import (
	"go.ajitem.com/donut/lexer"
	"testing"
)

func TestOperatorPrecedenceParsing(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "-a * b",
			expected: "((-a) * b)",
		},
		{
			input:    "!-a",
			expected: "(!(-a))",
		},
		{
			input:    "a + b + c",
			expected: "((a + b) + c)",
		},
		{
			input:    "a - b - c",
			expected: "((a - b) - c)",
		},
		{
			input:    "a * b * c",
			expected: "((a * b) * c)",
		},
		{
			input:    "a * b / c",
			expected: "((a * b) / c)",
		},
		{
			input:    "a + b / c",
			expected: "(a + (b / c))",
		},
		{
			input:    "a + b * c + d / e - f",
			expected: "(((a + (b * c)) + (d / e)) - f)",
		},
		{
			input:    "3 + 4; -5 * 5",
			expected: "(3 + 4)((-5) * 5)",
		},
		{
			input:    "5 > 4 == 3 < 4",
			expected: "((5 > 4) == (3 < 4))",
		},
		{
			input:    "5 < 4 != 3 > 4",
			expected: "((5 < 4) != (3 > 4))",
		},
		{
			input:    "3 + 4 * 5 == 3 * 1 + 4 * 5",
			expected: "((3 + (4 * 5)) == ((3 * 1) + (4 * 5)))",
		},
		{
			input:    "a += 5 * 2",
			expected: "(a += (5 * 2))",
		},
		{
			input:    "true",
			expected: "true",
		},
		{
			input:    "false",
			expected: "false",
		},
		{
			input:    "3 > 5 == false",
			expected: "((3 > 5) == false)",
		},
		{
			input:    "30 < 5 != true",
			expected: "((30 < 5) != true)",
		},
		{
			input: "1 + (2 + 3) + 4",
			expected: "((1 + (2 + 3)) + 4)",
		},
		{
			input: "(5 + 5) * 2",
			expected: "((5 + 5) * 2)",
		},
		{
			input: "2 / (5 + 5)",
			expected: "(2 / (5 + 5))",
		},
		{
			input: "-(5 + 5)",
			expected: "(-(5 + 5))",
		},
		{
			input: "!(true == true)",
			expected: "(!(true == true))",
		},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)

		p := NewParser(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		actual := program.String()
		if actual != tt.expected {
			t.Errorf("expected=%q, got=%q", tt.expected, actual)
		}
	}
}
