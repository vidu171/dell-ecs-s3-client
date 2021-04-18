package mediastoredata_test

import (
	"fmt"
	"log"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/session"
	"github.com/vidu171/dell-ecs-s3-client/service/mediastore"
	"github.com/vidu171/dell-ecs-s3-client/service/mediastoredata"
)

func ExampleMediaStoreData_describeEndpoint() {
	sess, err := session.NewSession(aws.NewConfig())
	if err != nil {
		log.Fatal("Failed to create aws session", err)
	}

	// we need to use a MediaStore client to get a media store container endpoint address
	ctrlSvc := mediastore.New(sess)
	descResp, err := ctrlSvc.DescribeContainer(&mediastore.DescribeContainerInput{
		// specify a container name
		ContainerName: aws.String("some-container"),
	})
	if err != nil {
		log.Fatal("failed to get media store container endpoint", err)
	}

	// create a MediaStoreData client and use the retrieved container endpoint
	dataSvc := mediastoredata.New(sess, &aws.Config{
		Endpoint: descResp.Container.Endpoint,
	})
	output, err := dataSvc.ListItems(&mediastoredata.ListItemsInput{})
	if err != nil {
		log.Fatal("failed to make mediastoredata API call", err)
	}

	// prints the string representation of ListItemsOutput
	fmt.Println(output.GoString())
}
