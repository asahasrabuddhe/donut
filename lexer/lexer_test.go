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
		{token.LET, "let"},
		{token.IDENT, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.ADD, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "result"},
		{token.ASSIGN, "="},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.SUB, "-"},
		{token.DIV, "/"},
		{token.MUL, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "else"},
		{token.LBRACE, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "=="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NOTEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LTE, "<="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.GTE, ">="},
		{token.INT, "1"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "five"},
		{token.INCR, "++"},
		{token.SEMICOLON, ";"},
		{token.INCR, "++"},
		{token.IDENT, "ten"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.DECR, "--"},
		{token.SEMICOLON, ";"},
		{token.DECR, "--"},
		{token.IDENT, "five"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "five"},
		{token.ADD_ASSIGN, "+="},
		{token.INT, "3"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.SUB_ASSIGN, "-="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "five"},
		{token.MUL_ASSIGN, "*="},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "ten"},
		{token.DIV_ASSIGN, "/="},
		{token.INT, "4"},
		{token.SEMICOLON, ";"},
		{token.IDENT, "five"},
		{token.REM_ASSIGN, "%="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "q"},
		{token.ASSIGN, "="},
		{token.INT, "9"},
		{token.REM, "%"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENT, "d"},
		{token.ASSIGN, "="},
		{token.FLOAT, "3.5"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
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
