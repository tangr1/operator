package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"
)

// NewGetRecommendationsIDParams creates a new GetRecommendationsIDParams object
// with the default values initialized.
func NewGetRecommendationsIDParams() *GetRecommendationsIDParams {
	return &GetRecommendationsIDParams{}
}

/*GetRecommendationsIDParams contains all the parameters to send to the API endpoint
for the get recommendations ID operation typically these are written to a http.Request
*/
type GetRecommendationsIDParams struct {

	/*ID
	  推荐ID

	*/
	ID int64
}

// WithID adds the id to the get recommendations ID params
func (o *GetRecommendationsIDParams) WithID(id int64) *GetRecommendationsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetRecommendationsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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
