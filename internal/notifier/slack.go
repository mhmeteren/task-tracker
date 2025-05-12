package notifier

import (
	"bytes"
	"encoding/json"
	"net/http"
	"task-tracker/internal/logger"
)

type SlackNotifier struct{}

func (s SlackNotifier) Send(n Notification) error {
	webhookURL := "https://hooks.slack.com/services/" + n.Token

	payload := map[string]string{"text": n.Message}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	resp, err := http.Post(webhookURL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		logger.GlobalLogger.Error("Slack notifier erorr", &logger.LogFields{"tags": []string{"notification", "slack"}, "status": resp.Status})
	}

	return nil
}
