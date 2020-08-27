// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package kinesisvideoarchivedmediaiface provides an interface to enable mocking the Amazon Kinesis Video Streams Archived Media service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package kinesisvideoarchivedmediaiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/kinesisvideoarchivedmedia"
)

// KinesisVideoArchivedMediaAPI provides an interface to enable mocking the
// kinesisvideoarchivedmedia.KinesisVideoArchivedMedia service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Kinesis Video Streams Archived Media.
//    func myFunc(svc kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI) bool {
//        // Make svc.GetClip request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := kinesisvideoarchivedmedia.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockKinesisVideoArchivedMediaClient struct {
//        kinesisvideoarchivedmediaiface.KinesisVideoArchivedMediaAPI
//    }
//    func (m *mockKinesisVideoArchivedMediaClient) GetClip(input *kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockKinesisVideoArchivedMediaClient{}
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
type KinesisVideoArchivedMediaAPI interface {
	GetClip(*kinesisvideoarchivedmedia.GetClipInput) (*kinesisvideoarchivedmedia.GetClipOutput, error)
	GetClipWithContext(aws.Context, *kinesisvideoarchivedmedia.GetClipInput, ...request.Option) (*kinesisvideoarchivedmedia.GetClipOutput, error)
	GetClipRequest(*kinesisvideoarchivedmedia.GetClipInput) (*request.Request, *kinesisvideoarchivedmedia.GetClipOutput)

	GetDASHStreamingSessionURL(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error)
	GetDASHStreamingSessionURLWithContext(aws.Context, *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput, ...request.Option) (*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput, error)
	GetDASHStreamingSessionURLRequest(*kinesisvideoarchivedmedia.GetDASHStreamingSessionURLInput) (*request.Request, *kinesisvideoarchivedmedia.GetDASHStreamingSessionURLOutput)

	GetHLSStreamingSessionURL(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error)
	GetHLSStreamingSessionURLWithContext(aws.Context, *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput, ...request.Option) (*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput, error)
	GetHLSStreamingSessionURLRequest(*kinesisvideoarchivedmedia.GetHLSStreamingSessionURLInput) (*request.Request, *kinesisvideoarchivedmedia.GetHLSStreamingSessionURLOutput)

	GetMediaForFragmentList(*kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error)
	GetMediaForFragmentListWithContext(aws.Context, *kinesisvideoarchivedmedia.GetMediaForFragmentListInput, ...request.Option) (*kinesisvideoarchivedmedia.GetMediaForFragmentListOutput, error)
	GetMediaForFragmentListRequest(*kinesisvideoarchivedmedia.GetMediaForFragmentListInput) (*request.Request, *kinesisvideoarchivedmedia.GetMediaForFragmentListOutput)

	ListFragments(*kinesisvideoarchivedmedia.ListFragmentsInput) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error)
	ListFragmentsWithContext(aws.Context, *kinesisvideoarchivedmedia.ListFragmentsInput, ...request.Option) (*kinesisvideoarchivedmedia.ListFragmentsOutput, error)
	ListFragmentsRequest(*kinesisvideoarchivedmedia.ListFragmentsInput) (*request.Request, *kinesisvideoarchivedmedia.ListFragmentsOutput)

	ListFragmentsPages(*kinesisvideoarchivedmedia.ListFragmentsInput, func(*kinesisvideoarchivedmedia.ListFragmentsOutput, bool) bool) error
	ListFragmentsPagesWithContext(aws.Context, *kinesisvideoarchivedmedia.ListFragmentsInput, func(*kinesisvideoarchivedmedia.ListFragmentsOutput, bool) bool, ...request.Option) error
}

var _ KinesisVideoArchivedMediaAPI = (*kinesisvideoarchivedmedia.KinesisVideoArchivedMedia)(nil)
