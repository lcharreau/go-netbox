// Code generated by go-swagger; DO NOT EDIT.

package dcim

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

// NewDcimConsolePortsReadParams creates a new DcimConsolePortsReadParams object
// with the default values initialized.
func NewDcimConsolePortsReadParams() *DcimConsolePortsReadParams {
	var ()
	return &DcimConsolePortsReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsolePortsReadParamsWithTimeout creates a new DcimConsolePortsReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsolePortsReadParamsWithTimeout(timeout time.Duration) *DcimConsolePortsReadParams {
	var ()
	return &DcimConsolePortsReadParams{

		timeout: timeout,
	}
}

// NewDcimConsolePortsReadParamsWithContext creates a new DcimConsolePortsReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsolePortsReadParamsWithContext(ctx context.Context) *DcimConsolePortsReadParams {
	var ()
	return &DcimConsolePortsReadParams{

		Context: ctx,
	}
}

// NewDcimConsolePortsReadParamsWithHTTPClient creates a new DcimConsolePortsReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsolePortsReadParamsWithHTTPClient(client *http.Client) *DcimConsolePortsReadParams {
	var ()
	return &DcimConsolePortsReadParams{
		HTTPClient: client,
	}
}

/*DcimConsolePortsReadParams contains all the parameters to send to the API endpoint
for the dcim console ports read operation typically these are written to a http.Request
*/
type DcimConsolePortsReadParams struct {

	/*ID
	  A unique integer value identifying this console port.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithTimeout(timeout time.Duration) *DcimConsolePortsReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithContext(ctx context.Context) *DcimConsolePortsReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithHTTPClient(client *http.Client) *DcimConsolePortsReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim console ports read params
func (o *DcimConsolePortsReadParams) WithID(id int64) *DcimConsolePortsReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console ports read params
func (o *DcimConsolePortsReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsolePortsReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
