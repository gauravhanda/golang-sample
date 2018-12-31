package main

import (
	"flag"
	"log"
)

var studentListSize *int

func init() {
	studentListSize = flag.Int("count", 1, "Please enter the size of array")
	flag.Parse()
}

func main() {

	var myslice = make([]int, *studentListSize, *studentListSize)
	for i := 0; i < len(myslice); i++ {
		log.Printf("Counter index = %d", i)
	}
}
