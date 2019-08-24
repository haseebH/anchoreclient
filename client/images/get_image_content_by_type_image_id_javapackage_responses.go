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

// GetImageContentByTypeImageIDJavapackageReader is a Reader for the GetImageContentByTypeImageIDJavapackage structure.
type GetImageContentByTypeImageIDJavapackageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetImageContentByTypeImageIDJavapackageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetImageContentByTypeImageIDJavapackageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewGetImageContentByTypeImageIDJavapackageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetImageContentByTypeImageIDJavapackageOK creates a GetImageContentByTypeImageIDJavapackageOK with default headers values
func NewGetImageContentByTypeImageIDJavapackageOK() *GetImageContentByTypeImageIDJavapackageOK {
	return &GetImageContentByTypeImageIDJavapackageOK{}
}

/*GetImageContentByTypeImageIDJavapackageOK handles this case with default header values.

Content of specified type from the image
*/
type GetImageContentByTypeImageIDJavapackageOK struct {
	Payload *models.ContentJAVAPackageResponse
}

func (o *GetImageContentByTypeImageIDJavapackageOK) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/content/java][%d] getImageContentByTypeImageIdJavapackageOK  %+v", 200, o.Payload)
}

func (o *GetImageContentByTypeImageIDJavapackageOK) GetPayload() *models.ContentJAVAPackageResponse {
	return o.Payload
}

func (o *GetImageContentByTypeImageIDJavapackageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ContentJAVAPackageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetImageContentByTypeImageIDJavapackageInternalServerError creates a GetImageContentByTypeImageIDJavapackageInternalServerError with default headers values
func NewGetImageContentByTypeImageIDJavapackageInternalServerError() *GetImageContentByTypeImageIDJavapackageInternalServerError {
	return &GetImageContentByTypeImageIDJavapackageInternalServerError{}
}

/*GetImageContentByTypeImageIDJavapackageInternalServerError handles this case with default header values.

Internal error
*/
type GetImageContentByTypeImageIDJavapackageInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *GetImageContentByTypeImageIDJavapackageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images/by_id/{imageId}/content/java][%d] getImageContentByTypeImageIdJavapackageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetImageContentByTypeImageIDJavapackageInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *GetImageContentByTypeImageIDJavapackageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}