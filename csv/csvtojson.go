package csv

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

type header []string
type contentRow map[string]interface{}
type content []contentRow

// ReadCSV file from first argument
func ReadCSV() ([]byte, string) {
	file, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	curDir := currentDir()

	filename := curDir + "/" + (os.Args[1]) + ".json"
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
	if err != nil {
		log.Fatal(err)
	}
	return r, filename
}

func SaveToJsonFile(filename string, r []byte) error {
	return ioutil.WriteFile(filename, r, 0666)
	// saveToFile(filename, r)
	// fmt.Println(filename, "saved to disk.")
}

// get current dir (dir where program was ran)
func currentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}
