package notifier

import (
	"fmt"
	"net/http"
	"net/url"
)

type TelegramNotifier struct{}

func (t TelegramNotifier) Send(n Notification) error {
	apiURL := "https://api.telegram.org/bot" + n.Token + "/sendMessage"

	data := url.Values{}
	data.Set("chat_id", n.ChatID)
	data.Set("text", n.Message)

	resp, err := http.PostForm(apiURL, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		return fmt.Errorf("telegram error: %s", resp.Status)
	}

	return nil
}
