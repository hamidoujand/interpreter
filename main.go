package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hamidoujand/interpreter/repl"
)

func main() {
	usr, err := user.Current()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("Hello %s, This is the Monkey programming language!\n", usr.Username)
	fmt.Println("Feel free to type commands")
	repl.Start(os.Stdin, os.Stdout)
}
