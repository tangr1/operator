package topic

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type DeleteTopicsIDReader struct {
	formats strfmt.Registry
}

func (o *DeleteTopicsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteTopicsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteTopicsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteTopicsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteTopicsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteTopicsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteTopicsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteTopicsIDOK creates a DeleteTopicsIDOK with default headers values
func NewDeleteTopicsIDOK() *DeleteTopicsIDOK {
	return &DeleteTopicsIDOK{}
}

/*DeleteTopicsIDOK

成功删除主题
*/
type DeleteTopicsIDOK struct {
}

func (o *DeleteTopicsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdOK ", 200)
}

func (o *DeleteTopicsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTopicsIDBadRequest creates a DeleteTopicsIDBadRequest with default headers values
func NewDeleteTopicsIDBadRequest() *DeleteTopicsIDBadRequest {
	return &DeleteTopicsIDBadRequest{}
}

/*DeleteTopicsIDBadRequest

无效请求
*/
type DeleteTopicsIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteTopicsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteTopicsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteTopicsIDUnauthorized creates a DeleteTopicsIDUnauthorized with default headers values
func NewDeleteTopicsIDUnauthorized() *DeleteTopicsIDUnauthorized {
	return &DeleteTopicsIDUnauthorized{}
}

/*DeleteTopicsIDUnauthorized

未认证
*/
type DeleteTopicsIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteTopicsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteTopicsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteTopicsIDForbidden creates a DeleteTopicsIDForbidden with default headers values
func NewDeleteTopicsIDForbidden() *DeleteTopicsIDForbidden {
	return &DeleteTopicsIDForbidden{}
}

/*DeleteTopicsIDForbidden

无访问权限
*/
type DeleteTopicsIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteTopicsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteTopicsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteTopicsIDNotFound creates a DeleteTopicsIDNotFound with default headers values
func NewDeleteTopicsIDNotFound() *DeleteTopicsIDNotFound {
	return &DeleteTopicsIDNotFound{}
}

/*DeleteTopicsIDNotFound

没找到指定对象
*/
type DeleteTopicsIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteTopicsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteTopicsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteTopicsIDInternalServerError creates a DeleteTopicsIDInternalServerError with default headers values
func NewDeleteTopicsIDInternalServerError() *DeleteTopicsIDInternalServerError {
	return &DeleteTopicsIDInternalServerError{}
}

/*DeleteTopicsIDInternalServerError

服务端内部错误
*/
type DeleteTopicsIDInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteTopicsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /topics/{id}/][%d] deleteTopicsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteTopicsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
