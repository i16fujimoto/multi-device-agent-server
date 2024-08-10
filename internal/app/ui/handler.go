package ui

import (
	"time"

	echo "github.com/labstack/echo/v4"

	"github.com/multi-device-agent-server/internal/app/usecase"
	storagegw "github.com/multi-device-agent-server/internal/pkg/gateway/storage"
	"github.com/multi-device-agent-server/internal/pkg/logger"
)

type Handler interface {
	GetHealth(c echo.Context) error

	SaveConversation(c echo.Context) error
	ListDailyConversations(c echo.Context) error
}

type handler struct {
	conversationUC usecase.Conversation
}

func NewHandler(storageClient storagegw.StorageClient, logger *logger.Logger) Handler {
	return &handler{
		conversationUC: usecase.NewConversationUC(storageClient, time.Now, logger),
	}
}
