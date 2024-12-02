package services

import (
	"context"
	"fmt"
	"io"
	"time"

	"simple-image-gallery/config"
	"simple-image-gallery/utils"
)

func GenerateAndUploadImage(width, height int, text string) (string, error) {
	// Generate image
	img, err := utils.CreateImage(width, height, text)
	if err != nil {
		return "", err
	}

	// Initialize Firebase
	app, err := config.InitFirebase()
	if err != nil {
		return "", err
	}

	// Get Storage client
	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", err
	}

	// Generate unique filename
	filename := fmt.Sprintf("images/%d.png", time.Now().UnixNano())
	obj := bucket.Object(filename)

	// Upload image
	writer := obj.NewWriter(ctx)
	if err := utils.SaveImageToPNG(img, writer); err != nil {
		return "", err
	}
	writer.Close()

	// Get public URL
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return "", err
	}

	return attrs.MediaLink, nil
}

func GetImageContents(imageId string) ([]byte, string, error) {
	// Initialize Firebase
	app, err := config.InitFirebase()
	if err != nil {
		return nil, "", err
	}

	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, "", err
	}

	obj := bucket.Object(fmt.Sprintf("images/%s.png", imageId))

	// Get the object reader
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()

	// Read all contents
	contents, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	// Return the contents and the content type
	return contents, "image/png", nil
}

func GetImageURL(imageId string) (string, error) {
	// Initialize Firebase
	app, err := config.InitFirebase()
	if err != nil {
		return "", err
	}

	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		return "", err
	}
	bucket, err := client.DefaultBucket()
	if err != nil {
		return "", err
	}

	obj := bucket.Object(fmt.Sprintf("images/%s.png", imageId))
	attrs, err := obj.Attrs(ctx)
	if err != nil {
		return "", err
	}

	return attrs.MediaLink, nil
}

func GetImageStream(imageId string) (io.ReadCloser, string, error) {
	app, err := config.InitFirebase()
	if err != nil {
		return nil, "", err
	}

	ctx := context.Background()
	client, err := app.Storage(ctx)
	if err != nil {
		return nil, "", err
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, "", err
	}

	obj := bucket.Object(fmt.Sprintf("images/%s.png", imageId))

	// Get the object reader
	reader, err := obj.NewReader(ctx)
	if err != nil {
		return nil, "", err
	}

	return reader, "image/png", nil
}
