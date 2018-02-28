// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// ExtrasExportTemplatesUpdateReader is a Reader for the ExtrasExportTemplatesUpdate structure.
type ExtrasExportTemplatesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasExportTemplatesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesUpdateOK creates a ExtrasExportTemplatesUpdateOK with default headers values
func NewExtrasExportTemplatesUpdateOK() *ExtrasExportTemplatesUpdateOK {
	return &ExtrasExportTemplatesUpdateOK{}
}

/*ExtrasExportTemplatesUpdateOK handles this case with default header values.

ExtrasExportTemplatesUpdateOK extras export templates update o k
*/
type ExtrasExportTemplatesUpdateOK struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /extras/export-templates/{id}/][%d] extrasExportTemplatesUpdateOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
