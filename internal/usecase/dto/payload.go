package dto

type Payload struct {
	Message  string `json:"message"`
	Username string `json:"username"`
	Type     string `json:"type"`
}
