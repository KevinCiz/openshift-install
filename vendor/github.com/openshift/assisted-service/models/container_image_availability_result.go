// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ContainerImageAvailabilityResult Image availability result.
//
// swagger:model container_image_availability_result
type ContainerImageAvailabilityResult string

func NewContainerImageAvailabilityResult(value ContainerImageAvailabilityResult) *ContainerImageAvailabilityResult {
	v := value
	return &v
}

const (

	// ContainerImageAvailabilityResultSuccess captures enum value "success"
	ContainerImageAvailabilityResultSuccess ContainerImageAvailabilityResult = "success"

	// ContainerImageAvailabilityResultFailure captures enum value "failure"
	ContainerImageAvailabilityResultFailure ContainerImageAvailabilityResult = "failure"
)

// for schema
var containerImageAvailabilityResultEnum []interface{}

func init() {
	var res []ContainerImageAvailabilityResult
	if err := json.Unmarshal([]byte(`["success","failure"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		containerImageAvailabilityResultEnum = append(containerImageAvailabilityResultEnum, v)
	}
}

func (m ContainerImageAvailabilityResult) validateContainerImageAvailabilityResultEnum(path, location string, value ContainerImageAvailabilityResult) error {
	if err := validate.EnumCase(path, location, value, containerImageAvailabilityResultEnum, true); err != nil {
		return err
	}
	return nil
}

// Validate validates this container image availability result
func (m ContainerImageAvailabilityResult) Validate(formats strfmt.Registry) error {
	var res []error

	// value enum
	if err := m.validateContainerImageAvailabilityResultEnum("", "body", m); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// ContextValidate validates this container image availability result based on context it is used
func (m ContainerImageAvailabilityResult) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}
