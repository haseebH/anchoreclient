// Code generated by go-swagger; DO NOT EDIT.

package archives

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// ArchiveImageAnalysisReader is a Reader for the ArchiveImageAnalysis structure.
type ArchiveImageAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ArchiveImageAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewArchiveImageAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewArchiveImageAnalysisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewArchiveImageAnalysisOK creates a ArchiveImageAnalysisOK with default headers values
func NewArchiveImageAnalysisOK() *ArchiveImageAnalysisOK {
	return &ArchiveImageAnalysisOK{}
}

/*ArchiveImageAnalysisOK handles this case with default header values.

Archive statuses
*/
type ArchiveImageAnalysisOK struct {
	Payload models.AddAnalysisArchiveResult
}

func (o *ArchiveImageAnalysisOK) Error() string {
	return fmt.Sprintf("[POST /archives/images][%d] archiveImageAnalysisOK  %+v", 200, o.Payload)
}

func (o *ArchiveImageAnalysisOK) GetPayload() models.AddAnalysisArchiveResult {
	return o.Payload
}

func (o *ArchiveImageAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewArchiveImageAnalysisInternalServerError creates a ArchiveImageAnalysisInternalServerError with default headers values
func NewArchiveImageAnalysisInternalServerError() *ArchiveImageAnalysisInternalServerError {
	return &ArchiveImageAnalysisInternalServerError{}
}

/*ArchiveImageAnalysisInternalServerError handles this case with default header values.

Internal error
*/
type ArchiveImageAnalysisInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *ArchiveImageAnalysisInternalServerError) Error() string {
	return fmt.Sprintf("[POST /archives/images][%d] archiveImageAnalysisInternalServerError  %+v", 500, o.Payload)
}

func (o *ArchiveImageAnalysisInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *ArchiveImageAnalysisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
