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

package tenancy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// TenancyTenantsBulkDeleteReader is a Reader for the TenancyTenantsBulkDelete structure.
type TenancyTenantsBulkDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TenancyTenantsBulkDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewTenancyTenantsBulkDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewTenancyTenantsBulkDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewTenancyTenantsBulkDeleteNoContent creates a TenancyTenantsBulkDeleteNoContent with default headers values
func NewTenancyTenantsBulkDeleteNoContent() *TenancyTenantsBulkDeleteNoContent {
	return &TenancyTenantsBulkDeleteNoContent{}
}

/*
TenancyTenantsBulkDeleteNoContent describes a response with status code 204, with default header values.

TenancyTenantsBulkDeleteNoContent tenancy tenants bulk delete no content
*/
type TenancyTenantsBulkDeleteNoContent struct {
}

// IsSuccess returns true when this tenancy tenants bulk delete no content response has a 2xx status code
func (o *TenancyTenantsBulkDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this tenancy tenants bulk delete no content response has a 3xx status code
func (o *TenancyTenantsBulkDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this tenancy tenants bulk delete no content response has a 4xx status code
func (o *TenancyTenantsBulkDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this tenancy tenants bulk delete no content response has a 5xx status code
func (o *TenancyTenantsBulkDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this tenancy tenants bulk delete no content response a status code equal to that given
func (o *TenancyTenantsBulkDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

func (o *TenancyTenantsBulkDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/][%d] tenancyTenantsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantsBulkDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/][%d] tenancyTenantsBulkDeleteNoContent ", 204)
}

func (o *TenancyTenantsBulkDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewTenancyTenantsBulkDeleteDefault creates a TenancyTenantsBulkDeleteDefault with default headers values
func NewTenancyTenantsBulkDeleteDefault(code int) *TenancyTenantsBulkDeleteDefault {
	return &TenancyTenantsBulkDeleteDefault{
		_statusCode: code,
	}
}

/*
TenancyTenantsBulkDeleteDefault describes a response with status code -1, with default header values.

TenancyTenantsBulkDeleteDefault tenancy tenants bulk delete default
*/
type TenancyTenantsBulkDeleteDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the tenancy tenants bulk delete default response
func (o *TenancyTenantsBulkDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this tenancy tenants bulk delete default response has a 2xx status code
func (o *TenancyTenantsBulkDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this tenancy tenants bulk delete default response has a 3xx status code
func (o *TenancyTenantsBulkDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this tenancy tenants bulk delete default response has a 4xx status code
func (o *TenancyTenantsBulkDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this tenancy tenants bulk delete default response has a 5xx status code
func (o *TenancyTenantsBulkDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this tenancy tenants bulk delete default response a status code equal to that given
func (o *TenancyTenantsBulkDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *TenancyTenantsBulkDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/][%d] tenancy_tenants_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsBulkDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /tenancy/tenants/][%d] tenancy_tenants_bulk_delete default  %+v", o._statusCode, o.Payload)
}

func (o *TenancyTenantsBulkDeleteDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *TenancyTenantsBulkDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
