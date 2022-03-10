package TodoFunctions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
)

type Catalog struct {
	IndividualLog string `json:"individualLog"`
	Status        bool   `json:"status"`
}

var L = Catalog{}
var todoList = []Catalog{}

func (t *Catalog) unmarshalCsv() {
	CsvData, _ := ioutil.ReadFile("todoList.csv")
	if string(CsvData) != "" {
		err := json.Unmarshal(CsvData, &todoList)
		if err != nil {
			log.Fatal(err)
		}
	}

}

func init() {
	L.unmarshalCsv()
}

func (t *Catalog) savetoCsv() {
	jsonData, err := json.Marshal(todoList)
	if err != nil {
		log.Fatal(err)
	}

	//CsvData = append(CsvData, jsonData...)
	err = ioutil.WriteFile("todoList.csv", jsonData, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func (t Catalog) Add(IndividualLog string) string {
	t.IndividualLog = IndividualLog
	todoList = append(todoList, t)
	t.savetoCsv()
	//match the output in tests
	return "item successfully added"

}
func (t *Catalog) Done(stringedSerialNo string) string {
	//unmarshalCsv()
	serialNo, err := strconv.Atoi(stringedSerialNo)
	if err != nil {
		panic(err)
	}
	if serialNo != 0 {
		for i := range todoList {
			if i == (serialNo - 1) {
				todoList[i].Status = true

				fmt.Println("Task completed")
			}
		}
		t.savetoCsv()
		return "item status updated"
	}
	fmt.Println("please enter a positive integer")
	return "item status updated"
}

//func (t *Catalog) UnDone(stringedSerialNo string) string {
//	serialNo, err := strconv.Atoi(stringedSerialNo)
//
//	if err != nil {
//		panic(err)
//	}
//	if serialNo != 0 {
//		for i := range todoList {
//			if i == (serialNo-1) && todoList[i].Status == true {
//				todoList[i].Status = false
//			}
//			t.savetoCsv()
//			fmt.Println("Your task has been undone")
//			return "Go to the cafeteria"
//		}
//		fmt.Println("item status is already Undone please enter the correct serial no")
//		return "Go to the cafeteria"
//	}
//	fmt.Println("please enter a positive integer")
//	return "Go to the cafeteria"
//}

func (t *Catalog) UnDone(stringedSerialNo string) string {
	serialNo, err := strconv.Atoi(stringedSerialNo)
	if err != nil {
		log.Fatal("error")
	}
	if serialNo == 0 {
		fmt.Println("please enter a positive integer")
		return "please enter a positive integer"
	}
	for i := range todoList {
		if i == (serialNo - 1) {
			todoList[i].Status = false
			fmt.Println("item status updated")
		}
		t.savetoCsv()
		return "kpomo"
	}
	return "Your task has been undone"
}

func (t *Catalog) CleanUp() bool {
	for i, value := range todoList {
		if value.Status == true {
			todoList = append(todoList[:i], todoList[1+1:]...)
			t.savetoCsv()
			fmt.Println("Your list has been deleted")
			return true
		}
	}
	//fmt.Println("please enter a valid input string")
	return false
}
func (t *Catalog) List() string {
	for i, value := range todoList {
		if value.Status == false {
			fmt.Printf("%v, %v\n", i+1, value.IndividualLog)
			continue

		}

	}
	return "A List of Todo Task"
}
