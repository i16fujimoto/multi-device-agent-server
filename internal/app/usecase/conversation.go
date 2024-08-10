package usecase

import (
	"context"
	"time"

	"github.com/multi-device-agent-server/internal/pkg/domain/entity"
	storagegw "github.com/multi-device-agent-server/internal/pkg/gateway/storage"
	"github.com/multi-device-agent-server/internal/pkg/logger"
)

type Conversation interface {
	Save(context.Context, []*entity.Conversation) error
	ListDaily(context.Context) ([]*entity.Conversation, error)
}

type conversationUC struct {
	storageClient storagegw.StorageClient
	timeNow       func() time.Time
	logger        *logger.Logger
}

func NewConversationUC(storageClient storagegw.StorageClient, timeNow func() time.Time, logger *logger.Logger) Conversation {
	return &conversationUC{
		storageClient: storageClient,
		timeNow:       timeNow,
		logger:        logger,
	}
}

func (uc *conversationUC) Save(ctx context.Context, conversations []*entity.Conversation) error {
	return nil
}

func (uc *conversationUC) ListDaily(ctx context.Context) ([]*entity.Conversation, error) {
	return nil, nil
}
