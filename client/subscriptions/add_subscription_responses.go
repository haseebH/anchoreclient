// Code generated by go-swagger; DO NOT EDIT.

package subscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/haseebh/anchoreclient/models"
)

// AddSubscriptionReader is a Reader for the AddSubscription structure.
type AddSubscriptionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddSubscriptionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAddSubscriptionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddSubscriptionOK creates a AddSubscriptionOK with default headers values
func NewAddSubscriptionOK() *AddSubscriptionOK {
	return &AddSubscriptionOK{}
}

/*AddSubscriptionOK handles this case with default header values.

Subscription add success
*/
type AddSubscriptionOK struct {
	Payload models.SubscriptionList
}

func (o *AddSubscriptionOK) Error() string {
	return fmt.Sprintf("[POST /subscriptions][%d] addSubscriptionOK  %+v", 200, o.Payload)
}

func (o *AddSubscriptionOK) GetPayload() models.SubscriptionList {
	return o.Payload
}

func (o *AddSubscriptionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
