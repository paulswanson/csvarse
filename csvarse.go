package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

var csvHeader = flag.Bool("h", true, "First line is header")
var ogIndexOverride = flag.Int("o", -1, "Specify the objectGUID index")

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Error: Please specifiy an input csv file.")
		return
	}

	file, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
		return
	}

	defer file.Close()

	r := csv.NewReader(file)
	var headerRow []string
	var ogIndex int

	if *csvHeader {
		headerRow, err = readLine(r)
		if err != nil {
			log.Fatal(err)
		}
		ogIndex = findFirst("objectGUID", headerRow)
	}

	if *ogIndexOverride >= 0 {
		ogIndex = *ogIndexOverride
	}

	out := csv.NewWriter(os.Stdout)
	out.Write(headerRow)
	out.Flush()

	var dataRow []string
	for {
		dataRow, err = readLine(r)
		if err != nil {
			return
		}
		dataRow[ogIndex] = getHex(dataRow[ogIndex])
		out.Write(dataRow)
		out.Flush()
	}

}

func readLine(r *csv.Reader) ([]string, error) {
	readFields, err := r.Read()
	if err != nil {
		return readFields, err
	}

	return readFields, nil
}

func findFirst(s string, slice []string) int {
	for p, v := range slice {
		if strings.TrimSpace(v) == s {
			return p
		}
	}
	return -1
}

func getHex(s string) string {
	return strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(s), "X'"), "'")
}
