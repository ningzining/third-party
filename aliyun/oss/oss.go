package oss

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"os"
	"path"
	"time"
)

var (
	endpoint   = os.Getenv("OSS_ENDPOINT")
	accessID   = os.Getenv("OSS_ACCESS_KEY_ID")
	accessKey  = os.Getenv("OSS_ACCESS_KEY_SECRET")
	bucketName = os.Getenv("OSS_BUCKET_NAME")
)

type Client struct {
	OSS        *oss.Client
	DomainName string
}

func New() (*Client, error) {
	initDefaultConfig()
	client, err := oss.New(endpoint, accessID, accessKey)
	if err != nil {
		return nil, err
	}

	return &Client{
		OSS:        client,
		DomainName: "https://entangled-cotton.oss-cn-hangzhou.aliyuncs.com",
	}, nil
}

func initDefaultConfig() {
	if endpoint == "" {
		endpoint = "oss-cn-hangzhou.aliyuncs.com"
	}
	if accessID == "" {
		accessID = ""
	}
	if accessKey == "" {
		accessKey = ""
	}
	if bucketName == "" {
		bucketName = ""
	}
}

func (c *Client) Upload(rd io.Reader, folder string, filename string) (string, error) {
	bucket, err := c.OSS.Bucket(bucketName)
	if err != nil {
		return "", err
	}
	filepath := path.Join(folder, time.Now().Local().Format(time.DateOnly), filename)

	if err := bucket.PutObject(filepath, rd); err != nil {
		return "", err
	}

	return path.Join(c.DomainName, filepath), nil
}
