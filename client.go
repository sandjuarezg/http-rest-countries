package main

import (
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
	var opc int
	var fields string = "name;capital;region;subregion;population;area;languages"
	var countries []country

	for {

		var url string = "https://restcountries.eu/rest/v2/region/"

		fmt.Println("0. Exit")
		fmt.Println("1. Africa")
		fmt.Println("2. Americas")
		fmt.Println("3. Asia")
		fmt.Println("4. Europe")
		fmt.Println("5. Oceania")
		fmt.Scan(&opc)

		switch opc {
		case 0:

			fmt.Println("E X I T I N G . . .")
			os.Exit(0)

		case 1:

			url = fmt.Sprintf("%safrica?fields=%s", url, fields)

		case 2:

			url = fmt.Sprintf("%samericas?fields=%s", url, fields)

		case 3:
			url = fmt.Sprintf("%sasia?fields=%s", url, fields)

		case 4:
			url = fmt.Sprintf("%seurope?fields=%s", url, fields)

		case 5:
			url = fmt.Sprintf("%soceania?fields=%s", url, fields)

		default:

			continue

		}

		func() {
			request, err := http.NewRequest("GET", url, nil)
			if err != nil {
				log.Fatal(err)
			}

			request.Header.Set("Accept", "application/json")

			response, err := client.Do(request)
			if err != nil {
				log.Fatal(err)
			}
			defer response.Body.Close()

			if response.StatusCode != 200 {
				log.Fatal(response.Status)
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
		}()
	}
}
