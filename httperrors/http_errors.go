// Package errors - Describe custom errors structure with default methods

package httperrors

import (
	"fmt"
	"net/http"
	"strings"
)

const (
	CONFLICT_STATUS_CODE              = http.StatusConflict
	FORBIDDEN_STATUS_CODE             = http.StatusForbidden
	UNPROCESSABLE_ENTITY_STATUS_CODE  = http.StatusUnprocessableEntity
	INTERNAL_SERVER_ERROR_STATUS_CODE = http.StatusInternalServerError
	BAD_REQUEST_STATUS_CODE           = http.StatusBadRequest
	UNAUTHORIZED_STATUS_CODE          = http.StatusUnauthorized
	NOT_FOUND_STATUS_CODE             = http.StatusNotFound
)

// Logger - Interface that describe a Logger Service
type Logger interface {
	Error(args ...interface{})
}

// IHTTPError - Interface that describe HTTP error behavior
type IHTTPError interface {
	StatusCode() int
	Error() string
}

// GenericError - Describe a GenericError provided by user
type GenericError struct {
	HTTPStatusCode int
	Message        string
}

// ConflictError - Describe a ConflictError HTTP error
type ConflictError struct {
	HTTPStatusCode int
	Message        string
}

// ForbiddenError - Describe a Forbidden HTTP error
type ForbiddenError struct {
	HTTPStatusCode int
	Message        string
}

// UnprocessableEntityError - Describe an UnprocessableEntity HTTP error
type UnprocessableEntityError struct {
	HTTPStatusCode int
	Message        string
}

// InternalServerError - Describe an InternalServer HTTP error
type InternalServerError struct {
	HTTPStatusCode int
	Message        string
}

// BadRequestError - Describe a BadRequest HTTP error
type BadRequestError struct {
	HTTPStatusCode int
	Message        string
}

// UnauthorizedError - Describe a BadRequest HTTP error
type UnauthorizedError struct {
	HTTPStatusCode int
	Message        string
}

// NotFoundError - Describe a NotFound HTTP error
type NotFoundError struct {
	HTTPStatusCode int
	Message        string
}

// BuildGenericError - Create pointer to new GenericError
func BuildGenericError(err error, httpStatusCode int) *GenericError {
	return &GenericError{
		HTTPStatusCode: httpStatusCode,
		Message:        err.Error(),
	}
}

// NewConflictError - Create pointer to new ConflictError
func NewConflictError(err error) *ConflictError {
	return &ConflictError{
		HTTPStatusCode: CONFLICT_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewForbiddenError - Create pointer to new ForbiddenError
func NewForbiddenError(err error) *ForbiddenError {
	return &ForbiddenError{
		HTTPStatusCode: FORBIDDEN_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewUnprocessableEntityError - Create pointer to new UnprocessableEntityError
func NewUnprocessableEntityError(err error) *UnprocessableEntityError {
	return &UnprocessableEntityError{
		HTTPStatusCode: UNPROCESSABLE_ENTITY_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewUnprocessableEntityError - Create pointer to new InternalServerError
func NewInternalServerError(err error) *InternalServerError {
	return &InternalServerError{
		HTTPStatusCode: INTERNAL_SERVER_ERROR_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewBadRequestError - Create pointer to new BadRequestError
func NewBadRequestError(err error) *BadRequestError {
	return &BadRequestError{
		HTTPStatusCode: BAD_REQUEST_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewUnauthorizedError - Create pointer to new UnauthorizedError
func NewUnauthorizedError(err error) *UnauthorizedError {
	return &UnauthorizedError{
		HTTPStatusCode: UNAUTHORIZED_STATUS_CODE,
		Message:        err.Error(),
	}
}

// NewNotFoundError - Create pointer to new NotFoundError
func NewNotFoundError(err error) *NotFoundError {
	return &NotFoundError{
		HTTPStatusCode: NOT_FOUND_STATUS_CODE,
		Message:        err.Error(),
	}
}

func (ge *GenericError) StatusCode() int {
	return ge.HTTPStatusCode
}

func (ge *GenericError) Error() string {
	return ge.Message
}

func (ce *ConflictError) StatusCode() int {
	return ce.HTTPStatusCode
}

func (ce *ConflictError) Error() string {
	return ce.Message
}

func (fe *ForbiddenError) StatusCode() int {
	return fe.HTTPStatusCode
}

func (fe *ForbiddenError) Error() string {
	return fe.Message
}

func (uee *UnprocessableEntityError) StatusCode() int {
	return uee.HTTPStatusCode
}

func (uee *UnprocessableEntityError) Error() string {
	return uee.Message
}

func (iee *InternalServerError) StatusCode() int {
	return iee.HTTPStatusCode
}

func (iee *InternalServerError) Error() string {
	return iee.Message
}

func (bre *BadRequestError) StatusCode() int {
	return bre.HTTPStatusCode
}

func (bre *BadRequestError) Error() string {
	return bre.Message
}

func (nfe *NotFoundError) StatusCode() int {
	return nfe.HTTPStatusCode
}

func (nfe *NotFoundError) Error() string {
	return nfe.Message
}

func (une *UnauthorizedError) StatusCode() int {
	return une.HTTPStatusCode
}

func (une *UnauthorizedError) Error() string {
	return une.Message
}

// LogHttpError - Log HTTP Errors into console
func LogHttpError(httpError IHTTPError, logger Logger) {
	messages := []string{fmt.Sprintf("HTTP_STATUS_CODE:%d", httpError.StatusCode()), httpError.Error()}
	logger.Error(strings.Join(messages, " | "))
}
