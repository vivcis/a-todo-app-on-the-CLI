package main

import "fmt"

func DeleteFromTodo(todo string) string {
	for i, v := range done {
		if todo == v.List {
			done = append(done[:i], done[i+1:]...)
			s := fmt.Sprintf("%s has been deleted from the done list", todo)
			openAndWriteDoneCSV()
			return s
		}
	}
	return "nil"
}
