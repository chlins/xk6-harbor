// Code generated by go-swagger; DO NOT EDIT.

package gc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/goharbor/xk6-harbor/pkg/harbor/models"
)

// GetGCHistoryReader is a Reader for the GetGCHistory structure.
type GetGCHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGCHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGCHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetGCHistoryUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetGCHistoryForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetGCHistoryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /system/gc] getGCHistory", response, response.Code())
	}
}

// NewGetGCHistoryOK creates a GetGCHistoryOK with default headers values
func NewGetGCHistoryOK() *GetGCHistoryOK {
	return &GetGCHistoryOK{}
}

/*
GetGCHistoryOK describes a response with status code 200, with default header values.

Get gc results successfully.
*/
type GetGCHistoryOK struct {

	/* Link refers to the previous page and next page
	 */
	Link string

	/* The total count of history
	 */
	XTotalCount int64

	Payload []*models.GCHistory
}

// IsSuccess returns true when this get Gc history o k response has a 2xx status code
func (o *GetGCHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get Gc history o k response has a 3xx status code
func (o *GetGCHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc history o k response has a 4xx status code
func (o *GetGCHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc history o k response has a 5xx status code
func (o *GetGCHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc history o k response a status code equal to that given
func (o *GetGCHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get Gc history o k response
func (o *GetGCHistoryOK) Code() int {
	return 200
}

func (o *GetGCHistoryOK) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryOK  %+v", 200, o.Payload)
}

func (o *GetGCHistoryOK) String() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryOK  %+v", 200, o.Payload)
}

func (o *GetGCHistoryOK) GetPayload() []*models.GCHistory {
	return o.Payload
}

func (o *GetGCHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// hydrates response header X-Total-Count
	hdrXTotalCount := response.GetHeader("X-Total-Count")

	if hdrXTotalCount != "" {
		valxTotalCount, err := swag.ConvertInt64(hdrXTotalCount)
		if err != nil {
			return errors.InvalidType("X-Total-Count", "header", "int64", hdrXTotalCount)
		}
		o.XTotalCount = valxTotalCount
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetGCHistoryUnauthorized creates a GetGCHistoryUnauthorized with default headers values
func NewGetGCHistoryUnauthorized() *GetGCHistoryUnauthorized {
	return &GetGCHistoryUnauthorized{}
}

/*
GetGCHistoryUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetGCHistoryUnauthorized struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc history unauthorized response has a 2xx status code
func (o *GetGCHistoryUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc history unauthorized response has a 3xx status code
func (o *GetGCHistoryUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc history unauthorized response has a 4xx status code
func (o *GetGCHistoryUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc history unauthorized response has a 5xx status code
func (o *GetGCHistoryUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc history unauthorized response a status code equal to that given
func (o *GetGCHistoryUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get Gc history unauthorized response
func (o *GetGCHistoryUnauthorized) Code() int {
	return 401
}

func (o *GetGCHistoryUnauthorized) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCHistoryUnauthorized) String() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryUnauthorized  %+v", 401, o.Payload)
}

func (o *GetGCHistoryUnauthorized) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCHistoryUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetGCHistoryForbidden creates a GetGCHistoryForbidden with default headers values
func NewGetGCHistoryForbidden() *GetGCHistoryForbidden {
	return &GetGCHistoryForbidden{}
}

/*
GetGCHistoryForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetGCHistoryForbidden struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc history forbidden response has a 2xx status code
func (o *GetGCHistoryForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc history forbidden response has a 3xx status code
func (o *GetGCHistoryForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc history forbidden response has a 4xx status code
func (o *GetGCHistoryForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get Gc history forbidden response has a 5xx status code
func (o *GetGCHistoryForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get Gc history forbidden response a status code equal to that given
func (o *GetGCHistoryForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get Gc history forbidden response
func (o *GetGCHistoryForbidden) Code() int {
	return 403
}

func (o *GetGCHistoryForbidden) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetGCHistoryForbidden) String() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryForbidden  %+v", 403, o.Payload)
}

func (o *GetGCHistoryForbidden) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCHistoryForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

// NewGetGCHistoryInternalServerError creates a GetGCHistoryInternalServerError with default headers values
func NewGetGCHistoryInternalServerError() *GetGCHistoryInternalServerError {
	return &GetGCHistoryInternalServerError{}
}

/*
GetGCHistoryInternalServerError describes a response with status code 500, with default header values.

Internal server error
*/
type GetGCHistoryInternalServerError struct {

	/* The ID of the corresponding request for the response
	 */
	XRequestID string

	Payload *models.Errors
}

// IsSuccess returns true when this get Gc history internal server error response has a 2xx status code
func (o *GetGCHistoryInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get Gc history internal server error response has a 3xx status code
func (o *GetGCHistoryInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get Gc history internal server error response has a 4xx status code
func (o *GetGCHistoryInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get Gc history internal server error response has a 5xx status code
func (o *GetGCHistoryInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get Gc history internal server error response a status code equal to that given
func (o *GetGCHistoryInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get Gc history internal server error response
func (o *GetGCHistoryInternalServerError) Code() int {
	return 500
}

func (o *GetGCHistoryInternalServerError) Error() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCHistoryInternalServerError) String() string {
	return fmt.Sprintf("[GET /system/gc][%d] getGcHistoryInternalServerError  %+v", 500, o.Payload)
}

func (o *GetGCHistoryInternalServerError) GetPayload() *models.Errors {
	return o.Payload
}

func (o *GetGCHistoryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
