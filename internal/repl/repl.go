// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package repl

import (
	"bufio"
	"fmt"
	"go.ajitem.com/donut/internal/lexer"
	"go.ajitem.com/donut/internal/parser"
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
			return nil
		}
		line := scanner.Text()

		l := lexer.NewLexer(line)
		p := parser.NewParser(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				_, _ = fmt.Fprintf(out, "\t%s\n", err.Error())
			}
			continue
		}

		_, _ = fmt.Fprintln(out, program)
	}
}
