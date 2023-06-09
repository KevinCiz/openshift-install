// Code generated by go-swagger; DO NOT EDIT.

package installer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/openshift/assisted-service/models"
)

// NewV2ImportClusterParams creates a new V2ImportClusterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewV2ImportClusterParams() *V2ImportClusterParams {
	return &V2ImportClusterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewV2ImportClusterParamsWithTimeout creates a new V2ImportClusterParams object
// with the ability to set a timeout on a request.
func NewV2ImportClusterParamsWithTimeout(timeout time.Duration) *V2ImportClusterParams {
	return &V2ImportClusterParams{
		timeout: timeout,
	}
}

// NewV2ImportClusterParamsWithContext creates a new V2ImportClusterParams object
// with the ability to set a context for a request.
func NewV2ImportClusterParamsWithContext(ctx context.Context) *V2ImportClusterParams {
	return &V2ImportClusterParams{
		Context: ctx,
	}
}

// NewV2ImportClusterParamsWithHTTPClient creates a new V2ImportClusterParams object
// with the ability to set a custom HTTPClient for a request.
func NewV2ImportClusterParamsWithHTTPClient(client *http.Client) *V2ImportClusterParams {
	return &V2ImportClusterParams{
		HTTPClient: client,
	}
}

/* V2ImportClusterParams contains all the parameters to send to the API endpoint
   for the v2 import cluster operation.

   Typically these are written to a http.Request.
*/
type V2ImportClusterParams struct {

	/* NewImportClusterParams.

	   Parameters for importing a OCP cluster for adding nodes.
	*/
	NewImportClusterParams *models.ImportClusterParams

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the v2 import cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ImportClusterParams) WithDefaults() *V2ImportClusterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the v2 import cluster params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *V2ImportClusterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the v2 import cluster params
func (o *V2ImportClusterParams) WithTimeout(timeout time.Duration) *V2ImportClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the v2 import cluster params
func (o *V2ImportClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the v2 import cluster params
func (o *V2ImportClusterParams) WithContext(ctx context.Context) *V2ImportClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the v2 import cluster params
func (o *V2ImportClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the v2 import cluster params
func (o *V2ImportClusterParams) WithHTTPClient(client *http.Client) *V2ImportClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the v2 import cluster params
func (o *V2ImportClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNewImportClusterParams adds the newImportClusterParams to the v2 import cluster params
func (o *V2ImportClusterParams) WithNewImportClusterParams(newImportClusterParams *models.ImportClusterParams) *V2ImportClusterParams {
	o.SetNewImportClusterParams(newImportClusterParams)
	return o
}

// SetNewImportClusterParams adds the newImportClusterParams to the v2 import cluster params
func (o *V2ImportClusterParams) SetNewImportClusterParams(newImportClusterParams *models.ImportClusterParams) {
	o.NewImportClusterParams = newImportClusterParams
}

// WriteToRequest writes these params to a swagger request
func (o *V2ImportClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.NewImportClusterParams != nil {
		if err := r.SetBodyParam(o.NewImportClusterParams); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
