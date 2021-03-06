// Code generated by go-swagger; DO NOT EDIT.

package summaries

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// ListImagetagsReader is a Reader for the ListImagetags structure.
type ListImagetagsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListImagetagsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListImagetagsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewListImagetagsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListImagetagsOK creates a ListImagetagsOK with default headers values
func NewListImagetagsOK() *ListImagetagsOK {
	return &ListImagetagsOK{}
}

/*ListImagetagsOK handles this case with default header values.

successful operation
*/
type ListImagetagsOK struct {
	Payload models.AnchoreImageTagSummaryList
}

func (o *ListImagetagsOK) Error() string {
	return fmt.Sprintf("[GET /summaries/imagetags][%d] listImagetagsOK  %+v", 200, o.Payload)
}

func (o *ListImagetagsOK) GetPayload() models.AnchoreImageTagSummaryList {
	return o.Payload
}

func (o *ListImagetagsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListImagetagsInternalServerError creates a ListImagetagsInternalServerError with default headers values
func NewListImagetagsInternalServerError() *ListImagetagsInternalServerError {
	return &ListImagetagsInternalServerError{}
}

/*ListImagetagsInternalServerError handles this case with default header values.

Internal Error
*/
type ListImagetagsInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *ListImagetagsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /summaries/imagetags][%d] listImagetagsInternalServerError  %+v", 500, o.Payload)
}

func (o *ListImagetagsInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *ListImagetagsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
