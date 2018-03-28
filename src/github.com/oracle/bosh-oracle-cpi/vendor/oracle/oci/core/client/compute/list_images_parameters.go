// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewListImagesParams creates a new ListImagesParams object
// with the default values initialized.
func NewListImagesParams() *ListImagesParams {
	var ()
	return &ListImagesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListImagesParamsWithTimeout creates a new ListImagesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListImagesParamsWithTimeout(timeout time.Duration) *ListImagesParams {
	var ()
	return &ListImagesParams{

		timeout: timeout,
	}
}

// NewListImagesParamsWithContext creates a new ListImagesParams object
// with the default values initialized, and the ability to set a context for a request
func NewListImagesParamsWithContext(ctx context.Context) *ListImagesParams {
	var ()
	return &ListImagesParams{

		Context: ctx,
	}
}

// NewListImagesParamsWithHTTPClient creates a new ListImagesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListImagesParamsWithHTTPClient(client *http.Client) *ListImagesParams {
	var ()
	return &ListImagesParams{
		HTTPClient: client,
	}
}

/*ListImagesParams contains all the parameters to send to the API endpoint
for the list images operation typically these are written to a http.Request
*/
type ListImagesParams struct {

	/*CompartmentID
	  The OCID of the compartment.

	*/
	CompartmentID string
	/*DisplayName
	  A filter to return only resources that match the given display name exactly.


	*/
	DisplayName *string
	/*LifecycleState
	  A filter to only return resources that match the given lifecycle state.  The state value is case-insensitive.


	*/
	LifecycleState *string
	/*Limit
	  The maximum number of items to return in a paginated "List" call.

	Example: `500`


	*/
	Limit *int64
	/*OperatingSystem
	  The image's operating system.

	Example: `Oracle Linux`


	*/
	OperatingSystem *string
	/*OperatingSystemVersion
	  The image's operating system version.

	Example: `7.2`


	*/
	OperatingSystemVersion *string
	/*Page
	  The value of the `opc-next-page` response header from the previous "List" call.


	*/
	Page *string
	/*Shape
	  Shape name.

	*/
	Shape *string
	/*SortBy
	  The field to sort by. You can provide one sort order (`sortOrder`). Default order for
	TIMECREATED is descending. Default order for DISPLAYNAME is ascending. The DISPLAYNAME
	sort order is case sensitive.

	**Note:** In general, some "List" operations (for example, `ListInstances`) let you
	optionally filter by Availability Domain if the scope of the resource type is within a
	single Availability Domain. If you call one of these "List" operations without specifying
	an Availability Domain, the resources are grouped by Availability Domain, then sorted.


	*/
	SortBy *string
	/*SortOrder
	  The sort order to use, either ascending (`ASC`) or descending (`DESC`). The DISPLAYNAME sort order
	is case sensitive.


	*/
	SortOrder *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list images params
func (o *ListImagesParams) WithTimeout(timeout time.Duration) *ListImagesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list images params
func (o *ListImagesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list images params
func (o *ListImagesParams) WithContext(ctx context.Context) *ListImagesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list images params
func (o *ListImagesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list images params
func (o *ListImagesParams) WithHTTPClient(client *http.Client) *ListImagesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list images params
func (o *ListImagesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCompartmentID adds the compartmentID to the list images params
func (o *ListImagesParams) WithCompartmentID(compartmentID string) *ListImagesParams {
	o.SetCompartmentID(compartmentID)
	return o
}

// SetCompartmentID adds the compartmentId to the list images params
func (o *ListImagesParams) SetCompartmentID(compartmentID string) {
	o.CompartmentID = compartmentID
}

// WithDisplayName adds the displayName to the list images params
func (o *ListImagesParams) WithDisplayName(displayName *string) *ListImagesParams {
	o.SetDisplayName(displayName)
	return o
}

// SetDisplayName adds the displayName to the list images params
func (o *ListImagesParams) SetDisplayName(displayName *string) {
	o.DisplayName = displayName
}

// WithLifecycleState adds the lifecycleState to the list images params
func (o *ListImagesParams) WithLifecycleState(lifecycleState *string) *ListImagesParams {
	o.SetLifecycleState(lifecycleState)
	return o
}

// SetLifecycleState adds the lifecycleState to the list images params
func (o *ListImagesParams) SetLifecycleState(lifecycleState *string) {
	o.LifecycleState = lifecycleState
}

// WithLimit adds the limit to the list images params
func (o *ListImagesParams) WithLimit(limit *int64) *ListImagesParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list images params
func (o *ListImagesParams) SetLimit(limit *int64) {
	o.Limit = limit
}

// WithOperatingSystem adds the operatingSystem to the list images params
func (o *ListImagesParams) WithOperatingSystem(operatingSystem *string) *ListImagesParams {
	o.SetOperatingSystem(operatingSystem)
	return o
}

// SetOperatingSystem adds the operatingSystem to the list images params
func (o *ListImagesParams) SetOperatingSystem(operatingSystem *string) {
	o.OperatingSystem = operatingSystem
}

// WithOperatingSystemVersion adds the operatingSystemVersion to the list images params
func (o *ListImagesParams) WithOperatingSystemVersion(operatingSystemVersion *string) *ListImagesParams {
	o.SetOperatingSystemVersion(operatingSystemVersion)
	return o
}

// SetOperatingSystemVersion adds the operatingSystemVersion to the list images params
func (o *ListImagesParams) SetOperatingSystemVersion(operatingSystemVersion *string) {
	o.OperatingSystemVersion = operatingSystemVersion
}

// WithPage adds the page to the list images params
func (o *ListImagesParams) WithPage(page *string) *ListImagesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the list images params
func (o *ListImagesParams) SetPage(page *string) {
	o.Page = page
}

// WithShape adds the shape to the list images params
func (o *ListImagesParams) WithShape(shape *string) *ListImagesParams {
	o.SetShape(shape)
	return o
}

// SetShape adds the shape to the list images params
func (o *ListImagesParams) SetShape(shape *string) {
	o.Shape = shape
}

// WithSortBy adds the sortBy to the list images params
func (o *ListImagesParams) WithSortBy(sortBy *string) *ListImagesParams {
	o.SetSortBy(sortBy)
	return o
}

// SetSortBy adds the sortBy to the list images params
func (o *ListImagesParams) SetSortBy(sortBy *string) {
	o.SortBy = sortBy
}

// WithSortOrder adds the sortOrder to the list images params
func (o *ListImagesParams) WithSortOrder(sortOrder *string) *ListImagesParams {
	o.SetSortOrder(sortOrder)
	return o
}

// SetSortOrder adds the sortOrder to the list images params
func (o *ListImagesParams) SetSortOrder(sortOrder *string) {
	o.SortOrder = sortOrder
}

// WriteToRequest writes these params to a swagger request
func (o *ListImagesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param compartmentId
	qrCompartmentID := o.CompartmentID
	qCompartmentID := qrCompartmentID
	if qCompartmentID != "" {
		if err := r.SetQueryParam("compartmentId", qCompartmentID); err != nil {
			return err
		}
	}

	if o.DisplayName != nil {

		// query param displayName
		var qrDisplayName string
		if o.DisplayName != nil {
			qrDisplayName = *o.DisplayName
		}
		qDisplayName := qrDisplayName
		if qDisplayName != "" {
			if err := r.SetQueryParam("displayName", qDisplayName); err != nil {
				return err
			}
		}

	}

	if o.LifecycleState != nil {

		// query param lifecycleState
		var qrLifecycleState string
		if o.LifecycleState != nil {
			qrLifecycleState = *o.LifecycleState
		}
		qLifecycleState := qrLifecycleState
		if qLifecycleState != "" {
			if err := r.SetQueryParam("lifecycleState", qLifecycleState); err != nil {
				return err
			}
		}

	}

	if o.Limit != nil {

		// query param limit
		var qrLimit int64
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt64(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	if o.OperatingSystem != nil {

		// query param operatingSystem
		var qrOperatingSystem string
		if o.OperatingSystem != nil {
			qrOperatingSystem = *o.OperatingSystem
		}
		qOperatingSystem := qrOperatingSystem
		if qOperatingSystem != "" {
			if err := r.SetQueryParam("operatingSystem", qOperatingSystem); err != nil {
				return err
			}
		}

	}

	if o.OperatingSystemVersion != nil {

		// query param operatingSystemVersion
		var qrOperatingSystemVersion string
		if o.OperatingSystemVersion != nil {
			qrOperatingSystemVersion = *o.OperatingSystemVersion
		}
		qOperatingSystemVersion := qrOperatingSystemVersion
		if qOperatingSystemVersion != "" {
			if err := r.SetQueryParam("operatingSystemVersion", qOperatingSystemVersion); err != nil {
				return err
			}
		}

	}

	if o.Page != nil {

		// query param page
		var qrPage string
		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := qrPage
		if qPage != "" {
			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}

	}

	if o.Shape != nil {

		// query param shape
		var qrShape string
		if o.Shape != nil {
			qrShape = *o.Shape
		}
		qShape := qrShape
		if qShape != "" {
			if err := r.SetQueryParam("shape", qShape); err != nil {
				return err
			}
		}

	}

	if o.SortBy != nil {

		// query param sortBy
		var qrSortBy string
		if o.SortBy != nil {
			qrSortBy = *o.SortBy
		}
		qSortBy := qrSortBy
		if qSortBy != "" {
			if err := r.SetQueryParam("sortBy", qSortBy); err != nil {
				return err
			}
		}

	}

	if o.SortOrder != nil {

		// query param sortOrder
		var qrSortOrder string
		if o.SortOrder != nil {
			qrSortOrder = *o.SortOrder
		}
		qSortOrder := qrSortOrder
		if qSortOrder != "" {
			if err := r.SetQueryParam("sortOrder", qSortOrder); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
