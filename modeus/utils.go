package modeus

import "regexp"

func correctEmail(email string) string {
	if email == "" {
		return ""
	}
	if !regexp.MustCompile(`@edu\.narfu\.ru$`).MatchString(email) {
		email += "@edu.narfu.ru"
	}
	return email
}
