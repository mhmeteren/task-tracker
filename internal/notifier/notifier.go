package notifier

import "log"

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
				log.Printf("[NOTIFICATION] message: %v", err.Error())
			}
		}
	}()
}
