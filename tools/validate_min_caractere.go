package tools

import "regexp"

func (t *Tools) ValidateMinCaractere(pass []string, minSize int) {

	sum := 0

	for _, caractere := range pass {
		ok := specials[caractere]
		if ok {
			sum++
		}

		r := regexp.MustCompile("[A-ZÀ-Ýa-zà-ÿ]")
		exists := r.MatchString(caractere)
		if exists {
			sum++
		}
	}

}
