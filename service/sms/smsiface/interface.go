// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package smsiface provides an interface to enable mocking the AWS Server Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package smsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/sms"
)

// SMSAPI provides an interface to enable mocking the
// sms.SMS service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Server Migration Service.
//    func myFunc(svc smsiface.SMSAPI) bool {
//        // Make svc.CreateApp request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := sms.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockSMSClient struct {
//        smsiface.SMSAPI
//    }
//    func (m *mockSMSClient) CreateApp(input *sms.CreateAppInput) (*sms.CreateAppOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockSMSClient{}
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
type SMSAPI interface {
	CreateApp(*sms.CreateAppInput) (*sms.CreateAppOutput, error)
	CreateAppWithContext(aws.Context, *sms.CreateAppInput, ...request.Option) (*sms.CreateAppOutput, error)
	CreateAppRequest(*sms.CreateAppInput) (*request.Request, *sms.CreateAppOutput)

	CreateReplicationJob(*sms.CreateReplicationJobInput) (*sms.CreateReplicationJobOutput, error)
	CreateReplicationJobWithContext(aws.Context, *sms.CreateReplicationJobInput, ...request.Option) (*sms.CreateReplicationJobOutput, error)
	CreateReplicationJobRequest(*sms.CreateReplicationJobInput) (*request.Request, *sms.CreateReplicationJobOutput)

	DeleteApp(*sms.DeleteAppInput) (*sms.DeleteAppOutput, error)
	DeleteAppWithContext(aws.Context, *sms.DeleteAppInput, ...request.Option) (*sms.DeleteAppOutput, error)
	DeleteAppRequest(*sms.DeleteAppInput) (*request.Request, *sms.DeleteAppOutput)

	DeleteAppLaunchConfiguration(*sms.DeleteAppLaunchConfigurationInput) (*sms.DeleteAppLaunchConfigurationOutput, error)
	DeleteAppLaunchConfigurationWithContext(aws.Context, *sms.DeleteAppLaunchConfigurationInput, ...request.Option) (*sms.DeleteAppLaunchConfigurationOutput, error)
	DeleteAppLaunchConfigurationRequest(*sms.DeleteAppLaunchConfigurationInput) (*request.Request, *sms.DeleteAppLaunchConfigurationOutput)

	DeleteAppReplicationConfiguration(*sms.DeleteAppReplicationConfigurationInput) (*sms.DeleteAppReplicationConfigurationOutput, error)
	DeleteAppReplicationConfigurationWithContext(aws.Context, *sms.DeleteAppReplicationConfigurationInput, ...request.Option) (*sms.DeleteAppReplicationConfigurationOutput, error)
	DeleteAppReplicationConfigurationRequest(*sms.DeleteAppReplicationConfigurationInput) (*request.Request, *sms.DeleteAppReplicationConfigurationOutput)

	DeleteAppValidationConfiguration(*sms.DeleteAppValidationConfigurationInput) (*sms.DeleteAppValidationConfigurationOutput, error)
	DeleteAppValidationConfigurationWithContext(aws.Context, *sms.DeleteAppValidationConfigurationInput, ...request.Option) (*sms.DeleteAppValidationConfigurationOutput, error)
	DeleteAppValidationConfigurationRequest(*sms.DeleteAppValidationConfigurationInput) (*request.Request, *sms.DeleteAppValidationConfigurationOutput)

	DeleteReplicationJob(*sms.DeleteReplicationJobInput) (*sms.DeleteReplicationJobOutput, error)
	DeleteReplicationJobWithContext(aws.Context, *sms.DeleteReplicationJobInput, ...request.Option) (*sms.DeleteReplicationJobOutput, error)
	DeleteReplicationJobRequest(*sms.DeleteReplicationJobInput) (*request.Request, *sms.DeleteReplicationJobOutput)

	DeleteServerCatalog(*sms.DeleteServerCatalogInput) (*sms.DeleteServerCatalogOutput, error)
	DeleteServerCatalogWithContext(aws.Context, *sms.DeleteServerCatalogInput, ...request.Option) (*sms.DeleteServerCatalogOutput, error)
	DeleteServerCatalogRequest(*sms.DeleteServerCatalogInput) (*request.Request, *sms.DeleteServerCatalogOutput)

	DisassociateConnector(*sms.DisassociateConnectorInput) (*sms.DisassociateConnectorOutput, error)
	DisassociateConnectorWithContext(aws.Context, *sms.DisassociateConnectorInput, ...request.Option) (*sms.DisassociateConnectorOutput, error)
	DisassociateConnectorRequest(*sms.DisassociateConnectorInput) (*request.Request, *sms.DisassociateConnectorOutput)

	GenerateChangeSet(*sms.GenerateChangeSetInput) (*sms.GenerateChangeSetOutput, error)
	GenerateChangeSetWithContext(aws.Context, *sms.GenerateChangeSetInput, ...request.Option) (*sms.GenerateChangeSetOutput, error)
	GenerateChangeSetRequest(*sms.GenerateChangeSetInput) (*request.Request, *sms.GenerateChangeSetOutput)

	GenerateTemplate(*sms.GenerateTemplateInput) (*sms.GenerateTemplateOutput, error)
	GenerateTemplateWithContext(aws.Context, *sms.GenerateTemplateInput, ...request.Option) (*sms.GenerateTemplateOutput, error)
	GenerateTemplateRequest(*sms.GenerateTemplateInput) (*request.Request, *sms.GenerateTemplateOutput)

	GetApp(*sms.GetAppInput) (*sms.GetAppOutput, error)
	GetAppWithContext(aws.Context, *sms.GetAppInput, ...request.Option) (*sms.GetAppOutput, error)
	GetAppRequest(*sms.GetAppInput) (*request.Request, *sms.GetAppOutput)

	GetAppLaunchConfiguration(*sms.GetAppLaunchConfigurationInput) (*sms.GetAppLaunchConfigurationOutput, error)
	GetAppLaunchConfigurationWithContext(aws.Context, *sms.GetAppLaunchConfigurationInput, ...request.Option) (*sms.GetAppLaunchConfigurationOutput, error)
	GetAppLaunchConfigurationRequest(*sms.GetAppLaunchConfigurationInput) (*request.Request, *sms.GetAppLaunchConfigurationOutput)

	GetAppReplicationConfiguration(*sms.GetAppReplicationConfigurationInput) (*sms.GetAppReplicationConfigurationOutput, error)
	GetAppReplicationConfigurationWithContext(aws.Context, *sms.GetAppReplicationConfigurationInput, ...request.Option) (*sms.GetAppReplicationConfigurationOutput, error)
	GetAppReplicationConfigurationRequest(*sms.GetAppReplicationConfigurationInput) (*request.Request, *sms.GetAppReplicationConfigurationOutput)

	GetAppValidationConfiguration(*sms.GetAppValidationConfigurationInput) (*sms.GetAppValidationConfigurationOutput, error)
	GetAppValidationConfigurationWithContext(aws.Context, *sms.GetAppValidationConfigurationInput, ...request.Option) (*sms.GetAppValidationConfigurationOutput, error)
	GetAppValidationConfigurationRequest(*sms.GetAppValidationConfigurationInput) (*request.Request, *sms.GetAppValidationConfigurationOutput)

	GetAppValidationOutput(*sms.GetAppValidationOutputInput) (*sms.GetAppValidationOutputOutput, error)
	GetAppValidationOutputWithContext(aws.Context, *sms.GetAppValidationOutputInput, ...request.Option) (*sms.GetAppValidationOutputOutput, error)
	GetAppValidationOutputRequest(*sms.GetAppValidationOutputInput) (*request.Request, *sms.GetAppValidationOutputOutput)

	GetConnectors(*sms.GetConnectorsInput) (*sms.GetConnectorsOutput, error)
	GetConnectorsWithContext(aws.Context, *sms.GetConnectorsInput, ...request.Option) (*sms.GetConnectorsOutput, error)
	GetConnectorsRequest(*sms.GetConnectorsInput) (*request.Request, *sms.GetConnectorsOutput)

	GetConnectorsPages(*sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool) error
	GetConnectorsPagesWithContext(aws.Context, *sms.GetConnectorsInput, func(*sms.GetConnectorsOutput, bool) bool, ...request.Option) error

	GetReplicationJobs(*sms.GetReplicationJobsInput) (*sms.GetReplicationJobsOutput, error)
	GetReplicationJobsWithContext(aws.Context, *sms.GetReplicationJobsInput, ...request.Option) (*sms.GetReplicationJobsOutput, error)
	GetReplicationJobsRequest(*sms.GetReplicationJobsInput) (*request.Request, *sms.GetReplicationJobsOutput)

	GetReplicationJobsPages(*sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool) error
	GetReplicationJobsPagesWithContext(aws.Context, *sms.GetReplicationJobsInput, func(*sms.GetReplicationJobsOutput, bool) bool, ...request.Option) error

	GetReplicationRuns(*sms.GetReplicationRunsInput) (*sms.GetReplicationRunsOutput, error)
	GetReplicationRunsWithContext(aws.Context, *sms.GetReplicationRunsInput, ...request.Option) (*sms.GetReplicationRunsOutput, error)
	GetReplicationRunsRequest(*sms.GetReplicationRunsInput) (*request.Request, *sms.GetReplicationRunsOutput)

	GetReplicationRunsPages(*sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool) error
	GetReplicationRunsPagesWithContext(aws.Context, *sms.GetReplicationRunsInput, func(*sms.GetReplicationRunsOutput, bool) bool, ...request.Option) error

	GetServers(*sms.GetServersInput) (*sms.GetServersOutput, error)
	GetServersWithContext(aws.Context, *sms.GetServersInput, ...request.Option) (*sms.GetServersOutput, error)
	GetServersRequest(*sms.GetServersInput) (*request.Request, *sms.GetServersOutput)

	GetServersPages(*sms.GetServersInput, func(*sms.GetServersOutput, bool) bool) error
	GetServersPagesWithContext(aws.Context, *sms.GetServersInput, func(*sms.GetServersOutput, bool) bool, ...request.Option) error

	ImportAppCatalog(*sms.ImportAppCatalogInput) (*sms.ImportAppCatalogOutput, error)
	ImportAppCatalogWithContext(aws.Context, *sms.ImportAppCatalogInput, ...request.Option) (*sms.ImportAppCatalogOutput, error)
	ImportAppCatalogRequest(*sms.ImportAppCatalogInput) (*request.Request, *sms.ImportAppCatalogOutput)

	ImportServerCatalog(*sms.ImportServerCatalogInput) (*sms.ImportServerCatalogOutput, error)
	ImportServerCatalogWithContext(aws.Context, *sms.ImportServerCatalogInput, ...request.Option) (*sms.ImportServerCatalogOutput, error)
	ImportServerCatalogRequest(*sms.ImportServerCatalogInput) (*request.Request, *sms.ImportServerCatalogOutput)

	LaunchApp(*sms.LaunchAppInput) (*sms.LaunchAppOutput, error)
	LaunchAppWithContext(aws.Context, *sms.LaunchAppInput, ...request.Option) (*sms.LaunchAppOutput, error)
	LaunchAppRequest(*sms.LaunchAppInput) (*request.Request, *sms.LaunchAppOutput)

	ListApps(*sms.ListAppsInput) (*sms.ListAppsOutput, error)
	ListAppsWithContext(aws.Context, *sms.ListAppsInput, ...request.Option) (*sms.ListAppsOutput, error)
	ListAppsRequest(*sms.ListAppsInput) (*request.Request, *sms.ListAppsOutput)

	NotifyAppValidationOutput(*sms.NotifyAppValidationOutputInput) (*sms.NotifyAppValidationOutputOutput, error)
	NotifyAppValidationOutputWithContext(aws.Context, *sms.NotifyAppValidationOutputInput, ...request.Option) (*sms.NotifyAppValidationOutputOutput, error)
	NotifyAppValidationOutputRequest(*sms.NotifyAppValidationOutputInput) (*request.Request, *sms.NotifyAppValidationOutputOutput)

	PutAppLaunchConfiguration(*sms.PutAppLaunchConfigurationInput) (*sms.PutAppLaunchConfigurationOutput, error)
	PutAppLaunchConfigurationWithContext(aws.Context, *sms.PutAppLaunchConfigurationInput, ...request.Option) (*sms.PutAppLaunchConfigurationOutput, error)
	PutAppLaunchConfigurationRequest(*sms.PutAppLaunchConfigurationInput) (*request.Request, *sms.PutAppLaunchConfigurationOutput)

	PutAppReplicationConfiguration(*sms.PutAppReplicationConfigurationInput) (*sms.PutAppReplicationConfigurationOutput, error)
	PutAppReplicationConfigurationWithContext(aws.Context, *sms.PutAppReplicationConfigurationInput, ...request.Option) (*sms.PutAppReplicationConfigurationOutput, error)
	PutAppReplicationConfigurationRequest(*sms.PutAppReplicationConfigurationInput) (*request.Request, *sms.PutAppReplicationConfigurationOutput)

	PutAppValidationConfiguration(*sms.PutAppValidationConfigurationInput) (*sms.PutAppValidationConfigurationOutput, error)
	PutAppValidationConfigurationWithContext(aws.Context, *sms.PutAppValidationConfigurationInput, ...request.Option) (*sms.PutAppValidationConfigurationOutput, error)
	PutAppValidationConfigurationRequest(*sms.PutAppValidationConfigurationInput) (*request.Request, *sms.PutAppValidationConfigurationOutput)

	StartAppReplication(*sms.StartAppReplicationInput) (*sms.StartAppReplicationOutput, error)
	StartAppReplicationWithContext(aws.Context, *sms.StartAppReplicationInput, ...request.Option) (*sms.StartAppReplicationOutput, error)
	StartAppReplicationRequest(*sms.StartAppReplicationInput) (*request.Request, *sms.StartAppReplicationOutput)

	StartOnDemandAppReplication(*sms.StartOnDemandAppReplicationInput) (*sms.StartOnDemandAppReplicationOutput, error)
	StartOnDemandAppReplicationWithContext(aws.Context, *sms.StartOnDemandAppReplicationInput, ...request.Option) (*sms.StartOnDemandAppReplicationOutput, error)
	StartOnDemandAppReplicationRequest(*sms.StartOnDemandAppReplicationInput) (*request.Request, *sms.StartOnDemandAppReplicationOutput)

	StartOnDemandReplicationRun(*sms.StartOnDemandReplicationRunInput) (*sms.StartOnDemandReplicationRunOutput, error)
	StartOnDemandReplicationRunWithContext(aws.Context, *sms.StartOnDemandReplicationRunInput, ...request.Option) (*sms.StartOnDemandReplicationRunOutput, error)
	StartOnDemandReplicationRunRequest(*sms.StartOnDemandReplicationRunInput) (*request.Request, *sms.StartOnDemandReplicationRunOutput)

	StopAppReplication(*sms.StopAppReplicationInput) (*sms.StopAppReplicationOutput, error)
	StopAppReplicationWithContext(aws.Context, *sms.StopAppReplicationInput, ...request.Option) (*sms.StopAppReplicationOutput, error)
	StopAppReplicationRequest(*sms.StopAppReplicationInput) (*request.Request, *sms.StopAppReplicationOutput)

	TerminateApp(*sms.TerminateAppInput) (*sms.TerminateAppOutput, error)
	TerminateAppWithContext(aws.Context, *sms.TerminateAppInput, ...request.Option) (*sms.TerminateAppOutput, error)
	TerminateAppRequest(*sms.TerminateAppInput) (*request.Request, *sms.TerminateAppOutput)

	UpdateApp(*sms.UpdateAppInput) (*sms.UpdateAppOutput, error)
	UpdateAppWithContext(aws.Context, *sms.UpdateAppInput, ...request.Option) (*sms.UpdateAppOutput, error)
	UpdateAppRequest(*sms.UpdateAppInput) (*request.Request, *sms.UpdateAppOutput)

	UpdateReplicationJob(*sms.UpdateReplicationJobInput) (*sms.UpdateReplicationJobOutput, error)
	UpdateReplicationJobWithContext(aws.Context, *sms.UpdateReplicationJobInput, ...request.Option) (*sms.UpdateReplicationJobOutput, error)
	UpdateReplicationJobRequest(*sms.UpdateReplicationJobInput) (*request.Request, *sms.UpdateReplicationJobOutput)
}

var _ SMSAPI = (*sms.SMS)(nil)
