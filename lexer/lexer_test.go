// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package lexer

import (
	"go.ajitem.com/donut/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	
	let ten = 10;

	let add = fn(x, y) {
		x + y;
	};

	let result = add(five, ten);
	!-/*5;
	5 < 10 > 5;

	if (5 < 10) {
		return true;
	} else {
		return false;
	}

	10 == 10;
	10 != 9;

	5 <= 10;
	5 >= 1;

	five++;
	++ten;

	ten--;
	--five;

	five += 3;
	ten -= 2;
	five *= 2;
	ten /= 4;
	five %= 5;
	let q = 9 % 5;
	let d = 3.5;
	`
	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Identifier, "five"},
		{token.Assign, "="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "ten"},
		{token.Assign, "="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.LeftParenthesis, "("},
		{token.Identifier, "x"},
		{token.Comma, ","},
		{token.Identifier, "y"},
		{token.RightParenthesis, ")"},
		{token.LeftBrace, "{"},
		{token.Identifier, "x"},
		{token.Add, "+"},
		{token.Identifier, "y"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "result"},
		{token.Assign, "="},
		{token.Identifier, "add"},
		{token.LeftParenthesis, "("},
		{token.Identifier, "five"},
		{token.Comma, ","},
		{token.Identifier, "ten"},
		{token.RightParenthesis, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Subtract, "-"},
		{token.Divide, "/"},
		{token.Multiply, "*"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.GreaterThan, ">"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.LeftParenthesis, "("},
		{token.Integer, "5"},
		{token.LessThan, "<"},
		{token.Integer, "10"},
		{token.RightParenthesis, ")"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Else, "else"},
		{token.LeftBrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.RightBrace, "}"},
		{token.Integer, "10"},
		{token.Equals, "=="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Integer, "10"},
		{token.NotEquals, "!="},
		{token.Integer, "9"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.LessThanOrEquals, "<="},
		{token.Integer, "10"},
		{token.Semicolon, ";"},
		{token.Integer, "5"},
		{token.GreaterThanOrEquals, ">="},
		{token.Integer, "1"},
		{token.Semicolon, ";"},
		{token.Identifier, "five"},
		{token.Increment, "++"},
		{token.Semicolon, ";"},
		{token.Increment, "++"},
		{token.Identifier, "ten"},
		{token.Semicolon, ";"},
		{token.Identifier, "ten"},
		{token.Decrement, "--"},
		{token.Semicolon, ";"},
		{token.Decrement, "--"},
		{token.Identifier, "five"},
		{token.Semicolon, ";"},
		{token.Identifier, "five"},
		{token.AddAssign, "+="},
		{token.Integer, "3"},
		{token.Semicolon, ";"},
		{token.Identifier, "ten"},
		{token.SubtractAssign, "-="},
		{token.Integer, "2"},
		{token.Semicolon, ";"},
		{token.Identifier, "five"},
		{token.MultiplyAssign, "*="},
		{token.Integer, "2"},
		{token.Semicolon, ";"},
		{token.Identifier, "ten"},
		{token.DivideAssign, "/="},
		{token.Integer, "4"},
		{token.Semicolon, ";"},
		{token.Identifier, "five"},
		{token.RemainderAssign, "%="},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "q"},
		{token.Assign, "="},
		{token.Integer, "9"},
		{token.Remainder, "%"},
		{token.Integer, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Identifier, "d"},
		{token.Assign, "="},
		{token.Float, "3.5"},
		{token.Semicolon, ";"},
		{token.Eof, ""},
	}

	l := NewLexer(input)

	for i, tt := range tests {
		tok := l.NextToken()
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
