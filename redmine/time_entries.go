package redmine

import (
	"github.com/parnurzeal/gorequest"
	"strconv"
)

// GetTimeEntriesByUserID return redmine time entries by userID and dates
func GetTimeEntriesByUserID(
	userID int,
	startDate, endDate string,
) (
	TEResult,
	error,
) {
	var result TEResult
	request := gorequest.New()
	path := getPath("/time_entries.json")
	_, _, errs := request.Get(path).
		Query("user_id=" + strconv.Itoa(userID)).
		Query("key=" + c.RedmineAPIKey).
		Query("from=" + startDate).
		Query("to=" + endDate).
		EndStruct(&result)
	if errs != nil {
		return TEResult{}, errs[0]
	}
	return result, nil
}
