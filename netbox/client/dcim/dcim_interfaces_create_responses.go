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

// DcimInterfacesCreateReader is a Reader for the DcimInterfacesCreate structure.
type DcimInterfacesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInterfacesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimInterfacesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimInterfacesCreateCreated creates a DcimInterfacesCreateCreated with default headers values
func NewDcimInterfacesCreateCreated() *DcimInterfacesCreateCreated {
	return &DcimInterfacesCreateCreated{}
}

/*DcimInterfacesCreateCreated handles this case with default header values.

DcimInterfacesCreateCreated dcim interfaces create created
*/
type DcimInterfacesCreateCreated struct {
	Payload *models.WritableInterface
}

func (o *DcimInterfacesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/interfaces/][%d] dcimInterfacesCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimInterfacesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableInterface)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
