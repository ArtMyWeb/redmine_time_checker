package config

import (
	"encoding/json"
	"github.com/caarlos0/env"
	"strconv"
)

var Configuration *Config

type (

	// Config structure for env settings
	Config struct {
		StrUserList   string `env:"USERS_LIST"`
		BotToken      string `env:"BOT_TOKEN"`
		UserList      []List
		RedmineURL    string `env:"REDMINE_URL"`
		RedmineAPIKey string `env:"REDMINE_API_KEY"`
	}

	// List structure where we store manager's chat id and list of developers
	List struct {
		ManagerChannelID int64 `json:"manager_channel_id"`
		DevelopersIDs    []int `json:"developers_ids"`
	}
)

func init() {
	Configuration = NewConfig()
}

// NewConfig load config from env
func NewConfig() *Config {
	c := Config{}
	env.Parse(&c)

	//unquote user list because in env it's stored already
	//like "[{\"manager_channel_id\": __manager_chanel_id__,\"developers_ids\":[104,229,158]}]"
	//but for structure we need to remove \"
	str, _ := strconv.Unquote(c.StrUserList)
	json.Unmarshal([]byte(str), &c.UserList)

	return &c
}
