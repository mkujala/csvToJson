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
	if len(os.Args) < 2 || len(os.Args) > 3 {
		fmt.Println("Give one .csv file as an argument!")
		os.Exit(1)
	}
	if len(os.Args) == 3 && os.Args[2] != "-p" {
		fmt.Printf("Invalid flag \"%v\" after filename %v\n", os.Args[2], os.Args[1])
		os.Exit(1)
	}
	data, filename := csv.ReadCSV()

	if len(os.Args) == 3 {
		csv.PrintJSON(data)
		return
	}
	csv.SaveToJSONFile(filename, data)
}
