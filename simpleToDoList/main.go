package main

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"main.go/funcs"
	"main.go/structs"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	storageOfStructs := structs.Storage{}
	storageOfStructs.CreateStorage()
	logger := funcs.Logger{}
	for {
		pp.Println("Введите команду")
		if ok := scanner.Scan(); !ok {
			fmt.Println("Произошла ошибка пользовательского ввода")
			return
		}
		text := scanner.Text()
		logger.AddEvent(text)
		fields := strings.Fields(text)
		if len(fields) == 0 {
			pp.Println("Вы ничего не ввели")
			continue
		}

		if fields[0] == "exit" {
			break
		} else if fields[0] == "help" {
			funcs.Help()
			continue
		} else if fields[0] == "add" {
			newTask := structs.CreateTask(strings.Join(fields[1:], " "))
			storageOfStructs.AddTask(&newTask)
		} else if fields[0] == "list" {
			storageOfStructs.ListAllTask()
		} else if fields[0] == "listTask" {
			storageOfStructs.ListTaskByTitle(fields[1])
		} else if fields[0] == "del" {
			storageOfStructs.DeleteTask(fields[1])
		} else if fields[0] == "done" {
			storageOfStructs.GetTaskByTitle(fields[1]).CompleteTask()
		} else if fields[0] == "events" {
			logger.PrintAllLogs()
		} else {
			pp.Println("Такой команды не существует")
		}
		text = scanner.Text()
		fields = strings.Fields(text)
	}
}
