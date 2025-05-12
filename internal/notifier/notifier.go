package notifier

import (
	"task-tracker/internal/logger"
)

type Notification struct {
	Message   string
	Recipient string
	Token     string
	Service   string
}

type Notifier interface {
	Send(n Notification) error
}

var Queue chan Notification

var services = map[string]Notifier{}

func NotifyInit() {
	Register("telegram", TelegramNotifier{})
	Register("slack", SlackNotifier{})

	StartWorker()
}

func Register(serviceName string, impl Notifier) {
	services[serviceName] = impl
}

func Enqueue(n Notification) {
	Queue <- n
}

func StartWorker() {
	Queue = make(chan Notification, 100)
	go func() {
		for msg := range Queue {
			service, ok := services[msg.Service]
			if !ok {
				continue
			}
			err := service.Send(msg)
			if err != nil {
				logger.GlobalLogger.Error("Notifier erorr", &logger.LogFields{"tags": []string{"notification"}, "error": err.Error()})
			}
		}
	}()
}
