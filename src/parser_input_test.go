package main

import (
	"testing"
)

func Test_parser_input_len(t *testing.T) {

	gene_slice := parse_input("../test_data/test_data.txt")

	if len(gene_slice) != 3 {
		t.Errorf("Expected the length of the returned slice to be 4 instead it equaled %d", len(gene_slice))
	}
}

func Test_parser_input_element(t *testing.T) {

	gene_slice := parse_input("../test_data/test_data.txt")

	if gene_slice[0] != "HPSE" {
		t.Errorf("Expected the first element of the returned slice to be APC instead it was %s", gene_slice[0])
	}
}
