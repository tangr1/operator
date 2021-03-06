package expert

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/strfmt"
)

// New creates a new expert API client.
func New(transport client.Transport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for expert API
*/
type Client struct {
	transport client.Transport
	formats   strfmt.Registry
}

/*删除专家

删除指定专家

*/
func (a *Client) DeleteExpertsID(params *DeleteExpertsIDParams, authInfo client.AuthInfoWriter) (*DeleteExpertsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteExpertsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "DeleteExpertsID",
		Params:   params,
		Reader:   &DeleteExpertsIDReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteExpertsIDOK), nil
}

/*查看专家列表

查看所有专家，支持分页和过滤器

*/
func (a *Client) GetExperts(params *GetExpertsParams, authInfo client.AuthInfoWriter) (*GetExpertsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpertsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "GetExperts",
		Params:   params,
		Reader:   &GetExpertsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExpertsOK), nil
}

/*查询专家

得到指定专家的具体信息

*/
func (a *Client) GetExpertsID(params *GetExpertsIDParams, authInfo client.AuthInfoWriter) (*GetExpertsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetExpertsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "GetExpertsID",
		Params:   params,
		Reader:   &GetExpertsIDReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetExpertsIDOK), nil
}

/*创建专家

根据请求创建专家

*/
func (a *Client) PostExperts(params *PostExpertsParams, authInfo client.AuthInfoWriter) (*PostExpertsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostExpertsParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PostExperts",
		Params:   params,
		Reader:   &PostExpertsReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostExpertsCreated), nil
}

/*更新专家

根据请求中指定要更新的属性对专家进行更新

Example:
1. 更新专家姓名和邮件
    ```json
    {
      "name":"newName",
      "email": "new@email.com"
    }
    ```

*/
func (a *Client) PutExpertsID(params *PutExpertsIDParams, authInfo client.AuthInfoWriter) (*PutExpertsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutExpertsIDParams()
	}

	result, err := a.transport.Submit(&client.Operation{
		ID:       "PutExpertsID",
		Params:   params,
		Reader:   &PutExpertsIDReader{formats: a.formats},
		AuthInfo: authInfo,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PutExpertsIDOK), nil
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport client.Transport) {
	a.transport = transport
}

// NewAPIError creates a new API error
func NewAPIError(opName string, response interface{}, code int) APIError {
	return APIError{
		OperationName: opName,
		Response:      response,
		Code:          code,
	}
}

// APIError wraps an error model and captures the status code
type APIError struct {
	OperationName string
	Response      interface{}
	Code          int
}

func (a APIError) Error() string {
	return fmt.Sprintf("%s (status %d): %+v ", a.OperationName, a.Code, a.Response)
}
