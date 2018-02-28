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

// DcimRacksUnitsReader is a Reader for the DcimRacksUnits structure.
type DcimRacksUnitsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksUnitsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRacksUnitsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksUnitsOK creates a DcimRacksUnitsOK with default headers values
func NewDcimRacksUnitsOK() *DcimRacksUnitsOK {
	return &DcimRacksUnitsOK{}
}

/*DcimRacksUnitsOK handles this case with default header values.

DcimRacksUnitsOK dcim racks units o k
*/
type DcimRacksUnitsOK struct {
	Payload *models.Rack
}

func (o *DcimRacksUnitsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/racks/{id}/units/][%d] dcimRacksUnitsOK  %+v", 200, o.Payload)
}

func (o *DcimRacksUnitsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Rack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
