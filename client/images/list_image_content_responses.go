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

// ListImageContentReader is a Reader for the ListImageContent structure.
type ListImageContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImageContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImageContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListImageContentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListImageContentOK creates a ListImageContentOK with default headers values
func NewListImageContentOK() *ListImageContentOK {
	return &ListImageContentOK{}
}

/*ListImageContentOK handles this case with default header values.

Content listing for the image
*/
type ListImageContentOK struct {
	Payload []string
}

func (o *ListImageContentOK) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/content][%d] listImageContentOK  %+v", 200, o.Payload)
}

func (o *ListImageContentOK) GetPayload() []string {
	return o.Payload
}

func (o *ListImageContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImageContentInternalServerError creates a ListImageContentInternalServerError with default headers values
func NewListImageContentInternalServerError() *ListImageContentInternalServerError {
	return &ListImageContentInternalServerError{}
}

/*ListImageContentInternalServerError handles this case with default header values.

Internal Error
*/
type ListImageContentInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *ListImageContentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/content][%d] listImageContentInternalServerError  %+v", 500, o.Payload)
}

func (o *ListImageContentInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *ListImageContentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
