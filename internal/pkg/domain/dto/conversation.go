package dto

import (
	"encoding/json"

	"github.com/multi-device-agent-server/internal/pkg/cerror"
)

type Conversations []*Conversation

type Conversation struct {
	Time   int64   `json:"time"`
	User   string  `json:"user"`
	Agent  string  `json:"agent"`
	Device *string `json:"device"`
}

func Unmarshal(body []byte) (Conversations, error) {
	var conversations Conversations
	err := json.Unmarshal(body, &conversations)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to unmarshal conversations",
			cerror.WithInternalCode(),
		)
	}

	return conversations, nil
}

func (c Conversations) Marshal(conversations Conversations) ([]byte, error) {
	body, err := json.Marshal(conversations)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to marshal conversations",
			cerror.WithInternalCode(),
		)
	}

	return body, nil
}
