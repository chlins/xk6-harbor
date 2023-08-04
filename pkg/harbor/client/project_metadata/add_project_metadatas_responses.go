// Code generated by go-swagger; DO NOT EDIT.

package project_metadata

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// AddProjectMetadatasReader is a Reader for the AddProjectMetadatas structure.
type AddProjectMetadatasReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddProjectMetadatasReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddProjectMetadatasOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewAddProjectMetadatasBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewAddProjectMetadatasUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewAddProjectMetadatasForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewAddProjectMetadatasNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewAddProjectMetadatasConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewAddProjectMetadatasInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /projects/{project_name_or_id}/metadatas/] addProjectMetadatas", response, response.Code())
	}
}

// NewAddProjectMetadatasOK creates a AddProjectMetadatasOK with default headers values
func NewAddProjectMetadatasOK() *AddProjectMetadatasOK {
	return &AddProjectMetadatasOK{}
}

/*
AddProjectMetadatasOK describes a response with status code 200, with default header values.

Success
*/
type AddProjectMetadatasOK struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string
}

// IsSuccess returns true when this add project metadatas o k response has a 2xx status code
func (o *AddProjectMetadatasOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this add project metadatas o k response has a 3xx status code
func (o *AddProjectMetadatasOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas o k response has a 4xx status code
func (o *AddProjectMetadatasOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this add project metadatas o k response has a 5xx status code
func (o *AddProjectMetadatasOK) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas o k response a status code equal to that given
func (o *AddProjectMetadatasOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the add project metadatas o k response
func (o *AddProjectMetadatasOK) Code() int {
	return 200
}

func (o *AddProjectMetadatasOK) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasOK ", 200)
}

func (o *AddProjectMetadatasOK) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasOK ", 200)
}

func (o *AddProjectMetadatasOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header X-Request-Id
	hdrXRequestID := response.GetHeader("X-Request-Id")

	if hdrXRequestID != "" {
		o.XRequestID = hdrXRequestID
	}

	return nil
}

// NewAddProjectMetadatasBadRequest creates a AddProjectMetadatasBadRequest with default headers values
func NewAddProjectMetadatasBadRequest() *AddProjectMetadatasBadRequest {
	return &AddProjectMetadatasBadRequest{}
}

/*
AddProjectMetadatasBadRequest describes a response with status code 400, with default header values.

Bad request
*/
type AddProjectMetadatasBadRequest struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas bad request response has a 2xx status code
func (o *AddProjectMetadatasBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas bad request response has a 3xx status code
func (o *AddProjectMetadatasBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas bad request response has a 4xx status code
func (o *AddProjectMetadatasBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project metadatas bad request response has a 5xx status code
func (o *AddProjectMetadatasBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas bad request response a status code equal to that given
func (o *AddProjectMetadatasBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the add project metadatas bad request response
func (o *AddProjectMetadatasBadRequest) Code() int {
	return 400
}

func (o *AddProjectMetadatasBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectMetadatasBadRequest) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasBadRequest  %+v", 400, o.Payload)
}

func (o *AddProjectMetadatasBadRequest) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddProjectMetadatasUnauthorized creates a AddProjectMetadatasUnauthorized with default headers values
func NewAddProjectMetadatasUnauthorized() *AddProjectMetadatasUnauthorized {
	return &AddProjectMetadatasUnauthorized{}
}

/*
AddProjectMetadatasUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type AddProjectMetadatasUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas unauthorized response has a 2xx status code
func (o *AddProjectMetadatasUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas unauthorized response has a 3xx status code
func (o *AddProjectMetadatasUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas unauthorized response has a 4xx status code
func (o *AddProjectMetadatasUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project metadatas unauthorized response has a 5xx status code
func (o *AddProjectMetadatasUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas unauthorized response a status code equal to that given
func (o *AddProjectMetadatasUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the add project metadatas unauthorized response
func (o *AddProjectMetadatasUnauthorized) Code() int {
	return 401
}

func (o *AddProjectMetadatasUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}

func (o *AddProjectMetadatasUnauthorized) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasUnauthorized  %+v", 401, o.Payload)
}

func (o *AddProjectMetadatasUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddProjectMetadatasForbidden creates a AddProjectMetadatasForbidden with default headers values
func NewAddProjectMetadatasForbidden() *AddProjectMetadatasForbidden {
	return &AddProjectMetadatasForbidden{}
}

/*
AddProjectMetadatasForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type AddProjectMetadatasForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas forbidden response has a 2xx status code
func (o *AddProjectMetadatasForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas forbidden response has a 3xx status code
func (o *AddProjectMetadatasForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas forbidden response has a 4xx status code
func (o *AddProjectMetadatasForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project metadatas forbidden response has a 5xx status code
func (o *AddProjectMetadatasForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas forbidden response a status code equal to that given
func (o *AddProjectMetadatasForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the add project metadatas forbidden response
func (o *AddProjectMetadatasForbidden) Code() int {
	return 403
}

func (o *AddProjectMetadatasForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasForbidden  %+v", 403, o.Payload)
}

func (o *AddProjectMetadatasForbidden) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasForbidden  %+v", 403, o.Payload)
}

func (o *AddProjectMetadatasForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddProjectMetadatasNotFound creates a AddProjectMetadatasNotFound with default headers values
func NewAddProjectMetadatasNotFound() *AddProjectMetadatasNotFound {
	return &AddProjectMetadatasNotFound{}
}

/*
AddProjectMetadatasNotFound describes a response with status code 404, with default header values.

Not found
*/
type AddProjectMetadatasNotFound struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas not found response has a 2xx status code
func (o *AddProjectMetadatasNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas not found response has a 3xx status code
func (o *AddProjectMetadatasNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas not found response has a 4xx status code
func (o *AddProjectMetadatasNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project metadatas not found response has a 5xx status code
func (o *AddProjectMetadatasNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas not found response a status code equal to that given
func (o *AddProjectMetadatasNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the add project metadatas not found response
func (o *AddProjectMetadatasNotFound) Code() int {
	return 404
}

func (o *AddProjectMetadatasNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *AddProjectMetadatasNotFound) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasNotFound  %+v", 404, o.Payload)
}

func (o *AddProjectMetadatasNotFound) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddProjectMetadatasConflict creates a AddProjectMetadatasConflict with default headers values
func NewAddProjectMetadatasConflict() *AddProjectMetadatasConflict {
	return &AddProjectMetadatasConflict{}
}

/*
AddProjectMetadatasConflict describes a response with status code 409, with default header values.

Conflict
*/
type AddProjectMetadatasConflict struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas conflict response has a 2xx status code
func (o *AddProjectMetadatasConflict) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas conflict response has a 3xx status code
func (o *AddProjectMetadatasConflict) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas conflict response has a 4xx status code
func (o *AddProjectMetadatasConflict) IsClientError() bool {
	return true
}

// IsServerError returns true when this add project metadatas conflict response has a 5xx status code
func (o *AddProjectMetadatasConflict) IsServerError() bool {
	return false
}

// IsCode returns true when this add project metadatas conflict response a status code equal to that given
func (o *AddProjectMetadatasConflict) IsCode(code int) bool {
	return code == 409
}

// Code gets the status code for the add project metadatas conflict response
func (o *AddProjectMetadatasConflict) Code() int {
	return 409
}

func (o *AddProjectMetadatasConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasConflict  %+v", 409, o.Payload)
}

func (o *AddProjectMetadatasConflict) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasConflict  %+v", 409, o.Payload)
}

func (o *AddProjectMetadatasConflict) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewAddProjectMetadatasInternalServerError creates a AddProjectMetadatasInternalServerError with default headers values
func NewAddProjectMetadatasInternalServerError() *AddProjectMetadatasInternalServerError {
	return &AddProjectMetadatasInternalServerError{}
}

/*
AddProjectMetadatasInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type AddProjectMetadatasInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this add project metadatas internal server error response has a 2xx status code
func (o *AddProjectMetadatasInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this add project metadatas internal server error response has a 3xx status code
func (o *AddProjectMetadatasInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this add project metadatas internal server error response has a 4xx status code
func (o *AddProjectMetadatasInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this add project metadatas internal server error response has a 5xx status code
func (o *AddProjectMetadatasInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this add project metadatas internal server error response a status code equal to that given
func (o *AddProjectMetadatasInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the add project metadatas internal server error response
func (o *AddProjectMetadatasInternalServerError) Code() int {
	return 500
}

func (o *AddProjectMetadatasInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}

func (o *AddProjectMetadatasInternalServerError) String() string {
	return fmt.Sprintf("[POST /projects/{project_name_or_id}/metadatas/][%d] addProjectMetadatasInternalServerError  %+v", 500, o.Payload)
}

func (o *AddProjectMetadatasInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *AddProjectMetadatasInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
