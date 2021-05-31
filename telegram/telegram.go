package telegram

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

type (

	// TgMessage struct for Full activity message
	TgMessage struct {
		Activity []Activity `json:"activity"`
		User     string     `json:"user"`
		Total    float32    `json:"total"`
	}

	// Activity struct for activity list
	Activity struct {
		Activity    string  `json:"activity"`
		ProjectID   int     `json:"project_id"`
		ProjectName string  `json:"project_name"`
		IssueID     int     `json:"issue_id"`
		ProjectLink string  `json:"project_link"`
		IssueLink   string  `json:"issue_link"`
		Comment     string  `json:"comment"`
		SpentHours  float32 `json:"spent_hours"`
	}
)

// SendMessage send message to telegram channel
// return error
func SendMessage(bot *tgbotapi.BotAPI, tgChannel int64, message string) error {
	msg := tgbotapi.NewMessage(tgChannel, message)
	msg.ParseMode = tgbotapi.ModeHTML

	_, err := bot.Send(msg)
	return err
}
