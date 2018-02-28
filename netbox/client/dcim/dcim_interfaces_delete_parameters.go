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

// NewDcimInterfacesDeleteParams creates a new DcimInterfacesDeleteParams object
// with the default values initialized.
func NewDcimInterfacesDeleteParams() *DcimInterfacesDeleteParams {
	var ()
	return &DcimInterfacesDeleteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimInterfacesDeleteParamsWithTimeout creates a new DcimInterfacesDeleteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimInterfacesDeleteParamsWithTimeout(timeout time.Duration) *DcimInterfacesDeleteParams {
	var ()
	return &DcimInterfacesDeleteParams{

		timeout: timeout,
	}
}

// NewDcimInterfacesDeleteParamsWithContext creates a new DcimInterfacesDeleteParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimInterfacesDeleteParamsWithContext(ctx context.Context) *DcimInterfacesDeleteParams {
	var ()
	return &DcimInterfacesDeleteParams{

		Context: ctx,
	}
}

// NewDcimInterfacesDeleteParamsWithHTTPClient creates a new DcimInterfacesDeleteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimInterfacesDeleteParamsWithHTTPClient(client *http.Client) *DcimInterfacesDeleteParams {
	var ()
	return &DcimInterfacesDeleteParams{
		HTTPClient: client,
	}
}

/*DcimInterfacesDeleteParams contains all the parameters to send to the API endpoint
for the dcim interfaces delete operation typically these are written to a http.Request
*/
type DcimInterfacesDeleteParams struct {

	/*ID
	  A unique integer value identifying this interface.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) WithTimeout(timeout time.Duration) *DcimInterfacesDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) WithContext(ctx context.Context) *DcimInterfacesDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) WithHTTPClient(client *http.Client) *DcimInterfacesDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) WithID(id int64) *DcimInterfacesDeleteParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim interfaces delete params
func (o *DcimInterfacesDeleteParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimInterfacesDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
