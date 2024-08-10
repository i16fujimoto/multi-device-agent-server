//go:generate mockgen -destination=./mock/$GOFILE -package=mock_$GOPACKAGE -source=$GOFILE
package storage

import "context"

// StorageClient ストレージへのアクセスinterface
type StorageClient interface {
	// Find オブジェクトを取得する
	Find(ctx context.Context, bucket, objectPath string) ([]byte, error)
	// Save オブジェクトを保存する
	Save(ctx context.Context, bucket, objectPath string, body []byte) error
	// Exist オブジェクトが存在するか確認する
	Exist(ctx context.Context, bucket, objectPath string) (bool, error)
}
