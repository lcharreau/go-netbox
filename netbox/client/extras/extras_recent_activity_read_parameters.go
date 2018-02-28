// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewExtrasRecentActivityReadParams creates a new ExtrasRecentActivityReadParams object
// with the default values initialized.
func NewExtrasRecentActivityReadParams() *ExtrasRecentActivityReadParams {
	var ()
	return &ExtrasRecentActivityReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasRecentActivityReadParamsWithTimeout creates a new ExtrasRecentActivityReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasRecentActivityReadParamsWithTimeout(timeout time.Duration) *ExtrasRecentActivityReadParams {
	var ()
	return &ExtrasRecentActivityReadParams{

		timeout: timeout,
	}
}

// NewExtrasRecentActivityReadParamsWithContext creates a new ExtrasRecentActivityReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasRecentActivityReadParamsWithContext(ctx context.Context) *ExtrasRecentActivityReadParams {
	var ()
	return &ExtrasRecentActivityReadParams{

		Context: ctx,
	}
}

// NewExtrasRecentActivityReadParamsWithHTTPClient creates a new ExtrasRecentActivityReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasRecentActivityReadParamsWithHTTPClient(client *http.Client) *ExtrasRecentActivityReadParams {
	var ()
	return &ExtrasRecentActivityReadParams{
		HTTPClient: client,
	}
}

/*ExtrasRecentActivityReadParams contains all the parameters to send to the API endpoint
for the extras recent activity read operation typically these are written to a http.Request
*/
type ExtrasRecentActivityReadParams struct {

	/*ID
	  A unique integer value identifying this user action.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) WithTimeout(timeout time.Duration) *ExtrasRecentActivityReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) WithContext(ctx context.Context) *ExtrasRecentActivityReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) WithHTTPClient(client *http.Client) *ExtrasRecentActivityReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) WithID(id int64) *ExtrasRecentActivityReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the extras recent activity read params
func (o *ExtrasRecentActivityReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasRecentActivityReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
