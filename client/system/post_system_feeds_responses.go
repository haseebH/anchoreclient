// Code generated by go-swagger; DO NOT EDIT.

package system

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// PostSystemFeedsReader is a Reader for the PostSystemFeeds structure.
type PostSystemFeedsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostSystemFeedsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostSystemFeedsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewPostSystemFeedsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostSystemFeedsOK creates a PostSystemFeedsOK with default headers values
func NewPostSystemFeedsOK() *PostSystemFeedsOK {
	return &PostSystemFeedsOK{}
}

/*PostSystemFeedsOK handles this case with default header values.

Feeds operation success
*/
type PostSystemFeedsOK struct {
	Payload models.FeedSyncResults
}

func (o *PostSystemFeedsOK) Error() string {
	return fmt.Sprintf("[POST /system/feeds][%d] postSystemFeedsOK  %+v", 200, o.Payload)
}

func (o *PostSystemFeedsOK) GetPayload() models.FeedSyncResults {
	return o.Payload
}

func (o *PostSystemFeedsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostSystemFeedsInternalServerError creates a PostSystemFeedsInternalServerError with default headers values
func NewPostSystemFeedsInternalServerError() *PostSystemFeedsInternalServerError {
	return &PostSystemFeedsInternalServerError{}
}

/*PostSystemFeedsInternalServerError handles this case with default header values.

Internal Error
*/
type PostSystemFeedsInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *PostSystemFeedsInternalServerError) Error() string {
	return fmt.Sprintf("[POST /system/feeds][%d] postSystemFeedsInternalServerError  %+v", 500, o.Payload)
}

func (o *PostSystemFeedsInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *PostSystemFeedsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
