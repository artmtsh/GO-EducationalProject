package todo

import (
	"fmt"
	"github.com/k0kubun/pp"
	"strings"
	"todolistapp/logger"
)

type Storage struct {
	storage map[string]*Task
}

func (storage *Storage) CreateStorage() *Storage {
	storage.storage = make(map[string]*Task)
	return storage
}

func (storage *Storage) CreateAndAddTask(logger *logger.Logger, text string) {
	fields := strings.Fields(text)
	newTask := CreateTask(strings.Join(fields[0:], " "))
	_, ok := storage.storage[newTask.title]
	if ok {
		fmt.Println(taskAlreadyExist)
		return
	}
	storage.storage[newTask.title] = &newTask
	logger.AddEvent("Пользователь создал задачу с заголовком " + newTask.title)
}

func (storage *Storage) DeleteTask(logger *logger.Logger, title string) {
	_, ok := storage.storage[title]
	if !ok {
		fmt.Println(taskDoNotExist, title)
		logger.AddEvent("Пользователю не удалось удалить задачу с заголовком" + title + " так как она не существует")
		return
	}
	delete(storage.storage, title)
	logger.AddEvent("Пользователь удалил задачу с заголовком" + title)
}

func (storage *Storage) ListTaskByTitle(logger *logger.Logger, title string) {
	if task, ok := storage.storage[title]; !ok {
		fmt.Println(taskDoNotExist, title)
		logger.AddEvent("Пользователю не удалось вывести задачу с заголовком" + title + " так как она не существует")
	} else {
		task.PrintTask()
		logger.AddEvent("Пользователь задачу с заголовком " + title)
	}
}

func (storage *Storage) ListAllTask(logger *logger.Logger) {
	if len(storage.storage) == 0 {
		fmt.Println(emptyList)
		logger.AddEvent("Пользователю не удалось вывести задачу так как список задач пуст")
		return
	}
	for _, value := range storage.storage {
		value.PrintTask()
	}
	logger.AddEvent("Пользователь вывел полный список задач")
}

func (storage *Storage) GetTaskByTitle(title string) *Task {
	return storage.storage[title]
}

func (storage *Storage) CompleteTask(logger *logger.Logger, title string) {
	if _, ok := storage.storage[title]; ok {
		storage.storage[title].CompleteTask()
	} else {
		pp.Println("Такой задачи не существует")
		logger.AddEvent("Не удалось завершить задачу с заголовком " + title + " так как такой задачи не существует")
	}
}
