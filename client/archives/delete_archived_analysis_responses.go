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

// DeleteArchivedAnalysisReader is a Reader for the DeleteArchivedAnalysis structure.
type DeleteArchivedAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteArchivedAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteArchivedAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteArchivedAnalysisInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteArchivedAnalysisOK creates a DeleteArchivedAnalysisOK with default headers values
func NewDeleteArchivedAnalysisOK() *DeleteArchivedAnalysisOK {
	return &DeleteArchivedAnalysisOK{}
}

/*DeleteArchivedAnalysisOK handles this case with default header values.

ArchivdImageAnalysis record but with status 'deleted'
*/
type DeleteArchivedAnalysisOK struct {
	Payload *models.ArchivedAnalysis
}

func (o *DeleteArchivedAnalysisOK) Error() string {
	return fmt.Sprintf("[DELETE /archives/images/{imageDigest}][%d] deleteArchivedAnalysisOK  %+v", 200, o.Payload)
}

func (o *DeleteArchivedAnalysisOK) GetPayload() *models.ArchivedAnalysis {
	return o.Payload
}

func (o *DeleteArchivedAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ArchivedAnalysis)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteArchivedAnalysisInternalServerError creates a DeleteArchivedAnalysisInternalServerError with default headers values
func NewDeleteArchivedAnalysisInternalServerError() *DeleteArchivedAnalysisInternalServerError {
	return &DeleteArchivedAnalysisInternalServerError{}
}

/*DeleteArchivedAnalysisInternalServerError handles this case with default header values.

Internal error
*/
type DeleteArchivedAnalysisInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *DeleteArchivedAnalysisInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /archives/images/{imageDigest}][%d] deleteArchivedAnalysisInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteArchivedAnalysisInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *DeleteArchivedAnalysisInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
