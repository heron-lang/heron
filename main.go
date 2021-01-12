package main

import (
	"fmt"
	"github.com/poseidoncoder/heron/src/compiler"
	"io/ioutil"
	"os"
	"strings"
)

var outputName string

func main() {
  if len(os.Args) == 1 {
		helpScreen()
		return
	}

	switch os.Args[1] {
	case "build":
		if len(os.Args) < 3 {
			fmt.Println("the build command requires an input file as an argument")
			return
		}

		outputName = strings.TrimSuffix(os.Args[2], ".he") + ".css"

		var inputFile, err = ioutil.ReadFile(os.Args[2])
		if err != nil {
			fmt.Println("There was an error reading that file")
			fmt.Println(err)
			return
		}

		compiled := compiler.Compile(inputFile, os.Args[2])
		fmt.Println("Successfully compiled!")

		createOutFile()
		writeOut(compiled)
	case "version":
		fmt.Println("heron1.0.0")
	case "help":
		helpScreen()
	default:
		fmt.Println(fmt.Sprintf("%v: unknown command", os.Args[1]))
		fmt.Println("see 'heron help'")
		return
	}
}

func helpScreen() {
	if len(os.Args) < 3 {
		fmt.Println("Heron is a tool for managing Heron source code")

		fmt.Println("\nUsage:")
		fmt.Println("\theron <command> [arguments]")

		fmt.Println("\nThe commands are:")
		fmt.Println("\tbuild      compile Heron code to CSS")
		fmt.Println("\tversion    print Heron version")

		fmt.Println("\nUse 'heron help <topic>' for more information about that topic")
		return
	}

	switch os.Args[2] {
	case "build":
		fmt.Println("Usage:")
		fmt.Println("\theron build INPUT_FILE")

		fmt.Println("\nDescription:")
		fmt.Println("\t'Build' will 'transform' the input file into CSS.")
		fmt.Println("\tIt will then 'pour' the corresponding CSS into a file with the exact same path and name but it will use the '.css' file extension.")
		fmt.Println("\tFor example, if you run 'heron build ./heron.he', Heron will output a file named 'heron.css' with the corresponding CSS.")
	case "version":
		fmt.Println("Usage:")
		fmt.Println("\theron version")

		fmt.Println("\nDescription")
		fmt.Println("\t'Version' prints the current version of Heron that you have installed.")
	case "help":
		fmt.Println("Usage:")
		fmt.Println("\theron help")

		fmt.Println("\nDescription")
		fmt.Println("\t'Help' returns documentation of the Heron CLI.")
	default:
		fmt.Println(fmt.Sprintf("%v: unknown command", os.Args[2]))
		fmt.Println("see 'heron help'")
	}
}

func createOutFile() {
	_, err := os.Stat(outputName)
	if os.IsNotExist(err) {
		var file, err = os.Create(outputName)

		fmt.Println("That output file was not found, creating it...")

		if err != nil {
			fmt.Println("there was an error in creating that file")
			fmt.Println(err)
			return
		}

		defer file.Close()
	}
}

func writeOut(output strings.Builder) {
	file, err := os.OpenFile(outputName, os.O_RDWR, 0644)
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
