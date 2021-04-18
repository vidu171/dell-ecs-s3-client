// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecr_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/awserr"
	"github.com/vidu171/dell-ecs-s3-client/aws/session"
	"github.com/vidu171/dell-ecs-s3-client/service/ecr"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To delete multiple images
//
// This example deletes images with the tags precise and trusty in a repository called
// ubuntu in the default registry for an account.
func ExampleECR_BatchDeleteImage_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.BatchDeleteImageInput{
		ImageIds: []*ecr.ImageIdentifier{
			{
				ImageTag: aws.String("precise"),
			},
		},
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.BatchDeleteImage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To obtain multiple images in a single request
//
// This example obtains information for an image with a specified image digest ID from
// the repository named ubuntu in the current account.
func ExampleECR_BatchGetImage_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.BatchGetImageInput{
		ImageIds: []*ecr.ImageIdentifier{
			{
				ImageTag: aws.String("precise"),
			},
		},
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.BatchGetImage(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new repository
//
// This example creates a repository called nginx-web-app inside the project-a namespace
// in the default registry for an account.
func ExampleECR_CreateRepository_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.CreateRepositoryInput{
		RepositoryName: aws.String("project-a/nginx-web-app"),
	}

	result, err := svc.CreateRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeInvalidTagParameterException:
				fmt.Println(ecr.ErrCodeInvalidTagParameterException, aerr.Error())
			case ecr.ErrCodeTooManyTagsException:
				fmt.Println(ecr.ErrCodeTooManyTagsException, aerr.Error())
			case ecr.ErrCodeRepositoryAlreadyExistsException:
				fmt.Println(ecr.ErrCodeRepositoryAlreadyExistsException, aerr.Error())
			case ecr.ErrCodeLimitExceededException:
				fmt.Println(ecr.ErrCodeLimitExceededException, aerr.Error())
			case ecr.ErrCodeKmsException:
				fmt.Println(ecr.ErrCodeKmsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To force delete a repository
//
// This example force deletes a repository named ubuntu in the default registry for
// an account. The force parameter is required if the repository contains images.
func ExampleECR_DeleteRepository_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.DeleteRepositoryInput{
		Force:          aws.Bool(true),
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.DeleteRepository(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			case ecr.ErrCodeRepositoryNotEmptyException:
				fmt.Println(ecr.ErrCodeRepositoryNotEmptyException, aerr.Error())
			case ecr.ErrCodeKmsException:
				fmt.Println(ecr.ErrCodeKmsException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete the policy associated with a repository
//
// This example deletes the policy associated with the repository named ubuntu in the
// current account.
func ExampleECR_DeleteRepositoryPolicy_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.DeleteRepositoryPolicyInput{
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.DeleteRepositoryPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			case ecr.ErrCodeRepositoryPolicyNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryPolicyNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe all repositories in the current account
//
// The following example obtains a list and description of all repositories in the default
// registry to which the current user has access.
func ExampleECR_DescribeRepositories_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.DescribeRepositoriesInput{}

	result, err := svc.DescribeRepositories(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To obtain an authorization token
//
// This example gets an authorization token for your default registry.
func ExampleECR_GetAuthorizationToken_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.GetAuthorizationTokenInput{}

	result, err := svc.GetAuthorizationToken(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To get the current policy for a repository
//
// This example obtains the repository policy for the repository named ubuntu.
func ExampleECR_GetRepositoryPolicy_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.GetRepositoryPolicyInput{
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.GetRepositoryPolicy(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			case ecr.ErrCodeRepositoryPolicyNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryPolicyNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list all images in a repository
//
// This example lists all of the images in the repository named ubuntu in the default
// registry in the current account.
func ExampleECR_ListImages_shared00() {
	svc := ecr.New(session.New())
	input := &ecr.ListImagesInput{
		RepositoryName: aws.String("ubuntu"),
	}

	result, err := svc.ListImages(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case ecr.ErrCodeServerException:
				fmt.Println(ecr.ErrCodeServerException, aerr.Error())
			case ecr.ErrCodeInvalidParameterException:
				fmt.Println(ecr.ErrCodeInvalidParameterException, aerr.Error())
			case ecr.ErrCodeRepositoryNotFoundException:
				fmt.Println(ecr.ErrCodeRepositoryNotFoundException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}
