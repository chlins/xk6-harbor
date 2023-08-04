// Code generated by go-swagger; DO NOT EDIT.

package webhook

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// GetWebhookPolicyOfProjectReader is a Reader for the GetWebhookPolicyOfProject structure.
type GetWebhookPolicyOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhookPolicyOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebhookPolicyOfProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebhookPolicyOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetWebhookPolicyOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetWebhookPolicyOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebhookPolicyOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetWebhookPolicyOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}] GetWebhookPolicyOfProject", response, response.Code())
	}
}

// NewGetWebhookPolicyOfProjectOK creates a GetWebhookPolicyOfProjectOK with default headers values
func NewGetWebhookPolicyOfProjectOK() *GetWebhookPolicyOfProjectOK {
	return &GetWebhookPolicyOfProjectOK{}
}

/*
GetWebhookPolicyOfProjectOK describes a response with status code 200, with default header values.

Get webhook policy successfully.
*/
type GetWebhookPolicyOfProjectOK struct {
	Payload *models.WebhookPolicy
}

// IsSuccess returns true when this get webhook policy of project o k response has a 2xx status code
func (o *GetWebhookPolicyOfProjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get webhook policy of project o k response has a 3xx status code
func (o *GetWebhookPolicyOfProjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project o k response has a 4xx status code
func (o *GetWebhookPolicyOfProjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webhook policy of project o k response has a 5xx status code
func (o *GetWebhookPolicyOfProjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook policy of project o k response a status code equal to that given
func (o *GetWebhookPolicyOfProjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get webhook policy of project o k response
func (o *GetWebhookPolicyOfProjectOK) Code() int {
	return 200
}

func (o *GetWebhookPolicyOfProjectOK) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetWebhookPolicyOfProjectOK) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectOK  %+v", 200, o.Payload)
}

func (o *GetWebhookPolicyOfProjectOK) GetPayload() *models.WebhookPolicy {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WebhookPolicy)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookPolicyOfProjectBadRequest creates a GetWebhookPolicyOfProjectBadRequest with default headers values
func NewGetWebhookPolicyOfProjectBadRequest() *GetWebhookPolicyOfProjectBadRequest {
	return &GetWebhookPolicyOfProjectBadRequest{}
}

/*
GetWebhookPolicyOfProjectBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type GetWebhookPolicyOfProjectBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get webhook policy of project bad request response has a 2xx status code
func (o *GetWebhookPolicyOfProjectBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhook policy of project bad request response has a 3xx status code
func (o *GetWebhookPolicyOfProjectBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project bad request response has a 4xx status code
func (o *GetWebhookPolicyOfProjectBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webhook policy of project bad request response has a 5xx status code
func (o *GetWebhookPolicyOfProjectBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook policy of project bad request response a status code equal to that given
func (o *GetWebhookPolicyOfProjectBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the get webhook policy of project bad request response
func (o *GetWebhookPolicyOfProjectBadRequest) Code() int {
	return 400
}

func (o *GetWebhookPolicyOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebhookPolicyOfProjectBadRequest) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebhookPolicyOfProjectBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebhookPolicyOfProjectUnauthorized creates a GetWebhookPolicyOfProjectUnauthorized with default headers values
func NewGetWebhookPolicyOfProjectUnauthorized() *GetWebhookPolicyOfProjectUnauthorized {
	return &GetWebhookPolicyOfProjectUnauthorized{}
}

/*
GetWebhookPolicyOfProjectUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetWebhookPolicyOfProjectUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get webhook policy of project unauthorized response has a 2xx status code
func (o *GetWebhookPolicyOfProjectUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhook policy of project unauthorized response has a 3xx status code
func (o *GetWebhookPolicyOfProjectUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project unauthorized response has a 4xx status code
func (o *GetWebhookPolicyOfProjectUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webhook policy of project unauthorized response has a 5xx status code
func (o *GetWebhookPolicyOfProjectUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook policy of project unauthorized response a status code equal to that given
func (o *GetWebhookPolicyOfProjectUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get webhook policy of project unauthorized response
func (o *GetWebhookPolicyOfProjectUnauthorized) Code() int {
	return 401
}

func (o *GetWebhookPolicyOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebhookPolicyOfProjectUnauthorized) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *GetWebhookPolicyOfProjectUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebhookPolicyOfProjectForbidden creates a GetWebhookPolicyOfProjectForbidden with default headers values
func NewGetWebhookPolicyOfProjectForbidden() *GetWebhookPolicyOfProjectForbidden {
	return &GetWebhookPolicyOfProjectForbidden{}
}

/*
GetWebhookPolicyOfProjectForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetWebhookPolicyOfProjectForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get webhook policy of project forbidden response has a 2xx status code
func (o *GetWebhookPolicyOfProjectForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhook policy of project forbidden response has a 3xx status code
func (o *GetWebhookPolicyOfProjectForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project forbidden response has a 4xx status code
func (o *GetWebhookPolicyOfProjectForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webhook policy of project forbidden response has a 5xx status code
func (o *GetWebhookPolicyOfProjectForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook policy of project forbidden response a status code equal to that given
func (o *GetWebhookPolicyOfProjectForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get webhook policy of project forbidden response
func (o *GetWebhookPolicyOfProjectForbidden) Code() int {
	return 403
}

func (o *GetWebhookPolicyOfProjectForbidden) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetWebhookPolicyOfProjectForbidden) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *GetWebhookPolicyOfProjectForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebhookPolicyOfProjectNotFound creates a GetWebhookPolicyOfProjectNotFound with default headers values
func NewGetWebhookPolicyOfProjectNotFound() *GetWebhookPolicyOfProjectNotFound {
	return &GetWebhookPolicyOfProjectNotFound{}
}

/*
GetWebhookPolicyOfProjectNotFound describes a response with status code 404, with default header values.

Not found
*/
type GetWebhookPolicyOfProjectNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get webhook policy of project not found response has a 2xx status code
func (o *GetWebhookPolicyOfProjectNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhook policy of project not found response has a 3xx status code
func (o *GetWebhookPolicyOfProjectNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project not found response has a 4xx status code
func (o *GetWebhookPolicyOfProjectNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get webhook policy of project not found response has a 5xx status code
func (o *GetWebhookPolicyOfProjectNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get webhook policy of project not found response a status code equal to that given
func (o *GetWebhookPolicyOfProjectNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get webhook policy of project not found response
func (o *GetWebhookPolicyOfProjectNotFound) Code() int {
	return 404
}

func (o *GetWebhookPolicyOfProjectNotFound) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetWebhookPolicyOfProjectNotFound) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *GetWebhookPolicyOfProjectNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetWebhookPolicyOfProjectInternalServerError creates a GetWebhookPolicyOfProjectInternalServerError with default headers values
func NewGetWebhookPolicyOfProjectInternalServerError() *GetWebhookPolicyOfProjectInternalServerError {
	return &GetWebhookPolicyOfProjectInternalServerError{}
}

/*
GetWebhookPolicyOfProjectInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetWebhookPolicyOfProjectInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get webhook policy of project internal server error response has a 2xx status code
func (o *GetWebhookPolicyOfProjectInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get webhook policy of project internal server error response has a 3xx status code
func (o *GetWebhookPolicyOfProjectInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get webhook policy of project internal server error response has a 4xx status code
func (o *GetWebhookPolicyOfProjectInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get webhook policy of project internal server error response has a 5xx status code
func (o *GetWebhookPolicyOfProjectInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get webhook policy of project internal server error response a status code equal to that given
func (o *GetWebhookPolicyOfProjectInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get webhook policy of project internal server error response
func (o *GetWebhookPolicyOfProjectInternalServerError) Code() int {
	return 500
}

func (o *GetWebhookPolicyOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebhookPolicyOfProjectInternalServerError) String() string {
	return fmt.Sprintf("[GET /projects/{project_name_or_id}/webhook/policies/{webhook_policy_id}][%d] getWebhookPolicyOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *GetWebhookPolicyOfProjectInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetWebhookPolicyOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
