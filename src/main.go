package main

import (
	"fmt"
)

func main() {
	fmt.Println("Initializing program to gather gene information from gnomAD...")
	//Parsing the inputs
	input_path, _, api_endpoint := parser()

	//going through the input file and getting a slice of all the genes
	gene_slice := parse_input(input_path)

	_ = fetch_response(api_endpoint, gene_slice)

	db, dbName, _ := initialize_db()

	make_table(db, dbName)
}
