package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/hertzcodes/compiler-in-go/src/interpreter/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hertz Interactive Shell, Welcome %s!\n", user.Username)
	repl.Start(os.Stdin,os.Stdout)
}