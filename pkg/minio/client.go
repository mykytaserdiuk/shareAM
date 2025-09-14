package minio

import (
	"context"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/mykytaserdiuk/shaream/pkg/models"
	"io"
)

type MinioDB struct {
	client *minio.Client
	bucket string
}

func NewMinio(endpoint, accessKey, secretKey string, useSSL bool) (*MinioDB, error) {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, fmt.Errorf("не удалось создать клиента MinIO: %w", err)
	}
	return &MinioDB{client: client}, nil
}

func (m *MinioDB) SetupBucket(ctx context.Context, bucket string) error {
	m.bucket = bucket
	exists, err := m.BucketExists(ctx, bucket)
	if err != nil {
		return err
	}
	if !exists {
		err = m.MakeBucket(ctx, bucket)
		if err != nil {
			return err
		}
		err = m.client.SetBucketPolicy(ctx, bucket, `{
				  "Version": "2012-10-17",
				  "Statement": [
					{
					  "Action": ["s3:GetObject"],
					  "Effect": "Allow",
					  "Principal": {"AWS": ["*"]},
					  "Resource": ["arn:aws:s3:::mybucket/*"]
					}
				  ]
				}`)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *MinioDB) Upload(ctx context.Context, file *models.File, reader io.Reader) error {
	_, err := m.client.PutObject(ctx, m.bucket, file.Name, reader, file.Size, minio.PutObjectOptions{
		UserMetadata: map[string]string{"UserToken": *file.UserToken},
	})
	return err
}

func (m *MinioDB) GetFileLink(_ context.Context, file *models.File) string {
	url := fmt.Sprintf("http://%s/%s/%s", m.client.EndpointURL().Host, m.bucket, file.Name)
	return url
}

func (m *MinioDB) BucketExists(ctx context.Context, bucket string) (bool, error) {
	return m.client.BucketExists(ctx, bucket)
}

func (m *MinioDB) MakeBucket(ctx context.Context, bucket string) error {
	return m.client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{Region: "us-east-1"})
}
