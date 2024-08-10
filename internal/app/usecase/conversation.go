package usecase

import (
	"context"
	"time"

	"github.com/multi-device-agent-server/config"
	"github.com/multi-device-agent-server/internal/pkg/cerror"
	"github.com/multi-device-agent-server/internal/pkg/domain/dto"
	"github.com/multi-device-agent-server/internal/pkg/domain/entity"
	storagegw "github.com/multi-device-agent-server/internal/pkg/gateway/storage"
	"github.com/multi-device-agent-server/internal/pkg/logger"
)

type Conversation interface {
	Save(context.Context, string, []*entity.Conversation) error
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

func (u *conversationUC) Save(ctx context.Context, device string, conversations []*entity.Conversation) error {
	bucket := config.GetEnv().BucketName
	today := u.timeNow().Format("2006-01-02")
	objectPath := today + ".json"

	var entities []*entity.Conversation
	exist, err := u.storageClient.Exist(ctx, bucket, objectPath)
	if err != nil {
		u.logger.Error("Failed to check exist", logger.Fstring("package", "usecase"), logger.Ferror(err))

		return cerror.Wrap(err, "usecase")
	}
	if exist {
		existConversation, err := u.find(ctx, bucket, objectPath)
		if err != nil {
			u.logger.Error("Failed to find", logger.Fstring("package", "usecase"), logger.Ferror(err))

			return cerror.Wrap(err, "usecase")
		}

		entities = append(existConversation, conversations...)
	}

	entities = append(entities, conversations...)

	if err := u.save(ctx, device, bucket, objectPath, entities); err != nil {
		u.logger.Error("Failed to save", logger.Fstring("package", "usecase"), logger.Ferror(err))

		return cerror.Wrap(err, "usecase")
	}

	return nil
}

func (u *conversationUC) ListDaily(ctx context.Context) ([]*entity.Conversation, error) {
	bucket := config.GetEnv().BucketName
	today := u.timeNow().Format("2006-01-02")
	objectPath := today + ".json"

	conversations, err := u.find(ctx, bucket, objectPath)
	if err != nil {
		return nil, cerror.Wrap(err, "usecase")
	}

	return conversations, nil
}

func (u *conversationUC) save(ctx context.Context, device, bucket, objectPath string, conversations []*entity.Conversation) error {
	conversationsDTO := make(dto.Conversations, len(conversations))
	for i, c := range conversations {
		if c.Device != nil {
			conversationsDTO[i] = &dto.Conversation{
				Time:   c.Time,
				User:   c.User,
				Agent:  c.Agent,
				Device: c.Device,
			}
		} else {
			conversationsDTO[i] = &dto.Conversation{
				Time:   c.Time,
				User:   c.User,
				Agent:  c.Agent,
				Device: &device,
			}
		}
	}

	body, err := conversationsDTO.Marshal(conversationsDTO)
	if err != nil {
		return err
	}

	if err := u.storageClient.Save(ctx, bucket, objectPath, body); err != nil {
		return err
	}

	return nil
}

func (u *conversationUC) find(ctx context.Context, bucket, objectPath string) ([]*entity.Conversation, error) {
	file, err := u.storageClient.Find(ctx, bucket, objectPath)
	if err != nil {
		return nil, err
	}

	conversations, err := dto.Unmarshal(file)
	if err != nil {
		return nil, err
	}

	entities := make([]*entity.Conversation, len(conversations))
	for i, c := range conversations {
		entities[i] = &entity.Conversation{
			Time:   c.Time,
			User:   c.User,
			Agent:  c.Agent,
			Device: c.Device,
		}
	}

	return entities, nil
}
