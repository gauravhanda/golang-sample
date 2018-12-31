package main

import (
	"encoding/json"
	"log"
)

const jsonstring = `
	{
		"firstName" : "gaurav", 
		"lastName" : "Handa",
		"country" : "US",
		"zip" : "12345"
	}`

//Employee : struct for represening json
type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Country   string `json:"-"`
	Zip       string `json:"zip,omitempty"`
}

func main() {
	var employee Employee
	log.Println("parsing json " + jsonstring)
	error := json.Unmarshal([]byte(jsonstring), &employee)

	if error == nil {
		log.Print("Json Parsed successfully. Employee Object := ")
		log.Println(employee)
		out, err := json.Marshal(employee)
		if err != nil {
			log.Print("Faild to generate json")
			log.Println(err)
		} else {
			log.Println("Json Generated successfully" + string(out))
		}

	} else {
		log.Print("Faild to parse json")
		log.Print(error)
	}

}
