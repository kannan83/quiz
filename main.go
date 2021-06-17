package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const csvFile = "problems.csv"

func main() {
	// open csv file
	csvfile, err := os.Open(csvFile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully opened CSV file")
	defer csvfile.Close()

	// read data from csv
	consoleReader := bufio.NewReader(os.Stdin)
	r := csv.NewReader(csvfile)
	var nbResOk, nbQuestions = 0, 0 // contains how much responses are correct

	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
			//fmt.Println(err)
		}

		nbQuestions += 1
		fmt.Print("What is ", strings.TrimSpace(rec[0]), " ?: ")
		ans, err := consoleReader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		ans = strings.TrimSpace(ans)
		if ans == strings.TrimSpace(rec[1]) {
			nbResOk += 1
		}
	}

	fmt.Println("Result: ", nbResOk, " good responses out of ", nbQuestions)

}
