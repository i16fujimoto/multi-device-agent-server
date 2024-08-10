package entity

type Conversation struct {
	Time   int64   `json:"time" validate:"required"`
	User   string  `json:"user" validate:"required"`
	Agent  string  `json:"agent" validate:"required"`
	Device *string `json:"device"`
}
