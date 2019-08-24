// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new images API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for images API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddImage submits a new image for analysis by the engine

Creates a new analysis task that is executed asynchronously
*/
func (a *Client) AddImage(params *AddImageParams, authInfo runtime.ClientAuthInfoWriter) (*AddImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "add_image",
		Method:             "POST",
		PathPattern:        "/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AddImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*AddImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for add_image: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteImage deletes an image analysis
*/
func (a *Client) DeleteImage(params *DeleteImageParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_image",
		Method:             "DELETE",
		PathPattern:        "/images/{imageDigest}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_image: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteImageByImageID deletes image by docker image Id
*/
func (a *Client) DeleteImageByImageID(params *DeleteImageByImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteImageByImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteImageByImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delete_image_by_imageId",
		Method:             "DELETE",
		PathPattern:        "/images/by_id/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteImageByImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteImageByImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_image_by_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImage gets image metadata
*/
func (a *Client) GetImage(params *GetImageParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageByImageID lookups image by docker image Id
*/
func (a *Client) GetImageByImageID(params *GetImageByImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageByImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageByImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_by_imageId",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageByImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageByImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_by_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByType gets the content of an image by type
*/
func (a *Client) GetImageContentByType(params *GetImageContentByTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/content/{ctype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByTypeFiles gets the content of an image by type files
*/
func (a *Client) GetImageContentByTypeFiles(params *GetImageContentByTypeFilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type_files",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/content/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type_files: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByTypeImageID gets the content of an image by type
*/
func (a *Client) GetImageContentByTypeImageID(params *GetImageContentByTypeImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type_imageId",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/content/{ctype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByTypeImageIDFiles gets the content of an image by type files
*/
func (a *Client) GetImageContentByTypeImageIDFiles(params *GetImageContentByTypeImageIDFilesParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeImageIDFilesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeImageIDFilesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type_imageId_files",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/content/files",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeImageIDFilesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeImageIDFilesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type_imageId_files: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByTypeImageIDJavapackage gets the content of an image by type java
*/
func (a *Client) GetImageContentByTypeImageIDJavapackage(params *GetImageContentByTypeImageIDJavapackageParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeImageIDJavapackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeImageIDJavapackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type_imageId_javapackage",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/content/java",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeImageIDJavapackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeImageIDJavapackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type_imageId_javapackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageContentByTypeJavapackage gets the content of an image by type java
*/
func (a *Client) GetImageContentByTypeJavapackage(params *GetImageContentByTypeJavapackageParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageContentByTypeJavapackageOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageContentByTypeJavapackageParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_content_by_type_javapackage",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/content/java",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageContentByTypeJavapackageReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageContentByTypeJavapackageOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_content_by_type_javapackage: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageMetadataByType gets the metadata of an image by type
*/
func (a *Client) GetImageMetadataByType(params *GetImageMetadataByTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageMetadataByTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageMetadataByTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_metadata_by_type",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/metadata/{mtype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageMetadataByTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageMetadataByTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_metadata_by_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImagePolicyCheck checks policy evaluation status for image

Get the policy evaluation for the given image
*/
func (a *Client) GetImagePolicyCheck(params *GetImagePolicyCheckParams, authInfo runtime.ClientAuthInfoWriter) (*GetImagePolicyCheckOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagePolicyCheckParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_policy_check",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImagePolicyCheckReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImagePolicyCheckOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_policy_check: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImagePolicyCheckByImageID checks policy evaluation status for image

Get the policy evaluation for the given image
*/
func (a *Client) GetImagePolicyCheckByImageID(params *GetImagePolicyCheckByImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImagePolicyCheckByImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImagePolicyCheckByImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_policy_check_by_imageId",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/check",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImagePolicyCheckByImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImagePolicyCheckByImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_policy_check_by_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageVulnerabilitiesByType gets vulnerabilities by type
*/
func (a *Client) GetImageVulnerabilitiesByType(params *GetImageVulnerabilitiesByTypeParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageVulnerabilitiesByTypeOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageVulnerabilitiesByTypeParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_vulnerabilities_by_type",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/vuln/{vtype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageVulnerabilitiesByTypeReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageVulnerabilitiesByTypeOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_vulnerabilities_by_type: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageVulnerabilitiesByTypeImageID gets vulnerabilities by type
*/
func (a *Client) GetImageVulnerabilitiesByTypeImageID(params *GetImageVulnerabilitiesByTypeImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageVulnerabilitiesByTypeImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageVulnerabilitiesByTypeImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_vulnerabilities_by_type_imageId",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/vuln/{vtype}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageVulnerabilitiesByTypeImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageVulnerabilitiesByTypeImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_vulnerabilities_by_type_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageVulnerabilityTypes gets vulnerability types
*/
func (a *Client) GetImageVulnerabilityTypes(params *GetImageVulnerabilityTypesParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageVulnerabilityTypesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageVulnerabilityTypesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_vulnerability_types",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/vuln",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageVulnerabilityTypesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageVulnerabilityTypesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_vulnerability_types: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetImageVulnerabilityTypesByImageID gets vulnerability types
*/
func (a *Client) GetImageVulnerabilityTypesByImageID(params *GetImageVulnerabilityTypesByImageIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetImageVulnerabilityTypesByImageIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetImageVulnerabilityTypesByImageIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "get_image_vulnerability_types_by_imageId",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/vuln",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetImageVulnerabilityTypesByImageIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetImageVulnerabilityTypesByImageIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_image_vulnerability_types_by_imageId: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListImageContent lists image content types
*/
func (a *Client) ListImageContent(params *ListImageContentParams, authInfo runtime.ClientAuthInfoWriter) (*ListImageContentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImageContentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_image_content",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/content",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImageContentReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImageContentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_image_content: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListImageContentByImageid lists image content types
*/
func (a *Client) ListImageContentByImageid(params *ListImageContentByImageidParams, authInfo runtime.ClientAuthInfoWriter) (*ListImageContentByImageidOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImageContentByImageidParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_image_content_by_imageid",
		Method:             "GET",
		PathPattern:        "/images/by_id/{imageId}/content",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImageContentByImageidReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImageContentByImageidOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_image_content_by_imageid: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListImageMetadata lists image metadata types
*/
func (a *Client) ListImageMetadata(params *ListImageMetadataParams, authInfo runtime.ClientAuthInfoWriter) (*ListImageMetadataOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImageMetadataParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_image_metadata",
		Method:             "GET",
		PathPattern:        "/images/{imageDigest}/metadata",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImageMetadataReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImageMetadataOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_image_metadata: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
ListImages lists all visible images

List all images visible to the user
*/
func (a *Client) ListImages(params *ListImagesParams, authInfo runtime.ClientAuthInfoWriter) (*ListImagesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListImagesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "list_images",
		Method:             "GET",
		PathPattern:        "/images",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListImagesReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListImagesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for list_images: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
