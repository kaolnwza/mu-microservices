package service

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"net/url"
	"os"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
	"github.com/kaolnwza/muniverse/storage/entity"
	"google.golang.org/api/option"
)

var (
	storageClient *storage.Client
)

func GCSUploadFile(ctx context.Context, userUUID uuid.UUID, file multipart.File, bucket string) (*entity.Upload, error) {
	var err error
	cred := os.Getenv("DEPLOY_GOOGLE_APPLICATION_CREDENTIALS")
	if cred == "" {
		cred = os.Getenv("GOOGLE_APPLICATION_CREDENTIALS")
	}

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile(cred))
	if err != nil {
		return nil, err
	}

	objectLocation := fmt.Sprintf(`image/%s`, uuid.New())

	sw := storageClient.Bucket(bucket).Object(objectLocation).NewWriter(ctx)

	if _, err := io.Copy(sw, file); err != nil {
		return nil, err
	}

	if err := sw.Close(); err != nil {
		return nil, err
	}

	u, err := url.Parse(sw.Attrs().Name)
	if err != nil {
		return nil, err
	}

	var upload entity.Upload
	upload.Bucket = bucket
	upload.Path = u.EscapedPath()
	upload.UserUUID = userUUID

	// if err := repository.UploadImage(ctx, &upload); err != nil {
	// 	return nil, err
	// }

	return &upload, nil
}
