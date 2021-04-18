// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package servicediscovery_test

import (
	"fmt"
	"strings"
	"time"

	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/awserr"
	"github.com/vidu171/dell-ecs-s3-client/aws/session"
	"github.com/vidu171/dell-ecs-s3-client/service/servicediscovery"
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

// CreateHttpNamespace example
//
// This example creates an HTTP namespace.
func ExampleServiceDiscovery_CreateHttpNamespace_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.CreateHttpNamespaceInput{
		CreatorRequestId: aws.String("example-creator-request-id-0001"),
		Description:      aws.String("Example.com AWS Cloud Map HTTP Namespace"),
		Name:             aws.String("example-http.com"),
	}

	result, err := svc.CreateHttpNamespace(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
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

// Example: Create private DNS namespace
//
// Example: Create private DNS namespace
func ExampleServiceDiscovery_CreatePrivateDnsNamespace_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.CreatePrivateDnsNamespaceInput{
		CreatorRequestId: aws.String("eedd6892-50f3-41b2-8af9-611d6e1d1a8c"),
		Name:             aws.String("example.com"),
		Vpc:              aws.String("vpc-1c56417b"),
	}

	result, err := svc.CreatePrivateDnsNamespace(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
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

// CreatePublicDnsNamespace example
//
// This example creates a public namespace based on DNS.
func ExampleServiceDiscovery_CreatePublicDnsNamespace_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.CreatePublicDnsNamespaceInput{
		CreatorRequestId: aws.String("example-creator-request-id-0003"),
		Description:      aws.String("Example.com AWS Cloud Map Public DNS Namespace"),
		Name:             aws.String("example-public-dns.com"),
	}

	result, err := svc.CreatePublicDnsNamespace(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeNamespaceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
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

// Example: Create service
//
// Example: Create service
func ExampleServiceDiscovery_CreateService_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.CreateServiceInput{
		CreatorRequestId: aws.String("567c1193-6b00-4308-bd57-ad38a8822d25"),
		DnsConfig: &servicediscovery.DnsConfig{
			DnsRecords: []*servicediscovery.DnsRecord{
				{
					TTL:  aws.Int64(60),
					Type: aws.String("A"),
				},
			},
			NamespaceId:   aws.String("ns-ylexjili4cdxy3xm"),
			RoutingPolicy: aws.String("MULTIVALUE"),
		},
		Name:        aws.String("myservice"),
		NamespaceId: aws.String("ns-ylexjili4cdxy3xm"),
	}

	result, err := svc.CreateService(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeServiceAlreadyExists:
				fmt.Println(servicediscovery.ErrCodeServiceAlreadyExists, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
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

// Example: Delete namespace
//
// Example: Delete namespace
func ExampleServiceDiscovery_DeleteNamespace_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.DeleteNamespaceInput{
		Id: aws.String("ns-ylexjili4cdxy3xm"),
	}

	result, err := svc.DeleteNamespace(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
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

// Example: Delete service
//
// Example: Delete service
func ExampleServiceDiscovery_DeleteService_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.DeleteServiceInput{
		Id: aws.String("srv-p5zdwlg5uvvzjita"),
	}

	result, err := svc.DeleteService(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
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

// Example: Deregister a service instance
//
// Example: Deregister a service instance
func ExampleServiceDiscovery_DeregisterInstance_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.DeregisterInstanceInput{
		InstanceId: aws.String("myservice-53"),
		ServiceId:  aws.String("srv-p5zdwlg5uvvzjita"),
	}

	result, err := svc.DeregisterInstance(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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

// Example: Discover registered instances
//
// Example: Discover registered instances
func ExampleServiceDiscovery_DiscoverInstances_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.DiscoverInstancesInput{
		HealthStatus:  aws.String("ALL"),
		MaxResults:    aws.Int64(10),
		NamespaceName: aws.String("example.com"),
		ServiceName:   aws.String("myservice"),
	}

	result, err := svc.DiscoverInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeRequestLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeRequestLimitExceeded, aerr.Error())
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

// GetInstance example
//
// This example gets information about a specified instance.
func ExampleServiceDiscovery_GetInstance_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.GetInstanceInput{
		InstanceId: aws.String("i-abcd1234"),
		ServiceId:  aws.String("srv-e4anhexample0004"),
	}

	result, err := svc.GetInstance(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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

// GetInstancesHealthStatus example
//
// This example gets the current health status of one or more instances that are associate
// with a specified service.
func ExampleServiceDiscovery_GetInstancesHealthStatus_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.GetInstancesHealthStatusInput{
		ServiceId: aws.String("srv-e4anhexample0004"),
	}

	result, err := svc.GetInstancesHealthStatus(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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

// GetNamespace example
//
// This example gets information about a specified namespace.
func ExampleServiceDiscovery_GetNamespace_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.GetNamespaceInput{
		Id: aws.String("ns-e4anhexample0004"),
	}

	result, err := svc.GetNamespace(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeNamespaceNotFound:
				fmt.Println(servicediscovery.ErrCodeNamespaceNotFound, aerr.Error())
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

// Example: Get operation result
//
// Example: Get operation result
func ExampleServiceDiscovery_GetOperation_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.GetOperationInput{
		OperationId: aws.String("gv4g5meo7ndmeh4fqskygvk23d2fijwa-k9302yzd"),
	}

	result, err := svc.GetOperation(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeOperationNotFound:
				fmt.Println(servicediscovery.ErrCodeOperationNotFound, aerr.Error())
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

// GetService Example
//
// This example gets the settings for a specified service.
func ExampleServiceDiscovery_GetService_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.GetServiceInput{
		Id: aws.String("srv-e4anhexample0004"),
	}

	result, err := svc.GetService(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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

// Example: List service instances
//
// Example: List service instances
func ExampleServiceDiscovery_ListInstances_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.ListInstancesInput{
		ServiceId: aws.String("srv-qzpwvt2tfqcegapy"),
	}

	result, err := svc.ListInstances(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// Example: List namespaces
//
// Example: List namespaces
func ExampleServiceDiscovery_ListNamespaces_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.ListNamespacesInput{}

	result, err := svc.ListNamespaces(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// ListOperations Example
//
// This example gets the operations that have a STATUS of either PENDING or SUCCESS.
func ExampleServiceDiscovery_ListOperations_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.ListOperationsInput{
		Filters: []*servicediscovery.OperationFilter{
			{
				Condition: aws.String("IN"),
				Name:      aws.String("STATUS"),
				Values: []*string{
					aws.String("PENDING"),
					aws.String("SUCCESS"),
				},
			},
		},
	}

	result, err := svc.ListOperations(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// Example: List services
//
// Example: List services
func ExampleServiceDiscovery_ListServices_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.ListServicesInput{}

	result, err := svc.ListServices(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// ListTagsForResource example
//
// This example lists the tags of a resource.
func ExampleServiceDiscovery_ListTagsForResource_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
	}

	result, err := svc.ListTagsForResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// Example: Register Instance
//
// Example: Register Instance
func ExampleServiceDiscovery_RegisterInstance_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.RegisterInstanceInput{
		Attributes: map[string]*string{
			"AWS_INSTANCE_IPV4": aws.String("172.2.1.3"),
			"AWS_INSTANCE_PORT": aws.String("808"),
		},
		CreatorRequestId: aws.String("7a48a98a-72e6-4849-bfa7-1a458e030d7b"),
		InstanceId:       aws.String("myservice-53"),
		ServiceId:        aws.String("srv-p5zdwlg5uvvzjita"),
	}

	result, err := svc.RegisterInstance(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeResourceInUse:
				fmt.Println(servicediscovery.ErrCodeResourceInUse, aerr.Error())
			case servicediscovery.ErrCodeResourceLimitExceeded:
				fmt.Println(servicediscovery.ErrCodeResourceLimitExceeded, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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

// TagResource example
//
// This example adds "Department" and "Project" tags to a resource.
func ExampleServiceDiscovery_TagResource_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.TagResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
		Tags: []*servicediscovery.Tag{
			{
				Key:   aws.String("Department"),
				Value: aws.String("Engineering"),
			},
			{
				Key:   aws.String("Project"),
				Value: aws.String("Zeta"),
			},
		},
	}

	result, err := svc.TagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeTooManyTagsException:
				fmt.Println(servicediscovery.ErrCodeTooManyTagsException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// UntagResource example
//
// This example removes the "Department" and "Project" tags from a resource.
func ExampleServiceDiscovery_UntagResource_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:servicediscovery:us-east-1:123456789012:namespace/ns-ylexjili4cdxy3xm"),
		TagKeys: []*string{
			aws.String("Project"),
			aws.String("Department"),
		},
	}

	result, err := svc.UntagResource(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeResourceNotFoundException:
				fmt.Println(servicediscovery.ErrCodeResourceNotFoundException, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// UpdateInstanceCustomHealthStatus Example
//
// This example submits a request to change the health status of an instance associated
// with a service with a custom health check to HEALTHY.
func ExampleServiceDiscovery_UpdateInstanceCustomHealthStatus_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.UpdateInstanceCustomHealthStatusInput{
		InstanceId: aws.String("i-abcd1234"),
		ServiceId:  aws.String("srv-e4anhexample0004"),
		Status:     aws.String("HEALTHY"),
	}

	result, err := svc.UpdateInstanceCustomHealthStatus(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeInstanceNotFound:
				fmt.Println(servicediscovery.ErrCodeInstanceNotFound, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
			case servicediscovery.ErrCodeCustomHealthNotFound:
				fmt.Println(servicediscovery.ErrCodeCustomHealthNotFound, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
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

// UpdateService Example
//
// This example submits a request to replace the DnsConfig and HealthCheckConfig settings
// of a specified service.
func ExampleServiceDiscovery_UpdateService_shared00() {
	svc := servicediscovery.New(session.New())
	input := &servicediscovery.UpdateServiceInput{
		Id: aws.String("srv-e4anhexample0004"),
		Service: &servicediscovery.ServiceChange{
			DnsConfig: &servicediscovery.DnsConfigChange{
				DnsRecords: []*servicediscovery.DnsRecord{
					{
						TTL:  aws.Int64(60),
						Type: aws.String("A"),
					},
				},
			},
			HealthCheckConfig: &servicediscovery.HealthCheckConfig{
				FailureThreshold: aws.Int64(2),
				ResourcePath:     aws.String("/"),
				Type:             aws.String("HTTP"),
			},
		},
	}

	result, err := svc.UpdateService(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case servicediscovery.ErrCodeDuplicateRequest:
				fmt.Println(servicediscovery.ErrCodeDuplicateRequest, aerr.Error())
			case servicediscovery.ErrCodeInvalidInput:
				fmt.Println(servicediscovery.ErrCodeInvalidInput, aerr.Error())
			case servicediscovery.ErrCodeServiceNotFound:
				fmt.Println(servicediscovery.ErrCodeServiceNotFound, aerr.Error())
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
