package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Person struct {
	Name string
	Age  string
}

type Persons []Person

var persons = Persons{
	Person{
		"Peter",
		"23",
	},
	Person{
		"Paul",
		"12",
	},
}

func main() {
	personsJson, err := json.Marshal(persons)
	if err != nil {
		log.Fatal("Cannot encode to JSON ", err)
	}
	fmt.Fprintf(os.Stdout, "%s", personsJson)
}
