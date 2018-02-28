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

// IPAMPrefixesPartialUpdateReader is a Reader for the IPAMPrefixesPartialUpdate structure.
type IPAMPrefixesPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMPrefixesPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMPrefixesPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMPrefixesPartialUpdateOK creates a IPAMPrefixesPartialUpdateOK with default headers values
func NewIPAMPrefixesPartialUpdateOK() *IPAMPrefixesPartialUpdateOK {
	return &IPAMPrefixesPartialUpdateOK{}
}

/*IPAMPrefixesPartialUpdateOK handles this case with default header values.

IPAMPrefixesPartialUpdateOK ipam prefixes partial update o k
*/
type IPAMPrefixesPartialUpdateOK struct {
	Payload *models.WritablePrefix
}

func (o *IPAMPrefixesPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /ipam/prefixes/{id}/][%d] ipamPrefixesPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMPrefixesPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritablePrefix)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
