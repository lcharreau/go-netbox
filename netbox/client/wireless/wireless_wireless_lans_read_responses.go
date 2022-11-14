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

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// WirelessWirelessLansReadReader is a Reader for the WirelessWirelessLansRead structure.
type WirelessWirelessLansReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *WirelessWirelessLansReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewWirelessWirelessLansReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewWirelessWirelessLansReadDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewWirelessWirelessLansReadOK creates a WirelessWirelessLansReadOK with default headers values
func NewWirelessWirelessLansReadOK() *WirelessWirelessLansReadOK {
	return &WirelessWirelessLansReadOK{}
}

/*
WirelessWirelessLansReadOK describes a response with status code 200, with default header values.

WirelessWirelessLansReadOK wireless wireless lans read o k
*/
type WirelessWirelessLansReadOK struct {
	Payload *models.WirelessLAN
}

// IsSuccess returns true when this wireless wireless lans read o k response has a 2xx status code
func (o *WirelessWirelessLansReadOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this wireless wireless lans read o k response has a 3xx status code
func (o *WirelessWirelessLansReadOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this wireless wireless lans read o k response has a 4xx status code
func (o *WirelessWirelessLansReadOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this wireless wireless lans read o k response has a 5xx status code
func (o *WirelessWirelessLansReadOK) IsServerError() bool {
	return false
}

// IsCode returns true when this wireless wireless lans read o k response a status code equal to that given
func (o *WirelessWirelessLansReadOK) IsCode(code int) bool {
	return code == 200
}

func (o *WirelessWirelessLansReadOK) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansReadOK) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wirelessWirelessLansReadOK  %+v", 200, o.Payload)
}

func (o *WirelessWirelessLansReadOK) GetPayload() *models.WirelessLAN {
	return o.Payload
}

func (o *WirelessWirelessLansReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WirelessLAN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewWirelessWirelessLansReadDefault creates a WirelessWirelessLansReadDefault with default headers values
func NewWirelessWirelessLansReadDefault(code int) *WirelessWirelessLansReadDefault {
	return &WirelessWirelessLansReadDefault{
		_statusCode: code,
	}
}

/*
WirelessWirelessLansReadDefault describes a response with status code -1, with default header values.

WirelessWirelessLansReadDefault wireless wireless lans read default
*/
type WirelessWirelessLansReadDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the wireless wireless lans read default response
func (o *WirelessWirelessLansReadDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this wireless wireless lans read default response has a 2xx status code
func (o *WirelessWirelessLansReadDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this wireless wireless lans read default response has a 3xx status code
func (o *WirelessWirelessLansReadDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this wireless wireless lans read default response has a 4xx status code
func (o *WirelessWirelessLansReadDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this wireless wireless lans read default response has a 5xx status code
func (o *WirelessWirelessLansReadDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this wireless wireless lans read default response a status code equal to that given
func (o *WirelessWirelessLansReadDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *WirelessWirelessLansReadDefault) Error() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wireless_wireless-lans_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansReadDefault) String() string {
	return fmt.Sprintf("[GET /wireless/wireless-lans/{id}/][%d] wireless_wireless-lans_read default  %+v", o._statusCode, o.Payload)
}

func (o *WirelessWirelessLansReadDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *WirelessWirelessLansReadDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
