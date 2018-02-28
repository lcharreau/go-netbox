// Code generated by go-swagger; DO NOT EDIT.

package secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new secrets API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for secrets API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
SecretsChoicesList secrets choices list API
*/
func (a *Client) SecretsChoicesList(params *SecretsChoicesListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsChoicesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsChoicesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets__choices_list",
		Method:             "GET",
		PathPattern:        "/secrets/_choices/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsChoicesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsChoicesListOK), nil

}

/*
SecretsChoicesRead secrets choices read API
*/
func (a *Client) SecretsChoicesRead(params *SecretsChoicesReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsChoicesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsChoicesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets__choices_read",
		Method:             "GET",
		PathPattern:        "/secrets/_choices/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsChoicesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsChoicesReadOK), nil

}

/*
SecretsGenerateRsaKeyPairList This endpoint can be used to generate a new RSA key pair. The keys are returned in PEM format.

    {
        "public_key": "<public key>",
        "private_key": "<private key>"
    }
*/
func (a *Client) SecretsGenerateRsaKeyPairList(params *SecretsGenerateRsaKeyPairListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGenerateRsaKeyPairListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsGenerateRsaKeyPairListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_generate-rsa-key-pair_list",
		Method:             "GET",
		PathPattern:        "/secrets/generate-rsa-key-pair/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsGenerateRsaKeyPairListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsGenerateRsaKeyPairListOK), nil

}

/*
SecretsGetSessionKeyCreate Retrieve a temporary session key to use for encrypting and decrypting secrets via the API. The user's private RSA
key is POSTed with the name `private_key`. An example:

    curl -v -X POST -H "Authorization: Token <token>" -H "Accept: application/json; indent=4" \
    --data-urlencode "private_key@<filename>" https://netbox/api/secrets/get-session-key/

This request will yield a base64-encoded session key to be included in an `X-Session-Key` header in future requests:

    {
        "session_key": "+8t4SI6XikgVmB5+/urhozx9O5qCQANyOk1MNe6taRf="
    }

This endpoint accepts one optional parameter: `preserve_key`. If True and a session key exists, the existing session
key will be returned instead of a new one.
*/
func (a *Client) SecretsGetSessionKeyCreate(params *SecretsGetSessionKeyCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsGetSessionKeyCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsGetSessionKeyCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_get-session-key_create",
		Method:             "POST",
		PathPattern:        "/secrets/get-session-key/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsGetSessionKeyCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsGetSessionKeyCreateCreated), nil

}

/*
SecretsSecretRolesCreate secrets secret roles create API
*/
func (a *Client) SecretsSecretRolesCreate(params *SecretsSecretRolesCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_create",
		Method:             "POST",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesCreateCreated), nil

}

/*
SecretsSecretRolesDelete secrets secret roles delete API
*/
func (a *Client) SecretsSecretRolesDelete(params *SecretsSecretRolesDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesDeleteNoContent), nil

}

/*
SecretsSecretRolesList secrets secret roles list API
*/
func (a *Client) SecretsSecretRolesList(params *SecretsSecretRolesListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_list",
		Method:             "GET",
		PathPattern:        "/secrets/secret-roles/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesListOK), nil

}

/*
SecretsSecretRolesPartialUpdate secrets secret roles partial update API
*/
func (a *Client) SecretsSecretRolesPartialUpdate(params *SecretsSecretRolesPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesPartialUpdateOK), nil

}

/*
SecretsSecretRolesRead secrets secret roles read API
*/
func (a *Client) SecretsSecretRolesRead(params *SecretsSecretRolesReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_read",
		Method:             "GET",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesReadOK), nil

}

/*
SecretsSecretRolesUpdate secrets secret roles update API
*/
func (a *Client) SecretsSecretRolesUpdate(params *SecretsSecretRolesUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretRolesUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretRolesUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secret-roles_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secret-roles/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretRolesUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretRolesUpdateOK), nil

}

/*
SecretsSecretsCreate secrets secrets create API
*/
func (a *Client) SecretsSecretsCreate(params *SecretsSecretsCreateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsCreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_create",
		Method:             "POST",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsCreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsCreateCreated), nil

}

/*
SecretsSecretsDelete secrets secrets delete API
*/
func (a *Client) SecretsSecretsDelete(params *SecretsSecretsDeleteParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsDeleteNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsDeleteParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_delete",
		Method:             "DELETE",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsDeleteReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsDeleteNoContent), nil

}

/*
SecretsSecretsList secrets secrets list API
*/
func (a *Client) SecretsSecretsList(params *SecretsSecretsListParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsListParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_list",
		Method:             "GET",
		PathPattern:        "/secrets/secrets/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsListReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsListOK), nil

}

/*
SecretsSecretsPartialUpdate secrets secrets partial update API
*/
func (a *Client) SecretsSecretsPartialUpdate(params *SecretsSecretsPartialUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsPartialUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsPartialUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_partial_update",
		Method:             "PATCH",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsPartialUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsPartialUpdateOK), nil

}

/*
SecretsSecretsRead secrets secrets read API
*/
func (a *Client) SecretsSecretsRead(params *SecretsSecretsReadParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsReadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsReadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_read",
		Method:             "GET",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsReadReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsReadOK), nil

}

/*
SecretsSecretsUpdate secrets secrets update API
*/
func (a *Client) SecretsSecretsUpdate(params *SecretsSecretsUpdateParams, authInfo runtime.ClientAuthInfoWriter) (*SecretsSecretsUpdateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSecretsSecretsUpdateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "secrets_secrets_update",
		Method:             "PUT",
		PathPattern:        "/secrets/secrets/{id}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SecretsSecretsUpdateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SecretsSecretsUpdateOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
