package main

import (
	"fmt"
	"os"
	user2 "os/user"

	"github.com/zeroFruit/zf-lang/repl"
)

const LOGO = `
         f)FFF l)L                         
        f)      l)                         
z)ZZZZZ f)FFF   l)  a)AAAA  n)NNNN   g)GGG 
    z)  f)      l)   a)AAA  n)   NN g)   GG
  z)    f)      l)  a)   A  n)   NN g)   GG
z)ZZZZZ f)     l)LL  a)AAAA n)   NN  g)GGGG
                                         GG
                                    g)GGGG 		

- author:     zerofruit 
- copyright:  WRITING AN INTERPRETER IN GO, Thorsten Ball


`

func main() {
	user, err := user2.Current()
	if err != nil {
		panic(err)
	}

	fmt.Println(LOGO)
	fmt.Printf("Hello %s! This is the Zf programming language!\n", user.Username)
	fmt.Printf("Feel free to type in commands\n")
	repl.Start(os.Stdin, os.Stdout)
}
