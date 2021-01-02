package main

import (
	"fmt"
	"heron/src/compiler"
	"io/ioutil"
	"os"
	"strings"
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

	compiled := compiler.Compile(inputFile)
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

func writeOut(output strings.Builder) {
	file, err := os.OpenFile(os.Args[2], os.O_RDWR, 0644)
	fmt.Println("Writing output...")

	if err != nil {
		fmt.Println("There was an error opening the output file")
		fmt.Println(err)
		return
	}

	if _, err = file.Write([]byte(output.String())); err != nil {
		fmt.Println("There was an error writing to that file")
		fmt.Println(err)
		return
	}

	defer file.Close()
}
