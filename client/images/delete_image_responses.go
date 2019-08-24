// Code generated by go-swagger; DO NOT EDIT.

package images

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// DeleteImageReader is a Reader for the DeleteImage structure.
type DeleteImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteImageOK creates a DeleteImageOK with default headers values
func NewDeleteImageOK() *DeleteImageOK {
	return &DeleteImageOK{}
}

/*DeleteImageOK handles this case with default header values.

Image deletion success
*/
type DeleteImageOK struct {
}

func (o *DeleteImageOK) Error() string {
	return fmt.Sprintf("[DELETE /images/{imageDigest}][%d] deleteImageOK ", 200)
}

func (o *DeleteImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}