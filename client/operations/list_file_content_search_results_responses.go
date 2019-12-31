// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// ListFileContentSearchResultsReader is a Reader for the ListFileContentSearchResults structure.
type ListFileContentSearchResultsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFileContentSearchResultsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFileContentSearchResultsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewListFileContentSearchResultsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListFileContentSearchResultsOK creates a ListFileContentSearchResultsOK with default headers values
func NewListFileContentSearchResultsOK() *ListFileContentSearchResultsOK {
	return &ListFileContentSearchResultsOK{}
}

/*ListFileContentSearchResultsOK handles this case with default header values.

List of file metadata objects
*/
type ListFileContentSearchResultsOK struct {
	Payload models.FileContentSearchList
}

func (o *ListFileContentSearchResultsOK) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/artifacts/file_content_search][%d] listFileContentSearchResultsOK  %+v", 200, o.Payload)
}

func (o *ListFileContentSearchResultsOK) GetPayload() models.FileContentSearchList {
	return o.Payload
}

func (o *ListFileContentSearchResultsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFileContentSearchResultsNotFound creates a ListFileContentSearchResultsNotFound with default headers values
func NewListFileContentSearchResultsNotFound() *ListFileContentSearchResultsNotFound {
	return &ListFileContentSearchResultsNotFound{}
}

/*ListFileContentSearchResultsNotFound handles this case with default header values.

Image not found in this service
*/
type ListFileContentSearchResultsNotFound struct {
}

func (o *ListFileContentSearchResultsNotFound) Error() string {
	return fmt.Sprintf("[GET /images/{imageDigest}/artifacts/file_content_search][%d] listFileContentSearchResultsNotFound ", 404)
}

func (o *ListFileContentSearchResultsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}