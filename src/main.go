package main

import (
	"fmt"
)

func main() {
	fmt.Println("Initializing the webscrapper")
	//Parsing the inputs
	input_path, _ := parser()

	//going through the input file and getting a slice of all the genes
	gene_slice := parse_input(input_path)

	fmt.Println(gene_slice)
}
