package main

import (
	"fmt"
	"go-interpreter/repl"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
  fmt.Printf("Hi, %s! This is the go interpreter!\n", user.Username)
	fmt.Printf("Feel free to type something\n")
	repl.Start(os.Stdin, os.Stdout)
}
