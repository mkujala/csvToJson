// csvToJson converts a csv file to json and saves it to disk.
// You can convert one file at time
//
// USAGE: csvToJson yourFile.csv
package main

import (
	"csvToJson/utils"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Give one .csv file as an argument!")
		os.Exit(1)
	}
	utils.ReadCSV()
}
