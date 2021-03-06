package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewDeleteInvitecodesIDParams creates a new DeleteInvitecodesIDParams object
// with the default values initialized.
func NewDeleteInvitecodesIDParams() *DeleteInvitecodesIDParams {
	return &DeleteInvitecodesIDParams{}
}

/*DeleteInvitecodesIDParams contains all the parameters to send to the API endpoint
for the delete invitecodes ID operation typically these are written to a http.Request
*/
type DeleteInvitecodesIDParams struct {

	/*ID
	  邀请码ID

	*/
	ID int64
}

// WithID adds the id to the delete invitecodes ID params
func (o *DeleteInvitecodesIDParams) WithID(id int64) *DeleteInvitecodesIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteInvitecodesIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
