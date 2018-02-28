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

// DcimPlatformsCreateReader is a Reader for the DcimPlatformsCreate structure.
type DcimPlatformsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewDcimPlatformsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPlatformsCreateCreated creates a DcimPlatformsCreateCreated with default headers values
func NewDcimPlatformsCreateCreated() *DcimPlatformsCreateCreated {
	return &DcimPlatformsCreateCreated{}
}

/*DcimPlatformsCreateCreated handles this case with default header values.

DcimPlatformsCreateCreated dcim platforms create created
*/
type DcimPlatformsCreateCreated struct {
	Payload *models.Platform
}

func (o *DcimPlatformsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /dcim/platforms/][%d] dcimPlatformsCreateCreated  %+v", 201, o.Payload)
}

func (o *DcimPlatformsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Platform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
