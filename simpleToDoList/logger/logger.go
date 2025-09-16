package logger

import (
	"fmt"
	"time"
)

func createLog(description string) Event {
	return Event{
		Text:        "",
		Description: description,
		EventTime:   time.Now().Format("2006-01-02 15:04:05"),
	}
}

type Logger struct {
	Events []Event
}

func AddTextToLog(log *Event, text string) {
	log.Text = text
}

func (logger *Logger) CreateLogger() *Logger {
	logger.Events = make([]Event, 0)
	return logger
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

func PrintAllLogs(logger *Logger) {
	logger.AddEvent("Пользователь вывел список событий")
	AddTextToLog(&logger.Events[len(logger.Events)-1], "events")
	for value := range logger.Events {
		fmt.Println("текст:", logger.Events[value].Text)
		fmt.Println("Описание:", logger.Events[value].Description)
		fmt.Println("Время события:", logger.Events[value].EventTime, "\n")
	}
}
