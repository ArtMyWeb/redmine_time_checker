package main

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron/v3"
	"log"
	"os"
	"os/signal"
	"redmine_time_checker/config"
	"redmine_time_checker/helpers"
	"redmine_time_checker/redmine"
	"redmine_time_checker/telegram"
	"time"
)

func main() {

	//run cron job daily at 22:00 Kiev time
	job := cron.New(cron.WithLogger(cron.VerbosePrintfLogger(log.New(os.Stdout, "cron: ", log.LstdFlags))))
	job.AddFunc("CRON_TZ=Europe/Kiev * 22 * * ?", run)
	job.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func run() {
	c := config.NewConfig()
	bot, err := tgbotapi.NewBotAPI(c.BotToken)
	helpers.CheckError("Create Telegram Bot: ", err)
	bot.Debug = false

	start := time.Now().Format("2006-01-02")
	end := time.Now().Format("2006-01-02")

	//run for every manager from the list
	for _, m := range c.UserList {
		//send report by every manager's developer
		for _, devID := range m.DevelopersIDs {
			data, err := redmine.GetTimeEntriesByUserID(devID, start, end)
			helpers.CheckError("Get Redmine data", err)

			//create message for telegram
			var message string
			if len(data.TimeEntry) > 0 {
				message = helpers.FullActivityMessage(data.TimeEntry)
			} else {
				user, err := redmine.GetUsersByID(devID)
				helpers.CheckError("Get Redmine User data", err)

				message = helpers.EmptyActivityMessage(user)
			}

			if message != "" {
				//send message
				err := telegram.SendMessage(bot, m.ManagerChannelID, message)
				helpers.CheckError("Send Message", err)
			}
		}
	}
}
