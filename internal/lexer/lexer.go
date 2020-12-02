// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package lexer

import (
	"go.ajitem.com/donut/internal/token"
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
			tok = token.Token{Type: token.Equals, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Assign, l.ch)
		}
	case '+':
		if l.peekChar() == '+' {
			tok = token.Token{Type: token.Increment, Literal: l.getPeekedLiteral()}
		} else if l.peekChar() == '=' {
			tok = token.Token{Type: token.AddAssign, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Add, l.ch)
		}
	case '-':
		if l.peekChar() == '-' {
			tok = token.Token{Type: token.Decrement, Literal: l.getPeekedLiteral()}
		} else if l.peekChar() == '=' {
			tok = token.Token{Type: token.SubtractAssign, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Subtract, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.MultiplyAssign, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Multiply, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.DivideAssign, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Divide, l.ch)
		}
	case '%':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.RemainderAssign, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Remainder, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.NotEquals, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.Bang, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.LessThanOrEquals, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.LessThan, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tok = token.Token{Type: token.GreaterThanOrEquals, Literal: l.getPeekedLiteral()}
		} else {
			tok = token.NewToken(token.GreaterThan, l.ch)
		}
	case ',':
		tok = token.NewToken(token.Comma, l.ch)
	case ';':
		tok = token.NewToken(token.Semicolon, l.ch)
	case '(':
		tok = token.NewToken(token.LeftParenthesis, l.ch)
	case ')':
		tok = token.NewToken(token.RightParenthesis, l.ch)
	case '{':
		tok = token.NewToken(token.LeftBrace, l.ch)
	case '}':
		tok = token.NewToken(token.RightBrace, l.ch)
	case '[':
		tok = token.NewToken(token.LeftBracket, l.ch)
	case ']':
		tok = token.NewToken(token.RightBracket, l.ch)
	case 0:
		tok = token.NewToken(token.Eof, l.ch)
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readIdentifier()
			tok.Type = token.LookupIndent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readNumber()

			if strings.Contains(tok.Literal, ".") {
				tok.Type = token.Float
			} else {
				tok.Type = token.Integer
			}

			return tok
		} else {
			tok = token.NewToken(token.Illegal, l.ch)
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
