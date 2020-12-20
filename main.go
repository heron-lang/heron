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
	if len(os.Args) != 3 {
		fmt.Println("Missing the input and output arguments. The command should look like this: heron <input> <output>")
		return
	}

	var inputFile, err = ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("There was an error reading that file")
		fmt.Println(err)
		return
	}

	compiled := compileAres(string(inputFile))
	fmt.Println("Successfully compiled!")

	createOutFile()
	writeOut(compiled)
}

func createOutFile() {
	_, err := os.Stat(os.Args[2])
	if os.IsNotExist(err) {
		var file, err = os.Create(os.Args[2])

		fmt.Println("That output file was not found, creating it...")

		if err != nil {
			fmt.Println("There was an error in creating that file")
			fmt.Println(err)
			return
		}

		defer file.Close()
	}
}

func writeOut(output string) {
	file, err := os.OpenFile(os.Args[2], os.O_RDWR, 0644)
	fmt.Println("Writing output...")

	if err != nil {
		fmt.Println("There was an error opening the output file")
		fmt.Println(err)
		return
	}

	if _, err = file.Write([]byte(output)); err != nil {
		fmt.Println("There was an error writing to that file")
		fmt.Println(err)
		return
	}

	defer file.Close()
}

func compileAres(input string) string {
	p := parser.New(lexer.New(input))
	tree := p.ParseProgram()

	generator := &gen.Gen{Program: tree}
	generator.Generate()

	return generator.Output
}
