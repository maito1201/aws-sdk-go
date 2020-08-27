// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotsecuretunnelingiface provides an interface to enable mocking the AWS IoT Secure Tunneling service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotsecuretunnelingiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotsecuretunneling"
)

// IoTSecureTunnelingAPI provides an interface to enable mocking the
// iotsecuretunneling.IoTSecureTunneling service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Secure Tunneling.
//    func myFunc(svc iotsecuretunnelingiface.IoTSecureTunnelingAPI) bool {
//        // Make svc.CloseTunnel request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := iotsecuretunneling.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTSecureTunnelingClient struct {
//        iotsecuretunnelingiface.IoTSecureTunnelingAPI
//    }
//    func (m *mockIoTSecureTunnelingClient) CloseTunnel(input *iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTSecureTunnelingClient{}
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
type IoTSecureTunnelingAPI interface {
	CloseTunnel(*iotsecuretunneling.CloseTunnelInput) (*iotsecuretunneling.CloseTunnelOutput, error)
	CloseTunnelWithContext(aws.Context, *iotsecuretunneling.CloseTunnelInput, ...request.Option) (*iotsecuretunneling.CloseTunnelOutput, error)
	CloseTunnelRequest(*iotsecuretunneling.CloseTunnelInput) (*request.Request, *iotsecuretunneling.CloseTunnelOutput)

	DescribeTunnel(*iotsecuretunneling.DescribeTunnelInput) (*iotsecuretunneling.DescribeTunnelOutput, error)
	DescribeTunnelWithContext(aws.Context, *iotsecuretunneling.DescribeTunnelInput, ...request.Option) (*iotsecuretunneling.DescribeTunnelOutput, error)
	DescribeTunnelRequest(*iotsecuretunneling.DescribeTunnelInput) (*request.Request, *iotsecuretunneling.DescribeTunnelOutput)

	ListTagsForResource(*iotsecuretunneling.ListTagsForResourceInput) (*iotsecuretunneling.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotsecuretunneling.ListTagsForResourceInput, ...request.Option) (*iotsecuretunneling.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotsecuretunneling.ListTagsForResourceInput) (*request.Request, *iotsecuretunneling.ListTagsForResourceOutput)

	ListTunnels(*iotsecuretunneling.ListTunnelsInput) (*iotsecuretunneling.ListTunnelsOutput, error)
	ListTunnelsWithContext(aws.Context, *iotsecuretunneling.ListTunnelsInput, ...request.Option) (*iotsecuretunneling.ListTunnelsOutput, error)
	ListTunnelsRequest(*iotsecuretunneling.ListTunnelsInput) (*request.Request, *iotsecuretunneling.ListTunnelsOutput)

	ListTunnelsPages(*iotsecuretunneling.ListTunnelsInput, func(*iotsecuretunneling.ListTunnelsOutput, bool) bool) error
	ListTunnelsPagesWithContext(aws.Context, *iotsecuretunneling.ListTunnelsInput, func(*iotsecuretunneling.ListTunnelsOutput, bool) bool, ...request.Option) error

	OpenTunnel(*iotsecuretunneling.OpenTunnelInput) (*iotsecuretunneling.OpenTunnelOutput, error)
	OpenTunnelWithContext(aws.Context, *iotsecuretunneling.OpenTunnelInput, ...request.Option) (*iotsecuretunneling.OpenTunnelOutput, error)
	OpenTunnelRequest(*iotsecuretunneling.OpenTunnelInput) (*request.Request, *iotsecuretunneling.OpenTunnelOutput)

	TagResource(*iotsecuretunneling.TagResourceInput) (*iotsecuretunneling.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotsecuretunneling.TagResourceInput, ...request.Option) (*iotsecuretunneling.TagResourceOutput, error)
	TagResourceRequest(*iotsecuretunneling.TagResourceInput) (*request.Request, *iotsecuretunneling.TagResourceOutput)

	UntagResource(*iotsecuretunneling.UntagResourceInput) (*iotsecuretunneling.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotsecuretunneling.UntagResourceInput, ...request.Option) (*iotsecuretunneling.UntagResourceOutput, error)
	UntagResourceRequest(*iotsecuretunneling.UntagResourceInput) (*request.Request, *iotsecuretunneling.UntagResourceOutput)
}

var _ IoTSecureTunnelingAPI = (*iotsecuretunneling.IoTSecureTunneling)(nil)
