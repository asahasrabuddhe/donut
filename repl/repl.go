// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

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

		for tok := l.NextToken(); tok.Type != token.Eof; tok = l.NextToken() {
			_, err := fmt.Fprintf(out, "%+v\n", tok)
			if err != nil {
				return err
			}
		}
	}
}
