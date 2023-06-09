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
)

// NewTransformClusterToDay2Params creates a new TransformClusterToDay2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTransformClusterToDay2Params() *TransformClusterToDay2Params {
	return &TransformClusterToDay2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewTransformClusterToDay2ParamsWithTimeout creates a new TransformClusterToDay2Params object
// with the ability to set a timeout on a request.
func NewTransformClusterToDay2ParamsWithTimeout(timeout time.Duration) *TransformClusterToDay2Params {
	return &TransformClusterToDay2Params{
		timeout: timeout,
	}
}

// NewTransformClusterToDay2ParamsWithContext creates a new TransformClusterToDay2Params object
// with the ability to set a context for a request.
func NewTransformClusterToDay2ParamsWithContext(ctx context.Context) *TransformClusterToDay2Params {
	return &TransformClusterToDay2Params{
		Context: ctx,
	}
}

// NewTransformClusterToDay2ParamsWithHTTPClient creates a new TransformClusterToDay2Params object
// with the ability to set a custom HTTPClient for a request.
func NewTransformClusterToDay2ParamsWithHTTPClient(client *http.Client) *TransformClusterToDay2Params {
	return &TransformClusterToDay2Params{
		HTTPClient: client,
	}
}

/* TransformClusterToDay2Params contains all the parameters to send to the API endpoint
   for the transform cluster to day2 operation.

   Typically these are written to a http.Request.
*/
type TransformClusterToDay2Params struct {

	/* ClusterID.

	   The cluster to transform to day2 and to allow adding hosts by this.

	   Format: uuid
	*/
	ClusterID strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the transform cluster to day2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransformClusterToDay2Params) WithDefaults() *TransformClusterToDay2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the transform cluster to day2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TransformClusterToDay2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) WithTimeout(timeout time.Duration) *TransformClusterToDay2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) WithContext(ctx context.Context) *TransformClusterToDay2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) WithHTTPClient(client *http.Client) *TransformClusterToDay2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) WithClusterID(clusterID strfmt.UUID) *TransformClusterToDay2Params {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the transform cluster to day2 params
func (o *TransformClusterToDay2Params) SetClusterID(clusterID strfmt.UUID) {
	o.ClusterID = clusterID
}

// WriteToRequest writes these params to a swagger request
func (o *TransformClusterToDay2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
