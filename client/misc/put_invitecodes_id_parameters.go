package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/tangr1/hicto/models"
)

// NewPutInvitecodesIDParams creates a new PutInvitecodesIDParams object
// with the default values initialized.
func NewPutInvitecodesIDParams() *PutInvitecodesIDParams {
	return &PutInvitecodesIDParams{}
}

/*PutInvitecodesIDParams contains all the parameters to send to the API endpoint
for the put invitecodes ID operation typically these are written to a http.Request
*/
type PutInvitecodesIDParams struct {

	/*Body*/
	Body *models.UpdateInviteCodeRequest
	/*ID
	  邀请码ID

	*/
	ID int64
}

// WithBody adds the body to the put invitecodes ID params
func (o *PutInvitecodesIDParams) WithBody(body *models.UpdateInviteCodeRequest) *PutInvitecodesIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put invitecodes ID params
func (o *PutInvitecodesIDParams) WithID(id int64) *PutInvitecodesIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutInvitecodesIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateInviteCodeRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
