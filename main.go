package main

import (
	"fmt"
	"os"
	user2 "os/user"

	"github.com/zeroFruit/zf-lang/repl"
)

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Hello %s! This is the Zf programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
