// csvToJson converts a csv file to json and saves it to disk.
// You can convert one file at time
//
// USAGE: csv2json yourFile.csv
package main

import (
	"csv2json/csv"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give one .csv file as an argument!")
		os.Exit(1)
	}
	json, filename := csv.ReadCSV()
	csv.SaveToJSONFile(filename, json)
}
