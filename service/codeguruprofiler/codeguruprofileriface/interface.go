// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codeguruprofileriface provides an interface to enable mocking the Amazon CodeGuru Profiler service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package codeguruprofileriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/codeguruprofiler"
)

// CodeGuruProfilerAPI provides an interface to enable mocking the
// codeguruprofiler.CodeGuruProfiler service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon CodeGuru Profiler.
//    func myFunc(svc codeguruprofileriface.CodeGuruProfilerAPI) bool {
//        // Make svc.AddNotificationChannels request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := codeguruprofiler.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCodeGuruProfilerClient struct {
//        codeguruprofileriface.CodeGuruProfilerAPI
//    }
//    func (m *mockCodeGuruProfilerClient) AddNotificationChannels(input *codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCodeGuruProfilerClient{}
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
type CodeGuruProfilerAPI interface {
	AddNotificationChannels(*codeguruprofiler.AddNotificationChannelsInput) (*codeguruprofiler.AddNotificationChannelsOutput, error)
	AddNotificationChannelsWithContext(aws.Context, *codeguruprofiler.AddNotificationChannelsInput, ...request.Option) (*codeguruprofiler.AddNotificationChannelsOutput, error)
	AddNotificationChannelsRequest(*codeguruprofiler.AddNotificationChannelsInput) (*request.Request, *codeguruprofiler.AddNotificationChannelsOutput)

	BatchGetFrameMetricData(*codeguruprofiler.BatchGetFrameMetricDataInput) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error)
	BatchGetFrameMetricDataWithContext(aws.Context, *codeguruprofiler.BatchGetFrameMetricDataInput, ...request.Option) (*codeguruprofiler.BatchGetFrameMetricDataOutput, error)
	BatchGetFrameMetricDataRequest(*codeguruprofiler.BatchGetFrameMetricDataInput) (*request.Request, *codeguruprofiler.BatchGetFrameMetricDataOutput)

	ConfigureAgent(*codeguruprofiler.ConfigureAgentInput) (*codeguruprofiler.ConfigureAgentOutput, error)
	ConfigureAgentWithContext(aws.Context, *codeguruprofiler.ConfigureAgentInput, ...request.Option) (*codeguruprofiler.ConfigureAgentOutput, error)
	ConfigureAgentRequest(*codeguruprofiler.ConfigureAgentInput) (*request.Request, *codeguruprofiler.ConfigureAgentOutput)

	CreateProfilingGroup(*codeguruprofiler.CreateProfilingGroupInput) (*codeguruprofiler.CreateProfilingGroupOutput, error)
	CreateProfilingGroupWithContext(aws.Context, *codeguruprofiler.CreateProfilingGroupInput, ...request.Option) (*codeguruprofiler.CreateProfilingGroupOutput, error)
	CreateProfilingGroupRequest(*codeguruprofiler.CreateProfilingGroupInput) (*request.Request, *codeguruprofiler.CreateProfilingGroupOutput)

	DeleteProfilingGroup(*codeguruprofiler.DeleteProfilingGroupInput) (*codeguruprofiler.DeleteProfilingGroupOutput, error)
	DeleteProfilingGroupWithContext(aws.Context, *codeguruprofiler.DeleteProfilingGroupInput, ...request.Option) (*codeguruprofiler.DeleteProfilingGroupOutput, error)
	DeleteProfilingGroupRequest(*codeguruprofiler.DeleteProfilingGroupInput) (*request.Request, *codeguruprofiler.DeleteProfilingGroupOutput)

	DescribeProfilingGroup(*codeguruprofiler.DescribeProfilingGroupInput) (*codeguruprofiler.DescribeProfilingGroupOutput, error)
	DescribeProfilingGroupWithContext(aws.Context, *codeguruprofiler.DescribeProfilingGroupInput, ...request.Option) (*codeguruprofiler.DescribeProfilingGroupOutput, error)
	DescribeProfilingGroupRequest(*codeguruprofiler.DescribeProfilingGroupInput) (*request.Request, *codeguruprofiler.DescribeProfilingGroupOutput)

	GetFindingsReportAccountSummary(*codeguruprofiler.GetFindingsReportAccountSummaryInput) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error)
	GetFindingsReportAccountSummaryWithContext(aws.Context, *codeguruprofiler.GetFindingsReportAccountSummaryInput, ...request.Option) (*codeguruprofiler.GetFindingsReportAccountSummaryOutput, error)
	GetFindingsReportAccountSummaryRequest(*codeguruprofiler.GetFindingsReportAccountSummaryInput) (*request.Request, *codeguruprofiler.GetFindingsReportAccountSummaryOutput)

	GetFindingsReportAccountSummaryPages(*codeguruprofiler.GetFindingsReportAccountSummaryInput, func(*codeguruprofiler.GetFindingsReportAccountSummaryOutput, bool) bool) error
	GetFindingsReportAccountSummaryPagesWithContext(aws.Context, *codeguruprofiler.GetFindingsReportAccountSummaryInput, func(*codeguruprofiler.GetFindingsReportAccountSummaryOutput, bool) bool, ...request.Option) error

	GetNotificationConfiguration(*codeguruprofiler.GetNotificationConfigurationInput) (*codeguruprofiler.GetNotificationConfigurationOutput, error)
	GetNotificationConfigurationWithContext(aws.Context, *codeguruprofiler.GetNotificationConfigurationInput, ...request.Option) (*codeguruprofiler.GetNotificationConfigurationOutput, error)
	GetNotificationConfigurationRequest(*codeguruprofiler.GetNotificationConfigurationInput) (*request.Request, *codeguruprofiler.GetNotificationConfigurationOutput)

	GetPolicy(*codeguruprofiler.GetPolicyInput) (*codeguruprofiler.GetPolicyOutput, error)
	GetPolicyWithContext(aws.Context, *codeguruprofiler.GetPolicyInput, ...request.Option) (*codeguruprofiler.GetPolicyOutput, error)
	GetPolicyRequest(*codeguruprofiler.GetPolicyInput) (*request.Request, *codeguruprofiler.GetPolicyOutput)

	GetProfile(*codeguruprofiler.GetProfileInput) (*codeguruprofiler.GetProfileOutput, error)
	GetProfileWithContext(aws.Context, *codeguruprofiler.GetProfileInput, ...request.Option) (*codeguruprofiler.GetProfileOutput, error)
	GetProfileRequest(*codeguruprofiler.GetProfileInput) (*request.Request, *codeguruprofiler.GetProfileOutput)

	GetRecommendations(*codeguruprofiler.GetRecommendationsInput) (*codeguruprofiler.GetRecommendationsOutput, error)
	GetRecommendationsWithContext(aws.Context, *codeguruprofiler.GetRecommendationsInput, ...request.Option) (*codeguruprofiler.GetRecommendationsOutput, error)
	GetRecommendationsRequest(*codeguruprofiler.GetRecommendationsInput) (*request.Request, *codeguruprofiler.GetRecommendationsOutput)

	ListFindingsReports(*codeguruprofiler.ListFindingsReportsInput) (*codeguruprofiler.ListFindingsReportsOutput, error)
	ListFindingsReportsWithContext(aws.Context, *codeguruprofiler.ListFindingsReportsInput, ...request.Option) (*codeguruprofiler.ListFindingsReportsOutput, error)
	ListFindingsReportsRequest(*codeguruprofiler.ListFindingsReportsInput) (*request.Request, *codeguruprofiler.ListFindingsReportsOutput)

	ListFindingsReportsPages(*codeguruprofiler.ListFindingsReportsInput, func(*codeguruprofiler.ListFindingsReportsOutput, bool) bool) error
	ListFindingsReportsPagesWithContext(aws.Context, *codeguruprofiler.ListFindingsReportsInput, func(*codeguruprofiler.ListFindingsReportsOutput, bool) bool, ...request.Option) error

	ListProfileTimes(*codeguruprofiler.ListProfileTimesInput) (*codeguruprofiler.ListProfileTimesOutput, error)
	ListProfileTimesWithContext(aws.Context, *codeguruprofiler.ListProfileTimesInput, ...request.Option) (*codeguruprofiler.ListProfileTimesOutput, error)
	ListProfileTimesRequest(*codeguruprofiler.ListProfileTimesInput) (*request.Request, *codeguruprofiler.ListProfileTimesOutput)

	ListProfileTimesPages(*codeguruprofiler.ListProfileTimesInput, func(*codeguruprofiler.ListProfileTimesOutput, bool) bool) error
	ListProfileTimesPagesWithContext(aws.Context, *codeguruprofiler.ListProfileTimesInput, func(*codeguruprofiler.ListProfileTimesOutput, bool) bool, ...request.Option) error

	ListProfilingGroups(*codeguruprofiler.ListProfilingGroupsInput) (*codeguruprofiler.ListProfilingGroupsOutput, error)
	ListProfilingGroupsWithContext(aws.Context, *codeguruprofiler.ListProfilingGroupsInput, ...request.Option) (*codeguruprofiler.ListProfilingGroupsOutput, error)
	ListProfilingGroupsRequest(*codeguruprofiler.ListProfilingGroupsInput) (*request.Request, *codeguruprofiler.ListProfilingGroupsOutput)

	ListProfilingGroupsPages(*codeguruprofiler.ListProfilingGroupsInput, func(*codeguruprofiler.ListProfilingGroupsOutput, bool) bool) error
	ListProfilingGroupsPagesWithContext(aws.Context, *codeguruprofiler.ListProfilingGroupsInput, func(*codeguruprofiler.ListProfilingGroupsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*codeguruprofiler.ListTagsForResourceInput) (*codeguruprofiler.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *codeguruprofiler.ListTagsForResourceInput, ...request.Option) (*codeguruprofiler.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*codeguruprofiler.ListTagsForResourceInput) (*request.Request, *codeguruprofiler.ListTagsForResourceOutput)

	PostAgentProfile(*codeguruprofiler.PostAgentProfileInput) (*codeguruprofiler.PostAgentProfileOutput, error)
	PostAgentProfileWithContext(aws.Context, *codeguruprofiler.PostAgentProfileInput, ...request.Option) (*codeguruprofiler.PostAgentProfileOutput, error)
	PostAgentProfileRequest(*codeguruprofiler.PostAgentProfileInput) (*request.Request, *codeguruprofiler.PostAgentProfileOutput)

	PutPermission(*codeguruprofiler.PutPermissionInput) (*codeguruprofiler.PutPermissionOutput, error)
	PutPermissionWithContext(aws.Context, *codeguruprofiler.PutPermissionInput, ...request.Option) (*codeguruprofiler.PutPermissionOutput, error)
	PutPermissionRequest(*codeguruprofiler.PutPermissionInput) (*request.Request, *codeguruprofiler.PutPermissionOutput)

	RemoveNotificationChannel(*codeguruprofiler.RemoveNotificationChannelInput) (*codeguruprofiler.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelWithContext(aws.Context, *codeguruprofiler.RemoveNotificationChannelInput, ...request.Option) (*codeguruprofiler.RemoveNotificationChannelOutput, error)
	RemoveNotificationChannelRequest(*codeguruprofiler.RemoveNotificationChannelInput) (*request.Request, *codeguruprofiler.RemoveNotificationChannelOutput)

	RemovePermission(*codeguruprofiler.RemovePermissionInput) (*codeguruprofiler.RemovePermissionOutput, error)
	RemovePermissionWithContext(aws.Context, *codeguruprofiler.RemovePermissionInput, ...request.Option) (*codeguruprofiler.RemovePermissionOutput, error)
	RemovePermissionRequest(*codeguruprofiler.RemovePermissionInput) (*request.Request, *codeguruprofiler.RemovePermissionOutput)

	SubmitFeedback(*codeguruprofiler.SubmitFeedbackInput) (*codeguruprofiler.SubmitFeedbackOutput, error)
	SubmitFeedbackWithContext(aws.Context, *codeguruprofiler.SubmitFeedbackInput, ...request.Option) (*codeguruprofiler.SubmitFeedbackOutput, error)
	SubmitFeedbackRequest(*codeguruprofiler.SubmitFeedbackInput) (*request.Request, *codeguruprofiler.SubmitFeedbackOutput)

	TagResource(*codeguruprofiler.TagResourceInput) (*codeguruprofiler.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *codeguruprofiler.TagResourceInput, ...request.Option) (*codeguruprofiler.TagResourceOutput, error)
	TagResourceRequest(*codeguruprofiler.TagResourceInput) (*request.Request, *codeguruprofiler.TagResourceOutput)

	UntagResource(*codeguruprofiler.UntagResourceInput) (*codeguruprofiler.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *codeguruprofiler.UntagResourceInput, ...request.Option) (*codeguruprofiler.UntagResourceOutput, error)
	UntagResourceRequest(*codeguruprofiler.UntagResourceInput) (*request.Request, *codeguruprofiler.UntagResourceOutput)

	UpdateProfilingGroup(*codeguruprofiler.UpdateProfilingGroupInput) (*codeguruprofiler.UpdateProfilingGroupOutput, error)
	UpdateProfilingGroupWithContext(aws.Context, *codeguruprofiler.UpdateProfilingGroupInput, ...request.Option) (*codeguruprofiler.UpdateProfilingGroupOutput, error)
	UpdateProfilingGroupRequest(*codeguruprofiler.UpdateProfilingGroupInput) (*request.Request, *codeguruprofiler.UpdateProfilingGroupOutput)
}

var _ CodeGuruProfilerAPI = (*codeguruprofiler.CodeGuruProfiler)(nil)
