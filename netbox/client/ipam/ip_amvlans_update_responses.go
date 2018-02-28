// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// IPAMVlansUpdateReader is a Reader for the IPAMVlansUpdate structure.
type IPAMVlansUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVlansUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMVlansUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVlansUpdateOK creates a IPAMVlansUpdateOK with default headers values
func NewIPAMVlansUpdateOK() *IPAMVlansUpdateOK {
	return &IPAMVlansUpdateOK{}
}

/*IPAMVlansUpdateOK handles this case with default header values.

IPAMVlansUpdateOK ipam vlans update o k
*/
type IPAMVlansUpdateOK struct {
	Payload *models.WritableVLAN
}

func (o *IPAMVlansUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vlans/{id}/][%d] ipamVlansUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMVlansUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableVLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
