package funcs

import (
	"fmt"
	"time"
)

func Help() {
	fmt.Println("- Список команд, которые должны быть доступны в приложении:\n" +
		"- help — эта команда позволяет узнать доступные команды и их формат\n" +
		"- add {заголовок задачи из одного слова} {текст задачи из одного или нескольких слов} — эта команда позволяет добавлять новые задачи в список задач\n" +
		"- list — эта команда позволяет получить полный список всех задач\n" +
		"- del {заголовок существующей задачи} — эта команда позволяет удалить задачу по её заголовку\n" +
		"- done {заголовок существующей задачи} — эта команда позволяет отменить задачу как выполненную\n" +
		"- events — эта команда позволяет получить список всех событий\n" +
		"- exit — эта команда позволяет завершить выполнение программы")
}

type Event struct {
	Text        string
	Description string
	EventTime   string
}

func createLog(text string) Event {
	return Event{
		Text:        text,
		Description: "",
		EventTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

type Logger struct {
	Events []Event
}

func (logger *Logger) AddEvent(text string) {
	logger.Events = append(logger.Events, createLog(text))
}

func (logger *Logger) PrintAllLogs() {
	for value := range logger.Events {
		fmt.Println("текст:", logger.Events[value].Text)
		fmt.Println("Описание:", logger.Events[value].Description)
		fmt.Println("Время события:", logger.Events[value].EventTime, "\n")
	}
}
