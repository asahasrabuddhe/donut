// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

func NewToken(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers and Literals
	IDENT = "IDENT"
	INT   = "INT"
	FLOAT = "FLOAT"

	// Operators
	ASSIGN     = "="
	ADD        = "+"
	ADD_ASSIGN = "+="
	INCR       = "++"
	SUB        = "-"
	SUB_ASSIGN = "-="
	DECR       = "--"
	BANG       = "!"
	MUL        = "*"
	MUL_ASSIGN = "*="
	DIV        = "/"
	DIV_ASSIGN = "/="
	REM        = "%"
	REM_ASSIGN = "%="
	LT         = "<"
	GT         = ">"
	EQ         = "=="
	NOTEQ      = "!="
	LTE        = "<="
	GTE        = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"
	LBRACK = "["
	RBRACK = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"func":   FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIndent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
