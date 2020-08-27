// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package iotanalyticsiface provides an interface to enable mocking the AWS IoT Analytics service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package iotanalyticsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/iotanalytics"
)

// IoTAnalyticsAPI provides an interface to enable mocking the
// iotanalytics.IoTAnalytics service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS IoT Analytics.
//    func myFunc(svc iotanalyticsiface.IoTAnalyticsAPI) bool {
//        // Make svc.BatchPutMessage request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := iotanalytics.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockIoTAnalyticsClient struct {
//        iotanalyticsiface.IoTAnalyticsAPI
//    }
//    func (m *mockIoTAnalyticsClient) BatchPutMessage(input *iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockIoTAnalyticsClient{}
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
type IoTAnalyticsAPI interface {
	BatchPutMessage(*iotanalytics.BatchPutMessageInput) (*iotanalytics.BatchPutMessageOutput, error)
	BatchPutMessageWithContext(aws.Context, *iotanalytics.BatchPutMessageInput, ...request.Option) (*iotanalytics.BatchPutMessageOutput, error)
	BatchPutMessageRequest(*iotanalytics.BatchPutMessageInput) (*request.Request, *iotanalytics.BatchPutMessageOutput)

	CancelPipelineReprocessing(*iotanalytics.CancelPipelineReprocessingInput) (*iotanalytics.CancelPipelineReprocessingOutput, error)
	CancelPipelineReprocessingWithContext(aws.Context, *iotanalytics.CancelPipelineReprocessingInput, ...request.Option) (*iotanalytics.CancelPipelineReprocessingOutput, error)
	CancelPipelineReprocessingRequest(*iotanalytics.CancelPipelineReprocessingInput) (*request.Request, *iotanalytics.CancelPipelineReprocessingOutput)

	CreateChannel(*iotanalytics.CreateChannelInput) (*iotanalytics.CreateChannelOutput, error)
	CreateChannelWithContext(aws.Context, *iotanalytics.CreateChannelInput, ...request.Option) (*iotanalytics.CreateChannelOutput, error)
	CreateChannelRequest(*iotanalytics.CreateChannelInput) (*request.Request, *iotanalytics.CreateChannelOutput)

	CreateDataset(*iotanalytics.CreateDatasetInput) (*iotanalytics.CreateDatasetOutput, error)
	CreateDatasetWithContext(aws.Context, *iotanalytics.CreateDatasetInput, ...request.Option) (*iotanalytics.CreateDatasetOutput, error)
	CreateDatasetRequest(*iotanalytics.CreateDatasetInput) (*request.Request, *iotanalytics.CreateDatasetOutput)

	CreateDatasetContent(*iotanalytics.CreateDatasetContentInput) (*iotanalytics.CreateDatasetContentOutput, error)
	CreateDatasetContentWithContext(aws.Context, *iotanalytics.CreateDatasetContentInput, ...request.Option) (*iotanalytics.CreateDatasetContentOutput, error)
	CreateDatasetContentRequest(*iotanalytics.CreateDatasetContentInput) (*request.Request, *iotanalytics.CreateDatasetContentOutput)

	CreateDatastore(*iotanalytics.CreateDatastoreInput) (*iotanalytics.CreateDatastoreOutput, error)
	CreateDatastoreWithContext(aws.Context, *iotanalytics.CreateDatastoreInput, ...request.Option) (*iotanalytics.CreateDatastoreOutput, error)
	CreateDatastoreRequest(*iotanalytics.CreateDatastoreInput) (*request.Request, *iotanalytics.CreateDatastoreOutput)

	CreatePipeline(*iotanalytics.CreatePipelineInput) (*iotanalytics.CreatePipelineOutput, error)
	CreatePipelineWithContext(aws.Context, *iotanalytics.CreatePipelineInput, ...request.Option) (*iotanalytics.CreatePipelineOutput, error)
	CreatePipelineRequest(*iotanalytics.CreatePipelineInput) (*request.Request, *iotanalytics.CreatePipelineOutput)

	DeleteChannel(*iotanalytics.DeleteChannelInput) (*iotanalytics.DeleteChannelOutput, error)
	DeleteChannelWithContext(aws.Context, *iotanalytics.DeleteChannelInput, ...request.Option) (*iotanalytics.DeleteChannelOutput, error)
	DeleteChannelRequest(*iotanalytics.DeleteChannelInput) (*request.Request, *iotanalytics.DeleteChannelOutput)

	DeleteDataset(*iotanalytics.DeleteDatasetInput) (*iotanalytics.DeleteDatasetOutput, error)
	DeleteDatasetWithContext(aws.Context, *iotanalytics.DeleteDatasetInput, ...request.Option) (*iotanalytics.DeleteDatasetOutput, error)
	DeleteDatasetRequest(*iotanalytics.DeleteDatasetInput) (*request.Request, *iotanalytics.DeleteDatasetOutput)

	DeleteDatasetContent(*iotanalytics.DeleteDatasetContentInput) (*iotanalytics.DeleteDatasetContentOutput, error)
	DeleteDatasetContentWithContext(aws.Context, *iotanalytics.DeleteDatasetContentInput, ...request.Option) (*iotanalytics.DeleteDatasetContentOutput, error)
	DeleteDatasetContentRequest(*iotanalytics.DeleteDatasetContentInput) (*request.Request, *iotanalytics.DeleteDatasetContentOutput)

	DeleteDatastore(*iotanalytics.DeleteDatastoreInput) (*iotanalytics.DeleteDatastoreOutput, error)
	DeleteDatastoreWithContext(aws.Context, *iotanalytics.DeleteDatastoreInput, ...request.Option) (*iotanalytics.DeleteDatastoreOutput, error)
	DeleteDatastoreRequest(*iotanalytics.DeleteDatastoreInput) (*request.Request, *iotanalytics.DeleteDatastoreOutput)

	DeletePipeline(*iotanalytics.DeletePipelineInput) (*iotanalytics.DeletePipelineOutput, error)
	DeletePipelineWithContext(aws.Context, *iotanalytics.DeletePipelineInput, ...request.Option) (*iotanalytics.DeletePipelineOutput, error)
	DeletePipelineRequest(*iotanalytics.DeletePipelineInput) (*request.Request, *iotanalytics.DeletePipelineOutput)

	DescribeChannel(*iotanalytics.DescribeChannelInput) (*iotanalytics.DescribeChannelOutput, error)
	DescribeChannelWithContext(aws.Context, *iotanalytics.DescribeChannelInput, ...request.Option) (*iotanalytics.DescribeChannelOutput, error)
	DescribeChannelRequest(*iotanalytics.DescribeChannelInput) (*request.Request, *iotanalytics.DescribeChannelOutput)

	DescribeDataset(*iotanalytics.DescribeDatasetInput) (*iotanalytics.DescribeDatasetOutput, error)
	DescribeDatasetWithContext(aws.Context, *iotanalytics.DescribeDatasetInput, ...request.Option) (*iotanalytics.DescribeDatasetOutput, error)
	DescribeDatasetRequest(*iotanalytics.DescribeDatasetInput) (*request.Request, *iotanalytics.DescribeDatasetOutput)

	DescribeDatastore(*iotanalytics.DescribeDatastoreInput) (*iotanalytics.DescribeDatastoreOutput, error)
	DescribeDatastoreWithContext(aws.Context, *iotanalytics.DescribeDatastoreInput, ...request.Option) (*iotanalytics.DescribeDatastoreOutput, error)
	DescribeDatastoreRequest(*iotanalytics.DescribeDatastoreInput) (*request.Request, *iotanalytics.DescribeDatastoreOutput)

	DescribeLoggingOptions(*iotanalytics.DescribeLoggingOptionsInput) (*iotanalytics.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsWithContext(aws.Context, *iotanalytics.DescribeLoggingOptionsInput, ...request.Option) (*iotanalytics.DescribeLoggingOptionsOutput, error)
	DescribeLoggingOptionsRequest(*iotanalytics.DescribeLoggingOptionsInput) (*request.Request, *iotanalytics.DescribeLoggingOptionsOutput)

	DescribePipeline(*iotanalytics.DescribePipelineInput) (*iotanalytics.DescribePipelineOutput, error)
	DescribePipelineWithContext(aws.Context, *iotanalytics.DescribePipelineInput, ...request.Option) (*iotanalytics.DescribePipelineOutput, error)
	DescribePipelineRequest(*iotanalytics.DescribePipelineInput) (*request.Request, *iotanalytics.DescribePipelineOutput)

	GetDatasetContent(*iotanalytics.GetDatasetContentInput) (*iotanalytics.GetDatasetContentOutput, error)
	GetDatasetContentWithContext(aws.Context, *iotanalytics.GetDatasetContentInput, ...request.Option) (*iotanalytics.GetDatasetContentOutput, error)
	GetDatasetContentRequest(*iotanalytics.GetDatasetContentInput) (*request.Request, *iotanalytics.GetDatasetContentOutput)

	ListChannels(*iotanalytics.ListChannelsInput) (*iotanalytics.ListChannelsOutput, error)
	ListChannelsWithContext(aws.Context, *iotanalytics.ListChannelsInput, ...request.Option) (*iotanalytics.ListChannelsOutput, error)
	ListChannelsRequest(*iotanalytics.ListChannelsInput) (*request.Request, *iotanalytics.ListChannelsOutput)

	ListChannelsPages(*iotanalytics.ListChannelsInput, func(*iotanalytics.ListChannelsOutput, bool) bool) error
	ListChannelsPagesWithContext(aws.Context, *iotanalytics.ListChannelsInput, func(*iotanalytics.ListChannelsOutput, bool) bool, ...request.Option) error

	ListDatasetContents(*iotanalytics.ListDatasetContentsInput) (*iotanalytics.ListDatasetContentsOutput, error)
	ListDatasetContentsWithContext(aws.Context, *iotanalytics.ListDatasetContentsInput, ...request.Option) (*iotanalytics.ListDatasetContentsOutput, error)
	ListDatasetContentsRequest(*iotanalytics.ListDatasetContentsInput) (*request.Request, *iotanalytics.ListDatasetContentsOutput)

	ListDatasetContentsPages(*iotanalytics.ListDatasetContentsInput, func(*iotanalytics.ListDatasetContentsOutput, bool) bool) error
	ListDatasetContentsPagesWithContext(aws.Context, *iotanalytics.ListDatasetContentsInput, func(*iotanalytics.ListDatasetContentsOutput, bool) bool, ...request.Option) error

	ListDatasets(*iotanalytics.ListDatasetsInput) (*iotanalytics.ListDatasetsOutput, error)
	ListDatasetsWithContext(aws.Context, *iotanalytics.ListDatasetsInput, ...request.Option) (*iotanalytics.ListDatasetsOutput, error)
	ListDatasetsRequest(*iotanalytics.ListDatasetsInput) (*request.Request, *iotanalytics.ListDatasetsOutput)

	ListDatasetsPages(*iotanalytics.ListDatasetsInput, func(*iotanalytics.ListDatasetsOutput, bool) bool) error
	ListDatasetsPagesWithContext(aws.Context, *iotanalytics.ListDatasetsInput, func(*iotanalytics.ListDatasetsOutput, bool) bool, ...request.Option) error

	ListDatastores(*iotanalytics.ListDatastoresInput) (*iotanalytics.ListDatastoresOutput, error)
	ListDatastoresWithContext(aws.Context, *iotanalytics.ListDatastoresInput, ...request.Option) (*iotanalytics.ListDatastoresOutput, error)
	ListDatastoresRequest(*iotanalytics.ListDatastoresInput) (*request.Request, *iotanalytics.ListDatastoresOutput)

	ListDatastoresPages(*iotanalytics.ListDatastoresInput, func(*iotanalytics.ListDatastoresOutput, bool) bool) error
	ListDatastoresPagesWithContext(aws.Context, *iotanalytics.ListDatastoresInput, func(*iotanalytics.ListDatastoresOutput, bool) bool, ...request.Option) error

	ListPipelines(*iotanalytics.ListPipelinesInput) (*iotanalytics.ListPipelinesOutput, error)
	ListPipelinesWithContext(aws.Context, *iotanalytics.ListPipelinesInput, ...request.Option) (*iotanalytics.ListPipelinesOutput, error)
	ListPipelinesRequest(*iotanalytics.ListPipelinesInput) (*request.Request, *iotanalytics.ListPipelinesOutput)

	ListPipelinesPages(*iotanalytics.ListPipelinesInput, func(*iotanalytics.ListPipelinesOutput, bool) bool) error
	ListPipelinesPagesWithContext(aws.Context, *iotanalytics.ListPipelinesInput, func(*iotanalytics.ListPipelinesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*iotanalytics.ListTagsForResourceInput) (*iotanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *iotanalytics.ListTagsForResourceInput, ...request.Option) (*iotanalytics.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*iotanalytics.ListTagsForResourceInput) (*request.Request, *iotanalytics.ListTagsForResourceOutput)

	PutLoggingOptions(*iotanalytics.PutLoggingOptionsInput) (*iotanalytics.PutLoggingOptionsOutput, error)
	PutLoggingOptionsWithContext(aws.Context, *iotanalytics.PutLoggingOptionsInput, ...request.Option) (*iotanalytics.PutLoggingOptionsOutput, error)
	PutLoggingOptionsRequest(*iotanalytics.PutLoggingOptionsInput) (*request.Request, *iotanalytics.PutLoggingOptionsOutput)

	RunPipelineActivity(*iotanalytics.RunPipelineActivityInput) (*iotanalytics.RunPipelineActivityOutput, error)
	RunPipelineActivityWithContext(aws.Context, *iotanalytics.RunPipelineActivityInput, ...request.Option) (*iotanalytics.RunPipelineActivityOutput, error)
	RunPipelineActivityRequest(*iotanalytics.RunPipelineActivityInput) (*request.Request, *iotanalytics.RunPipelineActivityOutput)

	SampleChannelData(*iotanalytics.SampleChannelDataInput) (*iotanalytics.SampleChannelDataOutput, error)
	SampleChannelDataWithContext(aws.Context, *iotanalytics.SampleChannelDataInput, ...request.Option) (*iotanalytics.SampleChannelDataOutput, error)
	SampleChannelDataRequest(*iotanalytics.SampleChannelDataInput) (*request.Request, *iotanalytics.SampleChannelDataOutput)

	StartPipelineReprocessing(*iotanalytics.StartPipelineReprocessingInput) (*iotanalytics.StartPipelineReprocessingOutput, error)
	StartPipelineReprocessingWithContext(aws.Context, *iotanalytics.StartPipelineReprocessingInput, ...request.Option) (*iotanalytics.StartPipelineReprocessingOutput, error)
	StartPipelineReprocessingRequest(*iotanalytics.StartPipelineReprocessingInput) (*request.Request, *iotanalytics.StartPipelineReprocessingOutput)

	TagResource(*iotanalytics.TagResourceInput) (*iotanalytics.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *iotanalytics.TagResourceInput, ...request.Option) (*iotanalytics.TagResourceOutput, error)
	TagResourceRequest(*iotanalytics.TagResourceInput) (*request.Request, *iotanalytics.TagResourceOutput)

	UntagResource(*iotanalytics.UntagResourceInput) (*iotanalytics.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *iotanalytics.UntagResourceInput, ...request.Option) (*iotanalytics.UntagResourceOutput, error)
	UntagResourceRequest(*iotanalytics.UntagResourceInput) (*request.Request, *iotanalytics.UntagResourceOutput)

	UpdateChannel(*iotanalytics.UpdateChannelInput) (*iotanalytics.UpdateChannelOutput, error)
	UpdateChannelWithContext(aws.Context, *iotanalytics.UpdateChannelInput, ...request.Option) (*iotanalytics.UpdateChannelOutput, error)
	UpdateChannelRequest(*iotanalytics.UpdateChannelInput) (*request.Request, *iotanalytics.UpdateChannelOutput)

	UpdateDataset(*iotanalytics.UpdateDatasetInput) (*iotanalytics.UpdateDatasetOutput, error)
	UpdateDatasetWithContext(aws.Context, *iotanalytics.UpdateDatasetInput, ...request.Option) (*iotanalytics.UpdateDatasetOutput, error)
	UpdateDatasetRequest(*iotanalytics.UpdateDatasetInput) (*request.Request, *iotanalytics.UpdateDatasetOutput)

	UpdateDatastore(*iotanalytics.UpdateDatastoreInput) (*iotanalytics.UpdateDatastoreOutput, error)
	UpdateDatastoreWithContext(aws.Context, *iotanalytics.UpdateDatastoreInput, ...request.Option) (*iotanalytics.UpdateDatastoreOutput, error)
	UpdateDatastoreRequest(*iotanalytics.UpdateDatastoreInput) (*request.Request, *iotanalytics.UpdateDatastoreOutput)

	UpdatePipeline(*iotanalytics.UpdatePipelineInput) (*iotanalytics.UpdatePipelineOutput, error)
	UpdatePipelineWithContext(aws.Context, *iotanalytics.UpdatePipelineInput, ...request.Option) (*iotanalytics.UpdatePipelineOutput, error)
	UpdatePipelineRequest(*iotanalytics.UpdatePipelineInput) (*request.Request, *iotanalytics.UpdatePipelineOutput)
}

var _ IoTAnalyticsAPI = (*iotanalytics.IoTAnalytics)(nil)
