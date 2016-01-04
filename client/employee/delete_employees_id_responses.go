package employee

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"
	"github.com/go-swagger/go-swagger/strfmt"

	"github.com/tangr1/hicto/models"
)

type DeleteEmployeesIDReader struct {
	formats strfmt.Registry
}

func (o *DeleteEmployeesIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteEmployeesIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewDeleteEmployeesIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewDeleteEmployeesIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewDeleteEmployeesIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteEmployeesIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewDeleteEmployeesIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteEmployeesIDOK creates a DeleteEmployeesIDOK with default headers values
func NewDeleteEmployeesIDOK() *DeleteEmployeesIDOK {
	return &DeleteEmployeesIDOK{}
}

/*DeleteEmployeesIDOK

成功删除创业公司员工
*/
type DeleteEmployeesIDOK struct {
}

func (o *DeleteEmployeesIDOK) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdOK ", 200)
}

func (o *DeleteEmployeesIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteEmployeesIDBadRequest creates a DeleteEmployeesIDBadRequest with default headers values
func NewDeleteEmployeesIDBadRequest() *DeleteEmployeesIDBadRequest {
	return &DeleteEmployeesIDBadRequest{}
}

/*DeleteEmployeesIDBadRequest

无效请求
*/
type DeleteEmployeesIDBadRequest struct {
	Payload *models.Error
}

func (o *DeleteEmployeesIDBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteEmployeesIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteEmployeesIDUnauthorized creates a DeleteEmployeesIDUnauthorized with default headers values
func NewDeleteEmployeesIDUnauthorized() *DeleteEmployeesIDUnauthorized {
	return &DeleteEmployeesIDUnauthorized{}
}

/*DeleteEmployeesIDUnauthorized

未认证
*/
type DeleteEmployeesIDUnauthorized struct {
	Payload *models.Error
}

func (o *DeleteEmployeesIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteEmployeesIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteEmployeesIDForbidden creates a DeleteEmployeesIDForbidden with default headers values
func NewDeleteEmployeesIDForbidden() *DeleteEmployeesIDForbidden {
	return &DeleteEmployeesIDForbidden{}
}

/*DeleteEmployeesIDForbidden

无访问权限
*/
type DeleteEmployeesIDForbidden struct {
	Payload *models.Error
}

func (o *DeleteEmployeesIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdForbidden  %+v", 403, o.Payload)
}

func (o *DeleteEmployeesIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteEmployeesIDNotFound creates a DeleteEmployeesIDNotFound with default headers values
func NewDeleteEmployeesIDNotFound() *DeleteEmployeesIDNotFound {
	return &DeleteEmployeesIDNotFound{}
}

/*DeleteEmployeesIDNotFound

没找到指定对象
*/
type DeleteEmployeesIDNotFound struct {
	Payload *models.Error
}

func (o *DeleteEmployeesIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdNotFound  %+v", 404, o.Payload)
}

func (o *DeleteEmployeesIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewDeleteEmployeesIDInternalServerError creates a DeleteEmployeesIDInternalServerError with default headers values
func NewDeleteEmployeesIDInternalServerError() *DeleteEmployeesIDInternalServerError {
	return &DeleteEmployeesIDInternalServerError{}
}

/*DeleteEmployeesIDInternalServerError

服务端内部错误
*/
type DeleteEmployeesIDInternalServerError struct {
	Payload *models.Error
}

func (o *DeleteEmployeesIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /employees/{id}/][%d] deleteEmployeesIdInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteEmployeesIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
