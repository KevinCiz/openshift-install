// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_images

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
	"github.com/go-openapi/swag"
)

// NewPcloudCloudinstancesStockimagesGetallParams creates a new PcloudCloudinstancesStockimagesGetallParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPcloudCloudinstancesStockimagesGetallParams() *PcloudCloudinstancesStockimagesGetallParams {
	return &PcloudCloudinstancesStockimagesGetallParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPcloudCloudinstancesStockimagesGetallParamsWithTimeout creates a new PcloudCloudinstancesStockimagesGetallParams object
// with the ability to set a timeout on a request.
func NewPcloudCloudinstancesStockimagesGetallParamsWithTimeout(timeout time.Duration) *PcloudCloudinstancesStockimagesGetallParams {
	return &PcloudCloudinstancesStockimagesGetallParams{
		timeout: timeout,
	}
}

// NewPcloudCloudinstancesStockimagesGetallParamsWithContext creates a new PcloudCloudinstancesStockimagesGetallParams object
// with the ability to set a context for a request.
func NewPcloudCloudinstancesStockimagesGetallParamsWithContext(ctx context.Context) *PcloudCloudinstancesStockimagesGetallParams {
	return &PcloudCloudinstancesStockimagesGetallParams{
		Context: ctx,
	}
}

// NewPcloudCloudinstancesStockimagesGetallParamsWithHTTPClient creates a new PcloudCloudinstancesStockimagesGetallParams object
// with the ability to set a custom HTTPClient for a request.
func NewPcloudCloudinstancesStockimagesGetallParamsWithHTTPClient(client *http.Client) *PcloudCloudinstancesStockimagesGetallParams {
	return &PcloudCloudinstancesStockimagesGetallParams{
		HTTPClient: client,
	}
}

/* PcloudCloudinstancesStockimagesGetallParams contains all the parameters to send to the API endpoint
   for the pcloud cloudinstances stockimages getall operation.

   Typically these are written to a http.Request.
*/
type PcloudCloudinstancesStockimagesGetallParams struct {

	/* CloudInstanceID.

	   Cloud Instance ID of a PCloud Instance
	*/
	CloudInstanceID string

	/* Sap.

	   Include SAP images with get available stock images
	*/
	Sap *bool

	/* Vtl.

	   Include VTL images with get available stock images
	*/
	Vtl *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the pcloud cloudinstances stockimages getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesStockimagesGetallParams) WithDefaults() *PcloudCloudinstancesStockimagesGetallParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the pcloud cloudinstances stockimages getall params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PcloudCloudinstancesStockimagesGetallParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithTimeout(timeout time.Duration) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithContext(ctx context.Context) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithHTTPClient(client *http.Client) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCloudInstanceID adds the cloudInstanceID to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithCloudInstanceID(cloudInstanceID string) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetCloudInstanceID(cloudInstanceID)
	return o
}

// SetCloudInstanceID adds the cloudInstanceId to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetCloudInstanceID(cloudInstanceID string) {
	o.CloudInstanceID = cloudInstanceID
}

// WithSap adds the sap to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithSap(sap *bool) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetSap(sap)
	return o
}

// SetSap adds the sap to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetSap(sap *bool) {
	o.Sap = sap
}

// WithVtl adds the vtl to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) WithVtl(vtl *bool) *PcloudCloudinstancesStockimagesGetallParams {
	o.SetVtl(vtl)
	return o
}

// SetVtl adds the vtl to the pcloud cloudinstances stockimages getall params
func (o *PcloudCloudinstancesStockimagesGetallParams) SetVtl(vtl *bool) {
	o.Vtl = vtl
}

// WriteToRequest writes these params to a swagger request
func (o *PcloudCloudinstancesStockimagesGetallParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cloud_instance_id
	if err := r.SetPathParam("cloud_instance_id", o.CloudInstanceID); err != nil {
		return err
	}

	if o.Sap != nil {

		// query param sap
		var qrSap bool

		if o.Sap != nil {
			qrSap = *o.Sap
		}
		qSap := swag.FormatBool(qrSap)
		if qSap != "" {

			if err := r.SetQueryParam("sap", qSap); err != nil {
				return err
			}
		}
	}

	if o.Vtl != nil {

		// query param vtl
		var qrVtl bool

		if o.Vtl != nil {
			qrVtl = *o.Vtl
		}
		qVtl := swag.FormatBool(qrVtl)
		if qVtl != "" {

			if err := r.SetQueryParam("vtl", qVtl); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
