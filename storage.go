package main

import (
	firebase "firebase.google.com/go/v4"
	"google.golang.org/api/option"
)

func uploadWorld(name string, data []byte) error {

	storageBucket, err := getEnv("ST_BUCKET")

	if err != nil {
		return err
	}

	var config = &firebase.Config{
		StorageBucket: storageBucket,
	}

	var opt = option.WithCredentialsJSON(credentialsFile)
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

	wc.ContentType = "application/x-zip-compressed"

	_, err = wc.Write(data)

	if err != nil {
		return err
	}

	if err = wc.Close(); err != nil {
		return err
	}

	return nil

}
