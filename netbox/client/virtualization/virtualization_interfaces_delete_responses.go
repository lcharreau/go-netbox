// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// VirtualizationInterfacesDeleteReader is a Reader for the VirtualizationInterfacesDelete structure.
type VirtualizationInterfacesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationInterfacesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewVirtualizationInterfacesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVirtualizationInterfacesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVirtualizationInterfacesDeleteNoContent creates a VirtualizationInterfacesDeleteNoContent with default headers values
func NewVirtualizationInterfacesDeleteNoContent() *VirtualizationInterfacesDeleteNoContent {
	return &VirtualizationInterfacesDeleteNoContent{}
}

/*
VirtualizationInterfacesDeleteNoContent describes a response with status code 204, with default header values.

VirtualizationInterfacesDeleteNoContent virtualization interfaces delete no content
*/
type VirtualizationInterfacesDeleteNoContent struct {
}

// IsSuccess returns true when this virtualization interfaces delete no content response has a 2xx status code
func (o *VirtualizationInterfacesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this virtualization interfaces delete no content response has a 3xx status code
func (o *VirtualizationInterfacesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this virtualization interfaces delete no content response has a 4xx status code
func (o *VirtualizationInterfacesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this virtualization interfaces delete no content response has a 5xx status code
func (o *VirtualizationInterfacesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this virtualization interfaces delete no content response a status code equal to that given
func (o *VirtualizationInterfacesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *VirtualizationInterfacesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualizationInterfacesDeleteNoContent ", 204)
}

func (o *VirtualizationInterfacesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualizationInterfacesDeleteNoContent ", 204)
}

func (o *VirtualizationInterfacesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewVirtualizationInterfacesDeleteDefault creates a VirtualizationInterfacesDeleteDefault with default headers values
func NewVirtualizationInterfacesDeleteDefault(code int) *VirtualizationInterfacesDeleteDefault {
	return &VirtualizationInterfacesDeleteDefault{
		_statusCode: code,
	}
}

/*
VirtualizationInterfacesDeleteDefault describes a response with status code -1, with default header values.

VirtualizationInterfacesDeleteDefault virtualization interfaces delete default
*/
type VirtualizationInterfacesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the virtualization interfaces delete default response
func (o *VirtualizationInterfacesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this virtualization interfaces delete default response has a 2xx status code
func (o *VirtualizationInterfacesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this virtualization interfaces delete default response has a 3xx status code
func (o *VirtualizationInterfacesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this virtualization interfaces delete default response has a 4xx status code
func (o *VirtualizationInterfacesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this virtualization interfaces delete default response has a 5xx status code
func (o *VirtualizationInterfacesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this virtualization interfaces delete default response a status code equal to that given
func (o *VirtualizationInterfacesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VirtualizationInterfacesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualization_interfaces_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationInterfacesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /virtualization/interfaces/{id}/][%d] virtualization_interfaces_delete default  %+v", o._statusCode, o.Payload)
}

func (o *VirtualizationInterfacesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *VirtualizationInterfacesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
