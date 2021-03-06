package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetFeedbacksParams creates a new GetFeedbacksParams object
// with the default values initialized.
func NewGetFeedbacksParams() *GetFeedbacksParams {
	return &GetFeedbacksParams{}
}

/*GetFeedbacksParams contains all the parameters to send to the API endpoint
for the get feedbacks operation typically these are written to a http.Request
*/
type GetFeedbacksParams struct {

	/*Page
	  当前页码

	*/
	Page int64
	/*Pagesize
	  每页项数

	*/
	Pagesize int64
}

// WithPage adds the page to the get feedbacks params
func (o *GetFeedbacksParams) WithPage(page int64) *GetFeedbacksParams {
	o.Page = page
	return o
}

// WithPagesize adds the pagesize to the get feedbacks params
func (o *GetFeedbacksParams) WithPagesize(pagesize int64) *GetFeedbacksParams {
	o.Pagesize = pagesize
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetFeedbacksParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// query param page
	qrPage := o.Page
	qPage := swag.FormatInt64(qrPage)
	if qPage != "" && qPage != "-1" {
		if err := r.SetQueryParam("page", qPage); err != nil {
			return err
		}
	}

	// query array param page

	// query param pagesize
	qrPagesize := o.Pagesize
	qPagesize := swag.FormatInt64(qrPagesize)
	if qPagesize != "" && qPagesize != "-1" {
		if err := r.SetQueryParam("pagesize", qPagesize); err != nil {
			return err
		}
	}

	// query array param pagesize

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
