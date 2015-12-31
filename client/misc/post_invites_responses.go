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

type PostInvitesReader struct {
	formats strfmt.Registry
}

func (o *PostInvitesReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPostInvitesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostInvitesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostInvitesUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostInvitesForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostInvitesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostInvitesOK creates a PostInvitesOK with default headers values
func NewPostInvitesOK() *PostInvitesOK {
	return &PostInvitesOK{}
}

/*PostInvitesOK

成功邀请
*/
type PostInvitesOK struct {
	Payload *models.InviteCode
}

func (o *PostInvitesOK) Error() string {
	return fmt.Sprintf("[POST /invites/][%d] postInvitesOK  %+v", 200, o.Payload)
}

func (o *PostInvitesOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InviteCode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostInvitesBadRequest creates a PostInvitesBadRequest with default headers values
func NewPostInvitesBadRequest() *PostInvitesBadRequest {
	return &PostInvitesBadRequest{}
}

/*PostInvitesBadRequest

无效请求
*/
type PostInvitesBadRequest struct {
	Payload *models.Error
}

func (o *PostInvitesBadRequest) Error() string {
	return fmt.Sprintf("[POST /invites/][%d] postInvitesBadRequest  %+v", 400, o.Payload)
}

func (o *PostInvitesBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostInvitesUnauthorized creates a PostInvitesUnauthorized with default headers values
func NewPostInvitesUnauthorized() *PostInvitesUnauthorized {
	return &PostInvitesUnauthorized{}
}

/*PostInvitesUnauthorized

未认证
*/
type PostInvitesUnauthorized struct {
	Payload *models.Error
}

func (o *PostInvitesUnauthorized) Error() string {
	return fmt.Sprintf("[POST /invites/][%d] postInvitesUnauthorized  %+v", 401, o.Payload)
}

func (o *PostInvitesUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostInvitesForbidden creates a PostInvitesForbidden with default headers values
func NewPostInvitesForbidden() *PostInvitesForbidden {
	return &PostInvitesForbidden{}
}

/*PostInvitesForbidden

无访问权限
*/
type PostInvitesForbidden struct {
	Payload *models.Error
}

func (o *PostInvitesForbidden) Error() string {
	return fmt.Sprintf("[POST /invites/][%d] postInvitesForbidden  %+v", 403, o.Payload)
}

func (o *PostInvitesForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostInvitesInternalServerError creates a PostInvitesInternalServerError with default headers values
func NewPostInvitesInternalServerError() *PostInvitesInternalServerError {
	return &PostInvitesInternalServerError{}
}

/*PostInvitesInternalServerError

服务端内部错误
*/
type PostInvitesInternalServerError struct {
	Payload *models.Error
}

func (o *PostInvitesInternalServerError) Error() string {
	return fmt.Sprintf("[POST /invites/][%d] postInvitesInternalServerError  %+v", 500, o.Payload)
}

func (o *PostInvitesInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
