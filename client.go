package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

type languages struct {
	Iso639_1   string `json:"iso_639___1"`
	Iso639_2   string `json:"iso_639___2"`
	Name       string `json:"name"`
	NativeName string `json:"native_name"`
}

type country struct {
	Name       string      `json:"name"`
	Capital    string      `json:"capital"`
	Region     string      `json:"region"`
	Subregion  string      `json:"subregion"`
	Population int         `json:"population"`
	Area       float32     `json:"area"`
	Languages  []languages `json:"languages"`
}

var client *http.Client = &http.Client{}

func main() {
	var rStdin *bufio.Reader = bufio.NewReader(os.Stdin)
	var url string
	var fields string = "name;capital;region;subregion;population;area;languages"
	var countries []country
	var ban bool = true

	for {
		fmt.Println("0. Exit")
		fmt.Println("1. Africa")
		fmt.Println("2. Americas")
		fmt.Println("3. Asia")
		fmt.Println("4. Europe")
		fmt.Println("5. Oceania")
		reply, _, err := rStdin.ReadLine()
		if err != nil {
			log.Fatal(err)
		}

		switch string(reply) {
		case "0":

			fmt.Println("E X I T I N G . . .")
			os.Exit(0)

		case "1":

			url = "https://restcountries.eu/rest/v2/region/africa?fields=" + fields
			ban = true

		case "2":

			url = "https://restcountries.eu/rest/v2/region/americas?fields=" + fields
			ban = true

		case "3":

			url = "https://restcountries.eu/rest/v2/region/asia?fields=" + fields
			ban = true

		case "4":

			url = "https://restcountries.eu/rest/v2/region/europe?fields=" + fields
			ban = true

		case "5":

			url = "https://restcountries.eu/rest/v2/region/oceania?fields=" + fields
			ban = true

		default:

			ban = false

		}

		if ban {

			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal(err)
			}

			request.Header.Set("Accept", "application/json")

			response, err := client.Do(request)
			if response.StatusCode != 200 {
				log.Fatal(response.Status)
			}
			if err != nil {
				log.Fatal(err)
			}

			err = json.NewDecoder(response.Body).Decode(&countries)
			if err != nil {
				log.Fatal(err)
			}

			for i := 0; i < len(countries); i++ {
				for j := 0; j < len(countries[i].Languages); j++ {
					fmt.Printf("Name: %s\n", countries[i].Name)
					fmt.Printf("Capital: %s\n", countries[i].Capital)
					fmt.Printf("Region: %s\n", countries[i].Region)
					fmt.Printf("Subregion: %s\n", countries[i].Subregion)
					fmt.Printf("Population: %d\n", countries[i].Population)
					fmt.Printf("Area: %02f\n", countries[i].Area)
					fmt.Printf("Language: %s\n", countries[i].Languages[j].Name)
					fmt.Println("_________________________________________________")
				}
			}

			response.Body.Close()

		}
	}
}
