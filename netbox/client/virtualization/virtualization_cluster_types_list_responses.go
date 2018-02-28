// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/davcamer/go-netbox/netbox/models"
)

// VirtualizationClusterTypesListReader is a Reader for the VirtualizationClusterTypesList structure.
type VirtualizationClusterTypesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterTypesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterTypesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterTypesListOK creates a VirtualizationClusterTypesListOK with default headers values
func NewVirtualizationClusterTypesListOK() *VirtualizationClusterTypesListOK {
	return &VirtualizationClusterTypesListOK{}
}

/*VirtualizationClusterTypesListOK handles this case with default header values.

VirtualizationClusterTypesListOK virtualization cluster types list o k
*/
type VirtualizationClusterTypesListOK struct {
	Payload *models.VirtualizationClusterTypesListOKBody
}

func (o *VirtualizationClusterTypesListOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-types/][%d] virtualizationClusterTypesListOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterTypesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualizationClusterTypesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
