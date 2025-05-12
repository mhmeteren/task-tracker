package notifier

import (
	"net/http"
	"net/url"
	"task-tracker/internal/logger"
)

type TelegramNotifier struct{}

func (t TelegramNotifier) Send(n Notification) error {
	apiURL := "https://api.telegram.org/bot" + n.Token + "/sendMessage"

	data := url.Values{}
	data.Set("chat_id", n.Recipient)
	data.Set("text", n.Message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		logger.GlobalLogger.Error("Telegram notifier erorr", &logger.LogFields{"tags": []string{"notification", "telegram"}, "status": resp.Status})
	}

	return nil
}
