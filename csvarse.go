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

//var sqlCreate = flag.Bool("sql", false, "Output as SQL create query")
//var sqlTableName = flag.String("t", "ad_users", "Specify default SQL table name")
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

	//	if *sqlCreate && {
	//		sqlCreateHeader()
	//	}

	var dataRow []string
	for {
		dataRow, err = readLine(r)
		if err != nil {
			return
		}
		dataRow[ogIndex] = getHex(dataRow[ogIndex])
		//		if *sqlCreate {
		//			sqlCreateValue(dataRow)
		//		} else {
		//			fmt.Printf("%v\n", dataRow)
		//fmt.Printf("%v\n", dataRow)
		out.Write(dataRow)
		out.Flush()
	}

	//	if *sqlCreate {
	//		sqlCreateFooter()
	//	}

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

//func sqlCreateHeader() {
//	fmt.Println("SET SQL_MODE = \"NO_AUTO_VALUE_ON_ZERO\";")
//	fmt.Printf("CREATE TABLE IF NOT EXISTS `%v` (\n", sqlTableName)
//	fmt.Printf("`DN` ...") // To be completed one day ...
//}
//func sqlCreateValue(s string) {
//}
//func sqlCreateFooter() {}
