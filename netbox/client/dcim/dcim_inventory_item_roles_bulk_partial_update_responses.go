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

// DcimInventoryItemRolesBulkPartialUpdateReader is a Reader for the DcimInventoryItemRolesBulkPartialUpdate structure.
type DcimInventoryItemRolesBulkPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimInventoryItemRolesBulkPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimInventoryItemRolesBulkPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDcimInventoryItemRolesBulkPartialUpdateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDcimInventoryItemRolesBulkPartialUpdateOK creates a DcimInventoryItemRolesBulkPartialUpdateOK with default headers values
func NewDcimInventoryItemRolesBulkPartialUpdateOK() *DcimInventoryItemRolesBulkPartialUpdateOK {
	return &DcimInventoryItemRolesBulkPartialUpdateOK{}
}

/* DcimInventoryItemRolesBulkPartialUpdateOK describes a response with status code 200, with default header values.

DcimInventoryItemRolesBulkPartialUpdateOK dcim inventory item roles bulk partial update o k
*/
type DcimInventoryItemRolesBulkPartialUpdateOK struct {
	Payload *models.InventoryItemRole
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcimInventoryItemRolesBulkPartialUpdateOK  %+v", 200, o.Payload)
}
func (o *DcimInventoryItemRolesBulkPartialUpdateOK) GetPayload() *models.InventoryItemRole {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InventoryItemRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDcimInventoryItemRolesBulkPartialUpdateDefault creates a DcimInventoryItemRolesBulkPartialUpdateDefault with default headers values
func NewDcimInventoryItemRolesBulkPartialUpdateDefault(code int) *DcimInventoryItemRolesBulkPartialUpdateDefault {
	return &DcimInventoryItemRolesBulkPartialUpdateDefault{
		_statusCode: code,
	}
}

/* DcimInventoryItemRolesBulkPartialUpdateDefault describes a response with status code -1, with default header values.

DcimInventoryItemRolesBulkPartialUpdateDefault dcim inventory item roles bulk partial update default
*/
type DcimInventoryItemRolesBulkPartialUpdateDefault struct {
	_statusCode int

	Payload interface{}
}

// Code gets the status code for the dcim inventory item roles bulk partial update default response
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) Code() int {
	return o._statusCode
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) Error() string {
	return fmt.Sprintf("[PATCH /dcim/inventory-item-roles/][%d] dcim_inventory-item-roles_bulk_partial_update default  %+v", o._statusCode, o.Payload)
}
func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) GetPayload() interface{} {
	return o.Payload
}

func (o *DcimInventoryItemRolesBulkPartialUpdateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}