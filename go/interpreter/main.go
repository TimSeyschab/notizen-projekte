package main

import (
	"fmt"
	"interpreter/repl"
	"os"
)

/*
ASCII Art: Pratt Parser

          +-------------------+
          |     Lexer         |
          +-------------------+
                   |
                   v
          +-------------------+
          |   Token Stream    |
          +-------------------+
                   |
                   v
          +-------------------+
          |     Parser        |
          +-------------------+
                   |
          +--------+---------+
          |                  |
          v                  v
  +-------------+     +---------------+
  | Nud (Null   |     | Led (Left     |
  | Denotation) |     | Denotation)   |
  +-------------+     +---------------+
          |                  |
          v                  v
  +-------------------+   +-------------------+
  | Parse Expressions |   | Parse Expressions |
  +-------------------+   +-------------------+
                   |
                   v
          +-------------------+
          |   Abstract Syntax |
          |      Tree (AST)   |
          +-------------------+
                   |
                   v
          +-------------------+
          |    Evaluation     |
          +-------------------+
*/

// It prompts the user to enter a line of code and then starts
// the read-eval-print loop (REPL) using the standard input and output.
func main() {
	fmt.Printf("Bitte Programmzeile eingeben: \n")
	repl.Start(os.Stdin, os.Stdout)
}
