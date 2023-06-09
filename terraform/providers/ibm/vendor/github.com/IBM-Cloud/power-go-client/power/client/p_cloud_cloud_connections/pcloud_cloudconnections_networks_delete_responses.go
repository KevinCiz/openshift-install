// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_cloud_connections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudconnectionsNetworksDeleteReader is a Reader for the PcloudCloudconnectionsNetworksDelete structure.
type PcloudCloudconnectionsNetworksDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsNetworksDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsNetworksDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsNetworksDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsNetworksDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsNetworksDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsNetworksDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsNetworksDeleteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudconnectionsNetworksDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsNetworksDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsNetworksDeleteOK creates a PcloudCloudconnectionsNetworksDeleteOK with default headers values
func NewPcloudCloudconnectionsNetworksDeleteOK() *PcloudCloudconnectionsNetworksDeleteOK {
	return &PcloudCloudconnectionsNetworksDeleteOK{}
}

/* PcloudCloudconnectionsNetworksDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsNetworksDeleteOK struct {
	Payload *models.CloudConnection
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteOK) GetPayload() *models.CloudConnection {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CloudConnection)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteAccepted creates a PcloudCloudconnectionsNetworksDeleteAccepted with default headers values
func NewPcloudCloudconnectionsNetworksDeleteAccepted() *PcloudCloudconnectionsNetworksDeleteAccepted {
	return &PcloudCloudconnectionsNetworksDeleteAccepted{}
}

/* PcloudCloudconnectionsNetworksDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsNetworksDeleteAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteBadRequest creates a PcloudCloudconnectionsNetworksDeleteBadRequest with default headers values
func NewPcloudCloudconnectionsNetworksDeleteBadRequest() *PcloudCloudconnectionsNetworksDeleteBadRequest {
	return &PcloudCloudconnectionsNetworksDeleteBadRequest{}
}

/* PcloudCloudconnectionsNetworksDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsNetworksDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteUnauthorized creates a PcloudCloudconnectionsNetworksDeleteUnauthorized with default headers values
func NewPcloudCloudconnectionsNetworksDeleteUnauthorized() *PcloudCloudconnectionsNetworksDeleteUnauthorized {
	return &PcloudCloudconnectionsNetworksDeleteUnauthorized{}
}

/* PcloudCloudconnectionsNetworksDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsNetworksDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteNotFound creates a PcloudCloudconnectionsNetworksDeleteNotFound with default headers values
func NewPcloudCloudconnectionsNetworksDeleteNotFound() *PcloudCloudconnectionsNetworksDeleteNotFound {
	return &PcloudCloudconnectionsNetworksDeleteNotFound{}
}

/* PcloudCloudconnectionsNetworksDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsNetworksDeleteNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteRequestTimeout creates a PcloudCloudconnectionsNetworksDeleteRequestTimeout with default headers values
func NewPcloudCloudconnectionsNetworksDeleteRequestTimeout() *PcloudCloudconnectionsNetworksDeleteRequestTimeout {
	return &PcloudCloudconnectionsNetworksDeleteRequestTimeout{}
}

/* PcloudCloudconnectionsNetworksDeleteRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsNetworksDeleteRequestTimeout struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteRequestTimeout  %+v", 408, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteGone creates a PcloudCloudconnectionsNetworksDeleteGone with default headers values
func NewPcloudCloudconnectionsNetworksDeleteGone() *PcloudCloudconnectionsNetworksDeleteGone {
	return &PcloudCloudconnectionsNetworksDeleteGone{}
}

/* PcloudCloudconnectionsNetworksDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudconnectionsNetworksDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteGone  %+v", 410, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsNetworksDeleteInternalServerError creates a PcloudCloudconnectionsNetworksDeleteInternalServerError with default headers values
func NewPcloudCloudconnectionsNetworksDeleteInternalServerError() *PcloudCloudconnectionsNetworksDeleteInternalServerError {
	return &PcloudCloudconnectionsNetworksDeleteInternalServerError{}
}

/* PcloudCloudconnectionsNetworksDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsNetworksDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}/networks/{network_id}][%d] pcloudCloudconnectionsNetworksDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsNetworksDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
