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

// GetCredentialsReader is a Reader for the GetCredentials structure.
type GetCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetCredentialsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetCredentialsOK creates a GetCredentialsOK with default headers values
func NewGetCredentialsOK() *GetCredentialsOK {
	return &GetCredentialsOK{}
}

/*GetCredentialsOK handles this case with default header values.

User credential listing
*/
type GetCredentialsOK struct {
	Payload models.CredentialList
}

func (o *GetCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /user/credentials][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetCredentialsOK) GetPayload() models.CredentialList {
	return o.Payload
}

func (o *GetCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsInternalServerError creates a GetCredentialsInternalServerError with default headers values
func NewGetCredentialsInternalServerError() *GetCredentialsInternalServerError {
	return &GetCredentialsInternalServerError{}
}

/*GetCredentialsInternalServerError handles this case with default header values.

Internal error
*/
type GetCredentialsInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetCredentialsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /user/credentials][%d] getCredentialsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetCredentialsInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetCredentialsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
