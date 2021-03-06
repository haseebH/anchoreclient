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

// GetImageContentByTypeFilesReader is a Reader for the GetImageContentByTypeFiles structure.
type GetImageContentByTypeFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageContentByTypeFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageContentByTypeFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetImageContentByTypeFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageContentByTypeFilesOK creates a GetImageContentByTypeFilesOK with default headers values
func NewGetImageContentByTypeFilesOK() *GetImageContentByTypeFilesOK {
	return &GetImageContentByTypeFilesOK{}
}

/*GetImageContentByTypeFilesOK handles this case with default header values.

Content of specified type from the image
*/
type GetImageContentByTypeFilesOK struct {
	Payload *models.ContentFilesResponse
}

func (o *GetImageContentByTypeFilesOK) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/content/files][%d] getImageContentByTypeFilesOK  %+v", 200, o.Payload)
}

func (o *GetImageContentByTypeFilesOK) GetPayload() *models.ContentFilesResponse {
	return o.Payload
}

func (o *GetImageContentByTypeFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageContentByTypeFilesInternalServerError creates a GetImageContentByTypeFilesInternalServerError with default headers values
func NewGetImageContentByTypeFilesInternalServerError() *GetImageContentByTypeFilesInternalServerError {
	return &GetImageContentByTypeFilesInternalServerError{}
}

/*GetImageContentByTypeFilesInternalServerError handles this case with default header values.

Internal error
*/
type GetImageContentByTypeFilesInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetImageContentByTypeFilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/content/files][%d] getImageContentByTypeFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImageContentByTypeFilesInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetImageContentByTypeFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
