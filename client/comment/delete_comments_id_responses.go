package comment

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type DeleteCommentsIDReader struct {
	formats strfmt.Registry
}

func (o *DeleteCommentsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteCommentsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteCommentsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteCommentsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteCommentsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteCommentsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteCommentsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCommentsIDOK creates a DeleteCommentsIDOK with default headers values
func NewDeleteCommentsIDOK() *DeleteCommentsIDOK {
	return &DeleteCommentsIDOK{}
}

/*DeleteCommentsIDOK

成功删除评论
*/
type DeleteCommentsIDOK struct {
	Payload *models.Comment
}

func (o *DeleteCommentsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdOK  %+v", 200, o.Payload)
}

func (o *DeleteCommentsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Comment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteCommentsIDBadRequest creates a DeleteCommentsIDBadRequest with default headers values
func NewDeleteCommentsIDBadRequest() *DeleteCommentsIDBadRequest {
	return &DeleteCommentsIDBadRequest{}
}

/*DeleteCommentsIDBadRequest

无效请求
*/
type DeleteCommentsIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteCommentsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCommentsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteCommentsIDUnauthorized creates a DeleteCommentsIDUnauthorized with default headers values
func NewDeleteCommentsIDUnauthorized() *DeleteCommentsIDUnauthorized {
	return &DeleteCommentsIDUnauthorized{}
}

/*DeleteCommentsIDUnauthorized

未认证
*/
type DeleteCommentsIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteCommentsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCommentsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteCommentsIDForbidden creates a DeleteCommentsIDForbidden with default headers values
func NewDeleteCommentsIDForbidden() *DeleteCommentsIDForbidden {
	return &DeleteCommentsIDForbidden{}
}

/*DeleteCommentsIDForbidden

无访问权限
*/
type DeleteCommentsIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteCommentsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCommentsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteCommentsIDNotFound creates a DeleteCommentsIDNotFound with default headers values
func NewDeleteCommentsIDNotFound() *DeleteCommentsIDNotFound {
	return &DeleteCommentsIDNotFound{}
}

/*DeleteCommentsIDNotFound

没找到指定对象
*/
type DeleteCommentsIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteCommentsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCommentsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteCommentsIDInternalServerError creates a DeleteCommentsIDInternalServerError with default headers values
func NewDeleteCommentsIDInternalServerError() *DeleteCommentsIDInternalServerError {
	return &DeleteCommentsIDInternalServerError{}
}

/*DeleteCommentsIDInternalServerError

服务端内部错误
*/
type DeleteCommentsIDInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteCommentsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /comments/{id}/][%d] deleteCommentsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCommentsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
