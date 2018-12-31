package main

import (
	"log"

	. "github.com/gauravhanda/golang-sample/mathutil"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Println("Recovered from error ", err)
		}
	}()
	log.Println("Hello World")
	log.Printf("PI = %f , Sin = %f \r\n", PI(), Sin(10))
	var result, err = Divide(0, 10)
	if err == nil {
		log.Printf("Quotient = %f", result)
	} else {
		log.Printf("Error %s", err)
	}
}
