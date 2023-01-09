package tools

var specials = map[string]bool{
	"!": true,
	"@": true,
	"#": true,
	"$": true,
	"%": true,
	"^": true,
	"&": true,
	"*": true,
	"(": true,
	")": true,
	"-": true,
	"+": true,
	`\`: true,
	"/": true,
	"{": true,
	"}": true,
	"[": true,
	"]": true,
}

func (t *Tools) ValidateSpecialCaractere(pass []string, minSpecialChars int) bool {

	sum := 0

	for _, caractere := range pass {
		exists := specials[caractere]
		if exists {
			sum++
		}
	}

	return sum >= minSpecialChars
}
