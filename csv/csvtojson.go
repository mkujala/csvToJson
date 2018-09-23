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
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

type header []string
type contentRow map[string]interface{}

// Content is type for the []map that can be marshalled to json
type Content []contentRow

// Read file from first argument
func Read() (Content, string) {
	inputFilename := os.Args[1]
	checkFileType(inputFilename)
	outputFilename := removeExtension(inputFilename)

	file, err := os.Open(inputFilename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	curDir := currentDir()
	filename := curDir + "/" + (outputFilename) + ".json"
	h := header{}
	data := Content{}
	reader := csv.NewReader(bufio.NewReader(file))

	for i := 0; ; i++ {
		// intialize new row on every iteration
		cr := contentRow{}
		line, err := reader.Read()

		// break when end of file reached
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}

		// save csv's first row fields as header fields
		if i == 0 {
			for _, name := range line {
				h = append(h, name)
			}
		}

		// iterate through csv content rows
		for j, value := range line {
			// try to convert string->float, if no err save as float else save as string
			floatVal, err := strconv.ParseFloat(value, 64)
			if err == nil {
				cr[h[j]] = floatVal
			} else {
				cr[h[j]] = value
			}
		}
		// append new row to json array
		data = append(data, cr)
	}
	return data, filename
}

// SaveToJSONFile saves Marshalled json to disk
func SaveToJSONFile(filename string, c Content) {
	bs, err := json.Marshal(c)
	if err != nil {
		log.Fatal(err)
	}
	err2 := ioutil.WriteFile(filename, bs, 0666)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(filename, "saved to disk.")
}

// PrintJSON makes pretty print of generated JSON
func PrintJSON(c Content) {
	prettyJSON, err := (json.MarshalIndent(c, "", "    "))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(prettyJSON))
}

// get current dir (dir where program was ran)
func currentDir() string {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	return dir
}

func checkFileType(fn string) {
	if filepath.Ext(fn) != ".csv" {
		log.Fatal("Error: input file type must be .csv")
	}
}

func removeExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}
