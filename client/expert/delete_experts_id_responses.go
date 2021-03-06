package expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type DeleteExpertsIDReader struct {
	formats strfmt.Registry
}

func (o *DeleteExpertsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteExpertsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteExpertsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteExpertsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteExpertsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteExpertsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteExpertsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteExpertsIDOK creates a DeleteExpertsIDOK with default headers values
func NewDeleteExpertsIDOK() *DeleteExpertsIDOK {
	return &DeleteExpertsIDOK{}
}

/*DeleteExpertsIDOK

成功删除专家
*/
type DeleteExpertsIDOK struct {
}

func (o *DeleteExpertsIDOK) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdOK ", 200)
}

func (o *DeleteExpertsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteExpertsIDBadRequest creates a DeleteExpertsIDBadRequest with default headers values
func NewDeleteExpertsIDBadRequest() *DeleteExpertsIDBadRequest {
	return &DeleteExpertsIDBadRequest{}
}

/*DeleteExpertsIDBadRequest

无效请求
*/
type DeleteExpertsIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteExpertsIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteExpertsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteExpertsIDUnauthorized creates a DeleteExpertsIDUnauthorized with default headers values
func NewDeleteExpertsIDUnauthorized() *DeleteExpertsIDUnauthorized {
	return &DeleteExpertsIDUnauthorized{}
}

/*DeleteExpertsIDUnauthorized

未认证
*/
type DeleteExpertsIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteExpertsIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteExpertsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteExpertsIDForbidden creates a DeleteExpertsIDForbidden with default headers values
func NewDeleteExpertsIDForbidden() *DeleteExpertsIDForbidden {
	return &DeleteExpertsIDForbidden{}
}

/*DeleteExpertsIDForbidden

无访问权限
*/
type DeleteExpertsIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteExpertsIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteExpertsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteExpertsIDNotFound creates a DeleteExpertsIDNotFound with default headers values
func NewDeleteExpertsIDNotFound() *DeleteExpertsIDNotFound {
	return &DeleteExpertsIDNotFound{}
}

/*DeleteExpertsIDNotFound

没找到指定对象
*/
type DeleteExpertsIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteExpertsIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteExpertsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteExpertsIDInternalServerError creates a DeleteExpertsIDInternalServerError with default headers values
func NewDeleteExpertsIDInternalServerError() *DeleteExpertsIDInternalServerError {
	return &DeleteExpertsIDInternalServerError{}
}

/*DeleteExpertsIDInternalServerError

服务端内部错误
*/
type DeleteExpertsIDInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteExpertsIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /experts/{id}/][%d] deleteExpertsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteExpertsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
