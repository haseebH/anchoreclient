// Code generated by go-swagger; DO NOT EDIT.

package registries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// UpdateRegistryReader is a Reader for the UpdateRegistry structure.
type UpdateRegistryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRegistryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateRegistryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewUpdateRegistryInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateRegistryOK creates a UpdateRegistryOK with default headers values
func NewUpdateRegistryOK() *UpdateRegistryOK {
	return &UpdateRegistryOK{}
}

/*UpdateRegistryOK handles this case with default header values.

Updated registry configuration
*/
type UpdateRegistryOK struct {
	Payload models.RegistryConfigurationList
}

func (o *UpdateRegistryOK) Error() string {
	return fmt.Sprintf("[PUT /registries/{registry}][%d] updateRegistryOK  %+v", 200, o.Payload)
}

func (o *UpdateRegistryOK) GetPayload() models.RegistryConfigurationList {
	return o.Payload
}

func (o *UpdateRegistryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRegistryInternalServerError creates a UpdateRegistryInternalServerError with default headers values
func NewUpdateRegistryInternalServerError() *UpdateRegistryInternalServerError {
	return &UpdateRegistryInternalServerError{}
}

/*UpdateRegistryInternalServerError handles this case with default header values.

Internal Error
*/
type UpdateRegistryInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *UpdateRegistryInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /registries/{registry}][%d] updateRegistryInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateRegistryInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *UpdateRegistryInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}