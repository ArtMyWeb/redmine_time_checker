package helpers

import (
	"bytes"
	"redmine_time_checker/redmine"
	"redmine_time_checker/telegram"
	"text/template"
)

// convertRedmineDataToTelegramMessage return time-entry template string based on redmine TimeEntry data
func convertRedmineDataToTelegramMessage(entries []redmine.TimeEntry) string {
	message := telegram.TgMessage{}
	var user string
	var total float32
	total = 0

	for _, entry := range entries {
		user = entry.User.Name
		total += entry.Hours

		tmp := telegram.Activity{
			Activity:    entry.Activity.Name,
			Comment:     entry.Comments,
			ProjectID:   entry.Project.ID,
			IssueID:     entry.Issue.ID,
			ProjectLink: entry.GetLinkToProject(),
			IssueLink:   entry.GetLinkToIssue(),
			ProjectName: entry.Project.Name,
			SpentHours:  entry.Hours,
		}
		message.Activity = append(message.Activity, tmp)
	}

	message.User = user
	message.Total = total

	return activityMessage(message)
}

// FullActivityMessage return time-entry template string
func FullActivityMessage(entries []redmine.TimeEntry) string {
	message := convertRedmineDataToTelegramMessage(entries)
	return message
}

// ActivityMessage return time-entry template string
func activityMessage(message telegram.TgMessage) string {
	return parseTemplate(message, "./telegram/templates/time-entries.html")
}

// EmptyActivityMessage return empty activity template string
func EmptyActivityMessage(user redmine.User) string {
	return parseTemplate(user, "./telegram/templates/empty-entries.html")
}

// parseTemplate return string by data and template path
func parseTemplate(data interface{}, path string) string {
	tmpl, err := template.ParseFiles(path)
	CheckError("Read template", err)

	buf := new(bytes.Buffer)
	err = tmpl.Execute(buf, data)
	CheckError("Execute template", err)

	return buf.String()
}
