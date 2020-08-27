// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisvideomediaiface provides an interface to enable mocking the Amazon Kinesis Video Streams Media service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisvideomediaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideomedia"
)

// KinesisVideoMediaAPI provides an interface to enable mocking the
// kinesisvideomedia.KinesisVideoMedia service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Video Streams Media.
//    func myFunc(svc kinesisvideomediaiface.KinesisVideoMediaAPI) bool {
//        // Make svc.GetMedia request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := kinesisvideomedia.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisVideoMediaClient struct {
//        kinesisvideomediaiface.KinesisVideoMediaAPI
//    }
//    func (m *mockKinesisVideoMediaClient) GetMedia(input *kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisVideoMediaClient{}
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
type KinesisVideoMediaAPI interface {
	GetMedia(*kinesisvideomedia.GetMediaInput) (*kinesisvideomedia.GetMediaOutput, error)
	GetMediaWithContext(aws.Context, *kinesisvideomedia.GetMediaInput, ...request.Option) (*kinesisvideomedia.GetMediaOutput, error)
	GetMediaRequest(*kinesisvideomedia.GetMediaInput) (*request.Request, *kinesisvideomedia.GetMediaOutput)
}

var _ KinesisVideoMediaAPI = (*kinesisvideomedia.KinesisVideoMedia)(nil)
