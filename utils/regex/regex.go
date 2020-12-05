package regex

import "regexp"

func Match(str string, regex string) bool {
	match, err := regexp.MatchString(regex, str)
	return match && err == nil
}
