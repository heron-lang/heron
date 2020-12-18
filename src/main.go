package main

import (
	"ares/src/gen"
	"ares/src/parser"
	"ares/src/scanner"
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

	tokens := scanner.New(string(file))

	p := parser.New(tokens)
	tree := p.ParseProgram()

	generator := &gen.Gen{}
	css := generator.Generate(tree)
	fmt.Println(css)
}
