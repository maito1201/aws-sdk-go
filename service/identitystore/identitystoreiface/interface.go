// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package identitystoreiface provides an interface to enable mocking the AWS SSO Identity Store service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package identitystoreiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/identitystore"
)

// IdentityStoreAPI provides an interface to enable mocking the
// identitystore.IdentityStore service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS SSO Identity Store.
//    func myFunc(svc identitystoreiface.IdentityStoreAPI) bool {
//        // Make svc.DescribeGroup request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := identitystore.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIdentityStoreClient struct {
//        identitystoreiface.IdentityStoreAPI
//    }
//    func (m *mockIdentityStoreClient) DescribeGroup(input *identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIdentityStoreClient{}
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
type IdentityStoreAPI interface {
	DescribeGroup(*identitystore.DescribeGroupInput) (*identitystore.DescribeGroupOutput, error)
	DescribeGroupWithContext(aws.Context, *identitystore.DescribeGroupInput, ...request.Option) (*identitystore.DescribeGroupOutput, error)
	DescribeGroupRequest(*identitystore.DescribeGroupInput) (*request.Request, *identitystore.DescribeGroupOutput)

	DescribeUser(*identitystore.DescribeUserInput) (*identitystore.DescribeUserOutput, error)
	DescribeUserWithContext(aws.Context, *identitystore.DescribeUserInput, ...request.Option) (*identitystore.DescribeUserOutput, error)
	DescribeUserRequest(*identitystore.DescribeUserInput) (*request.Request, *identitystore.DescribeUserOutput)

	ListGroups(*identitystore.ListGroupsInput) (*identitystore.ListGroupsOutput, error)
	ListGroupsWithContext(aws.Context, *identitystore.ListGroupsInput, ...request.Option) (*identitystore.ListGroupsOutput, error)
	ListGroupsRequest(*identitystore.ListGroupsInput) (*request.Request, *identitystore.ListGroupsOutput)

	ListGroupsPages(*identitystore.ListGroupsInput, func(*identitystore.ListGroupsOutput, bool) bool) error
	ListGroupsPagesWithContext(aws.Context, *identitystore.ListGroupsInput, func(*identitystore.ListGroupsOutput, bool) bool, ...request.Option) error

	ListUsers(*identitystore.ListUsersInput) (*identitystore.ListUsersOutput, error)
	ListUsersWithContext(aws.Context, *identitystore.ListUsersInput, ...request.Option) (*identitystore.ListUsersOutput, error)
	ListUsersRequest(*identitystore.ListUsersInput) (*request.Request, *identitystore.ListUsersOutput)

	ListUsersPages(*identitystore.ListUsersInput, func(*identitystore.ListUsersOutput, bool) bool) error
	ListUsersPagesWithContext(aws.Context, *identitystore.ListUsersInput, func(*identitystore.ListUsersOutput, bool) bool, ...request.Option) error
}

var _ IdentityStoreAPI = (*identitystore.IdentityStore)(nil)
