// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package costandusagereportservice

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeDuplicateReportNameException for service response error code
	// "DuplicateReportNameException".
	//
	// A report with the specified name already exists in the account. Specify a
	// different report name.
	ErrCodeDuplicateReportNameException = "DuplicateReportNameException"

	// ErrCodeInternalErrorException for service response error code
	// "InternalErrorException".
	//
	// An error on the server occurred during the processing of your request. Try
	// again later.
	ErrCodeInternalErrorException = "InternalErrorException"

	// ErrCodeReportLimitReachedException for service response error code
	// "ReportLimitReachedException".
	//
	// This account already has five reports defined. To define a new report, you
	// must delete an existing report.
	ErrCodeReportLimitReachedException = "ReportLimitReachedException"

	// ErrCodeValidationException for service response error code
	// "ValidationException".
	//
	// The input fails to satisfy the constraints specified by an AWS service.
	ErrCodeValidationException = "ValidationException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"DuplicateReportNameException": newErrorDuplicateReportNameException,
	"InternalErrorException":       newErrorInternalErrorException,
	"ReportLimitReachedException":  newErrorReportLimitReachedException,
	"ValidationException":          newErrorValidationException,
}
