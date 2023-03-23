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

func (mail *Mail) IsToAEmail() bool {
	r := regexp.MustCompile("^[a-z0-9._%+\\-]+@[a-z.\\-]+\\.[a-z]{2,4}$")
	println(mail.To, "is valid", r.MatchString(mail.To))
	return r.MatchString(mail.To)
}
