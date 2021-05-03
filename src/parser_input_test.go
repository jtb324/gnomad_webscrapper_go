package main

import (
	"testing"
)

func Test_parser_input(t *testing.T) {

	gene_slice := parse_input("../test_data/Book2.xlsx")

	if len(gene_slice) != 4 {
		t.Errorf("Expected the length of the returned slice to be 4 instead it equaled %d", len(gene_slice))
	}
}
