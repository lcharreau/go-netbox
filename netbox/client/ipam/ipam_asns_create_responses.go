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

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// IpamAsnsCreateReader is a Reader for the IpamAsnsCreate structure.
type IpamAsnsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpamAsnsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewIpamAsnsCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewIpamAsnsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewIpamAsnsCreateCreated creates a IpamAsnsCreateCreated with default headers values
func NewIpamAsnsCreateCreated() *IpamAsnsCreateCreated {
	return &IpamAsnsCreateCreated{}
}

/*
IpamAsnsCreateCreated describes a response with status code 201, with default header values.

IpamAsnsCreateCreated ipam asns create created
*/
type IpamAsnsCreateCreated struct {
	Payload *models.ASN
}

// IsSuccess returns true when this ipam asns create created response has a 2xx status code
func (o *IpamAsnsCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipam asns create created response has a 3xx status code
func (o *IpamAsnsCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipam asns create created response has a 4xx status code
func (o *IpamAsnsCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipam asns create created response has a 5xx status code
func (o *IpamAsnsCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this ipam asns create created response a status code equal to that given
func (o *IpamAsnsCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *IpamAsnsCreateCreated) Error() string {
	return fmt.Sprintf("[POST /ipam/asns/][%d] ipamAsnsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamAsnsCreateCreated) String() string {
	return fmt.Sprintf("[POST /ipam/asns/][%d] ipamAsnsCreateCreated  %+v", 201, o.Payload)
}

func (o *IpamAsnsCreateCreated) GetPayload() *models.ASN {
	return o.Payload
}

func (o *IpamAsnsCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ASN)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewIpamAsnsCreateDefault creates a IpamAsnsCreateDefault with default headers values
func NewIpamAsnsCreateDefault(code int) *IpamAsnsCreateDefault {
	return &IpamAsnsCreateDefault{
		_statusCode: code,
	}
}

/*
IpamAsnsCreateDefault describes a response with status code -1, with default header values.

IpamAsnsCreateDefault ipam asns create default
*/
type IpamAsnsCreateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the ipam asns create default response
func (o *IpamAsnsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this ipam asns create default response has a 2xx status code
func (o *IpamAsnsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this ipam asns create default response has a 3xx status code
func (o *IpamAsnsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this ipam asns create default response has a 4xx status code
func (o *IpamAsnsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this ipam asns create default response has a 5xx status code
func (o *IpamAsnsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this ipam asns create default response a status code equal to that given
func (o *IpamAsnsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *IpamAsnsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /ipam/asns/][%d] ipam_asns_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsCreateDefault) String() string {
	return fmt.Sprintf("[POST /ipam/asns/][%d] ipam_asns_create default  %+v", o._statusCode, o.Payload)
}

func (o *IpamAsnsCreateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *IpamAsnsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
