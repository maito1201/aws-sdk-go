// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package translateiface provides an interface to enable mocking the Amazon Translate service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package translateiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/translate"
)

// TranslateAPI provides an interface to enable mocking the
// translate.Translate service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Translate.
//    func myFunc(svc translateiface.TranslateAPI) bool {
//        // Make svc.DeleteTerminology request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := translate.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockTranslateClient struct {
//        translateiface.TranslateAPI
//    }
//    func (m *mockTranslateClient) DeleteTerminology(input *translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockTranslateClient{}
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
type TranslateAPI interface {
	DeleteTerminology(*translate.DeleteTerminologyInput) (*translate.DeleteTerminologyOutput, error)
	DeleteTerminologyWithContext(aws.Context, *translate.DeleteTerminologyInput, ...request.Option) (*translate.DeleteTerminologyOutput, error)
	DeleteTerminologyRequest(*translate.DeleteTerminologyInput) (*request.Request, *translate.DeleteTerminologyOutput)

	DescribeTextTranslationJob(*translate.DescribeTextTranslationJobInput) (*translate.DescribeTextTranslationJobOutput, error)
	DescribeTextTranslationJobWithContext(aws.Context, *translate.DescribeTextTranslationJobInput, ...request.Option) (*translate.DescribeTextTranslationJobOutput, error)
	DescribeTextTranslationJobRequest(*translate.DescribeTextTranslationJobInput) (*request.Request, *translate.DescribeTextTranslationJobOutput)

	GetTerminology(*translate.GetTerminologyInput) (*translate.GetTerminologyOutput, error)
	GetTerminologyWithContext(aws.Context, *translate.GetTerminologyInput, ...request.Option) (*translate.GetTerminologyOutput, error)
	GetTerminologyRequest(*translate.GetTerminologyInput) (*request.Request, *translate.GetTerminologyOutput)

	ImportTerminology(*translate.ImportTerminologyInput) (*translate.ImportTerminologyOutput, error)
	ImportTerminologyWithContext(aws.Context, *translate.ImportTerminologyInput, ...request.Option) (*translate.ImportTerminologyOutput, error)
	ImportTerminologyRequest(*translate.ImportTerminologyInput) (*request.Request, *translate.ImportTerminologyOutput)

	ListTerminologies(*translate.ListTerminologiesInput) (*translate.ListTerminologiesOutput, error)
	ListTerminologiesWithContext(aws.Context, *translate.ListTerminologiesInput, ...request.Option) (*translate.ListTerminologiesOutput, error)
	ListTerminologiesRequest(*translate.ListTerminologiesInput) (*request.Request, *translate.ListTerminologiesOutput)

	ListTerminologiesPages(*translate.ListTerminologiesInput, func(*translate.ListTerminologiesOutput, bool) bool) error
	ListTerminologiesPagesWithContext(aws.Context, *translate.ListTerminologiesInput, func(*translate.ListTerminologiesOutput, bool) bool, ...request.Option) error

	ListTextTranslationJobs(*translate.ListTextTranslationJobsInput) (*translate.ListTextTranslationJobsOutput, error)
	ListTextTranslationJobsWithContext(aws.Context, *translate.ListTextTranslationJobsInput, ...request.Option) (*translate.ListTextTranslationJobsOutput, error)
	ListTextTranslationJobsRequest(*translate.ListTextTranslationJobsInput) (*request.Request, *translate.ListTextTranslationJobsOutput)

	ListTextTranslationJobsPages(*translate.ListTextTranslationJobsInput, func(*translate.ListTextTranslationJobsOutput, bool) bool) error
	ListTextTranslationJobsPagesWithContext(aws.Context, *translate.ListTextTranslationJobsInput, func(*translate.ListTextTranslationJobsOutput, bool) bool, ...request.Option) error

	StartTextTranslationJob(*translate.StartTextTranslationJobInput) (*translate.StartTextTranslationJobOutput, error)
	StartTextTranslationJobWithContext(aws.Context, *translate.StartTextTranslationJobInput, ...request.Option) (*translate.StartTextTranslationJobOutput, error)
	StartTextTranslationJobRequest(*translate.StartTextTranslationJobInput) (*request.Request, *translate.StartTextTranslationJobOutput)

	StopTextTranslationJob(*translate.StopTextTranslationJobInput) (*translate.StopTextTranslationJobOutput, error)
	StopTextTranslationJobWithContext(aws.Context, *translate.StopTextTranslationJobInput, ...request.Option) (*translate.StopTextTranslationJobOutput, error)
	StopTextTranslationJobRequest(*translate.StopTextTranslationJobInput) (*request.Request, *translate.StopTextTranslationJobOutput)

	Text(*translate.TextInput) (*translate.TextOutput, error)
	TextWithContext(aws.Context, *translate.TextInput, ...request.Option) (*translate.TextOutput, error)
	TextRequest(*translate.TextInput) (*request.Request, *translate.TextOutput)
}

var _ TranslateAPI = (*translate.Translate)(nil)
