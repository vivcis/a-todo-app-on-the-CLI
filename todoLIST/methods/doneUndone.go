package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

var done []Todos

func TaskDone(index int) []Todos {
	for i, v := range All {
		if index == v.ID {
			done = append(done, v)
			All = append(All[:i], All[i+1:]...)

		}
	}
	overWriteCsv()
	openAndWriteDoneCSV()
	return All
}

func TaskNotDone(index int) []Todos {
	for i, v := range done {
		if index == v.ID {
			done = append(done[:i], done[i+1:]...)
			All = append(All, v)
		}
	}
	overWriteCsv()
	openAndWriteDoneCSV()
	return All
}

func createCSV() {
	csvFile, err := os.Create("done.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	for _, v := range done {
		row := []string{strconv.Itoa(v.ID), v.List}
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}

func OpenAndReadDoneCSV() []Todos {
	file, err := os.Open("done.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err)
	}
	for _, v := range records {
		IdInt, err := strconv.Atoi(v[0])
		if err != nil {
			log.Fatalf("An error occured: %v", err.Error())
		}
		data := Todos{
			ID:   IdInt,
			List: v[1],
		}
		done = append(done, data)
	}
	return done

}

func openAndWriteDoneCSV() []Todos {
	file, err := os.OpenFile("done.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	for _, v := range done {
		//fmt.Println()
		row := []string{strconv.Itoa(v.ID), v.List}
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

	return done

}
