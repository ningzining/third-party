package oss

import (
	"fmt"
	"os"
	"testing"
)

func TestUpload(t *testing.T) {
	client, err := New()
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	file, err := os.Open("hello.jpg")
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	upload, err := client.Upload(file, "temp", file.Name())
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("%s\n", upload)
}
