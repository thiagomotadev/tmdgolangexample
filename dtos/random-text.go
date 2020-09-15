package dtos

// RandomText ...
type RandomText struct {
	ID      uint   `json:"id"`
	Text    string `json:"text"`
	IsValid bool   `json:"isValid"`
}
