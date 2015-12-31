package startup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetStartupsParams creates a new GetStartupsParams object
// with the default values initialized.
func NewGetStartupsParams() *GetStartupsParams {
	return &GetStartupsParams{}
}

/*GetStartupsParams contains all the parameters to send to the API endpoint
for the get startups operation typically these are written to a http.Request
*/
type GetStartupsParams struct {

	/*Page
	  当前页码

	*/
	Page int64
	/*Pagesize
	  每页项数

	*/
	Pagesize int64
}

// WithPage adds the page to the get startups params
func (o *GetStartupsParams) WithPage(page int64) *GetStartupsParams {
	o.Page = page
	return o
}

// WithPagesize adds the pagesize to the get startups params
func (o *GetStartupsParams) WithPagesize(pagesize int64) *GetStartupsParams {
	o.Pagesize = pagesize
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetStartupsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param page
	if err := r.SetPathParam("page", swag.FormatInt64(o.Page)); err != nil {
		return err
	}

	// path param pagesize
	if err := r.SetPathParam("pagesize", swag.FormatInt64(o.Pagesize)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
