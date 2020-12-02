package main

import (
	"fmt"
	"github.com/asahasrabuddhe/donut/repl"
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
