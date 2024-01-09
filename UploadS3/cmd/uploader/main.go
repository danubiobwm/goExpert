package main

import (
	"fmt"
	"io"
	"os"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	s3Client *s3.S3
	s3Bucket string
	wg       sync.WaitGroup
)

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
		Credentials: credentials.NewStaticCredentials(
			"AKIA4OXCOHAXNRXNJEFB",
			"ijgGjpheCKDCW0bH9qwvYaBagh9KD6DS9au1ZAsc",
			"",
		),
	})
	if err != nil {
		panic(err)
	}
	s3Client = s3.New(sess)
	s3Bucket = "goexpert-bucket-golang"
}

func main() {
	dir, err := os.Open("./tmp")
	if err != nil {
		panic(err)
	}

	defer dir.Close()

	uploadControl := make(chan struct{}, 100)
	errorFileUpload := make(chan string)

	go func() {
		for {
			select {
			case file := <-errorFileUpload:
				uploadControl <- struct{}{}
				wg.Add(1)
				go uploadFile(file, uploadControl, errorFileUpload)
			}
		}
	}()

	for {
		files, err := dir.Readdir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			fmt.Printf("Error reading directory: %s\n", err)
			continue
		}
		wg.Add(1)
		uploadControl <- struct{}{}
		go uploadFile(files[0].Name(), uploadControl, errorFileUpload)
	}
	wg.Wait()

}

func uploadFile(fileName string, uploadControl <-chan struct{}, errorFileUpload chan<- string) {

	completeFileName := fmt.Sprintf("./tmp/%s", fileName)
	fmt.Printf("uploading file: %s\n to bucket %s", completeFileName, s3Bucket)

	f, err := os.Open(completeFileName)

	if err != nil {
		fmt.Printf("error opening file: %s\n", completeFileName)
		<-uploadControl
		errorFileUpload <- fileName
		return
	}

	defer f.Close()

	_, err = s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s3Bucket),
		Key:    aws.String(fileName),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("error uploading file: %s\n", completeFileName)
		return
	}

	fmt.Printf("file uploaded: %s\n", completeFileName)
}
