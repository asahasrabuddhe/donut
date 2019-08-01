// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package ast

import "go.ajitem.com/donut/token"

// <expression>;

type FloatLiteral struct {
	Token token.Token
	Value float64
}

func (f *FloatLiteral) expressionNode() {}

func (f *FloatLiteral) TokenLiteral() string {
	return f.Token.Literal
}

func (f *FloatLiteral) String() string {
	return f.Token.Literal
}
