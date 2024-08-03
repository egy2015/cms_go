package storage

import (
	"context"
	"io"
	"os"
	"strconv"
	"time"

	"cloud.google.com/go/storage"
)

var (
	bucketHandle *storage.BucketHandle
	URL = "https://storage.googleapis.com/example/"
)

func newBucketHandle(ctx context.Context) error {
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	bucketHandle = client.Bucket("example")
	return nil
}

func Upload(dec []byte, name string) (err error) {
	ctx := context.Background()
	if err = newBucketHandle(ctx); err != nil {
		return
	}
	f, err := os.Create("upload/" + name)
	if err != nil {
		return
	}
	if _, err = f.Write(dec); err != nil {
		return
	}
	writer := bucketHandle.Object(name).NewWriter(ctx)
	file, err := os.Open("upload/" + name)
	if err != nil {
		return
	}
	if _, err = io.Copy(writer, file); err != nil {
		return
	}
	return writer.Close()
}

func UnixNano() (unixNanoString string) {
	return strconv.FormatInt(time.Now().UnixNano(), 10)
}
