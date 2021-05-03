package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func format_query(gene_id string) map[string]string {
	jsonData := map[string]string{
		"query": `
            { 
                people {
                    firstname,
                    lastname,
                    website
                }
            }
        `,
	}
	return jsonData
}

func fetch_response(api_website string, gene_list []string) {

	for i := 0; i < len(gene_list); i++ {
		response, error := http.Get(api_website)

		if error != nil {
			log.Fatalln(error)
		}
		defer response.Body.Close()

		body, error := ioutil.ReadAll(response.Body)

		if error != nil {
			log.Fatalln(error)
		}
		fmt.Println(string(body))
	}
}
