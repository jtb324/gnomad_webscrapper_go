package main

import (
	"fmt"
	"log"
	"os"

	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

func parse_input(input string) []string {
	//creating a slice that will hold the gene names
	var gene_slice []string

	// opening the excel file
	file, err := excelize.OpenFile(input)

	// catching the error if it occurs
	if err != nil {
		log.Fatal(err)
	}

	// Get all the rows in the vegan section.
	rows, err := file.GetRows("Sheet1")

	if err != nil {
		log.Fatal(err)
	}

	// iterating through each row. After the first row (index == 0)
	// then it will append the gene name to the row
	for index, row := range rows {

		if index != 0 {

			gene_slice = append(gene_slice, row[0])
		}
	}

	//if there are no genes in the file then the program needs to exit successful
	if len(gene_slice) == 0 {
		fmt.Println("The were not gene names found in the provided text file")
		os.Exit(1)
	}
	return gene_slice
}
