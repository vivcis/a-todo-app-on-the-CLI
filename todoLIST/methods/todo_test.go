package main

import "testing"

//----------------addToTodo TESTING-------------------------------------------------------
func TestAddToTodo(t *testing.T) {
	cece := Todos{1, "hello"}
	addTodos := []struct {
		input1 int
		input2 string

		expected string
	}{
		{1, "sleep", "a new task has been added"},
		{2, "take my path", "a new task has been added"},
		//{"eat", 3, "a new task has been added"},
	}

	for _, val := range addTodos {
		got := cece.AddToTodo(val.input1, val.input2)
		if got != val.expected {
			t.Errorf("AddToTodo : expected output %v : ouput %v", val.expected, got)
		}
	}

}

//----------------------DoneTodo TESTING---------------------------------------------------------------
func TestDoneTodo(t *testing.T) {
	//cece := []Todos{{"sleep", 1}}
	doneTodos := []struct {
		input          int
		expectedOutput []Todos
	}{
		{1, []Todos{{2, "take my path"}}},
		{2, []Todos{}},
	}

	for _, v := range doneTodos {
		got := TaskDone(v.input)
		if len(got) != len(v.expectedOutput) {
			t.Errorf("DoneTodo : expected output %v : ouput %v", v.expectedOutput, got)
		}
	}
}

//-------------------NotDoneTodo TESTING-----------------------------------------------------------
func TestNotDoneTodo(t *testing.T) {
	//cece := []Todos{{"sleep", 1}}
	doneTodos := []struct {
		input          int
		expectedOutput []Todos
	}{
		{1, []Todos{{1, "sleep"}}},
		{2, []Todos{{1, "sleep"}, {2, "take my path"}}},
	}

	for _, v := range doneTodos {
		got := TaskNotDone(v.input)
		if len(got) != len(v.expectedOutput) {
			t.Errorf("notDoneTodo : expected output %v : ouput %v", v.expectedOutput, got)
		}
	}
}

//---------------------ListOfTodos TESTING------------------------------------------------------------
func TestListOfTodos(t *testing.T) {
	listTodos := []Todos{{1, "sleep"}, {2, "take my path"}}
	got := ListTodos()
	if len(got) != len(listTodos) {
		t.Errorf("listOfTodos : expected output %v : ouput %v", listTodos, got)
	}
}

//----------------deleteTodo TESTING-------------------------------------------------------------------
func TestDeleteTodo(t *testing.T) {
	deleteTodo := []struct {
		input          string
		expectedOutput string
	}{
		{"sleep", "sleep has been deleted from the done list"},
		{"take my path", "take my path has been deleted from the done list"},
	}

	for _, v := range deleteTodo {
		got := DeleteFromTodo(v.input)
		if got != v.expectedOutput {
			//log.Print(got)
			t.Errorf("DeleteTodo : expected output %v : ouput %v", v.expectedOutput, got)
		}
	}
}
