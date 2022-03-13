package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
)

type Todos struct {
	ID   int
	List string
}

var All []Todos

func init() {
	ReadAndWriteCsv()
	OpenAndReadDoneCSV()
	createCSV()
}

func (l *Todos) AddToTodo(id int, text string) string {
	t := Todos{
		ID:   id,
		List: text,
	}
	for _, v := range All {
		if id == v.ID {
			return "You cannot have duplicate IDs"
		}
	}
	All = append(All, t)

	CsvFileTodos()
	return "a new task has been added"
}

func CsvFileTodos() {
	csvFile, err := os.Create("allTodos.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}
	defer csvFile.Close()

	csvwriter := csv.NewWriter(csvFile)
	defer csvwriter.Flush()
	for _, v := range All {
		row := []string{strconv.Itoa(v.ID), v.List}
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}

}

func ReadAndWriteCsv() []Todos {
	file, err := os.Open("allTodos.csv")
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	records, err := csv.NewReader(file).ReadAll()
	if err != nil {
		log.Println(err)
	}
	for _, v := range records {
		IdInt, _ := strconv.Atoi(v[0])
		data := Todos{
			ID:   IdInt,
			List: v[1],
		}
		All = append(All, data)

	}
	return All

}

func overWriteCsv() []Todos {
	file, err := os.OpenFile("allTodos.csv", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err != nil {
		log.Println(err)
	}
	defer file.Close()

	csvwriter := csv.NewWriter(file)
	defer csvwriter.Flush()
	for _, v := range All {
		row := []string{strconv.Itoa(v.ID), v.List}
		if err := csvwriter.Write(row); err != nil {
			log.Fatalln("error writing record to file", err)
		}
	}
	return All

}
