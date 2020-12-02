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
	Illegal = "Illegal"
	Eof     = "Eof"

	// Identifiers and Literals
	Identifier = "Identifier"
	Integer    = "Integer"
	Float      = "Float"

	// Operators
	Assign              = "="
	Add                 = "+"
	AddAssign           = "+="
	Increment           = "++"
	Subtract            = "-"
	SubtractAssign      = "-="
	Decrement           = "--"
	Bang                = "!"
	Multiply            = "*"
	MultiplyAssign      = "*="
	Divide              = "/"
	DivideAssign        = "/="
	Remainder           = "%"
	RemainderAssign     = "%="
	LessThan            = "<"
	GreaterThan         = ">"
	Equals              = "=="
	NotEquals           = "!="
	LessThanOrEquals    = "<="
	GreaterThanOrEquals = ">="

	// Delimiters
	Comma     = ","
	Semicolon = ";"

	LeftParenthesis  = "("
	RightParenthesis = ")"
	LeftBrace        = "{"
	RightBrace       = "}"
	LeftBracket      = "["
	RightBracket     = "]"

	// Keywords
	Function = "Function"
	Let      = "Let"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

func LookupIndent(ident string) Type {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return Identifier
}
