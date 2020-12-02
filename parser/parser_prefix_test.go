// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package parser

import (
	"fmt"
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"testing"
)

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input    string
		operator string
		value    interface{}
	}{
		{input: "!5;", operator: "!", value: 5},
		{input: "-15.51;", operator: "-", value: 15.51},
		{input: "--x;", operator: "--", value: "x"},
		{input: "++y;", operator: "++", value: "y"},
		{input: "!true;", operator: "!", value: true},
		{input: "!false;", operator: "!", value: false},
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

		testLiteralExpression(t, exp.Right, tt.value)
	}
}

func testLiteralExpression(t *testing.T, expression ast.Expression, value interface{}) bool {
	switch val := value.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(val))
	case int8:
		return testIntegerLiteral(t, expression, int64(val))
	case int16:
		return testIntegerLiteral(t, expression, int64(val))
	case int32:
		return testIntegerLiteral(t, expression, int64(val))
	case int64:
		return testIntegerLiteral(t, expression, val)
	case float32:
		return testFloatLiteral(t, expression, float64(val))
	case float64:
		return testFloatLiteral(t, expression, val)
	case string:
		return testIdentifier(t, expression, val)
	case bool:
		return testBooleanLiteral(t, expression, val)
	default:
		t.Errorf("type of expression not handled. got=%T", expression)
		return false
	}
}

func testIntegerLiteral(t *testing.T, expression ast.Expression, value int64) bool {
	integer, ok := expression.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
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

func testFloatLiteral(t *testing.T, expression ast.Expression, value float64) bool {
	float, ok := expression.(*ast.FloatLiteral)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
		return false
	}
	if float.Value != value {
		t.Errorf("float.Value not %f. got=%f", value, float.Value)
		return false
	}
	if float.TokenLiteral() != fmt.Sprintf("%v", value) {
		t.Errorf("float.TokenLiteral not %f. got=%s", value,
			float.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	identifier, ok := expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got=%s", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.TokenLiteral not %s. got=%s", value,
			identifier.TokenLiteral())
		return false
	}

	return true
}

func testBooleanLiteral(t *testing.T, expression ast.Expression, value bool) bool {
	identifier, ok := expression.(*ast.BooleanLiteral)
	if !ok {
		t.Errorf("expression not *ast.IntegerLiteral. got=%T", expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %t. got=%t", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != fmt.Sprintf("%t", value) {
		t.Errorf("identifier.TokenLiteral not %t. got=%s", value,
			identifier.TokenLiteral())
		return false
	}

	return true
}
