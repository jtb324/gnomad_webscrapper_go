package main

import (
	"os"
	"testing"
)

func Test_write_to_file(t *testing.T) {
	//unit test to test the write_to_file function
	var gene_slice []Data

	data1 := Data{Gene_name: "test", Data: Gene{Gene: gene_info{Start: 1, Stop: 2, Omim_id: "1234", Name: "test gene", Chrom: "6"}}}

	data2 := Data{Gene_name: "test2", Data: Gene{Gene: gene_info{Start: 3, Stop: 4, Omim_id: "4444", Name: "test gene 2", Chrom: "8"}}}

	gene_slice = append(gene_slice, data1, data2)

	err := write_to_file(gene_slice, "./test.txt")

	if err != nil {
		t.Errorf("expected no error while written to a file but instead there was an error: %s", err)
	}

	_, err = os.Stat("./test.txt")

	if os.IsNotExist(err) {
		t.Errorf("Expected output file was not created")
	} else {
		os.Remove("./test.txt")
	}

}
