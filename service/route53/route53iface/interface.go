// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53iface provides an interface to enable mocking the Amazon Route 53 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53iface

import (
	"github.com/vidu171/dell-ecs-s3-client/aws"
	"github.com/vidu171/dell-ecs-s3-client/aws/request"
	"github.com/vidu171/dell-ecs-s3-client/service/route53"
)

// Route53API provides an interface to enable mocking the
// route53.Route53 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53.
//    func myFunc(svc route53iface.Route53API) bool {
//        // Make svc.ActivateKeySigningKey request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := route53.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53Client struct {
//        route53iface.Route53API
//    }
//    func (m *mockRoute53Client) ActivateKeySigningKey(input *route53.ActivateKeySigningKeyInput) (*route53.ActivateKeySigningKeyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53Client{}
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
type Route53API interface {
	ActivateKeySigningKey(*route53.ActivateKeySigningKeyInput) (*route53.ActivateKeySigningKeyOutput, error)
	ActivateKeySigningKeyWithContext(aws.Context, *route53.ActivateKeySigningKeyInput, ...request.Option) (*route53.ActivateKeySigningKeyOutput, error)
	ActivateKeySigningKeyRequest(*route53.ActivateKeySigningKeyInput) (*request.Request, *route53.ActivateKeySigningKeyOutput)

	AssociateVPCWithHostedZone(*route53.AssociateVPCWithHostedZoneInput) (*route53.AssociateVPCWithHostedZoneOutput, error)
	AssociateVPCWithHostedZoneWithContext(aws.Context, *route53.AssociateVPCWithHostedZoneInput, ...request.Option) (*route53.AssociateVPCWithHostedZoneOutput, error)
	AssociateVPCWithHostedZoneRequest(*route53.AssociateVPCWithHostedZoneInput) (*request.Request, *route53.AssociateVPCWithHostedZoneOutput)

	ChangeResourceRecordSets(*route53.ChangeResourceRecordSetsInput) (*route53.ChangeResourceRecordSetsOutput, error)
	ChangeResourceRecordSetsWithContext(aws.Context, *route53.ChangeResourceRecordSetsInput, ...request.Option) (*route53.ChangeResourceRecordSetsOutput, error)
	ChangeResourceRecordSetsRequest(*route53.ChangeResourceRecordSetsInput) (*request.Request, *route53.ChangeResourceRecordSetsOutput)

	ChangeTagsForResource(*route53.ChangeTagsForResourceInput) (*route53.ChangeTagsForResourceOutput, error)
	ChangeTagsForResourceWithContext(aws.Context, *route53.ChangeTagsForResourceInput, ...request.Option) (*route53.ChangeTagsForResourceOutput, error)
	ChangeTagsForResourceRequest(*route53.ChangeTagsForResourceInput) (*request.Request, *route53.ChangeTagsForResourceOutput)

	CreateHealthCheck(*route53.CreateHealthCheckInput) (*route53.CreateHealthCheckOutput, error)
	CreateHealthCheckWithContext(aws.Context, *route53.CreateHealthCheckInput, ...request.Option) (*route53.CreateHealthCheckOutput, error)
	CreateHealthCheckRequest(*route53.CreateHealthCheckInput) (*request.Request, *route53.CreateHealthCheckOutput)

	CreateHostedZone(*route53.CreateHostedZoneInput) (*route53.CreateHostedZoneOutput, error)
	CreateHostedZoneWithContext(aws.Context, *route53.CreateHostedZoneInput, ...request.Option) (*route53.CreateHostedZoneOutput, error)
	CreateHostedZoneRequest(*route53.CreateHostedZoneInput) (*request.Request, *route53.CreateHostedZoneOutput)

	CreateKeySigningKey(*route53.CreateKeySigningKeyInput) (*route53.CreateKeySigningKeyOutput, error)
	CreateKeySigningKeyWithContext(aws.Context, *route53.CreateKeySigningKeyInput, ...request.Option) (*route53.CreateKeySigningKeyOutput, error)
	CreateKeySigningKeyRequest(*route53.CreateKeySigningKeyInput) (*request.Request, *route53.CreateKeySigningKeyOutput)

	CreateQueryLoggingConfig(*route53.CreateQueryLoggingConfigInput) (*route53.CreateQueryLoggingConfigOutput, error)
	CreateQueryLoggingConfigWithContext(aws.Context, *route53.CreateQueryLoggingConfigInput, ...request.Option) (*route53.CreateQueryLoggingConfigOutput, error)
	CreateQueryLoggingConfigRequest(*route53.CreateQueryLoggingConfigInput) (*request.Request, *route53.CreateQueryLoggingConfigOutput)

	CreateReusableDelegationSet(*route53.CreateReusableDelegationSetInput) (*route53.CreateReusableDelegationSetOutput, error)
	CreateReusableDelegationSetWithContext(aws.Context, *route53.CreateReusableDelegationSetInput, ...request.Option) (*route53.CreateReusableDelegationSetOutput, error)
	CreateReusableDelegationSetRequest(*route53.CreateReusableDelegationSetInput) (*request.Request, *route53.CreateReusableDelegationSetOutput)

	CreateTrafficPolicy(*route53.CreateTrafficPolicyInput) (*route53.CreateTrafficPolicyOutput, error)
	CreateTrafficPolicyWithContext(aws.Context, *route53.CreateTrafficPolicyInput, ...request.Option) (*route53.CreateTrafficPolicyOutput, error)
	CreateTrafficPolicyRequest(*route53.CreateTrafficPolicyInput) (*request.Request, *route53.CreateTrafficPolicyOutput)

	CreateTrafficPolicyInstance(*route53.CreateTrafficPolicyInstanceInput) (*route53.CreateTrafficPolicyInstanceOutput, error)
	CreateTrafficPolicyInstanceWithContext(aws.Context, *route53.CreateTrafficPolicyInstanceInput, ...request.Option) (*route53.CreateTrafficPolicyInstanceOutput, error)
	CreateTrafficPolicyInstanceRequest(*route53.CreateTrafficPolicyInstanceInput) (*request.Request, *route53.CreateTrafficPolicyInstanceOutput)

	CreateTrafficPolicyVersion(*route53.CreateTrafficPolicyVersionInput) (*route53.CreateTrafficPolicyVersionOutput, error)
	CreateTrafficPolicyVersionWithContext(aws.Context, *route53.CreateTrafficPolicyVersionInput, ...request.Option) (*route53.CreateTrafficPolicyVersionOutput, error)
	CreateTrafficPolicyVersionRequest(*route53.CreateTrafficPolicyVersionInput) (*request.Request, *route53.CreateTrafficPolicyVersionOutput)

	CreateVPCAssociationAuthorization(*route53.CreateVPCAssociationAuthorizationInput) (*route53.CreateVPCAssociationAuthorizationOutput, error)
	CreateVPCAssociationAuthorizationWithContext(aws.Context, *route53.CreateVPCAssociationAuthorizationInput, ...request.Option) (*route53.CreateVPCAssociationAuthorizationOutput, error)
	CreateVPCAssociationAuthorizationRequest(*route53.CreateVPCAssociationAuthorizationInput) (*request.Request, *route53.CreateVPCAssociationAuthorizationOutput)

	DeactivateKeySigningKey(*route53.DeactivateKeySigningKeyInput) (*route53.DeactivateKeySigningKeyOutput, error)
	DeactivateKeySigningKeyWithContext(aws.Context, *route53.DeactivateKeySigningKeyInput, ...request.Option) (*route53.DeactivateKeySigningKeyOutput, error)
	DeactivateKeySigningKeyRequest(*route53.DeactivateKeySigningKeyInput) (*request.Request, *route53.DeactivateKeySigningKeyOutput)

	DeleteHealthCheck(*route53.DeleteHealthCheckInput) (*route53.DeleteHealthCheckOutput, error)
	DeleteHealthCheckWithContext(aws.Context, *route53.DeleteHealthCheckInput, ...request.Option) (*route53.DeleteHealthCheckOutput, error)
	DeleteHealthCheckRequest(*route53.DeleteHealthCheckInput) (*request.Request, *route53.DeleteHealthCheckOutput)

	DeleteHostedZone(*route53.DeleteHostedZoneInput) (*route53.DeleteHostedZoneOutput, error)
	DeleteHostedZoneWithContext(aws.Context, *route53.DeleteHostedZoneInput, ...request.Option) (*route53.DeleteHostedZoneOutput, error)
	DeleteHostedZoneRequest(*route53.DeleteHostedZoneInput) (*request.Request, *route53.DeleteHostedZoneOutput)

	DeleteKeySigningKey(*route53.DeleteKeySigningKeyInput) (*route53.DeleteKeySigningKeyOutput, error)
	DeleteKeySigningKeyWithContext(aws.Context, *route53.DeleteKeySigningKeyInput, ...request.Option) (*route53.DeleteKeySigningKeyOutput, error)
	DeleteKeySigningKeyRequest(*route53.DeleteKeySigningKeyInput) (*request.Request, *route53.DeleteKeySigningKeyOutput)

	DeleteQueryLoggingConfig(*route53.DeleteQueryLoggingConfigInput) (*route53.DeleteQueryLoggingConfigOutput, error)
	DeleteQueryLoggingConfigWithContext(aws.Context, *route53.DeleteQueryLoggingConfigInput, ...request.Option) (*route53.DeleteQueryLoggingConfigOutput, error)
	DeleteQueryLoggingConfigRequest(*route53.DeleteQueryLoggingConfigInput) (*request.Request, *route53.DeleteQueryLoggingConfigOutput)

	DeleteReusableDelegationSet(*route53.DeleteReusableDelegationSetInput) (*route53.DeleteReusableDelegationSetOutput, error)
	DeleteReusableDelegationSetWithContext(aws.Context, *route53.DeleteReusableDelegationSetInput, ...request.Option) (*route53.DeleteReusableDelegationSetOutput, error)
	DeleteReusableDelegationSetRequest(*route53.DeleteReusableDelegationSetInput) (*request.Request, *route53.DeleteReusableDelegationSetOutput)

	DeleteTrafficPolicy(*route53.DeleteTrafficPolicyInput) (*route53.DeleteTrafficPolicyOutput, error)
	DeleteTrafficPolicyWithContext(aws.Context, *route53.DeleteTrafficPolicyInput, ...request.Option) (*route53.DeleteTrafficPolicyOutput, error)
	DeleteTrafficPolicyRequest(*route53.DeleteTrafficPolicyInput) (*request.Request, *route53.DeleteTrafficPolicyOutput)

	DeleteTrafficPolicyInstance(*route53.DeleteTrafficPolicyInstanceInput) (*route53.DeleteTrafficPolicyInstanceOutput, error)
	DeleteTrafficPolicyInstanceWithContext(aws.Context, *route53.DeleteTrafficPolicyInstanceInput, ...request.Option) (*route53.DeleteTrafficPolicyInstanceOutput, error)
	DeleteTrafficPolicyInstanceRequest(*route53.DeleteTrafficPolicyInstanceInput) (*request.Request, *route53.DeleteTrafficPolicyInstanceOutput)

	DeleteVPCAssociationAuthorization(*route53.DeleteVPCAssociationAuthorizationInput) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
	DeleteVPCAssociationAuthorizationWithContext(aws.Context, *route53.DeleteVPCAssociationAuthorizationInput, ...request.Option) (*route53.DeleteVPCAssociationAuthorizationOutput, error)
	DeleteVPCAssociationAuthorizationRequest(*route53.DeleteVPCAssociationAuthorizationInput) (*request.Request, *route53.DeleteVPCAssociationAuthorizationOutput)

	DisableHostedZoneDNSSEC(*route53.DisableHostedZoneDNSSECInput) (*route53.DisableHostedZoneDNSSECOutput, error)
	DisableHostedZoneDNSSECWithContext(aws.Context, *route53.DisableHostedZoneDNSSECInput, ...request.Option) (*route53.DisableHostedZoneDNSSECOutput, error)
	DisableHostedZoneDNSSECRequest(*route53.DisableHostedZoneDNSSECInput) (*request.Request, *route53.DisableHostedZoneDNSSECOutput)

	DisassociateVPCFromHostedZone(*route53.DisassociateVPCFromHostedZoneInput) (*route53.DisassociateVPCFromHostedZoneOutput, error)
	DisassociateVPCFromHostedZoneWithContext(aws.Context, *route53.DisassociateVPCFromHostedZoneInput, ...request.Option) (*route53.DisassociateVPCFromHostedZoneOutput, error)
	DisassociateVPCFromHostedZoneRequest(*route53.DisassociateVPCFromHostedZoneInput) (*request.Request, *route53.DisassociateVPCFromHostedZoneOutput)

	EnableHostedZoneDNSSEC(*route53.EnableHostedZoneDNSSECInput) (*route53.EnableHostedZoneDNSSECOutput, error)
	EnableHostedZoneDNSSECWithContext(aws.Context, *route53.EnableHostedZoneDNSSECInput, ...request.Option) (*route53.EnableHostedZoneDNSSECOutput, error)
	EnableHostedZoneDNSSECRequest(*route53.EnableHostedZoneDNSSECInput) (*request.Request, *route53.EnableHostedZoneDNSSECOutput)

	GetAccountLimit(*route53.GetAccountLimitInput) (*route53.GetAccountLimitOutput, error)
	GetAccountLimitWithContext(aws.Context, *route53.GetAccountLimitInput, ...request.Option) (*route53.GetAccountLimitOutput, error)
	GetAccountLimitRequest(*route53.GetAccountLimitInput) (*request.Request, *route53.GetAccountLimitOutput)

	GetChange(*route53.GetChangeInput) (*route53.GetChangeOutput, error)
	GetChangeWithContext(aws.Context, *route53.GetChangeInput, ...request.Option) (*route53.GetChangeOutput, error)
	GetChangeRequest(*route53.GetChangeInput) (*request.Request, *route53.GetChangeOutput)

	GetCheckerIpRanges(*route53.GetCheckerIpRangesInput) (*route53.GetCheckerIpRangesOutput, error)
	GetCheckerIpRangesWithContext(aws.Context, *route53.GetCheckerIpRangesInput, ...request.Option) (*route53.GetCheckerIpRangesOutput, error)
	GetCheckerIpRangesRequest(*route53.GetCheckerIpRangesInput) (*request.Request, *route53.GetCheckerIpRangesOutput)

	GetDNSSEC(*route53.GetDNSSECInput) (*route53.GetDNSSECOutput, error)
	GetDNSSECWithContext(aws.Context, *route53.GetDNSSECInput, ...request.Option) (*route53.GetDNSSECOutput, error)
	GetDNSSECRequest(*route53.GetDNSSECInput) (*request.Request, *route53.GetDNSSECOutput)

	GetGeoLocation(*route53.GetGeoLocationInput) (*route53.GetGeoLocationOutput, error)
	GetGeoLocationWithContext(aws.Context, *route53.GetGeoLocationInput, ...request.Option) (*route53.GetGeoLocationOutput, error)
	GetGeoLocationRequest(*route53.GetGeoLocationInput) (*request.Request, *route53.GetGeoLocationOutput)

	GetHealthCheck(*route53.GetHealthCheckInput) (*route53.GetHealthCheckOutput, error)
	GetHealthCheckWithContext(aws.Context, *route53.GetHealthCheckInput, ...request.Option) (*route53.GetHealthCheckOutput, error)
	GetHealthCheckRequest(*route53.GetHealthCheckInput) (*request.Request, *route53.GetHealthCheckOutput)

	GetHealthCheckCount(*route53.GetHealthCheckCountInput) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckCountWithContext(aws.Context, *route53.GetHealthCheckCountInput, ...request.Option) (*route53.GetHealthCheckCountOutput, error)
	GetHealthCheckCountRequest(*route53.GetHealthCheckCountInput) (*request.Request, *route53.GetHealthCheckCountOutput)

	GetHealthCheckLastFailureReason(*route53.GetHealthCheckLastFailureReasonInput) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckLastFailureReasonWithContext(aws.Context, *route53.GetHealthCheckLastFailureReasonInput, ...request.Option) (*route53.GetHealthCheckLastFailureReasonOutput, error)
	GetHealthCheckLastFailureReasonRequest(*route53.GetHealthCheckLastFailureReasonInput) (*request.Request, *route53.GetHealthCheckLastFailureReasonOutput)

	GetHealthCheckStatus(*route53.GetHealthCheckStatusInput) (*route53.GetHealthCheckStatusOutput, error)
	GetHealthCheckStatusWithContext(aws.Context, *route53.GetHealthCheckStatusInput, ...request.Option) (*route53.GetHealthCheckStatusOutput, error)
	GetHealthCheckStatusRequest(*route53.GetHealthCheckStatusInput) (*request.Request, *route53.GetHealthCheckStatusOutput)

	GetHostedZone(*route53.GetHostedZoneInput) (*route53.GetHostedZoneOutput, error)
	GetHostedZoneWithContext(aws.Context, *route53.GetHostedZoneInput, ...request.Option) (*route53.GetHostedZoneOutput, error)
	GetHostedZoneRequest(*route53.GetHostedZoneInput) (*request.Request, *route53.GetHostedZoneOutput)

	GetHostedZoneCount(*route53.GetHostedZoneCountInput) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneCountWithContext(aws.Context, *route53.GetHostedZoneCountInput, ...request.Option) (*route53.GetHostedZoneCountOutput, error)
	GetHostedZoneCountRequest(*route53.GetHostedZoneCountInput) (*request.Request, *route53.GetHostedZoneCountOutput)

	GetHostedZoneLimit(*route53.GetHostedZoneLimitInput) (*route53.GetHostedZoneLimitOutput, error)
	GetHostedZoneLimitWithContext(aws.Context, *route53.GetHostedZoneLimitInput, ...request.Option) (*route53.GetHostedZoneLimitOutput, error)
	GetHostedZoneLimitRequest(*route53.GetHostedZoneLimitInput) (*request.Request, *route53.GetHostedZoneLimitOutput)

	GetQueryLoggingConfig(*route53.GetQueryLoggingConfigInput) (*route53.GetQueryLoggingConfigOutput, error)
	GetQueryLoggingConfigWithContext(aws.Context, *route53.GetQueryLoggingConfigInput, ...request.Option) (*route53.GetQueryLoggingConfigOutput, error)
	GetQueryLoggingConfigRequest(*route53.GetQueryLoggingConfigInput) (*request.Request, *route53.GetQueryLoggingConfigOutput)

	GetReusableDelegationSet(*route53.GetReusableDelegationSetInput) (*route53.GetReusableDelegationSetOutput, error)
	GetReusableDelegationSetWithContext(aws.Context, *route53.GetReusableDelegationSetInput, ...request.Option) (*route53.GetReusableDelegationSetOutput, error)
	GetReusableDelegationSetRequest(*route53.GetReusableDelegationSetInput) (*request.Request, *route53.GetReusableDelegationSetOutput)

	GetReusableDelegationSetLimit(*route53.GetReusableDelegationSetLimitInput) (*route53.GetReusableDelegationSetLimitOutput, error)
	GetReusableDelegationSetLimitWithContext(aws.Context, *route53.GetReusableDelegationSetLimitInput, ...request.Option) (*route53.GetReusableDelegationSetLimitOutput, error)
	GetReusableDelegationSetLimitRequest(*route53.GetReusableDelegationSetLimitInput) (*request.Request, *route53.GetReusableDelegationSetLimitOutput)

	GetTrafficPolicy(*route53.GetTrafficPolicyInput) (*route53.GetTrafficPolicyOutput, error)
	GetTrafficPolicyWithContext(aws.Context, *route53.GetTrafficPolicyInput, ...request.Option) (*route53.GetTrafficPolicyOutput, error)
	GetTrafficPolicyRequest(*route53.GetTrafficPolicyInput) (*request.Request, *route53.GetTrafficPolicyOutput)

	GetTrafficPolicyInstance(*route53.GetTrafficPolicyInstanceInput) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyInstanceWithContext(aws.Context, *route53.GetTrafficPolicyInstanceInput, ...request.Option) (*route53.GetTrafficPolicyInstanceOutput, error)
	GetTrafficPolicyInstanceRequest(*route53.GetTrafficPolicyInstanceInput) (*request.Request, *route53.GetTrafficPolicyInstanceOutput)

	GetTrafficPolicyInstanceCount(*route53.GetTrafficPolicyInstanceCountInput) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	GetTrafficPolicyInstanceCountWithContext(aws.Context, *route53.GetTrafficPolicyInstanceCountInput, ...request.Option) (*route53.GetTrafficPolicyInstanceCountOutput, error)
	GetTrafficPolicyInstanceCountRequest(*route53.GetTrafficPolicyInstanceCountInput) (*request.Request, *route53.GetTrafficPolicyInstanceCountOutput)

	ListGeoLocations(*route53.ListGeoLocationsInput) (*route53.ListGeoLocationsOutput, error)
	ListGeoLocationsWithContext(aws.Context, *route53.ListGeoLocationsInput, ...request.Option) (*route53.ListGeoLocationsOutput, error)
	ListGeoLocationsRequest(*route53.ListGeoLocationsInput) (*request.Request, *route53.ListGeoLocationsOutput)

	ListHealthChecks(*route53.ListHealthChecksInput) (*route53.ListHealthChecksOutput, error)
	ListHealthChecksWithContext(aws.Context, *route53.ListHealthChecksInput, ...request.Option) (*route53.ListHealthChecksOutput, error)
	ListHealthChecksRequest(*route53.ListHealthChecksInput) (*request.Request, *route53.ListHealthChecksOutput)

	ListHealthChecksPages(*route53.ListHealthChecksInput, func(*route53.ListHealthChecksOutput, bool) bool) error
	ListHealthChecksPagesWithContext(aws.Context, *route53.ListHealthChecksInput, func(*route53.ListHealthChecksOutput, bool) bool, ...request.Option) error

	ListHostedZones(*route53.ListHostedZonesInput) (*route53.ListHostedZonesOutput, error)
	ListHostedZonesWithContext(aws.Context, *route53.ListHostedZonesInput, ...request.Option) (*route53.ListHostedZonesOutput, error)
	ListHostedZonesRequest(*route53.ListHostedZonesInput) (*request.Request, *route53.ListHostedZonesOutput)

	ListHostedZonesPages(*route53.ListHostedZonesInput, func(*route53.ListHostedZonesOutput, bool) bool) error
	ListHostedZonesPagesWithContext(aws.Context, *route53.ListHostedZonesInput, func(*route53.ListHostedZonesOutput, bool) bool, ...request.Option) error

	ListHostedZonesByName(*route53.ListHostedZonesByNameInput) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesByNameWithContext(aws.Context, *route53.ListHostedZonesByNameInput, ...request.Option) (*route53.ListHostedZonesByNameOutput, error)
	ListHostedZonesByNameRequest(*route53.ListHostedZonesByNameInput) (*request.Request, *route53.ListHostedZonesByNameOutput)

	ListHostedZonesByVPC(*route53.ListHostedZonesByVPCInput) (*route53.ListHostedZonesByVPCOutput, error)
	ListHostedZonesByVPCWithContext(aws.Context, *route53.ListHostedZonesByVPCInput, ...request.Option) (*route53.ListHostedZonesByVPCOutput, error)
	ListHostedZonesByVPCRequest(*route53.ListHostedZonesByVPCInput) (*request.Request, *route53.ListHostedZonesByVPCOutput)

	ListQueryLoggingConfigs(*route53.ListQueryLoggingConfigsInput) (*route53.ListQueryLoggingConfigsOutput, error)
	ListQueryLoggingConfigsWithContext(aws.Context, *route53.ListQueryLoggingConfigsInput, ...request.Option) (*route53.ListQueryLoggingConfigsOutput, error)
	ListQueryLoggingConfigsRequest(*route53.ListQueryLoggingConfigsInput) (*request.Request, *route53.ListQueryLoggingConfigsOutput)

	ListQueryLoggingConfigsPages(*route53.ListQueryLoggingConfigsInput, func(*route53.ListQueryLoggingConfigsOutput, bool) bool) error
	ListQueryLoggingConfigsPagesWithContext(aws.Context, *route53.ListQueryLoggingConfigsInput, func(*route53.ListQueryLoggingConfigsOutput, bool) bool, ...request.Option) error

	ListResourceRecordSets(*route53.ListResourceRecordSetsInput) (*route53.ListResourceRecordSetsOutput, error)
	ListResourceRecordSetsWithContext(aws.Context, *route53.ListResourceRecordSetsInput, ...request.Option) (*route53.ListResourceRecordSetsOutput, error)
	ListResourceRecordSetsRequest(*route53.ListResourceRecordSetsInput) (*request.Request, *route53.ListResourceRecordSetsOutput)

	ListResourceRecordSetsPages(*route53.ListResourceRecordSetsInput, func(*route53.ListResourceRecordSetsOutput, bool) bool) error
	ListResourceRecordSetsPagesWithContext(aws.Context, *route53.ListResourceRecordSetsInput, func(*route53.ListResourceRecordSetsOutput, bool) bool, ...request.Option) error

	ListReusableDelegationSets(*route53.ListReusableDelegationSetsInput) (*route53.ListReusableDelegationSetsOutput, error)
	ListReusableDelegationSetsWithContext(aws.Context, *route53.ListReusableDelegationSetsInput, ...request.Option) (*route53.ListReusableDelegationSetsOutput, error)
	ListReusableDelegationSetsRequest(*route53.ListReusableDelegationSetsInput) (*request.Request, *route53.ListReusableDelegationSetsOutput)

	ListTagsForResource(*route53.ListTagsForResourceInput) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *route53.ListTagsForResourceInput, ...request.Option) (*route53.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*route53.ListTagsForResourceInput) (*request.Request, *route53.ListTagsForResourceOutput)

	ListTagsForResources(*route53.ListTagsForResourcesInput) (*route53.ListTagsForResourcesOutput, error)
	ListTagsForResourcesWithContext(aws.Context, *route53.ListTagsForResourcesInput, ...request.Option) (*route53.ListTagsForResourcesOutput, error)
	ListTagsForResourcesRequest(*route53.ListTagsForResourcesInput) (*request.Request, *route53.ListTagsForResourcesOutput)

	ListTrafficPolicies(*route53.ListTrafficPoliciesInput) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPoliciesWithContext(aws.Context, *route53.ListTrafficPoliciesInput, ...request.Option) (*route53.ListTrafficPoliciesOutput, error)
	ListTrafficPoliciesRequest(*route53.ListTrafficPoliciesInput) (*request.Request, *route53.ListTrafficPoliciesOutput)

	ListTrafficPolicyInstances(*route53.ListTrafficPolicyInstancesInput) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyInstancesWithContext(aws.Context, *route53.ListTrafficPolicyInstancesInput, ...request.Option) (*route53.ListTrafficPolicyInstancesOutput, error)
	ListTrafficPolicyInstancesRequest(*route53.ListTrafficPolicyInstancesInput) (*request.Request, *route53.ListTrafficPolicyInstancesOutput)

	ListTrafficPolicyInstancesByHostedZone(*route53.ListTrafficPolicyInstancesByHostedZoneInput) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByHostedZoneWithContext(aws.Context, *route53.ListTrafficPolicyInstancesByHostedZoneInput, ...request.Option) (*route53.ListTrafficPolicyInstancesByHostedZoneOutput, error)
	ListTrafficPolicyInstancesByHostedZoneRequest(*route53.ListTrafficPolicyInstancesByHostedZoneInput) (*request.Request, *route53.ListTrafficPolicyInstancesByHostedZoneOutput)

	ListTrafficPolicyInstancesByPolicy(*route53.ListTrafficPolicyInstancesByPolicyInput) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyInstancesByPolicyWithContext(aws.Context, *route53.ListTrafficPolicyInstancesByPolicyInput, ...request.Option) (*route53.ListTrafficPolicyInstancesByPolicyOutput, error)
	ListTrafficPolicyInstancesByPolicyRequest(*route53.ListTrafficPolicyInstancesByPolicyInput) (*request.Request, *route53.ListTrafficPolicyInstancesByPolicyOutput)

	ListTrafficPolicyVersions(*route53.ListTrafficPolicyVersionsInput) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListTrafficPolicyVersionsWithContext(aws.Context, *route53.ListTrafficPolicyVersionsInput, ...request.Option) (*route53.ListTrafficPolicyVersionsOutput, error)
	ListTrafficPolicyVersionsRequest(*route53.ListTrafficPolicyVersionsInput) (*request.Request, *route53.ListTrafficPolicyVersionsOutput)

	ListVPCAssociationAuthorizations(*route53.ListVPCAssociationAuthorizationsInput) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	ListVPCAssociationAuthorizationsWithContext(aws.Context, *route53.ListVPCAssociationAuthorizationsInput, ...request.Option) (*route53.ListVPCAssociationAuthorizationsOutput, error)
	ListVPCAssociationAuthorizationsRequest(*route53.ListVPCAssociationAuthorizationsInput) (*request.Request, *route53.ListVPCAssociationAuthorizationsOutput)

	TestDNSAnswer(*route53.TestDNSAnswerInput) (*route53.TestDNSAnswerOutput, error)
	TestDNSAnswerWithContext(aws.Context, *route53.TestDNSAnswerInput, ...request.Option) (*route53.TestDNSAnswerOutput, error)
	TestDNSAnswerRequest(*route53.TestDNSAnswerInput) (*request.Request, *route53.TestDNSAnswerOutput)

	UpdateHealthCheck(*route53.UpdateHealthCheckInput) (*route53.UpdateHealthCheckOutput, error)
	UpdateHealthCheckWithContext(aws.Context, *route53.UpdateHealthCheckInput, ...request.Option) (*route53.UpdateHealthCheckOutput, error)
	UpdateHealthCheckRequest(*route53.UpdateHealthCheckInput) (*request.Request, *route53.UpdateHealthCheckOutput)

	UpdateHostedZoneComment(*route53.UpdateHostedZoneCommentInput) (*route53.UpdateHostedZoneCommentOutput, error)
	UpdateHostedZoneCommentWithContext(aws.Context, *route53.UpdateHostedZoneCommentInput, ...request.Option) (*route53.UpdateHostedZoneCommentOutput, error)
	UpdateHostedZoneCommentRequest(*route53.UpdateHostedZoneCommentInput) (*request.Request, *route53.UpdateHostedZoneCommentOutput)

	UpdateTrafficPolicyComment(*route53.UpdateTrafficPolicyCommentInput) (*route53.UpdateTrafficPolicyCommentOutput, error)
	UpdateTrafficPolicyCommentWithContext(aws.Context, *route53.UpdateTrafficPolicyCommentInput, ...request.Option) (*route53.UpdateTrafficPolicyCommentOutput, error)
	UpdateTrafficPolicyCommentRequest(*route53.UpdateTrafficPolicyCommentInput) (*request.Request, *route53.UpdateTrafficPolicyCommentOutput)

	UpdateTrafficPolicyInstance(*route53.UpdateTrafficPolicyInstanceInput) (*route53.UpdateTrafficPolicyInstanceOutput, error)
	UpdateTrafficPolicyInstanceWithContext(aws.Context, *route53.UpdateTrafficPolicyInstanceInput, ...request.Option) (*route53.UpdateTrafficPolicyInstanceOutput, error)
	UpdateTrafficPolicyInstanceRequest(*route53.UpdateTrafficPolicyInstanceInput) (*request.Request, *route53.UpdateTrafficPolicyInstanceOutput)

	WaitUntilResourceRecordSetsChanged(*route53.GetChangeInput) error
	WaitUntilResourceRecordSetsChangedWithContext(aws.Context, *route53.GetChangeInput, ...request.WaiterOption) error
}

var _ Route53API = (*route53.Route53)(nil)
