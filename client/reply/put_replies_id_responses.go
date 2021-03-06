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

type PutRepliesIDReader struct {
	formats strfmt.Registry
}

func (o *PutRepliesIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutRepliesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPutRepliesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPutRepliesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPutRepliesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewPutRepliesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPutRepliesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutRepliesIDOK creates a PutRepliesIDOK with default headers values
func NewPutRepliesIDOK() *PutRepliesIDOK {
	return &PutRepliesIDOK{}
}

/*PutRepliesIDOK

成功更新回复
*/
type PutRepliesIDOK struct {
	Payload *models.Reply
}

func (o *PutRepliesIDOK) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdOK  %+v", 200, o.Payload)
}

func (o *PutRepliesIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Reply)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutRepliesIDBadRequest creates a PutRepliesIDBadRequest with default headers values
func NewPutRepliesIDBadRequest() *PutRepliesIDBadRequest {
	return &PutRepliesIDBadRequest{}
}

/*PutRepliesIDBadRequest

无效请求
*/
type PutRepliesIDBadRequest struct {
	Payload *models.Error
}

func (o *PutRepliesIDBadRequest) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdBadRequest  %+v", 400, o.Payload)
}

func (o *PutRepliesIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutRepliesIDUnauthorized creates a PutRepliesIDUnauthorized with default headers values
func NewPutRepliesIDUnauthorized() *PutRepliesIDUnauthorized {
	return &PutRepliesIDUnauthorized{}
}

/*PutRepliesIDUnauthorized

未认证
*/
type PutRepliesIDUnauthorized struct {
	Payload *models.Error
}

func (o *PutRepliesIDUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *PutRepliesIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutRepliesIDForbidden creates a PutRepliesIDForbidden with default headers values
func NewPutRepliesIDForbidden() *PutRepliesIDForbidden {
	return &PutRepliesIDForbidden{}
}

/*PutRepliesIDForbidden

无访问权限
*/
type PutRepliesIDForbidden struct {
	Payload *models.Error
}

func (o *PutRepliesIDForbidden) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdForbidden  %+v", 403, o.Payload)
}

func (o *PutRepliesIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutRepliesIDNotFound creates a PutRepliesIDNotFound with default headers values
func NewPutRepliesIDNotFound() *PutRepliesIDNotFound {
	return &PutRepliesIDNotFound{}
}

/*PutRepliesIDNotFound

没找到指定对象
*/
type PutRepliesIDNotFound struct {
	Payload *models.Error
}

func (o *PutRepliesIDNotFound) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdNotFound  %+v", 404, o.Payload)
}

func (o *PutRepliesIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPutRepliesIDInternalServerError creates a PutRepliesIDInternalServerError with default headers values
func NewPutRepliesIDInternalServerError() *PutRepliesIDInternalServerError {
	return &PutRepliesIDInternalServerError{}
}

/*PutRepliesIDInternalServerError

服务端内部错误
*/
type PutRepliesIDInternalServerError struct {
	Payload *models.Error
}

func (o *PutRepliesIDInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /replies/{id}/][%d] putRepliesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *PutRepliesIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
