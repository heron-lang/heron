package main

import (
	"ares/src/gen"
	"ares/src/lexer"
	"ares/src/parser"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	if len(os.Args) == 0 {
		fmt.Println("missing the file argument")
		return
	}

	file, err := ioutil.ReadFile(os.Args[1])

	if err != nil {
		fmt.Println("there was an error reading that file")
		fmt.Println(err)
		return
	}

	scanner := lexer.New(string(file))

	p := parser.New(scanner)
	tree := p.ParseProgram()

	//fmt.Println(tree)

	generator := &gen.Gen{Program: tree}
	generator.Generate()

	fmt.Println(generator.Output)
}
