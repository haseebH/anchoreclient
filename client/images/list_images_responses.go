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

// ListImagesReader is a Reader for the ListImages structure.
type ListImagesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImagesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListImagesInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListImagesOK creates a ListImagesOK with default headers values
func NewListImagesOK() *ListImagesOK {
	return &ListImagesOK{}
}

/*ListImagesOK handles this case with default header values.

successful operation
*/
type ListImagesOK struct {
	Payload models.AnchoreImageList
}

func (o *ListImagesOK) Error() string {
	return fmt.Sprintf("[GET /images][%d] listImagesOK  %+v", 200, o.Payload)
}

func (o *ListImagesOK) GetPayload() models.AnchoreImageList {
	return o.Payload
}

func (o *ListImagesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagesInternalServerError creates a ListImagesInternalServerError with default headers values
func NewListImagesInternalServerError() *ListImagesInternalServerError {
	return &ListImagesInternalServerError{}
}

/*ListImagesInternalServerError handles this case with default header values.

Internal Error
*/
type ListImagesInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *ListImagesInternalServerError) Error() string {
	return fmt.Sprintf("[GET /images][%d] listImagesInternalServerError  %+v", 500, o.Payload)
}

func (o *ListImagesInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *ListImagesInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}