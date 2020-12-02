// Copyright © 2020 Ajitem Sahasrabuddhe. All rights reserved.
// Use of this source code is governed by a MIT license
// details of which can be found in the LICENSE file.

package main

import (
	"fmt"
	"go.ajitem.com/donut/repl"
	"os"
	"os/user"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! Welcome to The Donut Programming Language!\n", usr.Username)
	fmt.Printf("Awaiting input...\n")

	repl.Start(os.Stdin, os.Stdout)
}
