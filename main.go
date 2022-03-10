/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package main

import (
	"log"
	"todo_list/cmd"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile | log.Lshortfile)
	cmd.Execute()
}
