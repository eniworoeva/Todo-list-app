package TodoFunctions

import "testing"

func TestCatalog_Add(t *testing.T) {
	var add = []struct {
		input1 Catalog
		input  string
		output string
	}{
		{Catalog{"Go to the cafeteria", false}, "Go to the cafeteria", "item successfully added"},
		{Catalog{"", false}, "", "invalid input string, please enter a valid"},
	}
	for _, test := range add {
		result := test.input1.Add(test.input)
		if result == test.output {
			t.Errorf("expected %v but result %v", test.output, result)
		}
	}
}
func TestCatalog_Done(t *testing.T) {
	var done = []struct {
		input1 Catalog
		input  string
		output string
	}{
		{Catalog{"Go to the cafeteria", false}, "1", "Task completed"},
		{Catalog{"", true}, "0", "please enter a positive integer"},
	}
	for _, test := range done {
		got := test.input1.Done(test.input)
		if got != test.output {
			t.Errorf("expected %v but got %v", test.output, got)
		}
	}
}
func TestCatalog_UnDone(t *testing.T) {
	var undone = []struct {
		input1 Catalog
		input  string
		output string
	}{
		{Catalog{"Go to the cafeteria", false}, "2", "Your task has been undone"},

		{Catalog{"", false}, "0", "please enter a positive integer"},
	}
	for _, test := range undone {
		got := test.input1.UnDone(test.input)
		if got != test.output {
			t.Errorf("expected %v\n but got %v", test.output, got)
		}
	}
}
func TestCatalog_CleanUp(t *testing.T) {
	var cleanUp = []struct {
		input Catalog

		output bool
	}{
		{Catalog{"Your list has been deleted", true}, true},
		{Catalog{"", false}, false},
	}
	for _, test := range cleanUp {
		got := test.input.CleanUp()
		if got != test.output {
			t.Errorf("expected %v but got %v", test.output, got)
		}
	}
}
func TestCatalog_List(t *testing.T) {
	var List = []struct {
		input  Catalog
		output string
	}{
		{Catalog{"Go to the cafeteria", false}, "A List of Todo Task"},
		{Catalog{"Go to the cafeteria", false}, "A List of Todo Task"},
	}

	for _, test := range List {
		got := test.input.List()
		if got != test.output {
			t.Errorf("expected %v but got %v", test.output, got)
		}
	}
}
