package todo

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	title          string
	text           string
	creationTime   time.Time
	done           bool
	completionTime time.Time
}

func CreateTask(text string) Task {
	words := strings.Fields(text)
	newTask := Task{
		title:        words[0],
		text:         strings.Join(words[1:len(words)], " "),
		creationTime: time.Now(),
		done:         false,
	}
	return newTask
}

func (task *Task) CompleteTask() {
	if task.done {
		fmt.Println(taskAlreadyDone)
		return
	}
	task.done = true
	task.completionTime = time.Now()
}

func (task *Task) PrintTask() {
	fmt.Println("Заголовок:", task.title)
	fmt.Println("Текст:", task.text)
	fmt.Println("Время создания:", task.creationTime.Format("2006-01-02 15:04:05"))
	if task.done {
		fmt.Println("Задача выполнена")
		fmt.Println("Время выполнения:", task.completionTime.Format("2006-01-02 15:04:05"), "\n")
	} else {
		fmt.Println("Задача еще не выполнена\n")
	}
}
