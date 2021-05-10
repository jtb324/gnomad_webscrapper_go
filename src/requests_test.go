package main

import (
	"testing"
)

func Test_fetch_response(t *testing.T) {

	var gene_list []string

	gene_list = append(gene_list, "EXT1", "CFTR")

	data_slice := fetch_response("https://gnomad.broadinstitute.org/api", gene_list, "./")

	if len(data_slice) != 2 {
		t.Errorf("Expectecd the returned slice to have a length of 2 instead it had a length of %d", len(data_slice))
	}
}

func Test_fetch_response_values(t *testing.T) {
	//unit test to test the actual values of the structures returned
	var gene_list []string

	gene_list = append(gene_list, "HPSE", "EXT1", "CFTR")

	data_slice := fetch_response("https://gnomad.broadinstitute.org/api", gene_list, "./")

	hpse_info := data_slice[0]

	if hpse_info.Data.Gene.Start != 84213614 {
		t.Errorf("Expected the value of the start column to be 84213614, instead it was %d", hpse_info.Data.Gene.Start)
	}
	if hpse_info.Data.Gene.Stop != 84256306 {
		t.Errorf("Expected the value of the start column to be 84256306, instead it was %d", hpse_info.Data.Gene.Stop)
	}
	if hpse_info.Data.Gene.Omim_id != "604724" {
		t.Errorf("Expected the value of the start column to be 604724, instead it was %s", hpse_info.Data.Gene.Omim_id)
	}
	if hpse_info.Data.Gene.Name != "heparanase" {
		t.Errorf("Expected the value of the start column to be heparanase, instead it was %s", hpse_info.Data.Gene.Name)
	}
	if hpse_info.Data.Gene.Chrom != "4" {
		t.Errorf("Expected the value of the start column to be heparanase, instead it was %s", hpse_info.Data.Gene.Chrom)
	}
}
