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

// NewDcimPowerOutletTemplatesReadParams creates a new DcimPowerOutletTemplatesReadParams object
// with the default values initialized.
func NewDcimPowerOutletTemplatesReadParams() *DcimPowerOutletTemplatesReadParams {
	var ()
	return &DcimPowerOutletTemplatesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithTimeout creates a new DcimPowerOutletTemplatesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletTemplatesReadParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesReadParams {
	var ()
	return &DcimPowerOutletTemplatesReadParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithContext creates a new DcimPowerOutletTemplatesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletTemplatesReadParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesReadParams {
	var ()
	return &DcimPowerOutletTemplatesReadParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesReadParamsWithHTTPClient creates a new DcimPowerOutletTemplatesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletTemplatesReadParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesReadParams {
	var ()
	return &DcimPowerOutletTemplatesReadParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletTemplatesReadParams contains all the parameters to send to the API endpoint
for the dcim power outlet templates read operation typically these are written to a http.Request
*/
type DcimPowerOutletTemplatesReadParams struct {

	/*ID
	  A unique integer value identifying this power outlet template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) WithID(id int64) *DcimPowerOutletTemplatesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim power outlet templates read params
func (o *DcimPowerOutletTemplatesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
