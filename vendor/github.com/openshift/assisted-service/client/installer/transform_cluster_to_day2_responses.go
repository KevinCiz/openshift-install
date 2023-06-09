// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// TransformClusterToDay2Reader is a Reader for the TransformClusterToDay2 structure.
type TransformClusterToDay2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TransformClusterToDay2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewTransformClusterToDay2Accepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewTransformClusterToDay2Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTransformClusterToDay2Forbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTransformClusterToDay2NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 405:
		result := NewTransformClusterToDay2MethodNotAllowed()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTransformClusterToDay2Conflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTransformClusterToDay2InternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewTransformClusterToDay2Accepted creates a TransformClusterToDay2Accepted with default headers values
func NewTransformClusterToDay2Accepted() *TransformClusterToDay2Accepted {
	return &TransformClusterToDay2Accepted{}
}

/* TransformClusterToDay2Accepted describes a response with status code 202, with default header values.

Success.
*/
type TransformClusterToDay2Accepted struct {
	Payload *models.Cluster
}

func (o *TransformClusterToDay2Accepted) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2Accepted  %+v", 202, o.Payload)
}
func (o *TransformClusterToDay2Accepted) GetPayload() *models.Cluster {
	return o.Payload
}

func (o *TransformClusterToDay2Accepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Cluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2Unauthorized creates a TransformClusterToDay2Unauthorized with default headers values
func NewTransformClusterToDay2Unauthorized() *TransformClusterToDay2Unauthorized {
	return &TransformClusterToDay2Unauthorized{}
}

/* TransformClusterToDay2Unauthorized describes a response with status code 401, with default header values.

Unauthorized.
*/
type TransformClusterToDay2Unauthorized struct {
	Payload *models.InfraError
}

func (o *TransformClusterToDay2Unauthorized) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2Unauthorized  %+v", 401, o.Payload)
}
func (o *TransformClusterToDay2Unauthorized) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *TransformClusterToDay2Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2Forbidden creates a TransformClusterToDay2Forbidden with default headers values
func NewTransformClusterToDay2Forbidden() *TransformClusterToDay2Forbidden {
	return &TransformClusterToDay2Forbidden{}
}

/* TransformClusterToDay2Forbidden describes a response with status code 403, with default header values.

Forbidden.
*/
type TransformClusterToDay2Forbidden struct {
	Payload *models.InfraError
}

func (o *TransformClusterToDay2Forbidden) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2Forbidden  %+v", 403, o.Payload)
}
func (o *TransformClusterToDay2Forbidden) GetPayload() *models.InfraError {
	return o.Payload
}

func (o *TransformClusterToDay2Forbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InfraError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2NotFound creates a TransformClusterToDay2NotFound with default headers values
func NewTransformClusterToDay2NotFound() *TransformClusterToDay2NotFound {
	return &TransformClusterToDay2NotFound{}
}

/* TransformClusterToDay2NotFound describes a response with status code 404, with default header values.

Error.
*/
type TransformClusterToDay2NotFound struct {
	Payload *models.Error
}

func (o *TransformClusterToDay2NotFound) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2NotFound  %+v", 404, o.Payload)
}
func (o *TransformClusterToDay2NotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransformClusterToDay2NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2MethodNotAllowed creates a TransformClusterToDay2MethodNotAllowed with default headers values
func NewTransformClusterToDay2MethodNotAllowed() *TransformClusterToDay2MethodNotAllowed {
	return &TransformClusterToDay2MethodNotAllowed{}
}

/* TransformClusterToDay2MethodNotAllowed describes a response with status code 405, with default header values.

Method Not Allowed.
*/
type TransformClusterToDay2MethodNotAllowed struct {
	Payload *models.Error
}

func (o *TransformClusterToDay2MethodNotAllowed) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2MethodNotAllowed  %+v", 405, o.Payload)
}
func (o *TransformClusterToDay2MethodNotAllowed) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransformClusterToDay2MethodNotAllowed) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2Conflict creates a TransformClusterToDay2Conflict with default headers values
func NewTransformClusterToDay2Conflict() *TransformClusterToDay2Conflict {
	return &TransformClusterToDay2Conflict{}
}

/* TransformClusterToDay2Conflict describes a response with status code 409, with default header values.

Error.
*/
type TransformClusterToDay2Conflict struct {
	Payload *models.Error
}

func (o *TransformClusterToDay2Conflict) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2Conflict  %+v", 409, o.Payload)
}
func (o *TransformClusterToDay2Conflict) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransformClusterToDay2Conflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTransformClusterToDay2InternalServerError creates a TransformClusterToDay2InternalServerError with default headers values
func NewTransformClusterToDay2InternalServerError() *TransformClusterToDay2InternalServerError {
	return &TransformClusterToDay2InternalServerError{}
}

/* TransformClusterToDay2InternalServerError describes a response with status code 500, with default header values.

Error.
*/
type TransformClusterToDay2InternalServerError struct {
	Payload *models.Error
}

func (o *TransformClusterToDay2InternalServerError) Error() string {
	return fmt.Sprintf("[POST /v2/clusters/{cluster_id}/actions/allow-add-workers][%d] transformClusterToDay2InternalServerError  %+v", 500, o.Payload)
}
func (o *TransformClusterToDay2InternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *TransformClusterToDay2InternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
