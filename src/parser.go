package main

import (
	"flag"
	"fmt"
	"os"
)

func parser() (string, string) {
	//creating a pointer for the input argument that leads to a file
	var input string
	var output string

	flag.StringVar(&input, "input", "0", "flag for the input file to get a list of genes from")
	//creating a pointer that has a string for the output
	flag.StringVar(&output, "output", "0", "flag for the output directory to write the output to")

	flag.Parse()

	if input == "0" || output == "0" {
		fmt.Println("invalid number of arguments passed to the program")
		fmt.Println("Expectd format: main -input 'input filepath' -output 'output directory'")
		os.Exit(1)
	}
	return input, output
}
