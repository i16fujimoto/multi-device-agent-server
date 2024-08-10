package storage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/multi-device-agent-server/internal/pkg/cerror"
)

type FileStorageClient struct {
	// BasePath ストレージの親になるパス
	BasePath string
}

func NewFileStorage(basePath string) *FileStorageClient {
	return &FileStorageClient{
		BasePath: basePath,
	}
}

// Find オブジェクトを取得する
func (s *FileStorageClient) Find(ctx context.Context, bucket, objectPath string) ([]byte, error) {
	path := fmt.Sprintf("%s/%s/%s", s.BasePath, bucket, objectPath)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, cerror.Wrap(
				err,
				"object is not found",
				cerror.WithNotFoundCode(),
			)
		}

		return nil, cerror.Wrap(
			err,
			"failed to get file",
			cerror.WithInternalCode(),
		)
	}

	f, err := os.Open(path)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to open file",
			cerror.WithInternalCode(),
		)
	}

	body, err := io.ReadAll(f)
	if err != nil {
		return nil, cerror.Wrap(
			err,
			"failed to read file",
			cerror.WithInternalCode(),
		)
	}

	return body, nil
}

// Save オブジェクトを保存する
func (s *FileStorageClient) Save(ctx context.Context, bucket, objectPath string, body []byte) error {
	path := fmt.Sprintf("%s/%s/%s", s.BasePath, bucket, objectPath)
	if err := os.MkdirAll(filepath.Dir(path), 0777); err != nil { //nolint:gomnd
		return cerror.Wrap(
			err,
			"failed to create directory",
			cerror.WithInternalCode(),
		)
	}

	f, err := os.Create(path)
	if err != nil {
		return cerror.Wrap(
			err,
			"failed to create file",
			cerror.WithInternalCode(),
		)
	}
	defer f.Close()

	if _, err = f.Write(body); err != nil {
		return cerror.Wrap(
			err,
			"failed to write file",
			cerror.WithInternalCode(),
		)
	}

	return nil
}

// Exist オブジェクトが存在するか確認する
func (s *FileStorageClient) Exist(ctx context.Context, bucket, objectPath string) (bool, error) {
	path := fmt.Sprintf("%s/%s/%s", s.BasePath, bucket, objectPath)
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		}

		return false, cerror.Wrap(
			err,
			"failed to get file",
			cerror.WithInternalCode(),
		)
	}

	return true, nil
}
