package startup

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type DeleteStartupsIDReader struct {
	formats strfmt.Registry
}

func (o *DeleteStartupsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteStartupsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteStartupsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteStartupsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteStartupsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteStartupsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteStartupsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteStartupsIDOK creates a DeleteStartupsIDOK with default headers values
func NewDeleteStartupsIDOK() *DeleteStartupsIDOK {
	return &DeleteStartupsIDOK{}
}

/*DeleteStartupsIDOK

成功删除创业公司
*/
type DeleteStartupsIDOK struct {
}

func (o *DeleteStartupsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdOK ", 200)
}

func (o *DeleteStartupsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteStartupsIDBadRequest creates a DeleteStartupsIDBadRequest with default headers values
func NewDeleteStartupsIDBadRequest() *DeleteStartupsIDBadRequest {
	return &DeleteStartupsIDBadRequest{}
}

/*DeleteStartupsIDBadRequest

无效请求
*/
type DeleteStartupsIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteStartupsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteStartupsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteStartupsIDUnauthorized creates a DeleteStartupsIDUnauthorized with default headers values
func NewDeleteStartupsIDUnauthorized() *DeleteStartupsIDUnauthorized {
	return &DeleteStartupsIDUnauthorized{}
}

/*DeleteStartupsIDUnauthorized

未认证
*/
type DeleteStartupsIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteStartupsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteStartupsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteStartupsIDForbidden creates a DeleteStartupsIDForbidden with default headers values
func NewDeleteStartupsIDForbidden() *DeleteStartupsIDForbidden {
	return &DeleteStartupsIDForbidden{}
}

/*DeleteStartupsIDForbidden

无访问权限
*/
type DeleteStartupsIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteStartupsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteStartupsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteStartupsIDNotFound creates a DeleteStartupsIDNotFound with default headers values
func NewDeleteStartupsIDNotFound() *DeleteStartupsIDNotFound {
	return &DeleteStartupsIDNotFound{}
}

/*DeleteStartupsIDNotFound

没找到指定对象
*/
type DeleteStartupsIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteStartupsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteStartupsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteStartupsIDInternalServerError creates a DeleteStartupsIDInternalServerError with default headers values
func NewDeleteStartupsIDInternalServerError() *DeleteStartupsIDInternalServerError {
	return &DeleteStartupsIDInternalServerError{}
}

/*DeleteStartupsIDInternalServerError

服务端内部错误
*/
type DeleteStartupsIDInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteStartupsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /startups/{id}/][%d] deleteStartupsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteStartupsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
