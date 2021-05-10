package main

import (
	"fmt"
	"log"
	"os"
)

func write_to_file(geneData []Data, output_path string) error {
	fmt.Println("writing the returned information to a text file at " + output_path)

	file, err := os.Create(output_path)

	if err != nil {
		log.Fatal(err)
		return err
	}
	// writing the header string to a file
	file.WriteString("gene_id\tgene_start\tgene_stop\tOMIM_id\tgene_name\tchromosome_number\n")

	defer file.Close()

	for i := 0; i < len(geneData); i++ {

		//extracting the struct from the geneData slice
		geneInfoStruct := geneData[i]

		geneInfoString := fmt.Sprintf("%s\t%d\t%d\t%s\t%s\t%s\n", geneInfoStruct.Gene_name, geneInfoStruct.Data.Gene.Start, geneInfoStruct.Data.Gene.Stop, geneInfoStruct.Data.Gene.Omim_id, geneInfoStruct.Data.Gene.Name, geneInfoStruct.Data.Gene.Chrom)

		file.WriteString(geneInfoString)
	}

	return nil
}
