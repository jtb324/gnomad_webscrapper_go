package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func format_query(gene_id string) []byte {
	/*function to format the graphQL query
	Parameters
	__________
	gene_id string
		string containing the id of the gene

	Returns
	_______
	[]byte
		returns a slice of bytes where the query dictionary was converted
		to a byte object
	*/
	jsonData := map[string]string{
		"query": `
			{
				gene(gene_symbol: "` + gene_id + `", reference_genome: GRCh37) {
		 			start
		 			stop
					omim_id
    				name
    				chrom
				}
		  	}
		  `,
	}

	jsonValue, _ := json.Marshal(jsonData)

	return jsonValue
}

func fetch_response(api_website string, gene_list []string) {
	/*function to fetch the reponse from the the bnomad api*/
	for i := 0; i < len(gene_list); i++ {

		jsonByteString := format_query(gene_list[i])
		// fmt.Println(jsonByteString)
		request, error := http.NewRequest("POST", api_website, bytes.NewBuffer(jsonByteString))

		if error != nil {
			log.Fatalln(error)
		}
		request.Header.Add("Content-Type", "application/json")

		client := &http.Client{Timeout: time.Second * 10}

		response, response_err := client.Do(request)

		if response_err != nil {
			log.Fatalf("The HTTP request failed with error %s\n", response_err)
		}

		//deferiing the responses close
		defer response.Body.Close()

		data, _ := ioutil.ReadAll(response.Body)
		fmt.Println(string(data))
		// body, error := ioutil.ReadAll(response.Body)

		// if error != nil {
		// 	log.Fatalln(error)
		// }
		// fmt.Println(string(body))
	}
}
