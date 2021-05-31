package redmine

import (
	"redmine_time_checker/config"
	"strconv"
)

var c *config.Config

type (

	// TEResult struct contains Redmine Time Entry Result data
	TEResult struct {
		TimeEntry []TimeEntry `json:"time_entries"`
	}

	// UserResult struct contains user data
	UserResult struct {
		User User `json:"user"`
	}

	// TimeEntry struct contains Redmine Time Entry data
	TimeEntry struct {
		ID       int      `json:"id"`
		Hours    float32  `json:"hours"`
		Comments string   `json:"comments"`
		SpentOn  string   `json:"spent_on"`
		Project  Project  `json:"project"`
		Issue    Issue    `json:"issue"`
		User     User     `json:"user"`
		Activity Activity `json:"activity"`
	}

	// Project struct contains Redmine Time Entry project data
	Project struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}

	// Issue struct contains Redmine Time Entry issue data
	Issue struct {
		ID int `json:"id"`
	}

	// User struct contains Redmine Time Entry user data
	User struct {
		ID        int    `json:"id"`
		Name      string `json:"name"`
		FirstName string `json:"firstname"`
		LastName  string `json:"lastname"`
	}

	// Activity struct contains Redmine Time Entry activity data
	Activity struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	}
)

func init() {
	c = config.NewConfig()
}

// getPath return path based on REDMINE_URL and link
func getPath(link string) string {
	return c.RedmineURL + link
}

// GetLinkToIssue return link to issue
func (self *TimeEntry) GetLinkToIssue() string {
	return getPath("/issue/" + strconv.Itoa(self.Issue.ID))
}

// GetLinkToProject return link to project
func (self *TimeEntry) GetLinkToProject() string {
	return getPath("/project/" + strconv.Itoa(self.Project.ID))
}
