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

// DcimRegionsUpdateReader is a Reader for the DcimRegionsUpdate structure.
type DcimRegionsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRegionsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRegionsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRegionsUpdateOK creates a DcimRegionsUpdateOK with default headers values
func NewDcimRegionsUpdateOK() *DcimRegionsUpdateOK {
	return &DcimRegionsUpdateOK{}
}

/*DcimRegionsUpdateOK handles this case with default header values.

DcimRegionsUpdateOK dcim regions update o k
*/
type DcimRegionsUpdateOK struct {
	Payload *models.WritableRegion
}

func (o *DcimRegionsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/regions/{id}/][%d] dcimRegionsUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRegionsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableRegion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
