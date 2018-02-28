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

	"github.com/davcamer/go-netbox/netbox/models"
)

// NewDcimConsoleServerPortTemplatesPartialUpdateParams creates a new DcimConsoleServerPortTemplatesPartialUpdateParams object
// with the default values initialized.
func NewDcimConsoleServerPortTemplatesPartialUpdateParams() *DcimConsoleServerPortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortTemplatesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithTimeout creates a new DcimConsoleServerPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortTemplatesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithContext creates a new DcimConsoleServerPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithContext(ctx context.Context) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortTemplatesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithHTTPClient creates a new DcimConsoleServerPortTemplatesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimConsoleServerPortTemplatesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	var ()
	return &DcimConsoleServerPortTemplatesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimConsoleServerPortTemplatesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim console server port templates partial update operation typically these are written to a http.Request
*/
type DcimConsoleServerPortTemplatesPartialUpdateParams struct {

	/*Data*/
	Data *models.WritableConsoleServerPortTemplate
	/*ID
	  A unique integer value identifying this console server port template.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WithContext(ctx context.Context) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WithData(data *models.WritableConsoleServerPortTemplate) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) SetData(data *models.WritableConsoleServerPortTemplate) {
	o.Data = data
}

// WithID adds the id to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WithID(id int64) *DcimConsoleServerPortTemplatesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim console server port templates partial update params
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimConsoleServerPortTemplatesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
