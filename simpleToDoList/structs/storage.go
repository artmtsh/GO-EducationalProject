package structs

import (
	"fmt"
)

type Storage struct {
	storage map[string]*Task
}

func (storage *Storage) CreateStorage() {
	storage.storage = make(map[string]*Task)
}

func (storage *Storage) AddTask(task *Task) {
	storage.storage[task.title] = task
}

func (storage *Storage) DeleteTask(title string) {
	_, ok := storage.storage[title]
	if !ok {
		fmt.Println("Задачи с заголовком", title, "не существует")
		return
	}
	delete(storage.storage, title)
}

func (storage *Storage) ListTaskByTitle(title string) {
	storage.storage[title].PrintTask()
}

func (storage *Storage) ListAllTask() {
	if len(storage.storage) == 0 {
		fmt.Println("Список пуст")
		return
	}
	for _, value := range storage.storage {
		value.PrintTask()
	}
}

func (storage *Storage) GetTaskByTitle(title string) *Task {
	return storage.storage[title]
}

func (storage *Storage) CompleteTask(title string) {
	storage.GetTaskByTitle(title).CompleteTask()
}
