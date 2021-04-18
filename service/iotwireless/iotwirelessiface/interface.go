// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotwirelessiface provides an interface to enable mocking the AWS IoT Wireless service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotwirelessiface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/iotwireless"
)

// IoTWirelessAPI provides an interface to enable mocking the
// iotwireless.IoTWireless service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Wireless.
//    func myFunc(svc iotwirelessiface.IoTWirelessAPI) bool {
//        // Make svc.AssociateAwsAccountWithPartnerAccount request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := iotwireless.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTWirelessClient struct {
//        iotwirelessiface.IoTWirelessAPI
//    }
//    func (m *mockIoTWirelessClient) AssociateAwsAccountWithPartnerAccount(input *iotwireless.AssociateAwsAccountWithPartnerAccountInput) (*iotwireless.AssociateAwsAccountWithPartnerAccountOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTWirelessClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type IoTWirelessAPI interface {
	AssociateAwsAccountWithPartnerAccount(*iotwireless.AssociateAwsAccountWithPartnerAccountInput) (*iotwireless.AssociateAwsAccountWithPartnerAccountOutput, error)
	AssociateAwsAccountWithPartnerAccountWithContext(aws.Context, *iotwireless.AssociateAwsAccountWithPartnerAccountInput, ...request.Option) (*iotwireless.AssociateAwsAccountWithPartnerAccountOutput, error)
	AssociateAwsAccountWithPartnerAccountRequest(*iotwireless.AssociateAwsAccountWithPartnerAccountInput) (*request.Request, *iotwireless.AssociateAwsAccountWithPartnerAccountOutput)

	AssociateWirelessDeviceWithThing(*iotwireless.AssociateWirelessDeviceWithThingInput) (*iotwireless.AssociateWirelessDeviceWithThingOutput, error)
	AssociateWirelessDeviceWithThingWithContext(aws.Context, *iotwireless.AssociateWirelessDeviceWithThingInput, ...request.Option) (*iotwireless.AssociateWirelessDeviceWithThingOutput, error)
	AssociateWirelessDeviceWithThingRequest(*iotwireless.AssociateWirelessDeviceWithThingInput) (*request.Request, *iotwireless.AssociateWirelessDeviceWithThingOutput)

	AssociateWirelessGatewayWithCertificate(*iotwireless.AssociateWirelessGatewayWithCertificateInput) (*iotwireless.AssociateWirelessGatewayWithCertificateOutput, error)
	AssociateWirelessGatewayWithCertificateWithContext(aws.Context, *iotwireless.AssociateWirelessGatewayWithCertificateInput, ...request.Option) (*iotwireless.AssociateWirelessGatewayWithCertificateOutput, error)
	AssociateWirelessGatewayWithCertificateRequest(*iotwireless.AssociateWirelessGatewayWithCertificateInput) (*request.Request, *iotwireless.AssociateWirelessGatewayWithCertificateOutput)

	AssociateWirelessGatewayWithThing(*iotwireless.AssociateWirelessGatewayWithThingInput) (*iotwireless.AssociateWirelessGatewayWithThingOutput, error)
	AssociateWirelessGatewayWithThingWithContext(aws.Context, *iotwireless.AssociateWirelessGatewayWithThingInput, ...request.Option) (*iotwireless.AssociateWirelessGatewayWithThingOutput, error)
	AssociateWirelessGatewayWithThingRequest(*iotwireless.AssociateWirelessGatewayWithThingInput) (*request.Request, *iotwireless.AssociateWirelessGatewayWithThingOutput)

	CreateDestination(*iotwireless.CreateDestinationInput) (*iotwireless.CreateDestinationOutput, error)
	CreateDestinationWithContext(aws.Context, *iotwireless.CreateDestinationInput, ...request.Option) (*iotwireless.CreateDestinationOutput, error)
	CreateDestinationRequest(*iotwireless.CreateDestinationInput) (*request.Request, *iotwireless.CreateDestinationOutput)

	CreateDeviceProfile(*iotwireless.CreateDeviceProfileInput) (*iotwireless.CreateDeviceProfileOutput, error)
	CreateDeviceProfileWithContext(aws.Context, *iotwireless.CreateDeviceProfileInput, ...request.Option) (*iotwireless.CreateDeviceProfileOutput, error)
	CreateDeviceProfileRequest(*iotwireless.CreateDeviceProfileInput) (*request.Request, *iotwireless.CreateDeviceProfileOutput)

	CreateServiceProfile(*iotwireless.CreateServiceProfileInput) (*iotwireless.CreateServiceProfileOutput, error)
	CreateServiceProfileWithContext(aws.Context, *iotwireless.CreateServiceProfileInput, ...request.Option) (*iotwireless.CreateServiceProfileOutput, error)
	CreateServiceProfileRequest(*iotwireless.CreateServiceProfileInput) (*request.Request, *iotwireless.CreateServiceProfileOutput)

	CreateWirelessDevice(*iotwireless.CreateWirelessDeviceInput) (*iotwireless.CreateWirelessDeviceOutput, error)
	CreateWirelessDeviceWithContext(aws.Context, *iotwireless.CreateWirelessDeviceInput, ...request.Option) (*iotwireless.CreateWirelessDeviceOutput, error)
	CreateWirelessDeviceRequest(*iotwireless.CreateWirelessDeviceInput) (*request.Request, *iotwireless.CreateWirelessDeviceOutput)

	CreateWirelessGateway(*iotwireless.CreateWirelessGatewayInput) (*iotwireless.CreateWirelessGatewayOutput, error)
	CreateWirelessGatewayWithContext(aws.Context, *iotwireless.CreateWirelessGatewayInput, ...request.Option) (*iotwireless.CreateWirelessGatewayOutput, error)
	CreateWirelessGatewayRequest(*iotwireless.CreateWirelessGatewayInput) (*request.Request, *iotwireless.CreateWirelessGatewayOutput)

	CreateWirelessGatewayTask(*iotwireless.CreateWirelessGatewayTaskInput) (*iotwireless.CreateWirelessGatewayTaskOutput, error)
	CreateWirelessGatewayTaskWithContext(aws.Context, *iotwireless.CreateWirelessGatewayTaskInput, ...request.Option) (*iotwireless.CreateWirelessGatewayTaskOutput, error)
	CreateWirelessGatewayTaskRequest(*iotwireless.CreateWirelessGatewayTaskInput) (*request.Request, *iotwireless.CreateWirelessGatewayTaskOutput)

	CreateWirelessGatewayTaskDefinition(*iotwireless.CreateWirelessGatewayTaskDefinitionInput) (*iotwireless.CreateWirelessGatewayTaskDefinitionOutput, error)
	CreateWirelessGatewayTaskDefinitionWithContext(aws.Context, *iotwireless.CreateWirelessGatewayTaskDefinitionInput, ...request.Option) (*iotwireless.CreateWirelessGatewayTaskDefinitionOutput, error)
	CreateWirelessGatewayTaskDefinitionRequest(*iotwireless.CreateWirelessGatewayTaskDefinitionInput) (*request.Request, *iotwireless.CreateWirelessGatewayTaskDefinitionOutput)

	DeleteDestination(*iotwireless.DeleteDestinationInput) (*iotwireless.DeleteDestinationOutput, error)
	DeleteDestinationWithContext(aws.Context, *iotwireless.DeleteDestinationInput, ...request.Option) (*iotwireless.DeleteDestinationOutput, error)
	DeleteDestinationRequest(*iotwireless.DeleteDestinationInput) (*request.Request, *iotwireless.DeleteDestinationOutput)

	DeleteDeviceProfile(*iotwireless.DeleteDeviceProfileInput) (*iotwireless.DeleteDeviceProfileOutput, error)
	DeleteDeviceProfileWithContext(aws.Context, *iotwireless.DeleteDeviceProfileInput, ...request.Option) (*iotwireless.DeleteDeviceProfileOutput, error)
	DeleteDeviceProfileRequest(*iotwireless.DeleteDeviceProfileInput) (*request.Request, *iotwireless.DeleteDeviceProfileOutput)

	DeleteServiceProfile(*iotwireless.DeleteServiceProfileInput) (*iotwireless.DeleteServiceProfileOutput, error)
	DeleteServiceProfileWithContext(aws.Context, *iotwireless.DeleteServiceProfileInput, ...request.Option) (*iotwireless.DeleteServiceProfileOutput, error)
	DeleteServiceProfileRequest(*iotwireless.DeleteServiceProfileInput) (*request.Request, *iotwireless.DeleteServiceProfileOutput)

	DeleteWirelessDevice(*iotwireless.DeleteWirelessDeviceInput) (*iotwireless.DeleteWirelessDeviceOutput, error)
	DeleteWirelessDeviceWithContext(aws.Context, *iotwireless.DeleteWirelessDeviceInput, ...request.Option) (*iotwireless.DeleteWirelessDeviceOutput, error)
	DeleteWirelessDeviceRequest(*iotwireless.DeleteWirelessDeviceInput) (*request.Request, *iotwireless.DeleteWirelessDeviceOutput)

	DeleteWirelessGateway(*iotwireless.DeleteWirelessGatewayInput) (*iotwireless.DeleteWirelessGatewayOutput, error)
	DeleteWirelessGatewayWithContext(aws.Context, *iotwireless.DeleteWirelessGatewayInput, ...request.Option) (*iotwireless.DeleteWirelessGatewayOutput, error)
	DeleteWirelessGatewayRequest(*iotwireless.DeleteWirelessGatewayInput) (*request.Request, *iotwireless.DeleteWirelessGatewayOutput)

	DeleteWirelessGatewayTask(*iotwireless.DeleteWirelessGatewayTaskInput) (*iotwireless.DeleteWirelessGatewayTaskOutput, error)
	DeleteWirelessGatewayTaskWithContext(aws.Context, *iotwireless.DeleteWirelessGatewayTaskInput, ...request.Option) (*iotwireless.DeleteWirelessGatewayTaskOutput, error)
	DeleteWirelessGatewayTaskRequest(*iotwireless.DeleteWirelessGatewayTaskInput) (*request.Request, *iotwireless.DeleteWirelessGatewayTaskOutput)

	DeleteWirelessGatewayTaskDefinition(*iotwireless.DeleteWirelessGatewayTaskDefinitionInput) (*iotwireless.DeleteWirelessGatewayTaskDefinitionOutput, error)
	DeleteWirelessGatewayTaskDefinitionWithContext(aws.Context, *iotwireless.DeleteWirelessGatewayTaskDefinitionInput, ...request.Option) (*iotwireless.DeleteWirelessGatewayTaskDefinitionOutput, error)
	DeleteWirelessGatewayTaskDefinitionRequest(*iotwireless.DeleteWirelessGatewayTaskDefinitionInput) (*request.Request, *iotwireless.DeleteWirelessGatewayTaskDefinitionOutput)

	DisassociateAwsAccountFromPartnerAccount(*iotwireless.DisassociateAwsAccountFromPartnerAccountInput) (*iotwireless.DisassociateAwsAccountFromPartnerAccountOutput, error)
	DisassociateAwsAccountFromPartnerAccountWithContext(aws.Context, *iotwireless.DisassociateAwsAccountFromPartnerAccountInput, ...request.Option) (*iotwireless.DisassociateAwsAccountFromPartnerAccountOutput, error)
	DisassociateAwsAccountFromPartnerAccountRequest(*iotwireless.DisassociateAwsAccountFromPartnerAccountInput) (*request.Request, *iotwireless.DisassociateAwsAccountFromPartnerAccountOutput)

	DisassociateWirelessDeviceFromThing(*iotwireless.DisassociateWirelessDeviceFromThingInput) (*iotwireless.DisassociateWirelessDeviceFromThingOutput, error)
	DisassociateWirelessDeviceFromThingWithContext(aws.Context, *iotwireless.DisassociateWirelessDeviceFromThingInput, ...request.Option) (*iotwireless.DisassociateWirelessDeviceFromThingOutput, error)
	DisassociateWirelessDeviceFromThingRequest(*iotwireless.DisassociateWirelessDeviceFromThingInput) (*request.Request, *iotwireless.DisassociateWirelessDeviceFromThingOutput)

	DisassociateWirelessGatewayFromCertificate(*iotwireless.DisassociateWirelessGatewayFromCertificateInput) (*iotwireless.DisassociateWirelessGatewayFromCertificateOutput, error)
	DisassociateWirelessGatewayFromCertificateWithContext(aws.Context, *iotwireless.DisassociateWirelessGatewayFromCertificateInput, ...request.Option) (*iotwireless.DisassociateWirelessGatewayFromCertificateOutput, error)
	DisassociateWirelessGatewayFromCertificateRequest(*iotwireless.DisassociateWirelessGatewayFromCertificateInput) (*request.Request, *iotwireless.DisassociateWirelessGatewayFromCertificateOutput)

	DisassociateWirelessGatewayFromThing(*iotwireless.DisassociateWirelessGatewayFromThingInput) (*iotwireless.DisassociateWirelessGatewayFromThingOutput, error)
	DisassociateWirelessGatewayFromThingWithContext(aws.Context, *iotwireless.DisassociateWirelessGatewayFromThingInput, ...request.Option) (*iotwireless.DisassociateWirelessGatewayFromThingOutput, error)
	DisassociateWirelessGatewayFromThingRequest(*iotwireless.DisassociateWirelessGatewayFromThingInput) (*request.Request, *iotwireless.DisassociateWirelessGatewayFromThingOutput)

	GetDestination(*iotwireless.GetDestinationInput) (*iotwireless.GetDestinationOutput, error)
	GetDestinationWithContext(aws.Context, *iotwireless.GetDestinationInput, ...request.Option) (*iotwireless.GetDestinationOutput, error)
	GetDestinationRequest(*iotwireless.GetDestinationInput) (*request.Request, *iotwireless.GetDestinationOutput)

	GetDeviceProfile(*iotwireless.GetDeviceProfileInput) (*iotwireless.GetDeviceProfileOutput, error)
	GetDeviceProfileWithContext(aws.Context, *iotwireless.GetDeviceProfileInput, ...request.Option) (*iotwireless.GetDeviceProfileOutput, error)
	GetDeviceProfileRequest(*iotwireless.GetDeviceProfileInput) (*request.Request, *iotwireless.GetDeviceProfileOutput)

	GetPartnerAccount(*iotwireless.GetPartnerAccountInput) (*iotwireless.GetPartnerAccountOutput, error)
	GetPartnerAccountWithContext(aws.Context, *iotwireless.GetPartnerAccountInput, ...request.Option) (*iotwireless.GetPartnerAccountOutput, error)
	GetPartnerAccountRequest(*iotwireless.GetPartnerAccountInput) (*request.Request, *iotwireless.GetPartnerAccountOutput)

	GetServiceEndpoint(*iotwireless.GetServiceEndpointInput) (*iotwireless.GetServiceEndpointOutput, error)
	GetServiceEndpointWithContext(aws.Context, *iotwireless.GetServiceEndpointInput, ...request.Option) (*iotwireless.GetServiceEndpointOutput, error)
	GetServiceEndpointRequest(*iotwireless.GetServiceEndpointInput) (*request.Request, *iotwireless.GetServiceEndpointOutput)

	GetServiceProfile(*iotwireless.GetServiceProfileInput) (*iotwireless.GetServiceProfileOutput, error)
	GetServiceProfileWithContext(aws.Context, *iotwireless.GetServiceProfileInput, ...request.Option) (*iotwireless.GetServiceProfileOutput, error)
	GetServiceProfileRequest(*iotwireless.GetServiceProfileInput) (*request.Request, *iotwireless.GetServiceProfileOutput)

	GetWirelessDevice(*iotwireless.GetWirelessDeviceInput) (*iotwireless.GetWirelessDeviceOutput, error)
	GetWirelessDeviceWithContext(aws.Context, *iotwireless.GetWirelessDeviceInput, ...request.Option) (*iotwireless.GetWirelessDeviceOutput, error)
	GetWirelessDeviceRequest(*iotwireless.GetWirelessDeviceInput) (*request.Request, *iotwireless.GetWirelessDeviceOutput)

	GetWirelessDeviceStatistics(*iotwireless.GetWirelessDeviceStatisticsInput) (*iotwireless.GetWirelessDeviceStatisticsOutput, error)
	GetWirelessDeviceStatisticsWithContext(aws.Context, *iotwireless.GetWirelessDeviceStatisticsInput, ...request.Option) (*iotwireless.GetWirelessDeviceStatisticsOutput, error)
	GetWirelessDeviceStatisticsRequest(*iotwireless.GetWirelessDeviceStatisticsInput) (*request.Request, *iotwireless.GetWirelessDeviceStatisticsOutput)

	GetWirelessGateway(*iotwireless.GetWirelessGatewayInput) (*iotwireless.GetWirelessGatewayOutput, error)
	GetWirelessGatewayWithContext(aws.Context, *iotwireless.GetWirelessGatewayInput, ...request.Option) (*iotwireless.GetWirelessGatewayOutput, error)
	GetWirelessGatewayRequest(*iotwireless.GetWirelessGatewayInput) (*request.Request, *iotwireless.GetWirelessGatewayOutput)

	GetWirelessGatewayCertificate(*iotwireless.GetWirelessGatewayCertificateInput) (*iotwireless.GetWirelessGatewayCertificateOutput, error)
	GetWirelessGatewayCertificateWithContext(aws.Context, *iotwireless.GetWirelessGatewayCertificateInput, ...request.Option) (*iotwireless.GetWirelessGatewayCertificateOutput, error)
	GetWirelessGatewayCertificateRequest(*iotwireless.GetWirelessGatewayCertificateInput) (*request.Request, *iotwireless.GetWirelessGatewayCertificateOutput)

	GetWirelessGatewayFirmwareInformation(*iotwireless.GetWirelessGatewayFirmwareInformationInput) (*iotwireless.GetWirelessGatewayFirmwareInformationOutput, error)
	GetWirelessGatewayFirmwareInformationWithContext(aws.Context, *iotwireless.GetWirelessGatewayFirmwareInformationInput, ...request.Option) (*iotwireless.GetWirelessGatewayFirmwareInformationOutput, error)
	GetWirelessGatewayFirmwareInformationRequest(*iotwireless.GetWirelessGatewayFirmwareInformationInput) (*request.Request, *iotwireless.GetWirelessGatewayFirmwareInformationOutput)

	GetWirelessGatewayStatistics(*iotwireless.GetWirelessGatewayStatisticsInput) (*iotwireless.GetWirelessGatewayStatisticsOutput, error)
	GetWirelessGatewayStatisticsWithContext(aws.Context, *iotwireless.GetWirelessGatewayStatisticsInput, ...request.Option) (*iotwireless.GetWirelessGatewayStatisticsOutput, error)
	GetWirelessGatewayStatisticsRequest(*iotwireless.GetWirelessGatewayStatisticsInput) (*request.Request, *iotwireless.GetWirelessGatewayStatisticsOutput)

	GetWirelessGatewayTask(*iotwireless.GetWirelessGatewayTaskInput) (*iotwireless.GetWirelessGatewayTaskOutput, error)
	GetWirelessGatewayTaskWithContext(aws.Context, *iotwireless.GetWirelessGatewayTaskInput, ...request.Option) (*iotwireless.GetWirelessGatewayTaskOutput, error)
	GetWirelessGatewayTaskRequest(*iotwireless.GetWirelessGatewayTaskInput) (*request.Request, *iotwireless.GetWirelessGatewayTaskOutput)

	GetWirelessGatewayTaskDefinition(*iotwireless.GetWirelessGatewayTaskDefinitionInput) (*iotwireless.GetWirelessGatewayTaskDefinitionOutput, error)
	GetWirelessGatewayTaskDefinitionWithContext(aws.Context, *iotwireless.GetWirelessGatewayTaskDefinitionInput, ...request.Option) (*iotwireless.GetWirelessGatewayTaskDefinitionOutput, error)
	GetWirelessGatewayTaskDefinitionRequest(*iotwireless.GetWirelessGatewayTaskDefinitionInput) (*request.Request, *iotwireless.GetWirelessGatewayTaskDefinitionOutput)

	ListDestinations(*iotwireless.ListDestinationsInput) (*iotwireless.ListDestinationsOutput, error)
	ListDestinationsWithContext(aws.Context, *iotwireless.ListDestinationsInput, ...request.Option) (*iotwireless.ListDestinationsOutput, error)
	ListDestinationsRequest(*iotwireless.ListDestinationsInput) (*request.Request, *iotwireless.ListDestinationsOutput)

	ListDestinationsPages(*iotwireless.ListDestinationsInput, func(*iotwireless.ListDestinationsOutput, bool) bool) error
	ListDestinationsPagesWithContext(aws.Context, *iotwireless.ListDestinationsInput, func(*iotwireless.ListDestinationsOutput, bool) bool, ...request.Option) error

	ListDeviceProfiles(*iotwireless.ListDeviceProfilesInput) (*iotwireless.ListDeviceProfilesOutput, error)
	ListDeviceProfilesWithContext(aws.Context, *iotwireless.ListDeviceProfilesInput, ...request.Option) (*iotwireless.ListDeviceProfilesOutput, error)
	ListDeviceProfilesRequest(*iotwireless.ListDeviceProfilesInput) (*request.Request, *iotwireless.ListDeviceProfilesOutput)

	ListDeviceProfilesPages(*iotwireless.ListDeviceProfilesInput, func(*iotwireless.ListDeviceProfilesOutput, bool) bool) error
	ListDeviceProfilesPagesWithContext(aws.Context, *iotwireless.ListDeviceProfilesInput, func(*iotwireless.ListDeviceProfilesOutput, bool) bool, ...request.Option) error

	ListPartnerAccounts(*iotwireless.ListPartnerAccountsInput) (*iotwireless.ListPartnerAccountsOutput, error)
	ListPartnerAccountsWithContext(aws.Context, *iotwireless.ListPartnerAccountsInput, ...request.Option) (*iotwireless.ListPartnerAccountsOutput, error)
	ListPartnerAccountsRequest(*iotwireless.ListPartnerAccountsInput) (*request.Request, *iotwireless.ListPartnerAccountsOutput)

	ListServiceProfiles(*iotwireless.ListServiceProfilesInput) (*iotwireless.ListServiceProfilesOutput, error)
	ListServiceProfilesWithContext(aws.Context, *iotwireless.ListServiceProfilesInput, ...request.Option) (*iotwireless.ListServiceProfilesOutput, error)
	ListServiceProfilesRequest(*iotwireless.ListServiceProfilesInput) (*request.Request, *iotwireless.ListServiceProfilesOutput)

	ListServiceProfilesPages(*iotwireless.ListServiceProfilesInput, func(*iotwireless.ListServiceProfilesOutput, bool) bool) error
	ListServiceProfilesPagesWithContext(aws.Context, *iotwireless.ListServiceProfilesInput, func(*iotwireless.ListServiceProfilesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*iotwireless.ListTagsForResourceInput) (*iotwireless.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotwireless.ListTagsForResourceInput, ...request.Option) (*iotwireless.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotwireless.ListTagsForResourceInput) (*request.Request, *iotwireless.ListTagsForResourceOutput)

	ListWirelessDevices(*iotwireless.ListWirelessDevicesInput) (*iotwireless.ListWirelessDevicesOutput, error)
	ListWirelessDevicesWithContext(aws.Context, *iotwireless.ListWirelessDevicesInput, ...request.Option) (*iotwireless.ListWirelessDevicesOutput, error)
	ListWirelessDevicesRequest(*iotwireless.ListWirelessDevicesInput) (*request.Request, *iotwireless.ListWirelessDevicesOutput)

	ListWirelessDevicesPages(*iotwireless.ListWirelessDevicesInput, func(*iotwireless.ListWirelessDevicesOutput, bool) bool) error
	ListWirelessDevicesPagesWithContext(aws.Context, *iotwireless.ListWirelessDevicesInput, func(*iotwireless.ListWirelessDevicesOutput, bool) bool, ...request.Option) error

	ListWirelessGatewayTaskDefinitions(*iotwireless.ListWirelessGatewayTaskDefinitionsInput) (*iotwireless.ListWirelessGatewayTaskDefinitionsOutput, error)
	ListWirelessGatewayTaskDefinitionsWithContext(aws.Context, *iotwireless.ListWirelessGatewayTaskDefinitionsInput, ...request.Option) (*iotwireless.ListWirelessGatewayTaskDefinitionsOutput, error)
	ListWirelessGatewayTaskDefinitionsRequest(*iotwireless.ListWirelessGatewayTaskDefinitionsInput) (*request.Request, *iotwireless.ListWirelessGatewayTaskDefinitionsOutput)

	ListWirelessGateways(*iotwireless.ListWirelessGatewaysInput) (*iotwireless.ListWirelessGatewaysOutput, error)
	ListWirelessGatewaysWithContext(aws.Context, *iotwireless.ListWirelessGatewaysInput, ...request.Option) (*iotwireless.ListWirelessGatewaysOutput, error)
	ListWirelessGatewaysRequest(*iotwireless.ListWirelessGatewaysInput) (*request.Request, *iotwireless.ListWirelessGatewaysOutput)

	ListWirelessGatewaysPages(*iotwireless.ListWirelessGatewaysInput, func(*iotwireless.ListWirelessGatewaysOutput, bool) bool) error
	ListWirelessGatewaysPagesWithContext(aws.Context, *iotwireless.ListWirelessGatewaysInput, func(*iotwireless.ListWirelessGatewaysOutput, bool) bool, ...request.Option) error

	SendDataToWirelessDevice(*iotwireless.SendDataToWirelessDeviceInput) (*iotwireless.SendDataToWirelessDeviceOutput, error)
	SendDataToWirelessDeviceWithContext(aws.Context, *iotwireless.SendDataToWirelessDeviceInput, ...request.Option) (*iotwireless.SendDataToWirelessDeviceOutput, error)
	SendDataToWirelessDeviceRequest(*iotwireless.SendDataToWirelessDeviceInput) (*request.Request, *iotwireless.SendDataToWirelessDeviceOutput)

	TagResource(*iotwireless.TagResourceInput) (*iotwireless.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotwireless.TagResourceInput, ...request.Option) (*iotwireless.TagResourceOutput, error)
	TagResourceRequest(*iotwireless.TagResourceInput) (*request.Request, *iotwireless.TagResourceOutput)

	TestWirelessDevice(*iotwireless.TestWirelessDeviceInput) (*iotwireless.TestWirelessDeviceOutput, error)
	TestWirelessDeviceWithContext(aws.Context, *iotwireless.TestWirelessDeviceInput, ...request.Option) (*iotwireless.TestWirelessDeviceOutput, error)
	TestWirelessDeviceRequest(*iotwireless.TestWirelessDeviceInput) (*request.Request, *iotwireless.TestWirelessDeviceOutput)

	UntagResource(*iotwireless.UntagResourceInput) (*iotwireless.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotwireless.UntagResourceInput, ...request.Option) (*iotwireless.UntagResourceOutput, error)
	UntagResourceRequest(*iotwireless.UntagResourceInput) (*request.Request, *iotwireless.UntagResourceOutput)

	UpdateDestination(*iotwireless.UpdateDestinationInput) (*iotwireless.UpdateDestinationOutput, error)
	UpdateDestinationWithContext(aws.Context, *iotwireless.UpdateDestinationInput, ...request.Option) (*iotwireless.UpdateDestinationOutput, error)
	UpdateDestinationRequest(*iotwireless.UpdateDestinationInput) (*request.Request, *iotwireless.UpdateDestinationOutput)

	UpdatePartnerAccount(*iotwireless.UpdatePartnerAccountInput) (*iotwireless.UpdatePartnerAccountOutput, error)
	UpdatePartnerAccountWithContext(aws.Context, *iotwireless.UpdatePartnerAccountInput, ...request.Option) (*iotwireless.UpdatePartnerAccountOutput, error)
	UpdatePartnerAccountRequest(*iotwireless.UpdatePartnerAccountInput) (*request.Request, *iotwireless.UpdatePartnerAccountOutput)

	UpdateWirelessDevice(*iotwireless.UpdateWirelessDeviceInput) (*iotwireless.UpdateWirelessDeviceOutput, error)
	UpdateWirelessDeviceWithContext(aws.Context, *iotwireless.UpdateWirelessDeviceInput, ...request.Option) (*iotwireless.UpdateWirelessDeviceOutput, error)
	UpdateWirelessDeviceRequest(*iotwireless.UpdateWirelessDeviceInput) (*request.Request, *iotwireless.UpdateWirelessDeviceOutput)

	UpdateWirelessGateway(*iotwireless.UpdateWirelessGatewayInput) (*iotwireless.UpdateWirelessGatewayOutput, error)
	UpdateWirelessGatewayWithContext(aws.Context, *iotwireless.UpdateWirelessGatewayInput, ...request.Option) (*iotwireless.UpdateWirelessGatewayOutput, error)
	UpdateWirelessGatewayRequest(*iotwireless.UpdateWirelessGatewayInput) (*request.Request, *iotwireless.UpdateWirelessGatewayOutput)
}

var _ IoTWirelessAPI = (*iotwireless.IoTWireless)(nil)
