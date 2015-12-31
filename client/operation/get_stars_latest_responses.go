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

type GetStarsLatestReader struct {
	formats strfmt.Registry
}

func (o *GetStarsLatestReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStarsLatestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetStarsLatestBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetStarsLatestUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetStarsLatestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetStarsLatestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetStarsLatestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarsLatestOK creates a GetStarsLatestOK with default headers values
func NewGetStarsLatestOK() *GetStarsLatestOK {
	return &GetStarsLatestOK{}
}

/*GetStarsLatestOK

成功查询明星
*/
type GetStarsLatestOK struct {
	Payload *models.Star
}

func (o *GetStarsLatestOK) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestOK  %+v", 200, o.Payload)
}

func (o *GetStarsLatestOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Star)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsLatestBadRequest creates a GetStarsLatestBadRequest with default headers values
func NewGetStarsLatestBadRequest() *GetStarsLatestBadRequest {
	return &GetStarsLatestBadRequest{}
}

/*GetStarsLatestBadRequest

无效请求
*/
type GetStarsLatestBadRequest struct {
	Payload *models.Error
}

func (o *GetStarsLatestBadRequest) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestBadRequest  %+v", 400, o.Payload)
}

func (o *GetStarsLatestBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsLatestUnauthorized creates a GetStarsLatestUnauthorized with default headers values
func NewGetStarsLatestUnauthorized() *GetStarsLatestUnauthorized {
	return &GetStarsLatestUnauthorized{}
}

/*GetStarsLatestUnauthorized

未认证
*/
type GetStarsLatestUnauthorized struct {
	Payload *models.Error
}

func (o *GetStarsLatestUnauthorized) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStarsLatestUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsLatestForbidden creates a GetStarsLatestForbidden with default headers values
func NewGetStarsLatestForbidden() *GetStarsLatestForbidden {
	return &GetStarsLatestForbidden{}
}

/*GetStarsLatestForbidden

无访问权限
*/
type GetStarsLatestForbidden struct {
	Payload *models.Error
}

func (o *GetStarsLatestForbidden) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestForbidden  %+v", 403, o.Payload)
}

func (o *GetStarsLatestForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsLatestNotFound creates a GetStarsLatestNotFound with default headers values
func NewGetStarsLatestNotFound() *GetStarsLatestNotFound {
	return &GetStarsLatestNotFound{}
}

/*GetStarsLatestNotFound

没找到指定对象
*/
type GetStarsLatestNotFound struct {
	Payload *models.Error
}

func (o *GetStarsLatestNotFound) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestNotFound  %+v", 404, o.Payload)
}

func (o *GetStarsLatestNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsLatestInternalServerError creates a GetStarsLatestInternalServerError with default headers values
func NewGetStarsLatestInternalServerError() *GetStarsLatestInternalServerError {
	return &GetStarsLatestInternalServerError{}
}

/*GetStarsLatestInternalServerError

服务端内部错误
*/
type GetStarsLatestInternalServerError struct {
	Payload *models.Error
}

func (o *GetStarsLatestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stars/latest][%d] getStarsLatestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStarsLatestInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}