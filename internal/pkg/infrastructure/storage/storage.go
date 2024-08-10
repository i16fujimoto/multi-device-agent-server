package storage

import (
	"context"
	"errors"
	"io"

	"cloud.google.com/go/storage"

	"github.com/multi-device-agent-server/internal/pkg/cerror"
	storagegw "github.com/multi-device-agent-server/internal/pkg/gateway/storage"
)

type storageClient struct{}

func New() storagegw.StorageClient {
	return &storageClient{}
}

func (s *storageClient) Find(ctx context.Context, bucket, objectPath string) ([]byte, error) {
	cli, err := s.newClient(ctx)
	if err != nil {
		return nil, cerror.Wrap(err, "storage")
	}
	defer cli.Close()

	obj := cli.Bucket(bucket).Object(objectPath)
	reader, err := obj.NewReader(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return nil, cerror.Wrap(
				err,
				"object is not found",
				cerror.WithNotFoundCode(),
			)
		}

		return nil, cerror.Wrap(
			err,
			"failed to get object reader",
			cerror.WithStorageAPICode(),
		)
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to read object",
			cerror.WithStorageAPICode(),
		)
	}

	return body, nil
}

func (s *storageClient) Save(ctx context.Context, bucket, objectPath string, body []byte) error {
	cli, err := s.newClient(ctx)
	if err != nil {
		return cerror.Wrap(err, "storage")
	}
	defer cli.Close()

	writer := cli.Bucket(bucket).Object(objectPath).NewWriter(ctx)
	defer writer.Close()
	if _, err = writer.Write(body); err != nil {
		return cerror.Wrap(
			err,
			"failed to write object",
			cerror.WithStorageAPICode(),
		)
	}

	return nil
}

func (s *storageClient) Exist(ctx context.Context, bucket, objectPath string) (bool, error) {
	cli, err := s.newClient(ctx)
	if err != nil {
		return false, cerror.Wrap(err, "storage")
	}
	defer cli.Close()

	// See https://pkg.go.dev/cloud.google.com/go/storage#example-ObjectHandle-Exists
	_, err = cli.Bucket(bucket).Object(objectPath).Attrs(ctx)
	if err != nil {
		if errors.Is(err, storage.ErrObjectNotExist) {
			return false, nil
		}

		return false, cerror.Wrap(
			err,
			"failed to get object attrs",
			cerror.WithStorageAPICode(),
		)
	}

	return true, nil
}

func (s *storageClient) newClient(ctx context.Context) (*storage.Client, error) {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to new storage client",
			cerror.WithStorageAPICode(),
		)
	}

	return client, nil
}
