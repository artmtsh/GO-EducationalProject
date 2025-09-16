package scanner

import (
	"bufio"
	"fmt"
	"github.com/k0kubun/pp"
	"os"
	"strings"
	"todolistapp/logger"
	"todolistapp/todo"
)

type Scanner struct {
	scanner *bufio.Scanner
	list    *todo.Storage
	logger  *logger.Logger
}

func NewScanner() Scanner {
	storage := todo.Storage{}
	logger := logger.Logger{}
	return Scanner{
		scanner: bufio.NewScanner(os.Stdin),
		list:    storage.CreateStorage(),
		logger:  logger.CreateLogger(),
	}
}

func Start() {
	newScanner := NewScanner()
	for {
		pp.Println("Введите команду")
		if ok := newScanner.scanner.Scan(); !ok {
			newScanner.logger.AddEvent(userInputError)
			fmt.Println(userInputError)
			return
		}
		text := newScanner.scanner.Text()
		fields := strings.Fields(text)
		if len(fields) == 0 {
			pp.Println("Вы ничего не ввели")
			newScanner.logger.AddEvent("Пользователь ничего не ввел")
			continue
		}
		mapOfFunctions := createMapOfFunctions(newScanner.list)
		mapOfFunctionsWithStringArg := createMapOfFunctionsWithStringArg(newScanner.list)
		if fields[0] == "exit" {
			break
		}
		if run, ok := mapOfFunctions[fields[0]]; !ok {
			if run1, ok1 := mapOfFunctionsWithStringArg[fields[0]]; !ok1 {
				pp.Println(noSuchCmdError)
				newScanner.logger.AddEvent(noSuchCmdError)
			} else {
				run1(newScanner.logger, strings.Join(fields[1:], " "))
			}
		} else {
			run(newScanner.logger)
		}
		logger.AddTextToLog(&newScanner.logger.Events[len(newScanner.logger.Events)-1], text)

	}
}

func createMapOfFunctions(storage *todo.Storage) map[string]func(*logger.Logger) {
	return map[string]func(*logger.Logger){
		"help":   Help,
		"list":   storage.ListAllTask,
		"events": logger.PrintAllLogs,
	}
}
func createMapOfFunctionsWithStringArg(storage *todo.Storage) map[string]func(*logger.Logger, string) {
	return map[string]func(*logger.Logger, string){
		"listTask": storage.ListTaskByTitle,
		"del":      storage.DeleteTask,
		"done":     storage.CompleteTask,
		"add":      storage.CreateAndAddTask,
	}
}

func Help(*logger.Logger) {
	fmt.Println("- Список команд, которые должны быть доступны в приложении:\n" +
		"- help — эта команда позволяет узнать доступные команды и их формат\n" +
		"- add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач\n" +
		"- list — эта команда позволяет получить полный список всех задач\n" +
		"- del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку\n" +
		"- done {заголовок существующей задачи} — эта команда позволяет отметить задачу как выполненную\n" +
		"- events — эта команда позволяет получить список всех событий\n" +
		"- exit — эта команда позволяет завершить выполнение программы")
}
