// Code generated by go-swagger; DO NOT EDIT.

package preheat

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// StopExecutionReader is a Reader for the StopExecution structure.
type StopExecutionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *StopExecutionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewStopExecutionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewStopExecutionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewStopExecutionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewStopExecutionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewStopExecutionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewStopExecutionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}] StopExecution", response, response.Code())
	}
}

// NewStopExecutionOK creates a StopExecutionOK with default headers values
func NewStopExecutionOK() *StopExecutionOK {
	return &StopExecutionOK{}
}

/*
StopExecutionOK describes a response with status code 200, with default header values.

Success
*/
type StopExecutionOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this stop execution o k response has a 2xx status code
func (o *StopExecutionOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this stop execution o k response has a 3xx status code
func (o *StopExecutionOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution o k response has a 4xx status code
func (o *StopExecutionOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop execution o k response has a 5xx status code
func (o *StopExecutionOK) IsServerError() bool {
	return false
}

// IsCode returns true when this stop execution o k response a status code equal to that given
func (o *StopExecutionOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the stop execution o k response
func (o *StopExecutionOK) Code() int {
	return 200
}

func (o *StopExecutionOK) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionOK ", 200)
}

func (o *StopExecutionOK) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionOK ", 200)
}

func (o *StopExecutionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewStopExecutionBadRequest creates a StopExecutionBadRequest with default headers values
func NewStopExecutionBadRequest() *StopExecutionBadRequest {
	return &StopExecutionBadRequest{}
}

/*
StopExecutionBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type StopExecutionBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop execution bad request response has a 2xx status code
func (o *StopExecutionBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop execution bad request response has a 3xx status code
func (o *StopExecutionBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution bad request response has a 4xx status code
func (o *StopExecutionBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop execution bad request response has a 5xx status code
func (o *StopExecutionBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this stop execution bad request response a status code equal to that given
func (o *StopExecutionBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the stop execution bad request response
func (o *StopExecutionBadRequest) Code() int {
	return 400
}

func (o *StopExecutionBadRequest) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *StopExecutionBadRequest) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionBadRequest  %+v", 400, o.Payload)
}

func (o *StopExecutionBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionUnauthorized creates a StopExecutionUnauthorized with default headers values
func NewStopExecutionUnauthorized() *StopExecutionUnauthorized {
	return &StopExecutionUnauthorized{}
}

/*
StopExecutionUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type StopExecutionUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop execution unauthorized response has a 2xx status code
func (o *StopExecutionUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop execution unauthorized response has a 3xx status code
func (o *StopExecutionUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution unauthorized response has a 4xx status code
func (o *StopExecutionUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop execution unauthorized response has a 5xx status code
func (o *StopExecutionUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this stop execution unauthorized response a status code equal to that given
func (o *StopExecutionUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the stop execution unauthorized response
func (o *StopExecutionUnauthorized) Code() int {
	return 401
}

func (o *StopExecutionUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *StopExecutionUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionUnauthorized  %+v", 401, o.Payload)
}

func (o *StopExecutionUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionForbidden creates a StopExecutionForbidden with default headers values
func NewStopExecutionForbidden() *StopExecutionForbidden {
	return &StopExecutionForbidden{}
}

/*
StopExecutionForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type StopExecutionForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop execution forbidden response has a 2xx status code
func (o *StopExecutionForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop execution forbidden response has a 3xx status code
func (o *StopExecutionForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution forbidden response has a 4xx status code
func (o *StopExecutionForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop execution forbidden response has a 5xx status code
func (o *StopExecutionForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this stop execution forbidden response a status code equal to that given
func (o *StopExecutionForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the stop execution forbidden response
func (o *StopExecutionForbidden) Code() int {
	return 403
}

func (o *StopExecutionForbidden) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionForbidden  %+v", 403, o.Payload)
}

func (o *StopExecutionForbidden) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionForbidden  %+v", 403, o.Payload)
}

func (o *StopExecutionForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionNotFound creates a StopExecutionNotFound with default headers values
func NewStopExecutionNotFound() *StopExecutionNotFound {
	return &StopExecutionNotFound{}
}

/*
StopExecutionNotFound describes a response with status code 404, with default header values.

Not found
*/
type StopExecutionNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop execution not found response has a 2xx status code
func (o *StopExecutionNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop execution not found response has a 3xx status code
func (o *StopExecutionNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution not found response has a 4xx status code
func (o *StopExecutionNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this stop execution not found response has a 5xx status code
func (o *StopExecutionNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this stop execution not found response a status code equal to that given
func (o *StopExecutionNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the stop execution not found response
func (o *StopExecutionNotFound) Code() int {
	return 404
}

func (o *StopExecutionNotFound) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionNotFound  %+v", 404, o.Payload)
}

func (o *StopExecutionNotFound) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionNotFound  %+v", 404, o.Payload)
}

func (o *StopExecutionNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewStopExecutionInternalServerError creates a StopExecutionInternalServerError with default headers values
func NewStopExecutionInternalServerError() *StopExecutionInternalServerError {
	return &StopExecutionInternalServerError{}
}

/*
StopExecutionInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type StopExecutionInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this stop execution internal server error response has a 2xx status code
func (o *StopExecutionInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this stop execution internal server error response has a 3xx status code
func (o *StopExecutionInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this stop execution internal server error response has a 4xx status code
func (o *StopExecutionInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this stop execution internal server error response has a 5xx status code
func (o *StopExecutionInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this stop execution internal server error response a status code equal to that given
func (o *StopExecutionInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the stop execution internal server error response
func (o *StopExecutionInternalServerError) Code() int {
	return 500
}

func (o *StopExecutionInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *StopExecutionInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /projects/{project_name}/preheat/policies/{preheat_policy_name}/executions/{execution_id}][%d] stopExecutionInternalServerError  %+v", 500, o.Payload)
}

func (o *StopExecutionInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *StopExecutionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	o.Payload = new(models.Errors)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
