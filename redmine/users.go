package redmine

import (
	"github.com/parnurzeal/gorequest"
	"strconv"
)

// GetUsersByID return redmine user by ID
func GetUsersByID(
	userID int,
) (
	User,
	error,
) {
	var user UserResult
	request := gorequest.New()
	path := getPath("/users/" + strconv.Itoa(userID) + ".json")
	_, _, errs := request.Get(path).
		Query("key=" + c.RedmineAPIKey).
		EndStruct(&user)
	if errs != nil {
		return User{}, errs[0]
	}
	return user.User, nil
}
