// Code generated by go-swagger; DO NOT EDIT.

package user_management

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// CreateUserCredentialReader is a Reader for the CreateUserCredential structure.
type CreateUserCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateUserCredentialReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCreateUserCredentialOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewCreateUserCredentialInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateUserCredentialOK creates a CreateUserCredentialOK with default headers values
func NewCreateUserCredentialOK() *CreateUserCredentialOK {
	return &CreateUserCredentialOK{}
}

/*CreateUserCredentialOK handles this case with default header values.

Add a credential, overwritting if already exists
*/
type CreateUserCredentialOK struct {
	Payload *models.User
}

func (o *CreateUserCredentialOK) Error() string {
	return fmt.Sprintf("[POST /accounts/{accountname}/users/{username}/credentials][%d] createUserCredentialOK  %+v", 200, o.Payload)
}

func (o *CreateUserCredentialOK) GetPayload() *models.User {
	return o.Payload
}

func (o *CreateUserCredentialOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.User)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateUserCredentialInternalServerError creates a CreateUserCredentialInternalServerError with default headers values
func NewCreateUserCredentialInternalServerError() *CreateUserCredentialInternalServerError {
	return &CreateUserCredentialInternalServerError{}
}

/*CreateUserCredentialInternalServerError handles this case with default header values.

Internal error
*/
type CreateUserCredentialInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *CreateUserCredentialInternalServerError) Error() string {
	return fmt.Sprintf("[POST /accounts/{accountname}/users/{username}/credentials][%d] createUserCredentialInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateUserCredentialInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *CreateUserCredentialInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}