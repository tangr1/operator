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

type PostStartupsReader struct {
	formats strfmt.Registry
}

func (o *PostStartupsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostStartupsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostStartupsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostStartupsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostStartupsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostStartupsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostStartupsCreated creates a PostStartupsCreated with default headers values
func NewPostStartupsCreated() *PostStartupsCreated {
	return &PostStartupsCreated{}
}

/*PostStartupsCreated

成功创建创业公司
*/
type PostStartupsCreated struct {
	Payload *models.Startup
}

func (o *PostStartupsCreated) Error() string {
	return fmt.Sprintf("[POST /startups/][%d] postStartupsCreated  %+v", 201, o.Payload)
}

func (o *PostStartupsCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Startup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostStartupsBadRequest creates a PostStartupsBadRequest with default headers values
func NewPostStartupsBadRequest() *PostStartupsBadRequest {
	return &PostStartupsBadRequest{}
}

/*PostStartupsBadRequest

无效请求
*/
type PostStartupsBadRequest struct {
	Payload *models.Error
}

func (o *PostStartupsBadRequest) Error() string {
	return fmt.Sprintf("[POST /startups/][%d] postStartupsBadRequest  %+v", 400, o.Payload)
}

func (o *PostStartupsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostStartupsUnauthorized creates a PostStartupsUnauthorized with default headers values
func NewPostStartupsUnauthorized() *PostStartupsUnauthorized {
	return &PostStartupsUnauthorized{}
}

/*PostStartupsUnauthorized

未认证
*/
type PostStartupsUnauthorized struct {
	Payload *models.Error
}

func (o *PostStartupsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /startups/][%d] postStartupsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostStartupsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostStartupsForbidden creates a PostStartupsForbidden with default headers values
func NewPostStartupsForbidden() *PostStartupsForbidden {
	return &PostStartupsForbidden{}
}

/*PostStartupsForbidden

无访问权限
*/
type PostStartupsForbidden struct {
	Payload *models.Error
}

func (o *PostStartupsForbidden) Error() string {
	return fmt.Sprintf("[POST /startups/][%d] postStartupsForbidden  %+v", 403, o.Payload)
}

func (o *PostStartupsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostStartupsInternalServerError creates a PostStartupsInternalServerError with default headers values
func NewPostStartupsInternalServerError() *PostStartupsInternalServerError {
	return &PostStartupsInternalServerError{}
}

/*PostStartupsInternalServerError

服务端内部错误
*/
type PostStartupsInternalServerError struct {
	Payload *models.Error
}

func (o *PostStartupsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /startups/][%d] postStartupsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostStartupsInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
