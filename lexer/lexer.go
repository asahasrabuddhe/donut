// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package lexer

import (
	"go.ajitem.com/donut/token"
	"strings"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.eatWhiteSpace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.EQ, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.ASSIGN, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			tok = token.Token{Type: token.INCR, Literal: l.getPeekedLiteral()}
		} else if l.peekChar() == '=' {
			tok = token.Token{Type: token.ADD_ASSIGN, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.ADD, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			tok = token.Token{Type: token.DECR, Literal: l.getPeekedLiteral()}
		} else if l.peekChar() == '=' {
			tok = token.Token{Type: token.SUB_ASSIGN, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.SUB, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.MUL_ASSIGN, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.MUL, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.DIV_ASSIGN, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.DIV, l.ch)
		}
	case '%':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.REM_ASSIGN, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.REM, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NOTEQ, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.BANG, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.LTE, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.GTE, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.GT, l.ch)
		}
	case ',':
		tok = token.NewToken(token.COMMA, l.ch)
	case ';':
		tok = token.NewToken(token.SEMICOLON, l.ch)
	case '(':
		tok = token.NewToken(token.LPAREN, l.ch)
	case ')':
		tok = token.NewToken(token.RPAREN, l.ch)
	case '{':
		tok = token.NewToken(token.LBRACE, l.ch)
	case '}':
		tok = token.NewToken(token.RBRACE, l.ch)
	case '[':
		tok = token.NewToken(token.LBRACK, l.ch)
	case ']':
		tok = token.NewToken(token.RBRACK, l.ch)
	case 0:
		tok = token.NewToken(token.EOF, l.ch)
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()

			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.FLOAT
			} else {
				tok.Type = token.INT
			}

			return tok
		} else {
			tok = token.NewToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

func (l *Lexer) getPeekedLiteral() string {
	ch := l.ch
	l.readChar()
	return string(ch) + string(l.ch)
}

func (l *Lexer) readIdentifier() string {
	position := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position

	for isDigit(l.ch) {
		if l.peekChar() == '.' {
			l.readChar()
			l.readChar()

			continue
		}
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) eatWhiteSpace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
