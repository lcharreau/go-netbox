// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DcimConsolePortsDeleteReader is a Reader for the DcimConsolePortsDelete structure.
type DcimConsolePortsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsolePortsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDcimConsolePortsDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimConsolePortsDeleteNoContent creates a DcimConsolePortsDeleteNoContent with default headers values
func NewDcimConsolePortsDeleteNoContent() *DcimConsolePortsDeleteNoContent {
	return &DcimConsolePortsDeleteNoContent{}
}

/*DcimConsolePortsDeleteNoContent handles this case with default header values.

DcimConsolePortsDeleteNoContent dcim console ports delete no content
*/
type DcimConsolePortsDeleteNoContent struct {
}

func (o *DcimConsolePortsDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-ports/{id}/][%d] dcimConsolePortsDeleteNoContent ", 204)
}

func (o *DcimConsolePortsDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
