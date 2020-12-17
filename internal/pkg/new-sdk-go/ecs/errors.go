// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"github.com/aws/aws-sdk-go/private/protocol"
)

const (

	// ErrCodeAccessDeniedException for service response error code
	// "AccessDeniedException".
	ErrCodeAccessDeniedException = "AccessDeniedException"

	// ErrCodeAttributeLimitExceededException for service response error code
	// "AttributeLimitExceededException".
	ErrCodeAttributeLimitExceededException = "AttributeLimitExceededException"

	// ErrCodeBlockedException for service response error code
	// "BlockedException".
	ErrCodeBlockedException = "BlockedException"

	// ErrCodeClientException for service response error code
	// "ClientException".
	ErrCodeClientException = "ClientException"

	// ErrCodeClusterContainsContainerInstancesException for service response error code
	// "ClusterContainsContainerInstancesException".
	ErrCodeClusterContainsContainerInstancesException = "ClusterContainsContainerInstancesException"

	// ErrCodeClusterContainsServicesException for service response error code
	// "ClusterContainsServicesException".
	ErrCodeClusterContainsServicesException = "ClusterContainsServicesException"

	// ErrCodeClusterContainsTasksException for service response error code
	// "ClusterContainsTasksException".
	ErrCodeClusterContainsTasksException = "ClusterContainsTasksException"

	// ErrCodeClusterNotFoundException for service response error code
	// "ClusterNotFoundException".
	ErrCodeClusterNotFoundException = "ClusterNotFoundException"

	// ErrCodeCredentialProviderInvocationException for service response error code
	// "CredentialProviderInvocationException".
	ErrCodeCredentialProviderInvocationException = "CredentialProviderInvocationException"

	// ErrCodeInvalidParameterException for service response error code
	// "InvalidParameterException".
	ErrCodeInvalidParameterException = "InvalidParameterException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeMissingVersionException for service response error code
	// "MissingVersionException".
	ErrCodeMissingVersionException = "MissingVersionException"

	// ErrCodeNoUpdateAvailableException for service response error code
	// "NoUpdateAvailableException".
	ErrCodeNoUpdateAvailableException = "NoUpdateAvailableException"

	// ErrCodePlatformTaskDefinitionIncompatibilityException for service response error code
	// "PlatformTaskDefinitionIncompatibilityException".
	ErrCodePlatformTaskDefinitionIncompatibilityException = "PlatformTaskDefinitionIncompatibilityException"

	// ErrCodePlatformUnknownException for service response error code
	// "PlatformUnknownException".
	ErrCodePlatformUnknownException = "PlatformUnknownException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeServerException for service response error code
	// "ServerException".
	ErrCodeServerException = "ServerException"

	// ErrCodeServiceNotActiveException for service response error code
	// "ServiceNotActiveException".
	ErrCodeServiceNotActiveException = "ServiceNotActiveException"

	// ErrCodeServiceNotFoundException for service response error code
	// "ServiceNotFoundException".
	ErrCodeServiceNotFoundException = "ServiceNotFoundException"

	// ErrCodeTargetNotConnectedException for service response error code
	// "TargetNotConnectedException".
	ErrCodeTargetNotConnectedException = "TargetNotConnectedException"

	// ErrCodeTargetNotFoundException for service response error code
	// "TargetNotFoundException".
	ErrCodeTargetNotFoundException = "TargetNotFoundException"

	// ErrCodeTaskSetNotFoundException for service response error code
	// "TaskSetNotFoundException".
	ErrCodeTaskSetNotFoundException = "TaskSetNotFoundException"

	// ErrCodeUnsupportedFeatureException for service response error code
	// "UnsupportedFeatureException".
	ErrCodeUnsupportedFeatureException = "UnsupportedFeatureException"

	// ErrCodeUpdateInProgressException for service response error code
	// "UpdateInProgressException".
	ErrCodeUpdateInProgressException = "UpdateInProgressException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"AccessDeniedException":                          newErrorAccessDeniedException,
	"AttributeLimitExceededException":                newErrorAttributeLimitExceededException,
	"BlockedException":                               newErrorBlockedException,
	"ClientException":                                newErrorClientException,
	"ClusterContainsContainerInstancesException":     newErrorClusterContainsContainerInstancesException,
	"ClusterContainsServicesException":               newErrorClusterContainsServicesException,
	"ClusterContainsTasksException":                  newErrorClusterContainsTasksException,
	"ClusterNotFoundException":                       newErrorClusterNotFoundException,
	"CredentialProviderInvocationException":          newErrorCredentialProviderInvocationException,
	"InvalidParameterException":                      newErrorInvalidParameterException,
	"LimitExceededException":                         newErrorLimitExceededException,
	"MissingVersionException":                        newErrorMissingVersionException,
	"NoUpdateAvailableException":                     newErrorNoUpdateAvailableException,
	"PlatformTaskDefinitionIncompatibilityException": newErrorPlatformTaskDefinitionIncompatibilityException,
	"PlatformUnknownException":                       newErrorPlatformUnknownException,
	"ResourceNotFoundException":                      newErrorResourceNotFoundException,
	"ServerException":                                newErrorServerException,
	"ServiceNotActiveException":                      newErrorServiceNotActiveException,
	"ServiceNotFoundException":                       newErrorServiceNotFoundException,
	"TargetNotConnectedException":                    newErrorTargetNotConnectedException,
	"TargetNotFoundException":                        newErrorTargetNotFoundException,
	"TaskSetNotFoundException":                       newErrorTaskSetNotFoundException,
	"UnsupportedFeatureException":                    newErrorUnsupportedFeatureException,
	"UpdateInProgressException":                      newErrorUpdateInProgressException,
}
