// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package ast

type Node interface {
	TokenLiteral() string
	String() string
}

type Statement interface {
	Node
	statementNode()
}

type Statements []Statement

type Expression interface {
	Node
	expressionNode()
}
