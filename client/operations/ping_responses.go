// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// PingReader is a Reader for the Ping structure.
type PingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPingOK creates a PingOK with default headers values
func NewPingOK() *PingOK {
	return &PingOK{}
}

/*PingOK handles this case with default header values.

Version check response, returns the api version prefix (e.g. 'v1')
*/
type PingOK struct {
	Payload string
}

func (o *PingOK) Error() string {
	return fmt.Sprintf("[GET /][%d] pingOK  %+v", 200, o.Payload)
}

func (o *PingOK) GetPayload() string {
	return o.Payload
}

func (o *PingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
