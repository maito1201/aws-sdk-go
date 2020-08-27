// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package storagegatewayiface provides an interface to enable mocking the AWS Storage Gateway service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package storagegatewayiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/storagegateway"
)

// StorageGatewayAPI provides an interface to enable mocking the
// storagegateway.StorageGateway service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Storage Gateway.
//    func myFunc(svc storagegatewayiface.StorageGatewayAPI) bool {
//        // Make svc.ActivateGateway request
//    }
//
//    func main() {
//        sess := session.Must(session.NewSession())
//        svc := storagegateway.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockStorageGatewayClient struct {
//        storagegatewayiface.StorageGatewayAPI
//    }
//    func (m *mockStorageGatewayClient) ActivateGateway(input *storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockStorageGatewayClient{}
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
type StorageGatewayAPI interface {
	ActivateGateway(*storagegateway.ActivateGatewayInput) (*storagegateway.ActivateGatewayOutput, error)
	ActivateGatewayWithContext(aws.Context, *storagegateway.ActivateGatewayInput, ...request.Option) (*storagegateway.ActivateGatewayOutput, error)
	ActivateGatewayRequest(*storagegateway.ActivateGatewayInput) (*request.Request, *storagegateway.ActivateGatewayOutput)

	AddCache(*storagegateway.AddCacheInput) (*storagegateway.AddCacheOutput, error)
	AddCacheWithContext(aws.Context, *storagegateway.AddCacheInput, ...request.Option) (*storagegateway.AddCacheOutput, error)
	AddCacheRequest(*storagegateway.AddCacheInput) (*request.Request, *storagegateway.AddCacheOutput)

	AddTagsToResource(*storagegateway.AddTagsToResourceInput) (*storagegateway.AddTagsToResourceOutput, error)
	AddTagsToResourceWithContext(aws.Context, *storagegateway.AddTagsToResourceInput, ...request.Option) (*storagegateway.AddTagsToResourceOutput, error)
	AddTagsToResourceRequest(*storagegateway.AddTagsToResourceInput) (*request.Request, *storagegateway.AddTagsToResourceOutput)

	AddUploadBuffer(*storagegateway.AddUploadBufferInput) (*storagegateway.AddUploadBufferOutput, error)
	AddUploadBufferWithContext(aws.Context, *storagegateway.AddUploadBufferInput, ...request.Option) (*storagegateway.AddUploadBufferOutput, error)
	AddUploadBufferRequest(*storagegateway.AddUploadBufferInput) (*request.Request, *storagegateway.AddUploadBufferOutput)

	AddWorkingStorage(*storagegateway.AddWorkingStorageInput) (*storagegateway.AddWorkingStorageOutput, error)
	AddWorkingStorageWithContext(aws.Context, *storagegateway.AddWorkingStorageInput, ...request.Option) (*storagegateway.AddWorkingStorageOutput, error)
	AddWorkingStorageRequest(*storagegateway.AddWorkingStorageInput) (*request.Request, *storagegateway.AddWorkingStorageOutput)

	AssignTapePool(*storagegateway.AssignTapePoolInput) (*storagegateway.AssignTapePoolOutput, error)
	AssignTapePoolWithContext(aws.Context, *storagegateway.AssignTapePoolInput, ...request.Option) (*storagegateway.AssignTapePoolOutput, error)
	AssignTapePoolRequest(*storagegateway.AssignTapePoolInput) (*request.Request, *storagegateway.AssignTapePoolOutput)

	AttachVolume(*storagegateway.AttachVolumeInput) (*storagegateway.AttachVolumeOutput, error)
	AttachVolumeWithContext(aws.Context, *storagegateway.AttachVolumeInput, ...request.Option) (*storagegateway.AttachVolumeOutput, error)
	AttachVolumeRequest(*storagegateway.AttachVolumeInput) (*request.Request, *storagegateway.AttachVolumeOutput)

	CancelArchival(*storagegateway.CancelArchivalInput) (*storagegateway.CancelArchivalOutput, error)
	CancelArchivalWithContext(aws.Context, *storagegateway.CancelArchivalInput, ...request.Option) (*storagegateway.CancelArchivalOutput, error)
	CancelArchivalRequest(*storagegateway.CancelArchivalInput) (*request.Request, *storagegateway.CancelArchivalOutput)

	CancelRetrieval(*storagegateway.CancelRetrievalInput) (*storagegateway.CancelRetrievalOutput, error)
	CancelRetrievalWithContext(aws.Context, *storagegateway.CancelRetrievalInput, ...request.Option) (*storagegateway.CancelRetrievalOutput, error)
	CancelRetrievalRequest(*storagegateway.CancelRetrievalInput) (*request.Request, *storagegateway.CancelRetrievalOutput)

	CreateCachediSCSIVolume(*storagegateway.CreateCachediSCSIVolumeInput) (*storagegateway.CreateCachediSCSIVolumeOutput, error)
	CreateCachediSCSIVolumeWithContext(aws.Context, *storagegateway.CreateCachediSCSIVolumeInput, ...request.Option) (*storagegateway.CreateCachediSCSIVolumeOutput, error)
	CreateCachediSCSIVolumeRequest(*storagegateway.CreateCachediSCSIVolumeInput) (*request.Request, *storagegateway.CreateCachediSCSIVolumeOutput)

	CreateNFSFileShare(*storagegateway.CreateNFSFileShareInput) (*storagegateway.CreateNFSFileShareOutput, error)
	CreateNFSFileShareWithContext(aws.Context, *storagegateway.CreateNFSFileShareInput, ...request.Option) (*storagegateway.CreateNFSFileShareOutput, error)
	CreateNFSFileShareRequest(*storagegateway.CreateNFSFileShareInput) (*request.Request, *storagegateway.CreateNFSFileShareOutput)

	CreateSMBFileShare(*storagegateway.CreateSMBFileShareInput) (*storagegateway.CreateSMBFileShareOutput, error)
	CreateSMBFileShareWithContext(aws.Context, *storagegateway.CreateSMBFileShareInput, ...request.Option) (*storagegateway.CreateSMBFileShareOutput, error)
	CreateSMBFileShareRequest(*storagegateway.CreateSMBFileShareInput) (*request.Request, *storagegateway.CreateSMBFileShareOutput)

	CreateSnapshot(*storagegateway.CreateSnapshotInput) (*storagegateway.CreateSnapshotOutput, error)
	CreateSnapshotWithContext(aws.Context, *storagegateway.CreateSnapshotInput, ...request.Option) (*storagegateway.CreateSnapshotOutput, error)
	CreateSnapshotRequest(*storagegateway.CreateSnapshotInput) (*request.Request, *storagegateway.CreateSnapshotOutput)

	CreateSnapshotFromVolumeRecoveryPoint(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)
	CreateSnapshotFromVolumeRecoveryPointWithContext(aws.Context, *storagegateway.CreateSnapshotFromVolumeRecoveryPointInput, ...request.Option) (*storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput, error)
	CreateSnapshotFromVolumeRecoveryPointRequest(*storagegateway.CreateSnapshotFromVolumeRecoveryPointInput) (*request.Request, *storagegateway.CreateSnapshotFromVolumeRecoveryPointOutput)

	CreateStorediSCSIVolume(*storagegateway.CreateStorediSCSIVolumeInput) (*storagegateway.CreateStorediSCSIVolumeOutput, error)
	CreateStorediSCSIVolumeWithContext(aws.Context, *storagegateway.CreateStorediSCSIVolumeInput, ...request.Option) (*storagegateway.CreateStorediSCSIVolumeOutput, error)
	CreateStorediSCSIVolumeRequest(*storagegateway.CreateStorediSCSIVolumeInput) (*request.Request, *storagegateway.CreateStorediSCSIVolumeOutput)

	CreateTapePool(*storagegateway.CreateTapePoolInput) (*storagegateway.CreateTapePoolOutput, error)
	CreateTapePoolWithContext(aws.Context, *storagegateway.CreateTapePoolInput, ...request.Option) (*storagegateway.CreateTapePoolOutput, error)
	CreateTapePoolRequest(*storagegateway.CreateTapePoolInput) (*request.Request, *storagegateway.CreateTapePoolOutput)

	CreateTapeWithBarcode(*storagegateway.CreateTapeWithBarcodeInput) (*storagegateway.CreateTapeWithBarcodeOutput, error)
	CreateTapeWithBarcodeWithContext(aws.Context, *storagegateway.CreateTapeWithBarcodeInput, ...request.Option) (*storagegateway.CreateTapeWithBarcodeOutput, error)
	CreateTapeWithBarcodeRequest(*storagegateway.CreateTapeWithBarcodeInput) (*request.Request, *storagegateway.CreateTapeWithBarcodeOutput)

	CreateTapes(*storagegateway.CreateTapesInput) (*storagegateway.CreateTapesOutput, error)
	CreateTapesWithContext(aws.Context, *storagegateway.CreateTapesInput, ...request.Option) (*storagegateway.CreateTapesOutput, error)
	CreateTapesRequest(*storagegateway.CreateTapesInput) (*request.Request, *storagegateway.CreateTapesOutput)

	DeleteAutomaticTapeCreationPolicy(*storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error)
	DeleteAutomaticTapeCreationPolicyWithContext(aws.Context, *storagegateway.DeleteAutomaticTapeCreationPolicyInput, ...request.Option) (*storagegateway.DeleteAutomaticTapeCreationPolicyOutput, error)
	DeleteAutomaticTapeCreationPolicyRequest(*storagegateway.DeleteAutomaticTapeCreationPolicyInput) (*request.Request, *storagegateway.DeleteAutomaticTapeCreationPolicyOutput)

	DeleteBandwidthRateLimit(*storagegateway.DeleteBandwidthRateLimitInput) (*storagegateway.DeleteBandwidthRateLimitOutput, error)
	DeleteBandwidthRateLimitWithContext(aws.Context, *storagegateway.DeleteBandwidthRateLimitInput, ...request.Option) (*storagegateway.DeleteBandwidthRateLimitOutput, error)
	DeleteBandwidthRateLimitRequest(*storagegateway.DeleteBandwidthRateLimitInput) (*request.Request, *storagegateway.DeleteBandwidthRateLimitOutput)

	DeleteChapCredentials(*storagegateway.DeleteChapCredentialsInput) (*storagegateway.DeleteChapCredentialsOutput, error)
	DeleteChapCredentialsWithContext(aws.Context, *storagegateway.DeleteChapCredentialsInput, ...request.Option) (*storagegateway.DeleteChapCredentialsOutput, error)
	DeleteChapCredentialsRequest(*storagegateway.DeleteChapCredentialsInput) (*request.Request, *storagegateway.DeleteChapCredentialsOutput)

	DeleteFileShare(*storagegateway.DeleteFileShareInput) (*storagegateway.DeleteFileShareOutput, error)
	DeleteFileShareWithContext(aws.Context, *storagegateway.DeleteFileShareInput, ...request.Option) (*storagegateway.DeleteFileShareOutput, error)
	DeleteFileShareRequest(*storagegateway.DeleteFileShareInput) (*request.Request, *storagegateway.DeleteFileShareOutput)

	DeleteGateway(*storagegateway.DeleteGatewayInput) (*storagegateway.DeleteGatewayOutput, error)
	DeleteGatewayWithContext(aws.Context, *storagegateway.DeleteGatewayInput, ...request.Option) (*storagegateway.DeleteGatewayOutput, error)
	DeleteGatewayRequest(*storagegateway.DeleteGatewayInput) (*request.Request, *storagegateway.DeleteGatewayOutput)

	DeleteSnapshotSchedule(*storagegateway.DeleteSnapshotScheduleInput) (*storagegateway.DeleteSnapshotScheduleOutput, error)
	DeleteSnapshotScheduleWithContext(aws.Context, *storagegateway.DeleteSnapshotScheduleInput, ...request.Option) (*storagegateway.DeleteSnapshotScheduleOutput, error)
	DeleteSnapshotScheduleRequest(*storagegateway.DeleteSnapshotScheduleInput) (*request.Request, *storagegateway.DeleteSnapshotScheduleOutput)

	DeleteTape(*storagegateway.DeleteTapeInput) (*storagegateway.DeleteTapeOutput, error)
	DeleteTapeWithContext(aws.Context, *storagegateway.DeleteTapeInput, ...request.Option) (*storagegateway.DeleteTapeOutput, error)
	DeleteTapeRequest(*storagegateway.DeleteTapeInput) (*request.Request, *storagegateway.DeleteTapeOutput)

	DeleteTapeArchive(*storagegateway.DeleteTapeArchiveInput) (*storagegateway.DeleteTapeArchiveOutput, error)
	DeleteTapeArchiveWithContext(aws.Context, *storagegateway.DeleteTapeArchiveInput, ...request.Option) (*storagegateway.DeleteTapeArchiveOutput, error)
	DeleteTapeArchiveRequest(*storagegateway.DeleteTapeArchiveInput) (*request.Request, *storagegateway.DeleteTapeArchiveOutput)

	DeleteTapePool(*storagegateway.DeleteTapePoolInput) (*storagegateway.DeleteTapePoolOutput, error)
	DeleteTapePoolWithContext(aws.Context, *storagegateway.DeleteTapePoolInput, ...request.Option) (*storagegateway.DeleteTapePoolOutput, error)
	DeleteTapePoolRequest(*storagegateway.DeleteTapePoolInput) (*request.Request, *storagegateway.DeleteTapePoolOutput)

	DeleteVolume(*storagegateway.DeleteVolumeInput) (*storagegateway.DeleteVolumeOutput, error)
	DeleteVolumeWithContext(aws.Context, *storagegateway.DeleteVolumeInput, ...request.Option) (*storagegateway.DeleteVolumeOutput, error)
	DeleteVolumeRequest(*storagegateway.DeleteVolumeInput) (*request.Request, *storagegateway.DeleteVolumeOutput)

	DescribeAvailabilityMonitorTest(*storagegateway.DescribeAvailabilityMonitorTestInput) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error)
	DescribeAvailabilityMonitorTestWithContext(aws.Context, *storagegateway.DescribeAvailabilityMonitorTestInput, ...request.Option) (*storagegateway.DescribeAvailabilityMonitorTestOutput, error)
	DescribeAvailabilityMonitorTestRequest(*storagegateway.DescribeAvailabilityMonitorTestInput) (*request.Request, *storagegateway.DescribeAvailabilityMonitorTestOutput)

	DescribeBandwidthRateLimit(*storagegateway.DescribeBandwidthRateLimitInput) (*storagegateway.DescribeBandwidthRateLimitOutput, error)
	DescribeBandwidthRateLimitWithContext(aws.Context, *storagegateway.DescribeBandwidthRateLimitInput, ...request.Option) (*storagegateway.DescribeBandwidthRateLimitOutput, error)
	DescribeBandwidthRateLimitRequest(*storagegateway.DescribeBandwidthRateLimitInput) (*request.Request, *storagegateway.DescribeBandwidthRateLimitOutput)

	DescribeCache(*storagegateway.DescribeCacheInput) (*storagegateway.DescribeCacheOutput, error)
	DescribeCacheWithContext(aws.Context, *storagegateway.DescribeCacheInput, ...request.Option) (*storagegateway.DescribeCacheOutput, error)
	DescribeCacheRequest(*storagegateway.DescribeCacheInput) (*request.Request, *storagegateway.DescribeCacheOutput)

	DescribeCachediSCSIVolumes(*storagegateway.DescribeCachediSCSIVolumesInput) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)
	DescribeCachediSCSIVolumesWithContext(aws.Context, *storagegateway.DescribeCachediSCSIVolumesInput, ...request.Option) (*storagegateway.DescribeCachediSCSIVolumesOutput, error)
	DescribeCachediSCSIVolumesRequest(*storagegateway.DescribeCachediSCSIVolumesInput) (*request.Request, *storagegateway.DescribeCachediSCSIVolumesOutput)

	DescribeChapCredentials(*storagegateway.DescribeChapCredentialsInput) (*storagegateway.DescribeChapCredentialsOutput, error)
	DescribeChapCredentialsWithContext(aws.Context, *storagegateway.DescribeChapCredentialsInput, ...request.Option) (*storagegateway.DescribeChapCredentialsOutput, error)
	DescribeChapCredentialsRequest(*storagegateway.DescribeChapCredentialsInput) (*request.Request, *storagegateway.DescribeChapCredentialsOutput)

	DescribeGatewayInformation(*storagegateway.DescribeGatewayInformationInput) (*storagegateway.DescribeGatewayInformationOutput, error)
	DescribeGatewayInformationWithContext(aws.Context, *storagegateway.DescribeGatewayInformationInput, ...request.Option) (*storagegateway.DescribeGatewayInformationOutput, error)
	DescribeGatewayInformationRequest(*storagegateway.DescribeGatewayInformationInput) (*request.Request, *storagegateway.DescribeGatewayInformationOutput)

	DescribeMaintenanceStartTime(*storagegateway.DescribeMaintenanceStartTimeInput) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)
	DescribeMaintenanceStartTimeWithContext(aws.Context, *storagegateway.DescribeMaintenanceStartTimeInput, ...request.Option) (*storagegateway.DescribeMaintenanceStartTimeOutput, error)
	DescribeMaintenanceStartTimeRequest(*storagegateway.DescribeMaintenanceStartTimeInput) (*request.Request, *storagegateway.DescribeMaintenanceStartTimeOutput)

	DescribeNFSFileShares(*storagegateway.DescribeNFSFileSharesInput) (*storagegateway.DescribeNFSFileSharesOutput, error)
	DescribeNFSFileSharesWithContext(aws.Context, *storagegateway.DescribeNFSFileSharesInput, ...request.Option) (*storagegateway.DescribeNFSFileSharesOutput, error)
	DescribeNFSFileSharesRequest(*storagegateway.DescribeNFSFileSharesInput) (*request.Request, *storagegateway.DescribeNFSFileSharesOutput)

	DescribeSMBFileShares(*storagegateway.DescribeSMBFileSharesInput) (*storagegateway.DescribeSMBFileSharesOutput, error)
	DescribeSMBFileSharesWithContext(aws.Context, *storagegateway.DescribeSMBFileSharesInput, ...request.Option) (*storagegateway.DescribeSMBFileSharesOutput, error)
	DescribeSMBFileSharesRequest(*storagegateway.DescribeSMBFileSharesInput) (*request.Request, *storagegateway.DescribeSMBFileSharesOutput)

	DescribeSMBSettings(*storagegateway.DescribeSMBSettingsInput) (*storagegateway.DescribeSMBSettingsOutput, error)
	DescribeSMBSettingsWithContext(aws.Context, *storagegateway.DescribeSMBSettingsInput, ...request.Option) (*storagegateway.DescribeSMBSettingsOutput, error)
	DescribeSMBSettingsRequest(*storagegateway.DescribeSMBSettingsInput) (*request.Request, *storagegateway.DescribeSMBSettingsOutput)

	DescribeSnapshotSchedule(*storagegateway.DescribeSnapshotScheduleInput) (*storagegateway.DescribeSnapshotScheduleOutput, error)
	DescribeSnapshotScheduleWithContext(aws.Context, *storagegateway.DescribeSnapshotScheduleInput, ...request.Option) (*storagegateway.DescribeSnapshotScheduleOutput, error)
	DescribeSnapshotScheduleRequest(*storagegateway.DescribeSnapshotScheduleInput) (*request.Request, *storagegateway.DescribeSnapshotScheduleOutput)

	DescribeStorediSCSIVolumes(*storagegateway.DescribeStorediSCSIVolumesInput) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)
	DescribeStorediSCSIVolumesWithContext(aws.Context, *storagegateway.DescribeStorediSCSIVolumesInput, ...request.Option) (*storagegateway.DescribeStorediSCSIVolumesOutput, error)
	DescribeStorediSCSIVolumesRequest(*storagegateway.DescribeStorediSCSIVolumesInput) (*request.Request, *storagegateway.DescribeStorediSCSIVolumesOutput)

	DescribeTapeArchives(*storagegateway.DescribeTapeArchivesInput) (*storagegateway.DescribeTapeArchivesOutput, error)
	DescribeTapeArchivesWithContext(aws.Context, *storagegateway.DescribeTapeArchivesInput, ...request.Option) (*storagegateway.DescribeTapeArchivesOutput, error)
	DescribeTapeArchivesRequest(*storagegateway.DescribeTapeArchivesInput) (*request.Request, *storagegateway.DescribeTapeArchivesOutput)

	DescribeTapeArchivesPages(*storagegateway.DescribeTapeArchivesInput, func(*storagegateway.DescribeTapeArchivesOutput, bool) bool) error
	DescribeTapeArchivesPagesWithContext(aws.Context, *storagegateway.DescribeTapeArchivesInput, func(*storagegateway.DescribeTapeArchivesOutput, bool) bool, ...request.Option) error

	DescribeTapeRecoveryPoints(*storagegateway.DescribeTapeRecoveryPointsInput) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)
	DescribeTapeRecoveryPointsWithContext(aws.Context, *storagegateway.DescribeTapeRecoveryPointsInput, ...request.Option) (*storagegateway.DescribeTapeRecoveryPointsOutput, error)
	DescribeTapeRecoveryPointsRequest(*storagegateway.DescribeTapeRecoveryPointsInput) (*request.Request, *storagegateway.DescribeTapeRecoveryPointsOutput)

	DescribeTapeRecoveryPointsPages(*storagegateway.DescribeTapeRecoveryPointsInput, func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool) error
	DescribeTapeRecoveryPointsPagesWithContext(aws.Context, *storagegateway.DescribeTapeRecoveryPointsInput, func(*storagegateway.DescribeTapeRecoveryPointsOutput, bool) bool, ...request.Option) error

	DescribeTapes(*storagegateway.DescribeTapesInput) (*storagegateway.DescribeTapesOutput, error)
	DescribeTapesWithContext(aws.Context, *storagegateway.DescribeTapesInput, ...request.Option) (*storagegateway.DescribeTapesOutput, error)
	DescribeTapesRequest(*storagegateway.DescribeTapesInput) (*request.Request, *storagegateway.DescribeTapesOutput)

	DescribeTapesPages(*storagegateway.DescribeTapesInput, func(*storagegateway.DescribeTapesOutput, bool) bool) error
	DescribeTapesPagesWithContext(aws.Context, *storagegateway.DescribeTapesInput, func(*storagegateway.DescribeTapesOutput, bool) bool, ...request.Option) error

	DescribeUploadBuffer(*storagegateway.DescribeUploadBufferInput) (*storagegateway.DescribeUploadBufferOutput, error)
	DescribeUploadBufferWithContext(aws.Context, *storagegateway.DescribeUploadBufferInput, ...request.Option) (*storagegateway.DescribeUploadBufferOutput, error)
	DescribeUploadBufferRequest(*storagegateway.DescribeUploadBufferInput) (*request.Request, *storagegateway.DescribeUploadBufferOutput)

	DescribeVTLDevices(*storagegateway.DescribeVTLDevicesInput) (*storagegateway.DescribeVTLDevicesOutput, error)
	DescribeVTLDevicesWithContext(aws.Context, *storagegateway.DescribeVTLDevicesInput, ...request.Option) (*storagegateway.DescribeVTLDevicesOutput, error)
	DescribeVTLDevicesRequest(*storagegateway.DescribeVTLDevicesInput) (*request.Request, *storagegateway.DescribeVTLDevicesOutput)

	DescribeVTLDevicesPages(*storagegateway.DescribeVTLDevicesInput, func(*storagegateway.DescribeVTLDevicesOutput, bool) bool) error
	DescribeVTLDevicesPagesWithContext(aws.Context, *storagegateway.DescribeVTLDevicesInput, func(*storagegateway.DescribeVTLDevicesOutput, bool) bool, ...request.Option) error

	DescribeWorkingStorage(*storagegateway.DescribeWorkingStorageInput) (*storagegateway.DescribeWorkingStorageOutput, error)
	DescribeWorkingStorageWithContext(aws.Context, *storagegateway.DescribeWorkingStorageInput, ...request.Option) (*storagegateway.DescribeWorkingStorageOutput, error)
	DescribeWorkingStorageRequest(*storagegateway.DescribeWorkingStorageInput) (*request.Request, *storagegateway.DescribeWorkingStorageOutput)

	DetachVolume(*storagegateway.DetachVolumeInput) (*storagegateway.DetachVolumeOutput, error)
	DetachVolumeWithContext(aws.Context, *storagegateway.DetachVolumeInput, ...request.Option) (*storagegateway.DetachVolumeOutput, error)
	DetachVolumeRequest(*storagegateway.DetachVolumeInput) (*request.Request, *storagegateway.DetachVolumeOutput)

	DisableGateway(*storagegateway.DisableGatewayInput) (*storagegateway.DisableGatewayOutput, error)
	DisableGatewayWithContext(aws.Context, *storagegateway.DisableGatewayInput, ...request.Option) (*storagegateway.DisableGatewayOutput, error)
	DisableGatewayRequest(*storagegateway.DisableGatewayInput) (*request.Request, *storagegateway.DisableGatewayOutput)

	JoinDomain(*storagegateway.JoinDomainInput) (*storagegateway.JoinDomainOutput, error)
	JoinDomainWithContext(aws.Context, *storagegateway.JoinDomainInput, ...request.Option) (*storagegateway.JoinDomainOutput, error)
	JoinDomainRequest(*storagegateway.JoinDomainInput) (*request.Request, *storagegateway.JoinDomainOutput)

	ListAutomaticTapeCreationPolicies(*storagegateway.ListAutomaticTapeCreationPoliciesInput) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error)
	ListAutomaticTapeCreationPoliciesWithContext(aws.Context, *storagegateway.ListAutomaticTapeCreationPoliciesInput, ...request.Option) (*storagegateway.ListAutomaticTapeCreationPoliciesOutput, error)
	ListAutomaticTapeCreationPoliciesRequest(*storagegateway.ListAutomaticTapeCreationPoliciesInput) (*request.Request, *storagegateway.ListAutomaticTapeCreationPoliciesOutput)

	ListFileShares(*storagegateway.ListFileSharesInput) (*storagegateway.ListFileSharesOutput, error)
	ListFileSharesWithContext(aws.Context, *storagegateway.ListFileSharesInput, ...request.Option) (*storagegateway.ListFileSharesOutput, error)
	ListFileSharesRequest(*storagegateway.ListFileSharesInput) (*request.Request, *storagegateway.ListFileSharesOutput)

	ListFileSharesPages(*storagegateway.ListFileSharesInput, func(*storagegateway.ListFileSharesOutput, bool) bool) error
	ListFileSharesPagesWithContext(aws.Context, *storagegateway.ListFileSharesInput, func(*storagegateway.ListFileSharesOutput, bool) bool, ...request.Option) error

	ListGateways(*storagegateway.ListGatewaysInput) (*storagegateway.ListGatewaysOutput, error)
	ListGatewaysWithContext(aws.Context, *storagegateway.ListGatewaysInput, ...request.Option) (*storagegateway.ListGatewaysOutput, error)
	ListGatewaysRequest(*storagegateway.ListGatewaysInput) (*request.Request, *storagegateway.ListGatewaysOutput)

	ListGatewaysPages(*storagegateway.ListGatewaysInput, func(*storagegateway.ListGatewaysOutput, bool) bool) error
	ListGatewaysPagesWithContext(aws.Context, *storagegateway.ListGatewaysInput, func(*storagegateway.ListGatewaysOutput, bool) bool, ...request.Option) error

	ListLocalDisks(*storagegateway.ListLocalDisksInput) (*storagegateway.ListLocalDisksOutput, error)
	ListLocalDisksWithContext(aws.Context, *storagegateway.ListLocalDisksInput, ...request.Option) (*storagegateway.ListLocalDisksOutput, error)
	ListLocalDisksRequest(*storagegateway.ListLocalDisksInput) (*request.Request, *storagegateway.ListLocalDisksOutput)

	ListTagsForResource(*storagegateway.ListTagsForResourceInput) (*storagegateway.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *storagegateway.ListTagsForResourceInput, ...request.Option) (*storagegateway.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*storagegateway.ListTagsForResourceInput) (*request.Request, *storagegateway.ListTagsForResourceOutput)

	ListTagsForResourcePages(*storagegateway.ListTagsForResourceInput, func(*storagegateway.ListTagsForResourceOutput, bool) bool) error
	ListTagsForResourcePagesWithContext(aws.Context, *storagegateway.ListTagsForResourceInput, func(*storagegateway.ListTagsForResourceOutput, bool) bool, ...request.Option) error

	ListTapePools(*storagegateway.ListTapePoolsInput) (*storagegateway.ListTapePoolsOutput, error)
	ListTapePoolsWithContext(aws.Context, *storagegateway.ListTapePoolsInput, ...request.Option) (*storagegateway.ListTapePoolsOutput, error)
	ListTapePoolsRequest(*storagegateway.ListTapePoolsInput) (*request.Request, *storagegateway.ListTapePoolsOutput)

	ListTapes(*storagegateway.ListTapesInput) (*storagegateway.ListTapesOutput, error)
	ListTapesWithContext(aws.Context, *storagegateway.ListTapesInput, ...request.Option) (*storagegateway.ListTapesOutput, error)
	ListTapesRequest(*storagegateway.ListTapesInput) (*request.Request, *storagegateway.ListTapesOutput)

	ListTapesPages(*storagegateway.ListTapesInput, func(*storagegateway.ListTapesOutput, bool) bool) error
	ListTapesPagesWithContext(aws.Context, *storagegateway.ListTapesInput, func(*storagegateway.ListTapesOutput, bool) bool, ...request.Option) error

	ListVolumeInitiators(*storagegateway.ListVolumeInitiatorsInput) (*storagegateway.ListVolumeInitiatorsOutput, error)
	ListVolumeInitiatorsWithContext(aws.Context, *storagegateway.ListVolumeInitiatorsInput, ...request.Option) (*storagegateway.ListVolumeInitiatorsOutput, error)
	ListVolumeInitiatorsRequest(*storagegateway.ListVolumeInitiatorsInput) (*request.Request, *storagegateway.ListVolumeInitiatorsOutput)

	ListVolumeRecoveryPoints(*storagegateway.ListVolumeRecoveryPointsInput) (*storagegateway.ListVolumeRecoveryPointsOutput, error)
	ListVolumeRecoveryPointsWithContext(aws.Context, *storagegateway.ListVolumeRecoveryPointsInput, ...request.Option) (*storagegateway.ListVolumeRecoveryPointsOutput, error)
	ListVolumeRecoveryPointsRequest(*storagegateway.ListVolumeRecoveryPointsInput) (*request.Request, *storagegateway.ListVolumeRecoveryPointsOutput)

	ListVolumes(*storagegateway.ListVolumesInput) (*storagegateway.ListVolumesOutput, error)
	ListVolumesWithContext(aws.Context, *storagegateway.ListVolumesInput, ...request.Option) (*storagegateway.ListVolumesOutput, error)
	ListVolumesRequest(*storagegateway.ListVolumesInput) (*request.Request, *storagegateway.ListVolumesOutput)

	ListVolumesPages(*storagegateway.ListVolumesInput, func(*storagegateway.ListVolumesOutput, bool) bool) error
	ListVolumesPagesWithContext(aws.Context, *storagegateway.ListVolumesInput, func(*storagegateway.ListVolumesOutput, bool) bool, ...request.Option) error

	NotifyWhenUploaded(*storagegateway.NotifyWhenUploadedInput) (*storagegateway.NotifyWhenUploadedOutput, error)
	NotifyWhenUploadedWithContext(aws.Context, *storagegateway.NotifyWhenUploadedInput, ...request.Option) (*storagegateway.NotifyWhenUploadedOutput, error)
	NotifyWhenUploadedRequest(*storagegateway.NotifyWhenUploadedInput) (*request.Request, *storagegateway.NotifyWhenUploadedOutput)

	RefreshCache(*storagegateway.RefreshCacheInput) (*storagegateway.RefreshCacheOutput, error)
	RefreshCacheWithContext(aws.Context, *storagegateway.RefreshCacheInput, ...request.Option) (*storagegateway.RefreshCacheOutput, error)
	RefreshCacheRequest(*storagegateway.RefreshCacheInput) (*request.Request, *storagegateway.RefreshCacheOutput)

	RemoveTagsFromResource(*storagegateway.RemoveTagsFromResourceInput) (*storagegateway.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceWithContext(aws.Context, *storagegateway.RemoveTagsFromResourceInput, ...request.Option) (*storagegateway.RemoveTagsFromResourceOutput, error)
	RemoveTagsFromResourceRequest(*storagegateway.RemoveTagsFromResourceInput) (*request.Request, *storagegateway.RemoveTagsFromResourceOutput)

	ResetCache(*storagegateway.ResetCacheInput) (*storagegateway.ResetCacheOutput, error)
	ResetCacheWithContext(aws.Context, *storagegateway.ResetCacheInput, ...request.Option) (*storagegateway.ResetCacheOutput, error)
	ResetCacheRequest(*storagegateway.ResetCacheInput) (*request.Request, *storagegateway.ResetCacheOutput)

	RetrieveTapeArchive(*storagegateway.RetrieveTapeArchiveInput) (*storagegateway.RetrieveTapeArchiveOutput, error)
	RetrieveTapeArchiveWithContext(aws.Context, *storagegateway.RetrieveTapeArchiveInput, ...request.Option) (*storagegateway.RetrieveTapeArchiveOutput, error)
	RetrieveTapeArchiveRequest(*storagegateway.RetrieveTapeArchiveInput) (*request.Request, *storagegateway.RetrieveTapeArchiveOutput)

	RetrieveTapeRecoveryPoint(*storagegateway.RetrieveTapeRecoveryPointInput) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)
	RetrieveTapeRecoveryPointWithContext(aws.Context, *storagegateway.RetrieveTapeRecoveryPointInput, ...request.Option) (*storagegateway.RetrieveTapeRecoveryPointOutput, error)
	RetrieveTapeRecoveryPointRequest(*storagegateway.RetrieveTapeRecoveryPointInput) (*request.Request, *storagegateway.RetrieveTapeRecoveryPointOutput)

	SetLocalConsolePassword(*storagegateway.SetLocalConsolePasswordInput) (*storagegateway.SetLocalConsolePasswordOutput, error)
	SetLocalConsolePasswordWithContext(aws.Context, *storagegateway.SetLocalConsolePasswordInput, ...request.Option) (*storagegateway.SetLocalConsolePasswordOutput, error)
	SetLocalConsolePasswordRequest(*storagegateway.SetLocalConsolePasswordInput) (*request.Request, *storagegateway.SetLocalConsolePasswordOutput)

	SetSMBGuestPassword(*storagegateway.SetSMBGuestPasswordInput) (*storagegateway.SetSMBGuestPasswordOutput, error)
	SetSMBGuestPasswordWithContext(aws.Context, *storagegateway.SetSMBGuestPasswordInput, ...request.Option) (*storagegateway.SetSMBGuestPasswordOutput, error)
	SetSMBGuestPasswordRequest(*storagegateway.SetSMBGuestPasswordInput) (*request.Request, *storagegateway.SetSMBGuestPasswordOutput)

	ShutdownGateway(*storagegateway.ShutdownGatewayInput) (*storagegateway.ShutdownGatewayOutput, error)
	ShutdownGatewayWithContext(aws.Context, *storagegateway.ShutdownGatewayInput, ...request.Option) (*storagegateway.ShutdownGatewayOutput, error)
	ShutdownGatewayRequest(*storagegateway.ShutdownGatewayInput) (*request.Request, *storagegateway.ShutdownGatewayOutput)

	StartAvailabilityMonitorTest(*storagegateway.StartAvailabilityMonitorTestInput) (*storagegateway.StartAvailabilityMonitorTestOutput, error)
	StartAvailabilityMonitorTestWithContext(aws.Context, *storagegateway.StartAvailabilityMonitorTestInput, ...request.Option) (*storagegateway.StartAvailabilityMonitorTestOutput, error)
	StartAvailabilityMonitorTestRequest(*storagegateway.StartAvailabilityMonitorTestInput) (*request.Request, *storagegateway.StartAvailabilityMonitorTestOutput)

	StartGateway(*storagegateway.StartGatewayInput) (*storagegateway.StartGatewayOutput, error)
	StartGatewayWithContext(aws.Context, *storagegateway.StartGatewayInput, ...request.Option) (*storagegateway.StartGatewayOutput, error)
	StartGatewayRequest(*storagegateway.StartGatewayInput) (*request.Request, *storagegateway.StartGatewayOutput)

	UpdateAutomaticTapeCreationPolicy(*storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error)
	UpdateAutomaticTapeCreationPolicyWithContext(aws.Context, *storagegateway.UpdateAutomaticTapeCreationPolicyInput, ...request.Option) (*storagegateway.UpdateAutomaticTapeCreationPolicyOutput, error)
	UpdateAutomaticTapeCreationPolicyRequest(*storagegateway.UpdateAutomaticTapeCreationPolicyInput) (*request.Request, *storagegateway.UpdateAutomaticTapeCreationPolicyOutput)

	UpdateBandwidthRateLimit(*storagegateway.UpdateBandwidthRateLimitInput) (*storagegateway.UpdateBandwidthRateLimitOutput, error)
	UpdateBandwidthRateLimitWithContext(aws.Context, *storagegateway.UpdateBandwidthRateLimitInput, ...request.Option) (*storagegateway.UpdateBandwidthRateLimitOutput, error)
	UpdateBandwidthRateLimitRequest(*storagegateway.UpdateBandwidthRateLimitInput) (*request.Request, *storagegateway.UpdateBandwidthRateLimitOutput)

	UpdateChapCredentials(*storagegateway.UpdateChapCredentialsInput) (*storagegateway.UpdateChapCredentialsOutput, error)
	UpdateChapCredentialsWithContext(aws.Context, *storagegateway.UpdateChapCredentialsInput, ...request.Option) (*storagegateway.UpdateChapCredentialsOutput, error)
	UpdateChapCredentialsRequest(*storagegateway.UpdateChapCredentialsInput) (*request.Request, *storagegateway.UpdateChapCredentialsOutput)

	UpdateGatewayInformation(*storagegateway.UpdateGatewayInformationInput) (*storagegateway.UpdateGatewayInformationOutput, error)
	UpdateGatewayInformationWithContext(aws.Context, *storagegateway.UpdateGatewayInformationInput, ...request.Option) (*storagegateway.UpdateGatewayInformationOutput, error)
	UpdateGatewayInformationRequest(*storagegateway.UpdateGatewayInformationInput) (*request.Request, *storagegateway.UpdateGatewayInformationOutput)

	UpdateGatewaySoftwareNow(*storagegateway.UpdateGatewaySoftwareNowInput) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)
	UpdateGatewaySoftwareNowWithContext(aws.Context, *storagegateway.UpdateGatewaySoftwareNowInput, ...request.Option) (*storagegateway.UpdateGatewaySoftwareNowOutput, error)
	UpdateGatewaySoftwareNowRequest(*storagegateway.UpdateGatewaySoftwareNowInput) (*request.Request, *storagegateway.UpdateGatewaySoftwareNowOutput)

	UpdateMaintenanceStartTime(*storagegateway.UpdateMaintenanceStartTimeInput) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)
	UpdateMaintenanceStartTimeWithContext(aws.Context, *storagegateway.UpdateMaintenanceStartTimeInput, ...request.Option) (*storagegateway.UpdateMaintenanceStartTimeOutput, error)
	UpdateMaintenanceStartTimeRequest(*storagegateway.UpdateMaintenanceStartTimeInput) (*request.Request, *storagegateway.UpdateMaintenanceStartTimeOutput)

	UpdateNFSFileShare(*storagegateway.UpdateNFSFileShareInput) (*storagegateway.UpdateNFSFileShareOutput, error)
	UpdateNFSFileShareWithContext(aws.Context, *storagegateway.UpdateNFSFileShareInput, ...request.Option) (*storagegateway.UpdateNFSFileShareOutput, error)
	UpdateNFSFileShareRequest(*storagegateway.UpdateNFSFileShareInput) (*request.Request, *storagegateway.UpdateNFSFileShareOutput)

	UpdateSMBFileShare(*storagegateway.UpdateSMBFileShareInput) (*storagegateway.UpdateSMBFileShareOutput, error)
	UpdateSMBFileShareWithContext(aws.Context, *storagegateway.UpdateSMBFileShareInput, ...request.Option) (*storagegateway.UpdateSMBFileShareOutput, error)
	UpdateSMBFileShareRequest(*storagegateway.UpdateSMBFileShareInput) (*request.Request, *storagegateway.UpdateSMBFileShareOutput)

	UpdateSMBSecurityStrategy(*storagegateway.UpdateSMBSecurityStrategyInput) (*storagegateway.UpdateSMBSecurityStrategyOutput, error)
	UpdateSMBSecurityStrategyWithContext(aws.Context, *storagegateway.UpdateSMBSecurityStrategyInput, ...request.Option) (*storagegateway.UpdateSMBSecurityStrategyOutput, error)
	UpdateSMBSecurityStrategyRequest(*storagegateway.UpdateSMBSecurityStrategyInput) (*request.Request, *storagegateway.UpdateSMBSecurityStrategyOutput)

	UpdateSnapshotSchedule(*storagegateway.UpdateSnapshotScheduleInput) (*storagegateway.UpdateSnapshotScheduleOutput, error)
	UpdateSnapshotScheduleWithContext(aws.Context, *storagegateway.UpdateSnapshotScheduleInput, ...request.Option) (*storagegateway.UpdateSnapshotScheduleOutput, error)
	UpdateSnapshotScheduleRequest(*storagegateway.UpdateSnapshotScheduleInput) (*request.Request, *storagegateway.UpdateSnapshotScheduleOutput)

	UpdateVTLDeviceType(*storagegateway.UpdateVTLDeviceTypeInput) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
	UpdateVTLDeviceTypeWithContext(aws.Context, *storagegateway.UpdateVTLDeviceTypeInput, ...request.Option) (*storagegateway.UpdateVTLDeviceTypeOutput, error)
	UpdateVTLDeviceTypeRequest(*storagegateway.UpdateVTLDeviceTypeInput) (*request.Request, *storagegateway.UpdateVTLDeviceTypeOutput)
}

var _ StorageGatewayAPI = (*storagegateway.StorageGateway)(nil)
