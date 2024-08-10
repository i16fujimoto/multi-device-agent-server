package ui

import (
	"net/http"
	"slices"

	echo "github.com/labstack/echo/v4"

	"github.com/multi-device-agent-server/internal/pkg/cerror"
	"github.com/multi-device-agent-server/internal/pkg/domain/entity"
	"github.com/multi-device-agent-server/internal/pkg/logger"
)

type saveConversationRequest struct {
	Device        string                 `json:"device" validate:"required"`
	Conversations []*entity.Conversation `json:"conversations" validate:"required"`
}

func (h *handler) SaveConversation(c echo.Context) error {
	ctx := c.Request().Context()

	request := new(saveConversationRequest)
	if err := c.Bind(request); err != nil {
		h.logger.Error("Failed to bind", logger.Fstring("package", "ui"), logger.Ferror(err))

		return cerror.Wrap(err, "ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Failed to bind"),
		)
	}

	if err := c.Validate(request); err != nil {
		h.logger.Error("Failed to validate", logger.Fstring("package", "ui"), logger.Ferror(err))

		return cerror.Wrap(err, "ui",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Failed to validate"),
		)
	}

	if err := validateDevice(request.Device); err != nil {
		return err
	}

	if err := h.conversationUC.Save(ctx, request.Device, request.Conversations); err != nil {
		return err
	}

	return c.NoContent(http.StatusNoContent)
}

func validateDevice(device string) error {
	if device == "" {
		return cerror.New(
			"device is empty",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Device is empty"),
		)
	}

	if !slices.Contains(entity.AllDevices(), device) {
		return cerror.New(
			"device is invalid",
			cerror.WithInvalidArgumentCode(),
			cerror.WithClientMsg("Device is invalid"),
		)
	}

	return nil
}

type listDailyConversationsResponse struct {
	Conversations []*entity.Conversation `json:"conversations" validate:"required"`
}

func (h *handler) ListDailyConversations(c echo.Context) error {
	ctx := c.Request().Context()

	conversations, err := h.conversationUC.ListDaily(ctx)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &listDailyConversationsResponse{
		Conversations: conversations,
	})
}
