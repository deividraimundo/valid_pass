package services

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/deividraimundo/valid_pass/models"
)

func (s *Service) Verify(w http.ResponseWriter, r *http.Request) {

	verify := models.Verify{}
	var validated []string

	if err := json.NewDecoder(r.Body).Decode(&verify); err != nil {
		s.t.Response(w, http.StatusInternalServerError, `{"message": "Algo deu errado!"}`)
		return
	}

	pass := strings.Split(verify.Password, "")

	for _, rules := range verify.Rules {

		switch rules.TypeRule {
		// case "minSize":
		// 	isValid := s.t.ValidMinCase(pass, rules.Value, "")
		// 	if !isValid {
		// 		validated = append(validated, rules.TypeRule)
		// 	}

		case "minUppercase":
			isValid := s.t.ValidateMinCaseOrDigit(pass, rules.Value, "[A-ZÀ-Ý]")
			if !isValid {
				validated = append(validated, rules.TypeRule)
			}

		case "minLowercase":
			isValid := s.t.ValidateMinCaseOrDigit(pass, rules.Value, "[a-zà-ÿ]")
			if !isValid {
				validated = append(validated, rules.TypeRule)
			}

		case "minDigit":
			isValid := s.t.ValidateMinCaseOrDigit(pass, rules.Value, "[0-9]")
			if !isValid {
				validated = append(validated, rules.TypeRule)
			}

		case "minSpecialChars":
			isValid := s.t.ValidateSpecialCaractere(pass, rules.Value)
			if !isValid {
				validated = append(validated, rules.TypeRule)
			}

		case "noRepeted":
			isValid := s.t.ValidateNoRepeted(pass)
			if !isValid {
				validated = append(validated, rules.TypeRule)
			}
		}
	}

	body := models.Response{
		Verify: true,
	}

	if validated != nil {
		body.Verify = false
		body.NoMatch = validated
	}

	s.t.Response(w, http.StatusOK, body)
}
