// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetGreetingReader is a Reader for the GetGreeting structure.
type GetGreetingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetGreetingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetGreetingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetGreetingOK creates a GetGreetingOK with default headers values
func NewGetGreetingOK() *GetGreetingOK {
	return &GetGreetingOK{}
}

/*
GetGreetingOK describes a response with status code 200, with default header values.

returns a greeting
*/
type GetGreetingOK struct {
	Payload string
}

// IsSuccess returns true when this get greeting o k response has a 2xx status code
func (o *GetGreetingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get greeting o k response has a 3xx status code
func (o *GetGreetingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get greeting o k response has a 4xx status code
func (o *GetGreetingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get greeting o k response has a 5xx status code
func (o *GetGreetingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get greeting o k response a status code equal to that given
func (o *GetGreetingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetGreetingOK) Error() string {
	return fmt.Sprintf("[GET /hello][%d] getGreetingOK  %+v", 200, o.Payload)
}

func (o *GetGreetingOK) String() string {
	return fmt.Sprintf("[GET /hello][%d] getGreetingOK  %+v", 200, o.Payload)
}

func (o *GetGreetingOK) GetPayload() string {
	return o.Payload
}

func (o *GetGreetingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}