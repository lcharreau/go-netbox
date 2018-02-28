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

// DcimRackReservationsReadReader is a Reader for the DcimRackReservationsRead structure.
type DcimRackReservationsReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackReservationsReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackReservationsReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackReservationsReadOK creates a DcimRackReservationsReadOK with default headers values
func NewDcimRackReservationsReadOK() *DcimRackReservationsReadOK {
	return &DcimRackReservationsReadOK{}
}

/*DcimRackReservationsReadOK handles this case with default header values.

DcimRackReservationsReadOK dcim rack reservations read o k
*/
type DcimRackReservationsReadOK struct {
	Payload *models.RackReservation
}

func (o *DcimRackReservationsReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rack-reservations/{id}/][%d] dcimRackReservationsReadOK  %+v", 200, o.Payload)
}

func (o *DcimRackReservationsReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackReservation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
