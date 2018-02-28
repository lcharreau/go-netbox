// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// IPAMChoicesReadReader is a Reader for the IPAMChoicesRead structure.
type IPAMChoicesReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMChoicesReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMChoicesReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMChoicesReadOK creates a IPAMChoicesReadOK with default headers values
func NewIPAMChoicesReadOK() *IPAMChoicesReadOK {
	return &IPAMChoicesReadOK{}
}

/*IPAMChoicesReadOK handles this case with default header values.

IPAMChoicesReadOK ipam choices read o k
*/
type IPAMChoicesReadOK struct {
}

func (o *IPAMChoicesReadOK) Error() string {
	return fmt.Sprintf("[GET /ipam/_choices/{id}/][%d] ipamChoicesReadOK ", 200)
}

func (o *IPAMChoicesReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
