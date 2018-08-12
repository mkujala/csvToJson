package utils

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

type header []string
type contentRow map[string]interface{}
type content []contentRow

// ReadCSV file from first argument
func ReadCSV() {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// create public folder if not exist
	if _, err := os.Stat("public"); os.IsNotExist(err) {
		os.Mkdir("public", 0775)
	}

	filename := "public/" + (os.Args[1]) + ".json"
	h := header{}
	c := content{}
	reader := csv.NewReader(bufio.NewReader(file))

	for i := 0; ; i++ {

		// intialize new row on every iteration
		cr := contentRow{}
		line, error := reader.Read()

		// break when end of file reached
		if error == io.EOF {
			break
		} else if error != nil {
			log.Fatal(error)
		}

		// save csv's first row fields as header fields
		if i == 0 {
			for _, name := range line {
				h = append(h, name)
			}
		} else {
			for j, value := range line {

				// try to convert string->float, if no err save as float else save as string
				floatVal, err := strconv.ParseFloat(value, 64)
				if err == nil {
					cr[h[j]] = floatVal
				} else {
					cr[h[j]] = value
				}
			}
			c = append(c, cr)
		}
	}
	r, err := json.Marshal(c)
	saveToFile(filename, r)
	fmt.Println(filename, "saved to disk.")
}

func saveToFile(filename string, r []byte) error {
	return ioutil.WriteFile(filename, r, 0666)
}
