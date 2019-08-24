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

// ListUserCredentialsReader is a Reader for the ListUserCredentials structure.
type ListUserCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUserCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUserCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListUserCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUserCredentialsOK creates a ListUserCredentialsOK with default headers values
func NewListUserCredentialsOK() *ListUserCredentialsOK {
	return &ListUserCredentialsOK{}
}

/*ListUserCredentialsOK handles this case with default header values.

User credential listing
*/
type ListUserCredentialsOK struct {
	Payload models.CredentialList
}

func (o *ListUserCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountname}/users/{username}/credentials][%d] listUserCredentialsOK  %+v", 200, o.Payload)
}

func (o *ListUserCredentialsOK) GetPayload() models.CredentialList {
	return o.Payload
}

func (o *ListUserCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUserCredentialsInternalServerError creates a ListUserCredentialsInternalServerError with default headers values
func NewListUserCredentialsInternalServerError() *ListUserCredentialsInternalServerError {
	return &ListUserCredentialsInternalServerError{}
}

/*ListUserCredentialsInternalServerError handles this case with default header values.

Internal error
*/
type ListUserCredentialsInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *ListUserCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /accounts/{accountname}/users/{username}/credentials][%d] listUserCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUserCredentialsInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *ListUserCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}