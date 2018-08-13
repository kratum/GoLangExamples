package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	//"strings"
	//"log"
	"bufio"
	"io/ioutil"
	"os"
)

type Person struct {
	Url  string `json:"url`
	Band string `json:"band`
	Jahr string `json:"jahr`
}

func main() {
	// Open our jsonFile
	csvFile, err := os.Open("./list.csv")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened list.csv")

	// defer the closing of our jsonFile so that we can parse it later on
	defer csvFile.Close()

	// ReadAll of jsonFile
	r := csv.NewReader(bufio.NewReader(csvFile))

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var personsRecord []Person

	// Über den Loop können die Werte aus dem CSV in ein struct gelesen werden
	for _, record := range records {
		data := Person{
			Url:  record[0],
			Band: record[1],
			Jahr: record[2],
		}
		personsRecord = append(personsRecord, data)
		//fmt.Println(data.Firstname + " " + data.Lastname)
	}

	// Die gelesenen Daten werden zu JSON konvertiert
	recordsJson, err := json.Marshal(personsRecord)

	// Das Json wird auf die Platte geschrieben
	err = ioutil.WriteFile("data.json", recordsJson, 0644)

	// und geprintet
	fmt.Fprintf(os.Stdout, "%s", recordsJson)

}
