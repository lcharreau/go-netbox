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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// NewExtrasTopologyMapsCreateParams creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized.
func NewExtrasTopologyMapsCreateParams() *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewExtrasTopologyMapsCreateParamsWithTimeout creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewExtrasTopologyMapsCreateParamsWithTimeout(timeout time.Duration) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		timeout: timeout,
	}
}

// NewExtrasTopologyMapsCreateParamsWithContext creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewExtrasTopologyMapsCreateParamsWithContext(ctx context.Context) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{

		Context: ctx,
	}
}

// NewExtrasTopologyMapsCreateParamsWithHTTPClient creates a new ExtrasTopologyMapsCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewExtrasTopologyMapsCreateParamsWithHTTPClient(client *http.Client) *ExtrasTopologyMapsCreateParams {
	var ()
	return &ExtrasTopologyMapsCreateParams{
		HTTPClient: client,
	}
}

/*ExtrasTopologyMapsCreateParams contains all the parameters to send to the API endpoint
for the extras topology maps create operation typically these are written to a http.Request
*/
type ExtrasTopologyMapsCreateParams struct {

	/*Data*/
	Data *models.WritableTopologyMap

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithTimeout(timeout time.Duration) *ExtrasTopologyMapsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithContext(ctx context.Context) *ExtrasTopologyMapsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithHTTPClient(client *http.Client) *ExtrasTopologyMapsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) WithData(data *models.WritableTopologyMap) *ExtrasTopologyMapsCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the extras topology maps create params
func (o *ExtrasTopologyMapsCreateParams) SetData(data *models.WritableTopologyMap) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *ExtrasTopologyMapsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
