package tools

import "regexp"

func Check_–°ompromising(data string) bool {
	ok, _ := regexp.MatchString("[^0-9a-zA-Z]", data)
	return !ok
}
