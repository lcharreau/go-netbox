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

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimFrontPortsPartialUpdateReader is a Reader for the DcimFrontPortsPartialUpdate structure.
type DcimFrontPortsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimFrontPortsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimFrontPortsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimFrontPortsPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimFrontPortsPartialUpdateOK creates a DcimFrontPortsPartialUpdateOK with default headers values
func NewDcimFrontPortsPartialUpdateOK() *DcimFrontPortsPartialUpdateOK {
	return &DcimFrontPortsPartialUpdateOK{}
}

/*
DcimFrontPortsPartialUpdateOK describes a response with status code 200, with default header values.

DcimFrontPortsPartialUpdateOK dcim front ports partial update o k
*/
type DcimFrontPortsPartialUpdateOK struct {
	Payload *models.FrontPort
}

// IsSuccess returns true when this dcim front ports partial update o k response has a 2xx status code
func (o *DcimFrontPortsPartialUpdateOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim front ports partial update o k response has a 3xx status code
func (o *DcimFrontPortsPartialUpdateOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim front ports partial update o k response has a 4xx status code
func (o *DcimFrontPortsPartialUpdateOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim front ports partial update o k response has a 5xx status code
func (o *DcimFrontPortsPartialUpdateOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim front ports partial update o k response a status code equal to that given
func (o *DcimFrontPortsPartialUpdateOK) IsCode(code int) bool {
	return code == 200
}

func (o *DcimFrontPortsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcimFrontPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsPartialUpdateOK) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcimFrontPortsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimFrontPortsPartialUpdateOK) GetPayload() *models.FrontPort {
	return o.Payload
}

func (o *DcimFrontPortsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FrontPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimFrontPortsPartialUpdateDefault creates a DcimFrontPortsPartialUpdateDefault with default headers values
func NewDcimFrontPortsPartialUpdateDefault(code int) *DcimFrontPortsPartialUpdateDefault {
	return &DcimFrontPortsPartialUpdateDefault{
		_statusCode: code,
	}
}

/*
DcimFrontPortsPartialUpdateDefault describes a response with status code -1, with default header values.

DcimFrontPortsPartialUpdateDefault dcim front ports partial update default
*/
type DcimFrontPortsPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim front ports partial update default response
func (o *DcimFrontPortsPartialUpdateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim front ports partial update default response has a 2xx status code
func (o *DcimFrontPortsPartialUpdateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim front ports partial update default response has a 3xx status code
func (o *DcimFrontPortsPartialUpdateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim front ports partial update default response has a 4xx status code
func (o *DcimFrontPortsPartialUpdateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim front ports partial update default response has a 5xx status code
func (o *DcimFrontPortsPartialUpdateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim front ports partial update default response a status code equal to that given
func (o *DcimFrontPortsPartialUpdateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimFrontPortsPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcim_front-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsPartialUpdateDefault) String() string {
	return fmt.Sprintf("[PATCH /dcim/front-ports/{id}/][%d] dcim_front-ports_partial_update default  %+v", o._statusCode, o.Payload)
}

func (o *DcimFrontPortsPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimFrontPortsPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
