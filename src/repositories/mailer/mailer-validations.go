package mailer

import "regexp"

func (mail *Mail) IsFiledsEmpty() bool {
	if mail.To == "" {
		return true
	}

	if mail.Subject == "" {
		return true
	}

	if mail.Message == "" {
		return true
	}

	if mail.DstName == "" {
		return true
	}

	return false
}

func IsToAEmail(to string) bool {
	r := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z.\\-]+\\.[a-z]{2,4}$")
	return r.MatchString(to)
}
