package tools

func (t *Tools) ValidateNoRepeted(pass []string) bool {

	var letter string

	for _, caractere := range pass {
		if letter == caractere {
			return false
		}

		letter = caractere
	}

	return true
}
