// Code generated by go-swagger; DO NOT EDIT.

package vulnerabilities

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// QueryVulnerabilitiesReader is a Reader for the QueryVulnerabilities structure.
type QueryVulnerabilitiesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QueryVulnerabilitiesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewQueryVulnerabilitiesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewQueryVulnerabilitiesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewQueryVulnerabilitiesOK creates a QueryVulnerabilitiesOK with default headers values
func NewQueryVulnerabilitiesOK() *QueryVulnerabilitiesOK {
	return &QueryVulnerabilitiesOK{}
}

/*QueryVulnerabilitiesOK handles this case with default header values.

Vulnerability listing paginated
*/
type QueryVulnerabilitiesOK struct {
	Payload *models.PaginatedVulnerabilityList
}

func (o *QueryVulnerabilitiesOK) Error() string {
	return fmt.Sprintf("[GET /query/vulnerabilities][%d] queryVulnerabilitiesOK  %+v", 200, o.Payload)
}

func (o *QueryVulnerabilitiesOK) GetPayload() *models.PaginatedVulnerabilityList {
	return o.Payload
}

func (o *QueryVulnerabilitiesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PaginatedVulnerabilityList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQueryVulnerabilitiesBadRequest creates a QueryVulnerabilitiesBadRequest with default headers values
func NewQueryVulnerabilitiesBadRequest() *QueryVulnerabilitiesBadRequest {
	return &QueryVulnerabilitiesBadRequest{}
}

/*QueryVulnerabilitiesBadRequest handles this case with default header values.

Invalid filter parameters
*/
type QueryVulnerabilitiesBadRequest struct {
	Payload *models.APIErrorResponse
}

func (o *QueryVulnerabilitiesBadRequest) Error() string {
	return fmt.Sprintf("[GET /query/vulnerabilities][%d] queryVulnerabilitiesBadRequest  %+v", 400, o.Payload)
}

func (o *QueryVulnerabilitiesBadRequest) GetPayload() *models.APIErrorResponse {
	return o.Payload
}

func (o *QueryVulnerabilitiesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
