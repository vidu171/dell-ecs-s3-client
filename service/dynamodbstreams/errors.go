// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodbstreams

import (
	"github.com/vidu171/dell-ecs-s3-client/private/protocol"
)

const (

	// ErrCodeExpiredIteratorException for service response error code
	// "ExpiredIteratorException".
	//
	// The shard iterator has expired and can no longer be used to retrieve stream
	// records. A shard iterator expires 15 minutes after it is retrieved using
	// the GetShardIterator action.
	ErrCodeExpiredIteratorException = "ExpiredIteratorException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// An error occurred on the server side.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// There is no limit to the number of daily on-demand backups that can be taken.
	//
	// Up to 50 simultaneous table operations are allowed per account. These operations
	// include CreateTable, UpdateTable, DeleteTable,UpdateTimeToLive, RestoreTableFromBackup,
	// and RestoreTableToPointInTime.
	//
	// The only exception is when you are creating a table with one or more secondary
	// indexes. You can have up to 25 such requests running at a time; however,
	// if the table or index specifications are complex, DynamoDB might temporarily
	// reduce the number of concurrent operations.
	//
	// There is a soft account quota of 256 tables.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The operation tried to access a nonexistent table or index. The resource
	// might not be specified correctly, or its status might not be ACTIVE.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTrimmedDataAccessException for service response error code
	// "TrimmedDataAccessException".
	//
	// The operation attempted to read past the oldest stream record in a shard.
	//
	// In DynamoDB Streams, there is a 24 hour limit on data retention. Stream records
	// whose age exceeds this limit are subject to removal (trimming) from the stream.
	// You might receive a TrimmedDataAccessException if:
	//
	//    * You request a shard iterator with a sequence number older than the trim
	//    point (24 hours).
	//
	//    * You obtain a shard iterator, but before you use the iterator in a GetRecords
	//    request, a stream record in the shard exceeds the 24 hour period and is
	//    trimmed. This causes the iterator to access a record that no longer exists.
	ErrCodeTrimmedDataAccessException = "TrimmedDataAccessException"
)

var exceptionFromCode = map[string]func(protocol.ResponseMetadata) error{
	"ExpiredIteratorException":   newErrorExpiredIteratorException,
	"InternalServerError":        newErrorInternalServerError,
	"LimitExceededException":     newErrorLimitExceededException,
	"ResourceNotFoundException":  newErrorResourceNotFoundException,
	"TrimmedDataAccessException": newErrorTrimmedDataAccessException,
}