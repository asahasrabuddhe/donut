package parser

import (
	"go.ajitem.com/donut/ast"
	"go.ajitem.com/donut/lexer"
	"testing"
)

func TestFunctionLiteralParsing(t *testing.T) {
	input := `func(x, y) { x + y; }`

	l := lexer.NewLexer(input)

	p := NewParser(l)

	program := p.ParseProgram()

	checkParserErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("program.Statements contains more than 1 statements. got=%d\n", len(program.Statements))
	}

	expressionStatement, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("program.Statements[0] is not an ast.ExpressionStatement. got=%T\n", program.Statements[0])
	}

	functionLiteral, ok := expressionStatement.Expression.(*ast.FunctionLiteral)
	if !ok {
		t.Fatalf("expressionStatement.Expression is not ast.FunctionLiteral. got=%T\n", expressionStatement.Expression)
	}

	if len(functionLiteral.Parameters) != 2 {
		t.Fatalf("function did not receive 2 parameters. got=%d\n", len(functionLiteral.Parameters))
	}

	testLiteralExpression(t, functionLiteral.Parameters[0], "x")
	testLiteralExpression(t, functionLiteral.Parameters[1], "y")

	if len(functionLiteral.Body.Statements) != 1 {
		t.Fatalf("functionLiteral.Body.Statements does not have 1 statement. got=%d\n", len(functionLiteral.Body.Statements))
	}

	bodyStatement, ok := functionLiteral.Body.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("function body statemenbt is not ast.ExpressionStatement. got=%T\n", functionLiteral.Body.Statements[0])
	}

	testInfixExpression(t, bodyStatement.Expression, "x", "+", "y")
}

func TestFunctionParameterParsing(t *testing.T) {
	tests := []struct {
		input          string
		expectedParams []string
	}{
		{input: `func() {};`, expectedParams: []string{}},
		{input: `func(x) {};`, expectedParams: []string{"x"}},
		{input: `func(x, y, z) {};`, expectedParams: []string{"x", "y", "z"}},
	}

	for _, tt := range tests {
		l := lexer.NewLexer(tt.input)

		p := NewParser(l)

		program := p.ParseProgram()

		checkParserErrors(t, p)

		expressionStatement := program.Statements[0].(*ast.ExpressionStatement)
		functionLiteral := expressionStatement.Expression.(*ast.FunctionLiteral)

		if len(functionLiteral.Parameters) != len(tt.expectedParams) {
			t.Fatalf(
				"length parameters are incorrect. want %d, got %d\n",
				len(tt.expectedParams), len(functionLiteral.Parameters),
			)
		}

		for i, identifier := range tt.expectedParams {
			testLiteralExpression(t, functionLiteral.Parameters[i], identifier)
		}
	}
}
