// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimRacksDeleteReader is a Reader for the DcimRacksDelete structure.
type DcimRacksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimRacksDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksDeleteNoContent creates a DcimRacksDeleteNoContent with default headers values
func NewDcimRacksDeleteNoContent() *DcimRacksDeleteNoContent {
	return &DcimRacksDeleteNoContent{}
}

/*DcimRacksDeleteNoContent handles this case with default header values.

DcimRacksDeleteNoContent dcim racks delete no content
*/
type DcimRacksDeleteNoContent struct {
}

func (o *DcimRacksDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/racks/{id}/][%d] dcimRacksDeleteNoContent ", 204)
}

func (o *DcimRacksDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
