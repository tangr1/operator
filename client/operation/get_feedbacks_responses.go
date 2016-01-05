package operation

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type GetFeedbacksReader struct {
	formats strfmt.Registry
}

func (o *GetFeedbacksReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFeedbacksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFeedbacksOK creates a GetFeedbacksOK with default headers values
func NewGetFeedbacksOK() *GetFeedbacksOK {
	return &GetFeedbacksOK{}
}

/*GetFeedbacksOK

反馈列表
*/
type GetFeedbacksOK struct {
	Payload *models.FeedbackPageableResult
}

func (o *GetFeedbacksOK) Error() string {
	return fmt.Sprintf("[GET /feedbacks/][%d] getFeedbacksOK  %+v", 200, o.Payload)
}

func (o *GetFeedbacksOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FeedbackPageableResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
