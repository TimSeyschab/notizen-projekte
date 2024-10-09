package main

import (
	"fmt"
	"interpreter/repl"
	"os"
)

func main() {
	fmt.Printf("Bitte Programmzeile eingeben: \n")
	repl.Start(os.Stdin, os.Stdout)
}
