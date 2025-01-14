// Code generated by go-swagger; DO NOT EDIT.

package registry

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewListRegistryProviderTypesParams creates a new ListRegistryProviderTypesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListRegistryProviderTypesParams() *ListRegistryProviderTypesParams {
	return &ListRegistryProviderTypesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListRegistryProviderTypesParamsWithTimeout creates a new ListRegistryProviderTypesParams object
// with the ability to set a timeout on a request.
func NewListRegistryProviderTypesParamsWithTimeout(timeout time.Duration) *ListRegistryProviderTypesParams {
	return &ListRegistryProviderTypesParams{
		timeout: timeout,
	}
}

// NewListRegistryProviderTypesParamsWithContext creates a new ListRegistryProviderTypesParams object
// with the ability to set a context for a request.
func NewListRegistryProviderTypesParamsWithContext(ctx context.Context) *ListRegistryProviderTypesParams {
	return &ListRegistryProviderTypesParams{
		Context: ctx,
	}
}

// NewListRegistryProviderTypesParamsWithHTTPClient creates a new ListRegistryProviderTypesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListRegistryProviderTypesParamsWithHTTPClient(client *http.Client) *ListRegistryProviderTypesParams {
	return &ListRegistryProviderTypesParams{
		HTTPClient: client,
	}
}

/* ListRegistryProviderTypesParams contains all the parameters to send to the API endpoint
   for the list registry provider types operation.

   Typically these are written to a http.Request.
*/
type ListRegistryProviderTypesParams struct {

	/* XRequestID.

	   An unique ID for the request
	*/
	XRequestID *string `js:"xRequestID"`

	timeout    time.Duration
	Context    context.Context `js:"context"`
	HTTPClient *http.Client    `js:"httpClient"`
}

// WithDefaults hydrates default values in the list registry provider types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRegistryProviderTypesParams) WithDefaults() *ListRegistryProviderTypesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list registry provider types params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListRegistryProviderTypesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list registry provider types params
func (o *ListRegistryProviderTypesParams) WithTimeout(timeout time.Duration) *ListRegistryProviderTypesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list registry provider types params
func (o *ListRegistryProviderTypesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list registry provider types params
func (o *ListRegistryProviderTypesParams) WithContext(ctx context.Context) *ListRegistryProviderTypesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list registry provider types params
func (o *ListRegistryProviderTypesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list registry provider types params
func (o *ListRegistryProviderTypesParams) WithHTTPClient(client *http.Client) *ListRegistryProviderTypesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list registry provider types params
func (o *ListRegistryProviderTypesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the list registry provider types params
func (o *ListRegistryProviderTypesParams) WithXRequestID(xRequestID *string) *ListRegistryProviderTypesParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the list registry provider types params
func (o *ListRegistryProviderTypesParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WriteToRequest writes these params to a swagger request
func (o *ListRegistryProviderTypesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
