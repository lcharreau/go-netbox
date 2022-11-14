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
)

// DcimConsoleServerPortTemplatesDeleteReader is a Reader for the DcimConsoleServerPortTemplatesDelete structure.
type DcimConsoleServerPortTemplatesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimConsoleServerPortTemplatesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDcimConsoleServerPortTemplatesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimConsoleServerPortTemplatesDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimConsoleServerPortTemplatesDeleteNoContent creates a DcimConsoleServerPortTemplatesDeleteNoContent with default headers values
func NewDcimConsoleServerPortTemplatesDeleteNoContent() *DcimConsoleServerPortTemplatesDeleteNoContent {
	return &DcimConsoleServerPortTemplatesDeleteNoContent{}
}

/*
DcimConsoleServerPortTemplatesDeleteNoContent describes a response with status code 204, with default header values.

DcimConsoleServerPortTemplatesDeleteNoContent dcim console server port templates delete no content
*/
type DcimConsoleServerPortTemplatesDeleteNoContent struct {
}

// IsSuccess returns true when this dcim console server port templates delete no content response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dcim console server port templates delete no content response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dcim console server port templates delete no content response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this dcim console server port templates delete no content response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this dcim console server port templates delete no content response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcimConsoleServerPortTemplatesDeleteNoContent ", 204)
}

func (o *DcimConsoleServerPortTemplatesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDcimConsoleServerPortTemplatesDeleteDefault creates a DcimConsoleServerPortTemplatesDeleteDefault with default headers values
func NewDcimConsoleServerPortTemplatesDeleteDefault(code int) *DcimConsoleServerPortTemplatesDeleteDefault {
	return &DcimConsoleServerPortTemplatesDeleteDefault{
		_statusCode: code,
	}
}

/*
DcimConsoleServerPortTemplatesDeleteDefault describes a response with status code -1, with default header values.

DcimConsoleServerPortTemplatesDeleteDefault dcim console server port templates delete default
*/
type DcimConsoleServerPortTemplatesDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim console server port templates delete default response
func (o *DcimConsoleServerPortTemplatesDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dcim console server port templates delete default response has a 2xx status code
func (o *DcimConsoleServerPortTemplatesDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dcim console server port templates delete default response has a 3xx status code
func (o *DcimConsoleServerPortTemplatesDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dcim console server port templates delete default response has a 4xx status code
func (o *DcimConsoleServerPortTemplatesDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dcim console server port templates delete default response has a 5xx status code
func (o *DcimConsoleServerPortTemplatesDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dcim console server port templates delete default response a status code equal to that given
func (o *DcimConsoleServerPortTemplatesDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DcimConsoleServerPortTemplatesDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /dcim/console-server-port-templates/{id}/][%d] dcim_console-server-port-templates_delete default  %+v", o._statusCode, o.Payload)
}

func (o *DcimConsoleServerPortTemplatesDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimConsoleServerPortTemplatesDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
