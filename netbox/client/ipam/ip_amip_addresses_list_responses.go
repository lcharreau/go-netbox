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

// IPAMIPAddressesListReader is a Reader for the IPAMIPAddressesList structure.
type IPAMIPAddressesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMIPAddressesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMIPAddressesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMIPAddressesListOK creates a IPAMIPAddressesListOK with default headers values
func NewIPAMIPAddressesListOK() *IPAMIPAddressesListOK {
	return &IPAMIPAddressesListOK{}
}

/*IPAMIPAddressesListOK handles this case with default header values.

IPAMIPAddressesListOK ipam Ip addresses list o k
*/
type IPAMIPAddressesListOK struct {
	Payload *models.IPAMIPAddressesListOKBody
}

func (o *IPAMIPAddressesListOK) Error() string {
	return fmt.Sprintf("[GET /ipam/ip-addresses/][%d] ipamIpAddressesListOK  %+v", 200, o.Payload)
}

func (o *IPAMIPAddressesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.IPAMIPAddressesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
