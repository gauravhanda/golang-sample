package main

import (
	"log"
	"os"
)

func main() {
	var fileHandle, err = os.Open("README.md")
	switch err {
	case nil:
		log.Println("File read successsfully")
		var bytesRead = 1
		data := make([]byte, 100)

		for bytesRead > 0 {
			bytesRead, err := fileHandle.Read(data)
			if err != nil {
				log.Println(err)
				break
			} else {
				log.Printf("Bytes read %d", bytesRead)
				log.Println(string(data))
			}
		}
		fileHandle.Close()
	default:
		log.Print("Unable to read file")
		log.Println(err)
	}

}
