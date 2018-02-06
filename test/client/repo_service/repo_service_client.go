// Code generated by go-swagger; DO NOT EDIT.

package repo_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new repo service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for repo service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateRepo create repo API
*/
func (a *Client) CreateRepo(params *CreateRepoParams) (*CreateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "CreateRepo",
		Method:             "POST",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRepoOK), nil

}

/*
DeleteRepo delete repo API
*/
func (a *Client) DeleteRepo(params *DeleteRepoParams) (*DeleteRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "DeleteRepo",
		Method:             "DELETE",
		PathPattern:        "/v1/repos/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRepoOK), nil

}

/*
GetRepo get repo API
*/
func (a *Client) GetRepo(params *GetRepoParams) (*GetRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRepo",
		Method:             "GET",
		PathPattern:        "/v1/repos/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoOK), nil

}

/*
GetRepoList get repo list API
*/
func (a *Client) GetRepoList(params *GetRepoListParams) (*GetRepoListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRepoListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "GetRepoList",
		Method:             "GET",
		PathPattern:        "/v1/repos",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetRepoListReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRepoListOK), nil

}

/*
UpdateRepo update repo API
*/
func (a *Client) UpdateRepo(params *UpdateRepoParams) (*UpdateRepoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateRepoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "UpdateRepo",
		Method:             "POST",
		PathPattern:        "/v1/repos/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateRepoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*UpdateRepoOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}