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

type GetRecommendationsIDReader struct {
	formats strfmt.Registry
}

func (o *GetRecommendationsIDReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRecommendationsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetRecommendationsIDBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetRecommendationsIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetRecommendationsIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetRecommendationsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetRecommendationsIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetRecommendationsIDOK creates a GetRecommendationsIDOK with default headers values
func NewGetRecommendationsIDOK() *GetRecommendationsIDOK {
	return &GetRecommendationsIDOK{}
}

/*GetRecommendationsIDOK

成功查询推荐
*/
type GetRecommendationsIDOK struct {
	Payload *models.Recommendation
}

func (o *GetRecommendationsIDOK) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdOK  %+v", 200, o.Payload)
}

func (o *GetRecommendationsIDOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Recommendation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetRecommendationsIDBadRequest creates a GetRecommendationsIDBadRequest with default headers values
func NewGetRecommendationsIDBadRequest() *GetRecommendationsIDBadRequest {
	return &GetRecommendationsIDBadRequest{}
}

/*GetRecommendationsIDBadRequest

无效请求
*/
type GetRecommendationsIDBadRequest struct {
	Payload *models.Error
}

func (o *GetRecommendationsIDBadRequest) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdBadRequest  %+v", 400, o.Payload)
}

func (o *GetRecommendationsIDBadRequest) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetRecommendationsIDUnauthorized creates a GetRecommendationsIDUnauthorized with default headers values
func NewGetRecommendationsIDUnauthorized() *GetRecommendationsIDUnauthorized {
	return &GetRecommendationsIDUnauthorized{}
}

/*GetRecommendationsIDUnauthorized

未认证
*/
type GetRecommendationsIDUnauthorized struct {
	Payload *models.Error
}

func (o *GetRecommendationsIDUnauthorized) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdUnauthorized  %+v", 401, o.Payload)
}

func (o *GetRecommendationsIDUnauthorized) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetRecommendationsIDForbidden creates a GetRecommendationsIDForbidden with default headers values
func NewGetRecommendationsIDForbidden() *GetRecommendationsIDForbidden {
	return &GetRecommendationsIDForbidden{}
}

/*GetRecommendationsIDForbidden

无访问权限
*/
type GetRecommendationsIDForbidden struct {
	Payload *models.Error
}

func (o *GetRecommendationsIDForbidden) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdForbidden  %+v", 403, o.Payload)
}

func (o *GetRecommendationsIDForbidden) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetRecommendationsIDNotFound creates a GetRecommendationsIDNotFound with default headers values
func NewGetRecommendationsIDNotFound() *GetRecommendationsIDNotFound {
	return &GetRecommendationsIDNotFound{}
}

/*GetRecommendationsIDNotFound

没找到指定对象
*/
type GetRecommendationsIDNotFound struct {
	Payload *models.Error
}

func (o *GetRecommendationsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetRecommendationsIDNotFound) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}

// NewGetRecommendationsIDInternalServerError creates a GetRecommendationsIDInternalServerError with default headers values
func NewGetRecommendationsIDInternalServerError() *GetRecommendationsIDInternalServerError {
	return &GetRecommendationsIDInternalServerError{}
}

/*GetRecommendationsIDInternalServerError

服务端内部错误
*/
type GetRecommendationsIDInternalServerError struct {
	Payload *models.Error
}

func (o *GetRecommendationsIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /recommendations/{id}/][%d] getRecommendationsIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetRecommendationsIDInternalServerError) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil {
		return err
	}

	return nil
}
