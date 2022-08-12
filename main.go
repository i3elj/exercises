package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
)

func report(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	filename := flag.String("file", "test.csv",
		"a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)
	report(err)
	
	r := csv.NewReader(file)
	lines, err := r.ReadAll()
	report(err)
	
	fmt.Println(lines)
}
