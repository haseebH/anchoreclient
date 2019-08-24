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

// DeleteAnalysisArchiveRuleReader is a Reader for the DeleteAnalysisArchiveRule structure.
type DeleteAnalysisArchiveRuleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteAnalysisArchiveRuleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteAnalysisArchiveRuleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 500:
		result := NewDeleteAnalysisArchiveRuleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteAnalysisArchiveRuleOK creates a DeleteAnalysisArchiveRuleOK with default headers values
func NewDeleteAnalysisArchiveRuleOK() *DeleteAnalysisArchiveRuleOK {
	return &DeleteAnalysisArchiveRuleOK{}
}

/*DeleteAnalysisArchiveRuleOK handles this case with default header values.

Archive transition rule
*/
type DeleteAnalysisArchiveRuleOK struct {
	Payload *models.AnalysisArchiveTransitionRule
}

func (o *DeleteAnalysisArchiveRuleOK) Error() string {
	return fmt.Sprintf("[DELETE /archives/rules/{ruleId}][%d] deleteAnalysisArchiveRuleOK  %+v", 200, o.Payload)
}

func (o *DeleteAnalysisArchiveRuleOK) GetPayload() *models.AnalysisArchiveTransitionRule {
	return o.Payload
}

func (o *DeleteAnalysisArchiveRuleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AnalysisArchiveTransitionRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteAnalysisArchiveRuleInternalServerError creates a DeleteAnalysisArchiveRuleInternalServerError with default headers values
func NewDeleteAnalysisArchiveRuleInternalServerError() *DeleteAnalysisArchiveRuleInternalServerError {
	return &DeleteAnalysisArchiveRuleInternalServerError{}
}

/*DeleteAnalysisArchiveRuleInternalServerError handles this case with default header values.

Internal error
*/
type DeleteAnalysisArchiveRuleInternalServerError struct {
	Payload *models.APIErrorResponse
}

func (o *DeleteAnalysisArchiveRuleInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /archives/rules/{ruleId}][%d] deleteAnalysisArchiveRuleInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteAnalysisArchiveRuleInternalServerError) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *DeleteAnalysisArchiveRuleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
