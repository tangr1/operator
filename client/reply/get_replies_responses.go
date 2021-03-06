package reply

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type GetRepliesReader struct {
	formats strfmt.Registry
}

func (o *GetRepliesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRepliesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRepliesOK creates a GetRepliesOK with default headers values
func NewGetRepliesOK() *GetRepliesOK {
	return &GetRepliesOK{}
}

/*GetRepliesOK

回复列表
*/
type GetRepliesOK struct {
	Payload *models.ReplyPageableResult
}

func (o *GetRepliesOK) Error() string {
	return fmt.Sprintf("[GET /replies/][%d] getRepliesOK  %+v", 200, o.Payload)
}

func (o *GetRepliesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReplyPageableResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
