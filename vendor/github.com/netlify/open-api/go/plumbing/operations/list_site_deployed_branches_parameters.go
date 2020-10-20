// Code generated by go-swagger; DO NOT EDIT.

package operations

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

// NewListSiteDeployedBranchesParams creates a new ListSiteDeployedBranchesParams object
// with the default values initialized.
func NewListSiteDeployedBranchesParams() *ListSiteDeployedBranchesParams {
	var ()
	return &ListSiteDeployedBranchesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListSiteDeployedBranchesParamsWithTimeout creates a new ListSiteDeployedBranchesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListSiteDeployedBranchesParamsWithTimeout(timeout time.Duration) *ListSiteDeployedBranchesParams {
	var ()
	return &ListSiteDeployedBranchesParams{

		timeout: timeout,
	}
}

// NewListSiteDeployedBranchesParamsWithContext creates a new ListSiteDeployedBranchesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListSiteDeployedBranchesParamsWithContext(ctx context.Context) *ListSiteDeployedBranchesParams {
	var ()
	return &ListSiteDeployedBranchesParams{

		Context: ctx,
	}
}

// NewListSiteDeployedBranchesParamsWithHTTPClient creates a new ListSiteDeployedBranchesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListSiteDeployedBranchesParamsWithHTTPClient(client *http.Client) *ListSiteDeployedBranchesParams {
	var ()
	return &ListSiteDeployedBranchesParams{
		HTTPClient: client,
	}
}

/*ListSiteDeployedBranchesParams contains all the parameters to send to the API endpoint
for the list site deployed branches operation typically these are written to a http.Request
*/
type ListSiteDeployedBranchesParams struct {

	/*SiteID*/
	SiteID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) WithTimeout(timeout time.Duration) *ListSiteDeployedBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) WithContext(ctx context.Context) *ListSiteDeployedBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) WithHTTPClient(client *http.Client) *ListSiteDeployedBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithSiteID adds the siteID to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) WithSiteID(siteID string) *ListSiteDeployedBranchesParams {
	o.SetSiteID(siteID)
	return o
}

// SetSiteID adds the siteId to the list site deployed branches params
func (o *ListSiteDeployedBranchesParams) SetSiteID(siteID string) {
	o.SiteID = siteID
}

// WriteToRequest writes these params to a swagger request
func (o *ListSiteDeployedBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param site_id
	if err := r.SetPathParam("site_id", o.SiteID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
