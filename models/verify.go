package models

type Verify struct {
	Password string  `json:"password"`
	Rules    []Rules `json:"rules"`
}

type Rules struct {
	TypeRule string `json:"rule"`
	Value    int    `json:"value"`
}
