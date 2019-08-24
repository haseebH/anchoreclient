// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// GetImagePolicyCheckByImageIDReader is a Reader for the GetImagePolicyCheckByImageID structure.
type GetImagePolicyCheckByImageIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImagePolicyCheckByImageIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImagePolicyCheckByImageIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetImagePolicyCheckByImageIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImagePolicyCheckByImageIDOK creates a GetImagePolicyCheckByImageIDOK with default headers values
func NewGetImagePolicyCheckByImageIDOK() *GetImagePolicyCheckByImageIDOK {
	return &GetImagePolicyCheckByImageIDOK{}
}

/*GetImagePolicyCheckByImageIDOK handles this case with default header values.

Policy evaluation success
*/
type GetImagePolicyCheckByImageIDOK struct {
	Payload models.PolicyEvaluationList
}

func (o *GetImagePolicyCheckByImageIDOK) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/check][%d] getImagePolicyCheckByImageIdOK  %+v", 200, o.Payload)
}

func (o *GetImagePolicyCheckByImageIDOK) GetPayload() models.PolicyEvaluationList {
	return o.Payload
}

func (o *GetImagePolicyCheckByImageIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImagePolicyCheckByImageIDInternalServerError creates a GetImagePolicyCheckByImageIDInternalServerError with default headers values
func NewGetImagePolicyCheckByImageIDInternalServerError() *GetImagePolicyCheckByImageIDInternalServerError {
	return &GetImagePolicyCheckByImageIDInternalServerError{}
}

/*GetImagePolicyCheckByImageIDInternalServerError handles this case with default header values.

Internal Error
*/
type GetImagePolicyCheckByImageIDInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetImagePolicyCheckByImageIDInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/check][%d] getImagePolicyCheckByImageIdInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImagePolicyCheckByImageIDInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetImagePolicyCheckByImageIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
