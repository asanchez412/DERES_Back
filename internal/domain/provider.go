package domain

type Provider struct {
	Name  string  `json:"name"`
	RUT   string  `json:"rut"`
	Type  string  `json:"type"`
	Score float32 `json:"score"`
}
