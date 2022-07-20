//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

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

// NewSchemaObjectsSnapshotsRestoreParams creates a new SchemaObjectsSnapshotsRestoreParams object
// with the default values initialized.
func NewSchemaObjectsSnapshotsRestoreParams() *SchemaObjectsSnapshotsRestoreParams {
	var ()
	return &SchemaObjectsSnapshotsRestoreParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSchemaObjectsSnapshotsRestoreParamsWithTimeout creates a new SchemaObjectsSnapshotsRestoreParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSchemaObjectsSnapshotsRestoreParamsWithTimeout(timeout time.Duration) *SchemaObjectsSnapshotsRestoreParams {
	var ()
	return &SchemaObjectsSnapshotsRestoreParams{

		timeout: timeout,
	}
}

// NewSchemaObjectsSnapshotsRestoreParamsWithContext creates a new SchemaObjectsSnapshotsRestoreParams object
// with the default values initialized, and the ability to set a context for a request
func NewSchemaObjectsSnapshotsRestoreParamsWithContext(ctx context.Context) *SchemaObjectsSnapshotsRestoreParams {
	var ()
	return &SchemaObjectsSnapshotsRestoreParams{

		Context: ctx,
	}
}

// NewSchemaObjectsSnapshotsRestoreParamsWithHTTPClient creates a new SchemaObjectsSnapshotsRestoreParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSchemaObjectsSnapshotsRestoreParamsWithHTTPClient(client *http.Client) *SchemaObjectsSnapshotsRestoreParams {
	var ()
	return &SchemaObjectsSnapshotsRestoreParams{
		HTTPClient: client,
	}
}

/*SchemaObjectsSnapshotsRestoreParams contains all the parameters to send to the API endpoint
for the schema objects snapshots restore operation typically these are written to a http.Request
*/
type SchemaObjectsSnapshotsRestoreParams struct {

	/*ClassName
	  The name of the class

	*/
	ClassName string
	/*ID
	  The Id of the snapshot

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) WithTimeout(timeout time.Duration) *SchemaObjectsSnapshotsRestoreParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) WithContext(ctx context.Context) *SchemaObjectsSnapshotsRestoreParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) WithHTTPClient(client *http.Client) *SchemaObjectsSnapshotsRestoreParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClassName adds the className to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) WithClassName(className string) *SchemaObjectsSnapshotsRestoreParams {
	o.SetClassName(className)
	return o
}

// SetClassName adds the className to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) SetClassName(className string) {
	o.ClassName = className
}

// WithID adds the id to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) WithID(id string) *SchemaObjectsSnapshotsRestoreParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the schema objects snapshots restore params
func (o *SchemaObjectsSnapshotsRestoreParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *SchemaObjectsSnapshotsRestoreParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param className
	if err := r.SetPathParam("className", o.ClassName); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
