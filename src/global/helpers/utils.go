package helpers

import "regexp"

func IsFiledsEmpty(fileds []string) bool {

	for _, f := range fileds {
		if f == "" {
			return true
		}
	}

	return false
}

func IsAEmail(v string) bool {
	r := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z.\\-]+\\.[a-z]{2,4}$")
	return r.MatchString(v)
}
