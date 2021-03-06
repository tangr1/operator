package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/tangr1/hicto/models"
)

// NewPutRecommendationsIDParams creates a new PutRecommendationsIDParams object
// with the default values initialized.
func NewPutRecommendationsIDParams() *PutRecommendationsIDParams {
	return &PutRecommendationsIDParams{}
}

/*PutRecommendationsIDParams contains all the parameters to send to the API endpoint
for the put recommendations ID operation typically these are written to a http.Request
*/
type PutRecommendationsIDParams struct {

	/*Body*/
	Body *models.UpdateRecommendationRequest
	/*ID
	  推荐ID

	*/
	ID int64
}

// WithBody adds the body to the put recommendations ID params
func (o *PutRecommendationsIDParams) WithBody(body *models.UpdateRecommendationRequest) *PutRecommendationsIDParams {
	o.Body = body
	return o
}

// WithID adds the id to the put recommendations ID params
func (o *PutRecommendationsIDParams) WithID(id int64) *PutRecommendationsIDParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PutRecommendationsIDParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.UpdateRecommendationRequest)
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
