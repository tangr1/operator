package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

// NewPostTopicsParams creates a new PostTopicsParams object
// with the default values initialized.
func NewPostTopicsParams() *PostTopicsParams {
	return &PostTopicsParams{}
}

/*PostTopicsParams contains all the parameters to send to the API endpoint
for the post topics operation typically these are written to a http.Request
*/
type PostTopicsParams struct {

	/*Body*/
	Body *models.CreateTopicRequest
}

// WithBody adds the body to the post topics params
func (o *PostTopicsParams) WithBody(body *models.CreateTopicRequest) *PostTopicsParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostTopicsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.CreateTopicRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
