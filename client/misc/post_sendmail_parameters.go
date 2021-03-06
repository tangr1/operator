package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

// NewPostSendmailParams creates a new PostSendmailParams object
// with the default values initialized.
func NewPostSendmailParams() *PostSendmailParams {
	return &PostSendmailParams{}
}

/*PostSendmailParams contains all the parameters to send to the API endpoint
for the post sendmail operation typically these are written to a http.Request
*/
type PostSendmailParams struct {

	/*Body*/
	Body *models.SendMailRequest
}

// WithBody adds the body to the post sendmail params
func (o *PostSendmailParams) WithBody(body *models.SendMailRequest) *PostSendmailParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostSendmailParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.SendMailRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
