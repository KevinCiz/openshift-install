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

// PcloudCloudconnectionsDeleteReader is a Reader for the PcloudCloudconnectionsDelete structure.
type PcloudCloudconnectionsDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudconnectionsDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudconnectionsDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 202:
		result := NewPcloudCloudconnectionsDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudconnectionsDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudconnectionsDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudconnectionsDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 408:
		result := NewPcloudCloudconnectionsDeleteRequestTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewPcloudCloudconnectionsDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudconnectionsDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudconnectionsDeleteOK creates a PcloudCloudconnectionsDeleteOK with default headers values
func NewPcloudCloudconnectionsDeleteOK() *PcloudCloudconnectionsDeleteOK {
	return &PcloudCloudconnectionsDeleteOK{}
}

/* PcloudCloudconnectionsDeleteOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudconnectionsDeleteOK struct {
	Payload models.Object
}

func (o *PcloudCloudconnectionsDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteOK  %+v", 200, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteOK) GetPayload() models.Object {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteAccepted creates a PcloudCloudconnectionsDeleteAccepted with default headers values
func NewPcloudCloudconnectionsDeleteAccepted() *PcloudCloudconnectionsDeleteAccepted {
	return &PcloudCloudconnectionsDeleteAccepted{}
}

/* PcloudCloudconnectionsDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type PcloudCloudconnectionsDeleteAccepted struct {
	Payload *models.JobReference
}

func (o *PcloudCloudconnectionsDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteAccepted  %+v", 202, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteAccepted) GetPayload() *models.JobReference {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobReference)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteBadRequest creates a PcloudCloudconnectionsDeleteBadRequest with default headers values
func NewPcloudCloudconnectionsDeleteBadRequest() *PcloudCloudconnectionsDeleteBadRequest {
	return &PcloudCloudconnectionsDeleteBadRequest{}
}

/* PcloudCloudconnectionsDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudconnectionsDeleteBadRequest struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteBadRequest  %+v", 400, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteUnauthorized creates a PcloudCloudconnectionsDeleteUnauthorized with default headers values
func NewPcloudCloudconnectionsDeleteUnauthorized() *PcloudCloudconnectionsDeleteUnauthorized {
	return &PcloudCloudconnectionsDeleteUnauthorized{}
}

/* PcloudCloudconnectionsDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudconnectionsDeleteUnauthorized struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteUnauthorized  %+v", 401, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteNotFound creates a PcloudCloudconnectionsDeleteNotFound with default headers values
func NewPcloudCloudconnectionsDeleteNotFound() *PcloudCloudconnectionsDeleteNotFound {
	return &PcloudCloudconnectionsDeleteNotFound{}
}

/* PcloudCloudconnectionsDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudconnectionsDeleteNotFound struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteNotFound) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteNotFound  %+v", 404, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteRequestTimeout creates a PcloudCloudconnectionsDeleteRequestTimeout with default headers values
func NewPcloudCloudconnectionsDeleteRequestTimeout() *PcloudCloudconnectionsDeleteRequestTimeout {
	return &PcloudCloudconnectionsDeleteRequestTimeout{}
}

/* PcloudCloudconnectionsDeleteRequestTimeout describes a response with status code 408, with default header values.

Request Timeout
*/
type PcloudCloudconnectionsDeleteRequestTimeout struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteRequestTimeout  %+v", 408, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteRequestTimeout) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteRequestTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteGone creates a PcloudCloudconnectionsDeleteGone with default headers values
func NewPcloudCloudconnectionsDeleteGone() *PcloudCloudconnectionsDeleteGone {
	return &PcloudCloudconnectionsDeleteGone{}
}

/* PcloudCloudconnectionsDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type PcloudCloudconnectionsDeleteGone struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteGone) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteGone  %+v", 410, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudconnectionsDeleteInternalServerError creates a PcloudCloudconnectionsDeleteInternalServerError with default headers values
func NewPcloudCloudconnectionsDeleteInternalServerError() *PcloudCloudconnectionsDeleteInternalServerError {
	return &PcloudCloudconnectionsDeleteInternalServerError{}
}

/* PcloudCloudconnectionsDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudconnectionsDeleteInternalServerError struct {
	Payload *models.Error
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /pcloud/v1/cloud-instances/{cloud_instance_id}/cloud-connections/{cloud_connection_id}][%d] pcloudCloudconnectionsDeleteInternalServerError  %+v", 500, o.Payload)
}
func (o *PcloudCloudconnectionsDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudconnectionsDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
