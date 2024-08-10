package ui

import (
	"net/http"

	echo "github.com/labstack/echo/v4"

	"github.com/multi-device-agent-server/internal/pkg/cerror"
	"github.com/multi-device-agent-server/internal/pkg/domain/entity"
)

type saveConversationRequest struct {
	Device        string                `json:"device" validate:"required"`
	Conversations []*entity.Conversation `json:"conversations" validate:"required"`
}

func (h *handler) SaveConversation(c echo.Context) error {
	// ctx := c.Request().Context()

	request := new(saveConversationRequest)
	if err := c.Bind(request); err != nil {
		return cerror.Wrap(err, "ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Failed to bind"),
		)
	}

	if err := c.Validate(request); err != nil {
		return cerror.Wrap(err, "ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Failed to validate"),
		)
	}

	return c.NoContent(http.StatusNoContent)
}

type listDailyConversationsResponse struct {
	Conversations []*entity.Conversation `json:"conversations" validate:"required"`
}

func (h *handler) ListDailyConversations(c echo.Context) error {
	// ctx := c.Request().Context()

	return c.JSON(http.StatusOK, &listDailyConversationsResponse{})
}
