package main

import (
	"flag"
	"fmt"
	"os"
)

func parser() (string, string, string) {
	//creating a pointer for the input argument that leads to a file
	var input string
	var output string
	var api_string string

	flag.StringVar(&input, "input", "0", "flag for the input file to get a list of genes from")
	//creating a pointer that has a string for the output
	flag.StringVar(&output, "output", "0", "flag for the output directory to write the output to")
	//creating a pointer to a string that has the api endpoint
	flag.StringVar(&api_string, "api", "0", "flag for the hyperlink to the api endpoint")
	flag.Parse()

	// Checking to make sure that a value was passed to the three flags
	if input == "0" || output == "0" || api_string == "0" {
		fmt.Println("invalid number of arguments passed to the program")
		fmt.Println("Expectd format: main -input 'input filepath' -output 'output directory' -api 'api endpoint'")
		os.Exit(1)
	}
	return input, output, api_string
}
