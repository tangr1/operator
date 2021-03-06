package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

// NewPostEmployeesParams creates a new PostEmployeesParams object
// with the default values initialized.
func NewPostEmployeesParams() *PostEmployeesParams {
	return &PostEmployeesParams{}
}

/*PostEmployeesParams contains all the parameters to send to the API endpoint
for the post employees operation typically these are written to a http.Request
*/
type PostEmployeesParams struct {

	/*Body*/
	Body *models.AuthenticationResponse
}

// WithBody adds the body to the post employees params
func (o *PostEmployeesParams) WithBody(body *models.AuthenticationResponse) *PostEmployeesParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostEmployeesParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.AuthenticationResponse)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
