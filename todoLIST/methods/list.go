package main

import "fmt"

func ListTodos() []Todos {
	for _, v := range All {
		fmt.Printf("%v\t%s\n", v.ID, v.List)
	}
	return All
}
