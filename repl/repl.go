// Copyright (c) 2019 Ajitem Sahasrabuddhe <me@ajitem.com>
// All rights reserved. Use of this source code is governed
// by the MIT license that can be found in the LICENSE file.

package repl

import (
	"bufio"
	"fmt"
	"go.ajitem.com/donut/lexer"
	"go.ajitem.com/donut/token"
	"io"
)

const PROMPT = "🍩 "

func Start(in io.Reader, out io.Writer) error {
	scanner := bufio.NewScanner(in)

	for {
		_, err := fmt.Fprint(out, PROMPT)
		if err != nil {
			return err
		}

		scanned := scanner.Scan()

		if !scanned {
			continue
		}

		line := scanner.Text()
		l := lexer.NewLexer(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				return err
			}
		}
	}
}
