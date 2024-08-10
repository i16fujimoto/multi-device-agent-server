package dto

type Conversations []*Conversation

type Conversation struct {
	Time   int64   `json:"time"`
	User   string  `json:"user"`
	Agent  string  `json:"agent"`
	Device *string `json:"device"`
}
