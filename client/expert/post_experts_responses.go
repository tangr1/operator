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

type PostExpertsReader struct {
	formats strfmt.Registry
}

func (o *PostExpertsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewPostExpertsCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewPostExpertsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewPostExpertsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewPostExpertsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewPostExpertsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostExpertsCreated creates a PostExpertsCreated with default headers values
func NewPostExpertsCreated() *PostExpertsCreated {
	return &PostExpertsCreated{}
}

/*PostExpertsCreated

成功创建专家
*/
type PostExpertsCreated struct {
	Payload *models.AuthenticationResponse
}

func (o *PostExpertsCreated) Error() string {
	return fmt.Sprintf("[POST /experts/][%d] postExpertsCreated  %+v", 201, o.Payload)
}

func (o *PostExpertsCreated) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AuthenticationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostExpertsBadRequest creates a PostExpertsBadRequest with default headers values
func NewPostExpertsBadRequest() *PostExpertsBadRequest {
	return &PostExpertsBadRequest{}
}

/*PostExpertsBadRequest

无效请求
*/
type PostExpertsBadRequest struct {
	Payload *models.Error
}

func (o *PostExpertsBadRequest) Error() string {
	return fmt.Sprintf("[POST /experts/][%d] postExpertsBadRequest  %+v", 400, o.Payload)
}

func (o *PostExpertsBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostExpertsUnauthorized creates a PostExpertsUnauthorized with default headers values
func NewPostExpertsUnauthorized() *PostExpertsUnauthorized {
	return &PostExpertsUnauthorized{}
}

/*PostExpertsUnauthorized

未认证
*/
type PostExpertsUnauthorized struct {
	Payload *models.Error
}

func (o *PostExpertsUnauthorized) Error() string {
	return fmt.Sprintf("[POST /experts/][%d] postExpertsUnauthorized  %+v", 401, o.Payload)
}

func (o *PostExpertsUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostExpertsForbidden creates a PostExpertsForbidden with default headers values
func NewPostExpertsForbidden() *PostExpertsForbidden {
	return &PostExpertsForbidden{}
}

/*PostExpertsForbidden

无访问权限
*/
type PostExpertsForbidden struct {
	Payload *models.Error
}

func (o *PostExpertsForbidden) Error() string {
	return fmt.Sprintf("[POST /experts/][%d] postExpertsForbidden  %+v", 403, o.Payload)
}

func (o *PostExpertsForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewPostExpertsInternalServerError creates a PostExpertsInternalServerError with default headers values
func NewPostExpertsInternalServerError() *PostExpertsInternalServerError {
	return &PostExpertsInternalServerError{}
}

/*PostExpertsInternalServerError

服务端内部错误
*/
type PostExpertsInternalServerError struct {
	Payload *models.Error
}

func (o *PostExpertsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /experts/][%d] postExpertsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostExpertsInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}