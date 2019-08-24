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

// GetImageContentByTypeImageIDFilesReader is a Reader for the GetImageContentByTypeImageIDFiles structure.
type GetImageContentByTypeImageIDFilesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageContentByTypeImageIDFilesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageContentByTypeImageIDFilesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetImageContentByTypeImageIDFilesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageContentByTypeImageIDFilesOK creates a GetImageContentByTypeImageIDFilesOK with default headers values
func NewGetImageContentByTypeImageIDFilesOK() *GetImageContentByTypeImageIDFilesOK {
	return &GetImageContentByTypeImageIDFilesOK{}
}

/*GetImageContentByTypeImageIDFilesOK handles this case with default header values.

Content of specified type from the image
*/
type GetImageContentByTypeImageIDFilesOK struct {
	Payload *models.ContentFilesResponse
}

func (o *GetImageContentByTypeImageIDFilesOK) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/content/files][%d] getImageContentByTypeImageIdFilesOK  %+v", 200, o.Payload)
}

func (o *GetImageContentByTypeImageIDFilesOK) GetPayload() *models.ContentFilesResponse {
	return o.Payload
}

func (o *GetImageContentByTypeImageIDFilesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentFilesResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageContentByTypeImageIDFilesInternalServerError creates a GetImageContentByTypeImageIDFilesInternalServerError with default headers values
func NewGetImageContentByTypeImageIDFilesInternalServerError() *GetImageContentByTypeImageIDFilesInternalServerError {
	return &GetImageContentByTypeImageIDFilesInternalServerError{}
}

/*GetImageContentByTypeImageIDFilesInternalServerError handles this case with default header values.

Internal error
*/
type GetImageContentByTypeImageIDFilesInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetImageContentByTypeImageIDFilesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/content/files][%d] getImageContentByTypeImageIdFilesInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImageContentByTypeImageIDFilesInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetImageContentByTypeImageIDFilesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
