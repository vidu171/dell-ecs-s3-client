# dell-ecs-s3-client
![alt text](https://raw.githubusercontent.com/vidu171/dell-ecs-s3-client/main/logo.png | width=100 "Dell")

Golang client for Dell EMC s3 storage

The dell-ecs-s3-client is a modified module for Delll ECS s3 storage.
This package has been created by making some modifications on the [aws-sdk-go](https://github.com/aws/aws-sdk-go) package. To meet the requirements of dell ECS s3 storage.

# Installing
To use the package just run the following command
```sh
go get github.com/vidu171/dell-ecs-s3-client
```

# Usage
```go
func WriteS3ObjectToFile(bucket, fileName string) {

	// Create S3 Session
	sess := GetS3Session(
		"<ENDPOINT>",
		"<ACCESS_ID>",
		"<SECRET_KEY>",
		"",
	)
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("Unable to open file %q, %v", fileName, err)
	}

	defer file.Close()

	downloader := s3manager.NewDownloader(sess)
	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String(fileName),
		})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

```
