package main

import (
	"os"
	"testing"
)

func Test_parser(t *testing.T) {
	//saving the state of the args before the unit test
	oldArgs := os.Args
	defer func() { os.Args = oldArgs }()

	//setting the arguments
	os.Args = []string{"cmd", "-input=./test/", "-output=./test_output/", "-api=api_url"}

	input_path, output_path, api_endpoint := parser()

	if input_path != "./test/" {
		t.Errorf("Expected the input path to equal ./test/ insteasd it equaled: %s", input_path)
	}
	if output_path != "./test_output/" {
		t.Errorf("Expected the output path to equal ./test/ insteasd it equaled: %s", output_path)
	}
	if api_endpoint != "api_url" {
		t.Errorf("Expected the api_endpoint variable to equal ./test/ insteasd it equaled: %s", api_endpoint)
	}
}
