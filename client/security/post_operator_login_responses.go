package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type PostOperatorLoginReader struct {
	formats strfmt.Registry
}

func (o *PostOperatorLoginReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostOperatorLoginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostOperatorLoginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostOperatorLoginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostOperatorLoginOK creates a PostOperatorLoginOK with default headers values
func NewPostOperatorLoginOK() *PostOperatorLoginOK {
	return &PostOperatorLoginOK{}
}

/*PostOperatorLoginOK

成功获取token
*/
type PostOperatorLoginOK struct {
	Payload *models.AuthenticationResponse
}

func (o *PostOperatorLoginOK) Error() string {
	return fmt.Sprintf("[POST /operator/login/][%d] postOperatorLoginOK  %+v", 200, o.Payload)
}

func (o *PostOperatorLoginOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostOperatorLoginBadRequest creates a PostOperatorLoginBadRequest with default headers values
func NewPostOperatorLoginBadRequest() *PostOperatorLoginBadRequest {
	return &PostOperatorLoginBadRequest{}
}

/*PostOperatorLoginBadRequest

请求无效
*/
type PostOperatorLoginBadRequest struct {
	Payload *models.Error
}

func (o *PostOperatorLoginBadRequest) Error() string {
	return fmt.Sprintf("[POST /operator/login/][%d] postOperatorLoginBadRequest  %+v", 400, o.Payload)
}

func (o *PostOperatorLoginBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostOperatorLoginUnauthorized creates a PostOperatorLoginUnauthorized with default headers values
func NewPostOperatorLoginUnauthorized() *PostOperatorLoginUnauthorized {
	return &PostOperatorLoginUnauthorized{}
}

/*PostOperatorLoginUnauthorized

用户名或密码有问题
*/
type PostOperatorLoginUnauthorized struct {
	Payload *models.Error
}

func (o *PostOperatorLoginUnauthorized) Error() string {
	return fmt.Sprintf("[POST /operator/login/][%d] postOperatorLoginUnauthorized  %+v", 401, o.Payload)
}

func (o *PostOperatorLoginUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
