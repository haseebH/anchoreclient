// Code generated by go-swagger; DO NOT EDIT.

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// GetUsersAccountReader is a Reader for the GetUsersAccount structure.
type GetUsersAccountReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersAccountReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUsersAccountOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetUsersAccountInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUsersAccountOK creates a GetUsersAccountOK with default headers values
func NewGetUsersAccountOK() *GetUsersAccountOK {
	return &GetUsersAccountOK{}
}

/*GetUsersAccountOK handles this case with default header values.

User details for caller's user
*/
type GetUsersAccountOK struct {
	Payload *models.Account
}

func (o *GetUsersAccountOK) Error() string {
	return fmt.Sprintf("[GET /account][%d] getUsersAccountOK  %+v", 200, o.Payload)
}

func (o *GetUsersAccountOK) GetPayload() *models.Account {
	return o.Payload
}

func (o *GetUsersAccountOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Account)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersAccountInternalServerError creates a GetUsersAccountInternalServerError with default headers values
func NewGetUsersAccountInternalServerError() *GetUsersAccountInternalServerError {
	return &GetUsersAccountInternalServerError{}
}

/*GetUsersAccountInternalServerError handles this case with default header values.

Internal error
*/
type GetUsersAccountInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetUsersAccountInternalServerError) Error() string {
	return fmt.Sprintf("[GET /account][%d] getUsersAccountInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUsersAccountInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetUsersAccountInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}