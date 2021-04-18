package sqs

import "github.com/vidu171/dell-ecs-s3-client/aws/request"

func init() {
	initRequest = func(r *request.Request) {
		setupChecksumValidation(r)
	}
}
