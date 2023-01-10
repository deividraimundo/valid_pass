package tools

import "regexp"

// ValidateForRegex valida se o número mínimo tanto de caractere quanto de digítos estão dentro do minimo exigido.
func (t *Tools) ValidateForRegex(pass []string, min int, compile string) bool {

	sum := 0

	for _, caractere := range pass {
		r := regexp.MustCompile(compile)
		exists := r.MatchString(caractere)
		if exists {
			sum++
		}
	}

	return sum >= min
}
