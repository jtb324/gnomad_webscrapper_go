package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func parse_input(input string) []string {
	//creating a slice that will hold the gene names
	var gene_slice []string

	file_buffer, err := os.Open(input)

	if err != nil {
		log.Fatal(err)
	}

	//defer the file's closings
	defer func() {
		if err = file_buffer.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	snl := bufio.NewScanner(file_buffer)
	// opening the excel file

	i := 0
	for snl.Scan() {
		if i != 0 {
			gene_slice = append(gene_slice, snl.Text())
		}
		i++
	}

	// catching the error if it occurs
	err = snl.Err()
	if err != nil {
		log.Fatal(err)
	}

	//if there are no genes in the file then the program needs to exit successful
	if len(gene_slice) == 0 {
		fmt.Println("The were not gene names found in the provided text file")
		os.Exit(1)
	}
	return gene_slice
}
