package misc

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type GetInvitecodesIDReader struct {
	formats strfmt.Registry
}

func (o *GetInvitecodesIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInvitecodesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInvitecodesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInvitecodesIDOK creates a GetInvitecodesIDOK with default headers values
func NewGetInvitecodesIDOK() *GetInvitecodesIDOK {
	return &GetInvitecodesIDOK{}
}

/*GetInvitecodesIDOK

成功查询邀请码
*/
type GetInvitecodesIDOK struct {
	Payload *models.InviteCode
}

func (o *GetInvitecodesIDOK) Error() string {
	return fmt.Sprintf("[GET /invitecodes/{id}/][%d] getInvitecodesIdOK  %+v", 200, o.Payload)
}

func (o *GetInvitecodesIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InviteCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetInvitecodesIDNotFound creates a GetInvitecodesIDNotFound with default headers values
func NewGetInvitecodesIDNotFound() *GetInvitecodesIDNotFound {
	return &GetInvitecodesIDNotFound{}
}

/*GetInvitecodesIDNotFound

没找到指定对象
*/
type GetInvitecodesIDNotFound struct {
	Payload *models.Error
}

func (o *GetInvitecodesIDNotFound) Error() string {
	return fmt.Sprintf("[GET /invitecodes/{id}/][%d] getInvitecodesIdNotFound  %+v", 404, o.Payload)
}

func (o *GetInvitecodesIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
