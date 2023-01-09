package tools

import "regexp"

func (t *Tools) ValidateMinCaseOrDigit(pass []string, minUppercase int, compile string) bool {

	sum := 0

	for _, caractere := range pass {
		r := regexp.MustCompile(compile)
		exists := r.MatchString(caractere)
		if exists {
			sum++
		}
	}

	return sum >= minUppercase
}
