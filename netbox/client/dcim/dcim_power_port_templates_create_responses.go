// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// DcimPowerPortTemplatesCreateReader is a Reader for the DcimPowerPortTemplatesCreate structure.
type DcimPowerPortTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPowerPortTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimPowerPortTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPowerPortTemplatesCreateCreated creates a DcimPowerPortTemplatesCreateCreated with default headers values
func NewDcimPowerPortTemplatesCreateCreated() *DcimPowerPortTemplatesCreateCreated {
	return &DcimPowerPortTemplatesCreateCreated{}
}

/*DcimPowerPortTemplatesCreateCreated handles this case with default header values.

DcimPowerPortTemplatesCreateCreated dcim power port templates create created
*/
type DcimPowerPortTemplatesCreateCreated struct {
	Payload *models.WritablePowerPortTemplate
}

func (o *DcimPowerPortTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/power-port-templates/][%d] dcimPowerPortTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPowerPortTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritablePowerPortTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
