package main

import (
	"encoding/json"
	"fmt"
	//"log"
	"io/ioutil"
	"os"
)

type Person struct {
	Name string `json:"Name`
	Age  string `json:"Age`
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("./data.json")

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened data.json")

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// ReadAll of jsonFile
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// Initialize persons as Slice of Persons
	var persons []Person

	// Unmarshal the byteValues to persons
	json.Unmarshal([]byte(byteValue), &persons)

	fmt.Print(persons, "\n")
	// for each Person print Name and age
	for i := 0; i < len(persons); i++ {
		fmt.Printf("Name: %v, Age: %v \n", persons[i].Name, persons[i].Age)
	}

}
