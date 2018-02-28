// Code generated by go-swagger; DO NOT EDIT.

package virtualization

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

// NewVirtualizationClusterTypesReadParams creates a new VirtualizationClusterTypesReadParams object
// with the default values initialized.
func NewVirtualizationClusterTypesReadParams() *VirtualizationClusterTypesReadParams {
	var ()
	return &VirtualizationClusterTypesReadParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewVirtualizationClusterTypesReadParamsWithTimeout creates a new VirtualizationClusterTypesReadParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewVirtualizationClusterTypesReadParamsWithTimeout(timeout time.Duration) *VirtualizationClusterTypesReadParams {
	var ()
	return &VirtualizationClusterTypesReadParams{

		timeout: timeout,
	}
}

// NewVirtualizationClusterTypesReadParamsWithContext creates a new VirtualizationClusterTypesReadParams object
// with the default values initialized, and the ability to set a context for a request
func NewVirtualizationClusterTypesReadParamsWithContext(ctx context.Context) *VirtualizationClusterTypesReadParams {
	var ()
	return &VirtualizationClusterTypesReadParams{

		Context: ctx,
	}
}

// NewVirtualizationClusterTypesReadParamsWithHTTPClient creates a new VirtualizationClusterTypesReadParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewVirtualizationClusterTypesReadParamsWithHTTPClient(client *http.Client) *VirtualizationClusterTypesReadParams {
	var ()
	return &VirtualizationClusterTypesReadParams{
		HTTPClient: client,
	}
}

/*VirtualizationClusterTypesReadParams contains all the parameters to send to the API endpoint
for the virtualization cluster types read operation typically these are written to a http.Request
*/
type VirtualizationClusterTypesReadParams struct {

	/*ID
	  A unique integer value identifying this cluster type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithTimeout(timeout time.Duration) *VirtualizationClusterTypesReadParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithContext(ctx context.Context) *VirtualizationClusterTypesReadParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithHTTPClient(client *http.Client) *VirtualizationClusterTypesReadParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) WithID(id int64) *VirtualizationClusterTypesReadParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the virtualization cluster types read params
func (o *VirtualizationClusterTypesReadParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *VirtualizationClusterTypesReadParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
