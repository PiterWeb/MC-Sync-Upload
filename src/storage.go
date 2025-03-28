package src

import (
	"context"

	firebase "firebase.google.com/go/v4"
	"github.com/BurntSushi/toml"
	"google.golang.org/api/option"
)

type StorageConfig struct {
	StorageBucket string `toml:"ST_BUCKET"`
}

func UploadWorld(rawConfig []byte, name string, data []byte) error {

	ctx := context.Background()

	var storageConfig StorageConfig

	_, err := toml.Decode(string(rawConfig), &storageConfig)

	if err != nil {
		return err
	}

	var config = &firebase.Config{
		StorageBucket: storageConfig.StorageBucket,
	}

	var opt = option.WithCredentialsFile("./serviceAccount.json")
	firebaseApp, err := firebase.NewApp(ctx, config, opt)

	if err != nil {
		return err
	}

	client, err := firebaseApp.Storage(ctx)

	if err != nil {
		return err
	}

	bucket, err := client.DefaultBucket()

	if err != nil {
		return err
	}

	wc := bucket.Object(name).NewWriter(ctx)

	wc.ContentType = "application/octet-stream"

	_, err = wc.Write(data)

	if err != nil {
		return err
	}

	if err = wc.Close(); err != nil {
		return err
	}

	return nil

}
