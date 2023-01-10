package tools

import "regexp"

// ValidateMinCaractere valida se o número de caracteres está dentro do valor minimo.
// Obs: Essa função não considera dígitos, apenas caracteres e caracteres especiais.
func (t *Tools) ValidateMinCaractere(pass []string, minSize int) bool {

	sum := 0

	for _, caractere := range pass {
		// Verificando caracteres especiais
		ok := specials[caractere]
		if ok {
			sum++
			continue
		}

		// Verificando caracteres
		r := regexp.MustCompile("[A-ZÀ-Ýa-zà-ÿ]")
		exists := r.MatchString(caractere)
		if exists {
			sum++
			continue
		}
	}

	return sum >= minSize
}
