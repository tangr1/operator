package comment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

// NewPostCommentsParams creates a new PostCommentsParams object
// with the default values initialized.
func NewPostCommentsParams() *PostCommentsParams {
	return &PostCommentsParams{}
}

/*PostCommentsParams contains all the parameters to send to the API endpoint
for the post comments operation typically these are written to a http.Request
*/
type PostCommentsParams struct {

	/*Body*/
	Body *models.CreateCommentRequest
}

// WithBody adds the body to the post comments params
func (o *PostCommentsParams) WithBody(body *models.CreateCommentRequest) *PostCommentsParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostCommentsParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.CreateCommentRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
