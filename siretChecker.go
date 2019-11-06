package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func main() {

	type Company struct {
		statut          string
		codeClient      string
		typeRangAdresse string
		nomClient       string
	}
	OBSMap := make(map[int]Company)
	OBSFileName := os.Args[2]
	OBSFile, err := os.Open(OBSFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	OBSReader := csv.NewReader(bufio.NewReader(OBSFile))
	OBSReader.Comma = ';'

	for i := 0; ; i++ {
		record, err := OBSReader.Read()
		if err == io.EOF || len(record[0]) == 0 {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		key, error := strconv.Atoi(record[0])
		if error == nil {
			OBSMap[key] = Company{"X", record[2], record[3], record[4]}
		}
	}

	inseeFileName := os.Args[1]
	inseeFile, err := os.Open(inseeFileName)
	if err != nil {
		log.Fatalf("Couldn't open %v\n", err)
	}
	inseeReader := csv.NewReader(bufio.NewReader(inseeFile))
	inseeReader.Comma = ','

	for i := 0; ; i++ {
		record, err := inseeReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		key, _ := strconv.Atoi(record[2])
		_, exist := OBSMap[key]
		if exist {
			OBSMap[key] = Company{record[40], OBSMap[key].codeClient, OBSMap[key].typeRangAdresse, OBSMap[key].nomClient}
		}
	}
	//fmt.Printf("%v\n", OBSMap)
	for key, company := range OBSMap {
		fmt.Printf("%v;%v;%v;%v;%v\n", key, company.statut, company.codeClient, company.typeRangAdresse, company.nomClient)
	}
}
