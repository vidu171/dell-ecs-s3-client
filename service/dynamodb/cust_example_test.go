package dynamodb_test

import (
	"log"

	"github.com/vidu171/dell-ecs-s3-client/awstesting/unit"
	"github.com/vidu171/dell-ecs-s3-client/service/dynamodb"
)

func ExampleDynamoDB_TransactWriteItems_transactionCanceledException() {
	client := dynamodb.New(unit.Session)

	_, err := client.TransactWriteItems(&dynamodb.TransactWriteItemsInput{})
	if err != nil {
		switch t := err.(type) {
		case *dynamodb.TransactionCanceledException:
			log.Fatalf("failed to write items: %s\n%v",
				t.Message(), t.CancellationReasons)
		default:
			log.Fatalf("failed to write items: %v", err)
		}
	}
}
