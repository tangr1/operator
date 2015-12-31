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

type GetStarsIDReader struct {
	formats strfmt.Registry
}

func (o *GetStarsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetStarsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetStarsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetStarsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetStarsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetStarsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetStarsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetStarsIDOK creates a GetStarsIDOK with default headers values
func NewGetStarsIDOK() *GetStarsIDOK {
	return &GetStarsIDOK{}
}

/*GetStarsIDOK

成功查询明星
*/
type GetStarsIDOK struct {
	Payload *models.Star
}

func (o *GetStarsIDOK) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdOK  %+v", 200, o.Payload)
}

func (o *GetStarsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Star)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsIDBadRequest creates a GetStarsIDBadRequest with default headers values
func NewGetStarsIDBadRequest() *GetStarsIDBadRequest {
	return &GetStarsIDBadRequest{}
}

/*GetStarsIDBadRequest

无效请求
*/
type GetStarsIDBadRequest struct {
	Payload *models.Error
}

func (o *GetStarsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetStarsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsIDUnauthorized creates a GetStarsIDUnauthorized with default headers values
func NewGetStarsIDUnauthorized() *GetStarsIDUnauthorized {
	return &GetStarsIDUnauthorized{}
}

/*GetStarsIDUnauthorized

未认证
*/
type GetStarsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetStarsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetStarsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsIDForbidden creates a GetStarsIDForbidden with default headers values
func NewGetStarsIDForbidden() *GetStarsIDForbidden {
	return &GetStarsIDForbidden{}
}

/*GetStarsIDForbidden

无访问权限
*/
type GetStarsIDForbidden struct {
	Payload *models.Error
}

func (o *GetStarsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetStarsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsIDNotFound creates a GetStarsIDNotFound with default headers values
func NewGetStarsIDNotFound() *GetStarsIDNotFound {
	return &GetStarsIDNotFound{}
}

/*GetStarsIDNotFound

没找到指定对象
*/
type GetStarsIDNotFound struct {
	Payload *models.Error
}

func (o *GetStarsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetStarsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetStarsIDInternalServerError creates a GetStarsIDInternalServerError with default headers values
func NewGetStarsIDInternalServerError() *GetStarsIDInternalServerError {
	return &GetStarsIDInternalServerError{}
}

/*GetStarsIDInternalServerError

服务端内部错误
*/
type GetStarsIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetStarsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /stars/{id}/][%d] getStarsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetStarsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
