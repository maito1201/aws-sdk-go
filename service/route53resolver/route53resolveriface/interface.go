// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package route53resolveriface provides an interface to enable mocking the Amazon Route 53 Resolver service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package route53resolveriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/route53resolver"
)

// Route53ResolverAPI provides an interface to enable mocking the
// route53resolver.Route53Resolver service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Route 53 Resolver.
//    func myFunc(svc route53resolveriface.Route53ResolverAPI) bool {
//        // Make svc.AssociateResolverEndpointIpAddress request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := route53resolver.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRoute53ResolverClient struct {
//        route53resolveriface.Route53ResolverAPI
//    }
//    func (m *mockRoute53ResolverClient) AssociateResolverEndpointIpAddress(input *route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRoute53ResolverClient{}
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
type Route53ResolverAPI interface {
	AssociateResolverEndpointIpAddress(*route53resolver.AssociateResolverEndpointIpAddressInput) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error)
	AssociateResolverEndpointIpAddressWithContext(aws.Context, *route53resolver.AssociateResolverEndpointIpAddressInput, ...request.Option) (*route53resolver.AssociateResolverEndpointIpAddressOutput, error)
	AssociateResolverEndpointIpAddressRequest(*route53resolver.AssociateResolverEndpointIpAddressInput) (*request.Request, *route53resolver.AssociateResolverEndpointIpAddressOutput)

	AssociateResolverQueryLogConfig(*route53resolver.AssociateResolverQueryLogConfigInput) (*route53resolver.AssociateResolverQueryLogConfigOutput, error)
	AssociateResolverQueryLogConfigWithContext(aws.Context, *route53resolver.AssociateResolverQueryLogConfigInput, ...request.Option) (*route53resolver.AssociateResolverQueryLogConfigOutput, error)
	AssociateResolverQueryLogConfigRequest(*route53resolver.AssociateResolverQueryLogConfigInput) (*request.Request, *route53resolver.AssociateResolverQueryLogConfigOutput)

	AssociateResolverRule(*route53resolver.AssociateResolverRuleInput) (*route53resolver.AssociateResolverRuleOutput, error)
	AssociateResolverRuleWithContext(aws.Context, *route53resolver.AssociateResolverRuleInput, ...request.Option) (*route53resolver.AssociateResolverRuleOutput, error)
	AssociateResolverRuleRequest(*route53resolver.AssociateResolverRuleInput) (*request.Request, *route53resolver.AssociateResolverRuleOutput)

	CreateResolverEndpoint(*route53resolver.CreateResolverEndpointInput) (*route53resolver.CreateResolverEndpointOutput, error)
	CreateResolverEndpointWithContext(aws.Context, *route53resolver.CreateResolverEndpointInput, ...request.Option) (*route53resolver.CreateResolverEndpointOutput, error)
	CreateResolverEndpointRequest(*route53resolver.CreateResolverEndpointInput) (*request.Request, *route53resolver.CreateResolverEndpointOutput)

	CreateResolverQueryLogConfig(*route53resolver.CreateResolverQueryLogConfigInput) (*route53resolver.CreateResolverQueryLogConfigOutput, error)
	CreateResolverQueryLogConfigWithContext(aws.Context, *route53resolver.CreateResolverQueryLogConfigInput, ...request.Option) (*route53resolver.CreateResolverQueryLogConfigOutput, error)
	CreateResolverQueryLogConfigRequest(*route53resolver.CreateResolverQueryLogConfigInput) (*request.Request, *route53resolver.CreateResolverQueryLogConfigOutput)

	CreateResolverRule(*route53resolver.CreateResolverRuleInput) (*route53resolver.CreateResolverRuleOutput, error)
	CreateResolverRuleWithContext(aws.Context, *route53resolver.CreateResolverRuleInput, ...request.Option) (*route53resolver.CreateResolverRuleOutput, error)
	CreateResolverRuleRequest(*route53resolver.CreateResolverRuleInput) (*request.Request, *route53resolver.CreateResolverRuleOutput)

	DeleteResolverEndpoint(*route53resolver.DeleteResolverEndpointInput) (*route53resolver.DeleteResolverEndpointOutput, error)
	DeleteResolverEndpointWithContext(aws.Context, *route53resolver.DeleteResolverEndpointInput, ...request.Option) (*route53resolver.DeleteResolverEndpointOutput, error)
	DeleteResolverEndpointRequest(*route53resolver.DeleteResolverEndpointInput) (*request.Request, *route53resolver.DeleteResolverEndpointOutput)

	DeleteResolverQueryLogConfig(*route53resolver.DeleteResolverQueryLogConfigInput) (*route53resolver.DeleteResolverQueryLogConfigOutput, error)
	DeleteResolverQueryLogConfigWithContext(aws.Context, *route53resolver.DeleteResolverQueryLogConfigInput, ...request.Option) (*route53resolver.DeleteResolverQueryLogConfigOutput, error)
	DeleteResolverQueryLogConfigRequest(*route53resolver.DeleteResolverQueryLogConfigInput) (*request.Request, *route53resolver.DeleteResolverQueryLogConfigOutput)

	DeleteResolverRule(*route53resolver.DeleteResolverRuleInput) (*route53resolver.DeleteResolverRuleOutput, error)
	DeleteResolverRuleWithContext(aws.Context, *route53resolver.DeleteResolverRuleInput, ...request.Option) (*route53resolver.DeleteResolverRuleOutput, error)
	DeleteResolverRuleRequest(*route53resolver.DeleteResolverRuleInput) (*request.Request, *route53resolver.DeleteResolverRuleOutput)

	DisassociateResolverEndpointIpAddress(*route53resolver.DisassociateResolverEndpointIpAddressInput) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error)
	DisassociateResolverEndpointIpAddressWithContext(aws.Context, *route53resolver.DisassociateResolverEndpointIpAddressInput, ...request.Option) (*route53resolver.DisassociateResolverEndpointIpAddressOutput, error)
	DisassociateResolverEndpointIpAddressRequest(*route53resolver.DisassociateResolverEndpointIpAddressInput) (*request.Request, *route53resolver.DisassociateResolverEndpointIpAddressOutput)

	DisassociateResolverQueryLogConfig(*route53resolver.DisassociateResolverQueryLogConfigInput) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error)
	DisassociateResolverQueryLogConfigWithContext(aws.Context, *route53resolver.DisassociateResolverQueryLogConfigInput, ...request.Option) (*route53resolver.DisassociateResolverQueryLogConfigOutput, error)
	DisassociateResolverQueryLogConfigRequest(*route53resolver.DisassociateResolverQueryLogConfigInput) (*request.Request, *route53resolver.DisassociateResolverQueryLogConfigOutput)

	DisassociateResolverRule(*route53resolver.DisassociateResolverRuleInput) (*route53resolver.DisassociateResolverRuleOutput, error)
	DisassociateResolverRuleWithContext(aws.Context, *route53resolver.DisassociateResolverRuleInput, ...request.Option) (*route53resolver.DisassociateResolverRuleOutput, error)
	DisassociateResolverRuleRequest(*route53resolver.DisassociateResolverRuleInput) (*request.Request, *route53resolver.DisassociateResolverRuleOutput)

	GetResolverEndpoint(*route53resolver.GetResolverEndpointInput) (*route53resolver.GetResolverEndpointOutput, error)
	GetResolverEndpointWithContext(aws.Context, *route53resolver.GetResolverEndpointInput, ...request.Option) (*route53resolver.GetResolverEndpointOutput, error)
	GetResolverEndpointRequest(*route53resolver.GetResolverEndpointInput) (*request.Request, *route53resolver.GetResolverEndpointOutput)

	GetResolverQueryLogConfig(*route53resolver.GetResolverQueryLogConfigInput) (*route53resolver.GetResolverQueryLogConfigOutput, error)
	GetResolverQueryLogConfigWithContext(aws.Context, *route53resolver.GetResolverQueryLogConfigInput, ...request.Option) (*route53resolver.GetResolverQueryLogConfigOutput, error)
	GetResolverQueryLogConfigRequest(*route53resolver.GetResolverQueryLogConfigInput) (*request.Request, *route53resolver.GetResolverQueryLogConfigOutput)

	GetResolverQueryLogConfigAssociation(*route53resolver.GetResolverQueryLogConfigAssociationInput) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error)
	GetResolverQueryLogConfigAssociationWithContext(aws.Context, *route53resolver.GetResolverQueryLogConfigAssociationInput, ...request.Option) (*route53resolver.GetResolverQueryLogConfigAssociationOutput, error)
	GetResolverQueryLogConfigAssociationRequest(*route53resolver.GetResolverQueryLogConfigAssociationInput) (*request.Request, *route53resolver.GetResolverQueryLogConfigAssociationOutput)

	GetResolverQueryLogConfigPolicy(*route53resolver.GetResolverQueryLogConfigPolicyInput) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error)
	GetResolverQueryLogConfigPolicyWithContext(aws.Context, *route53resolver.GetResolverQueryLogConfigPolicyInput, ...request.Option) (*route53resolver.GetResolverQueryLogConfigPolicyOutput, error)
	GetResolverQueryLogConfigPolicyRequest(*route53resolver.GetResolverQueryLogConfigPolicyInput) (*request.Request, *route53resolver.GetResolverQueryLogConfigPolicyOutput)

	GetResolverRule(*route53resolver.GetResolverRuleInput) (*route53resolver.GetResolverRuleOutput, error)
	GetResolverRuleWithContext(aws.Context, *route53resolver.GetResolverRuleInput, ...request.Option) (*route53resolver.GetResolverRuleOutput, error)
	GetResolverRuleRequest(*route53resolver.GetResolverRuleInput) (*request.Request, *route53resolver.GetResolverRuleOutput)

	GetResolverRuleAssociation(*route53resolver.GetResolverRuleAssociationInput) (*route53resolver.GetResolverRuleAssociationOutput, error)
	GetResolverRuleAssociationWithContext(aws.Context, *route53resolver.GetResolverRuleAssociationInput, ...request.Option) (*route53resolver.GetResolverRuleAssociationOutput, error)
	GetResolverRuleAssociationRequest(*route53resolver.GetResolverRuleAssociationInput) (*request.Request, *route53resolver.GetResolverRuleAssociationOutput)

	GetResolverRulePolicy(*route53resolver.GetResolverRulePolicyInput) (*route53resolver.GetResolverRulePolicyOutput, error)
	GetResolverRulePolicyWithContext(aws.Context, *route53resolver.GetResolverRulePolicyInput, ...request.Option) (*route53resolver.GetResolverRulePolicyOutput, error)
	GetResolverRulePolicyRequest(*route53resolver.GetResolverRulePolicyInput) (*request.Request, *route53resolver.GetResolverRulePolicyOutput)

	ListResolverEndpointIpAddresses(*route53resolver.ListResolverEndpointIpAddressesInput) (*route53resolver.ListResolverEndpointIpAddressesOutput, error)
	ListResolverEndpointIpAddressesWithContext(aws.Context, *route53resolver.ListResolverEndpointIpAddressesInput, ...request.Option) (*route53resolver.ListResolverEndpointIpAddressesOutput, error)
	ListResolverEndpointIpAddressesRequest(*route53resolver.ListResolverEndpointIpAddressesInput) (*request.Request, *route53resolver.ListResolverEndpointIpAddressesOutput)

	ListResolverEndpointIpAddressesPages(*route53resolver.ListResolverEndpointIpAddressesInput, func(*route53resolver.ListResolverEndpointIpAddressesOutput, bool) bool) error
	ListResolverEndpointIpAddressesPagesWithContext(aws.Context, *route53resolver.ListResolverEndpointIpAddressesInput, func(*route53resolver.ListResolverEndpointIpAddressesOutput, bool) bool, ...request.Option) error

	ListResolverEndpoints(*route53resolver.ListResolverEndpointsInput) (*route53resolver.ListResolverEndpointsOutput, error)
	ListResolverEndpointsWithContext(aws.Context, *route53resolver.ListResolverEndpointsInput, ...request.Option) (*route53resolver.ListResolverEndpointsOutput, error)
	ListResolverEndpointsRequest(*route53resolver.ListResolverEndpointsInput) (*request.Request, *route53resolver.ListResolverEndpointsOutput)

	ListResolverEndpointsPages(*route53resolver.ListResolverEndpointsInput, func(*route53resolver.ListResolverEndpointsOutput, bool) bool) error
	ListResolverEndpointsPagesWithContext(aws.Context, *route53resolver.ListResolverEndpointsInput, func(*route53resolver.ListResolverEndpointsOutput, bool) bool, ...request.Option) error

	ListResolverQueryLogConfigAssociations(*route53resolver.ListResolverQueryLogConfigAssociationsInput) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error)
	ListResolverQueryLogConfigAssociationsWithContext(aws.Context, *route53resolver.ListResolverQueryLogConfigAssociationsInput, ...request.Option) (*route53resolver.ListResolverQueryLogConfigAssociationsOutput, error)
	ListResolverQueryLogConfigAssociationsRequest(*route53resolver.ListResolverQueryLogConfigAssociationsInput) (*request.Request, *route53resolver.ListResolverQueryLogConfigAssociationsOutput)

	ListResolverQueryLogConfigAssociationsPages(*route53resolver.ListResolverQueryLogConfigAssociationsInput, func(*route53resolver.ListResolverQueryLogConfigAssociationsOutput, bool) bool) error
	ListResolverQueryLogConfigAssociationsPagesWithContext(aws.Context, *route53resolver.ListResolverQueryLogConfigAssociationsInput, func(*route53resolver.ListResolverQueryLogConfigAssociationsOutput, bool) bool, ...request.Option) error

	ListResolverQueryLogConfigs(*route53resolver.ListResolverQueryLogConfigsInput) (*route53resolver.ListResolverQueryLogConfigsOutput, error)
	ListResolverQueryLogConfigsWithContext(aws.Context, *route53resolver.ListResolverQueryLogConfigsInput, ...request.Option) (*route53resolver.ListResolverQueryLogConfigsOutput, error)
	ListResolverQueryLogConfigsRequest(*route53resolver.ListResolverQueryLogConfigsInput) (*request.Request, *route53resolver.ListResolverQueryLogConfigsOutput)

	ListResolverQueryLogConfigsPages(*route53resolver.ListResolverQueryLogConfigsInput, func(*route53resolver.ListResolverQueryLogConfigsOutput, bool) bool) error
	ListResolverQueryLogConfigsPagesWithContext(aws.Context, *route53resolver.ListResolverQueryLogConfigsInput, func(*route53resolver.ListResolverQueryLogConfigsOutput, bool) bool, ...request.Option) error

	ListResolverRuleAssociations(*route53resolver.ListResolverRuleAssociationsInput) (*route53resolver.ListResolverRuleAssociationsOutput, error)
	ListResolverRuleAssociationsWithContext(aws.Context, *route53resolver.ListResolverRuleAssociationsInput, ...request.Option) (*route53resolver.ListResolverRuleAssociationsOutput, error)
	ListResolverRuleAssociationsRequest(*route53resolver.ListResolverRuleAssociationsInput) (*request.Request, *route53resolver.ListResolverRuleAssociationsOutput)

	ListResolverRuleAssociationsPages(*route53resolver.ListResolverRuleAssociationsInput, func(*route53resolver.ListResolverRuleAssociationsOutput, bool) bool) error
	ListResolverRuleAssociationsPagesWithContext(aws.Context, *route53resolver.ListResolverRuleAssociationsInput, func(*route53resolver.ListResolverRuleAssociationsOutput, bool) bool, ...request.Option) error

	ListResolverRules(*route53resolver.ListResolverRulesInput) (*route53resolver.ListResolverRulesOutput, error)
	ListResolverRulesWithContext(aws.Context, *route53resolver.ListResolverRulesInput, ...request.Option) (*route53resolver.ListResolverRulesOutput, error)
	ListResolverRulesRequest(*route53resolver.ListResolverRulesInput) (*request.Request, *route53resolver.ListResolverRulesOutput)

	ListResolverRulesPages(*route53resolver.ListResolverRulesInput, func(*route53resolver.ListResolverRulesOutput, bool) bool) error
	ListResolverRulesPagesWithContext(aws.Context, *route53resolver.ListResolverRulesInput, func(*route53resolver.ListResolverRulesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*route53resolver.ListTagsForResourceInput) (*route53resolver.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *route53resolver.ListTagsForResourceInput, ...request.Option) (*route53resolver.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*route53resolver.ListTagsForResourceInput) (*request.Request, *route53resolver.ListTagsForResourceOutput)

	ListTagsForResourcePages(*route53resolver.ListTagsForResourceInput, func(*route53resolver.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *route53resolver.ListTagsForResourceInput, func(*route53resolver.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	PutResolverQueryLogConfigPolicy(*route53resolver.PutResolverQueryLogConfigPolicyInput) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error)
	PutResolverQueryLogConfigPolicyWithContext(aws.Context, *route53resolver.PutResolverQueryLogConfigPolicyInput, ...request.Option) (*route53resolver.PutResolverQueryLogConfigPolicyOutput, error)
	PutResolverQueryLogConfigPolicyRequest(*route53resolver.PutResolverQueryLogConfigPolicyInput) (*request.Request, *route53resolver.PutResolverQueryLogConfigPolicyOutput)

	PutResolverRulePolicy(*route53resolver.PutResolverRulePolicyInput) (*route53resolver.PutResolverRulePolicyOutput, error)
	PutResolverRulePolicyWithContext(aws.Context, *route53resolver.PutResolverRulePolicyInput, ...request.Option) (*route53resolver.PutResolverRulePolicyOutput, error)
	PutResolverRulePolicyRequest(*route53resolver.PutResolverRulePolicyInput) (*request.Request, *route53resolver.PutResolverRulePolicyOutput)

	TagResource(*route53resolver.TagResourceInput) (*route53resolver.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *route53resolver.TagResourceInput, ...request.Option) (*route53resolver.TagResourceOutput, error)
	TagResourceRequest(*route53resolver.TagResourceInput) (*request.Request, *route53resolver.TagResourceOutput)

	UntagResource(*route53resolver.UntagResourceInput) (*route53resolver.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *route53resolver.UntagResourceInput, ...request.Option) (*route53resolver.UntagResourceOutput, error)
	UntagResourceRequest(*route53resolver.UntagResourceInput) (*request.Request, *route53resolver.UntagResourceOutput)

	UpdateResolverEndpoint(*route53resolver.UpdateResolverEndpointInput) (*route53resolver.UpdateResolverEndpointOutput, error)
	UpdateResolverEndpointWithContext(aws.Context, *route53resolver.UpdateResolverEndpointInput, ...request.Option) (*route53resolver.UpdateResolverEndpointOutput, error)
	UpdateResolverEndpointRequest(*route53resolver.UpdateResolverEndpointInput) (*request.Request, *route53resolver.UpdateResolverEndpointOutput)

	UpdateResolverRule(*route53resolver.UpdateResolverRuleInput) (*route53resolver.UpdateResolverRuleOutput, error)
	UpdateResolverRuleWithContext(aws.Context, *route53resolver.UpdateResolverRuleInput, ...request.Option) (*route53resolver.UpdateResolverRuleOutput, error)
	UpdateResolverRuleRequest(*route53resolver.UpdateResolverRuleInput) (*request.Request, *route53resolver.UpdateResolverRuleOutput)
}

var _ Route53ResolverAPI = (*route53resolver.Route53Resolver)(nil)
