// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netlify/open-api/go/models"
)

// ListSiteBuildHooksReader is a Reader for the ListSiteBuildHooks structure.
type ListSiteBuildHooksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListSiteBuildHooksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListSiteBuildHooksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewListSiteBuildHooksDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewListSiteBuildHooksOK creates a ListSiteBuildHooksOK with default headers values
func NewListSiteBuildHooksOK() *ListSiteBuildHooksOK {
	return &ListSiteBuildHooksOK{}
}

/*ListSiteBuildHooksOK handles this case with default header values.

OK
*/
type ListSiteBuildHooksOK struct {
	Payload []*models.BuildHook
}

func (o *ListSiteBuildHooksOK) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/build_hooks][%d] listSiteBuildHooksOK  %+v", 200, o.Payload)
}

func (o *ListSiteBuildHooksOK) GetPayload() []*models.BuildHook {
	return o.Payload
}

func (o *ListSiteBuildHooksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListSiteBuildHooksDefault creates a ListSiteBuildHooksDefault with default headers values
func NewListSiteBuildHooksDefault(code int) *ListSiteBuildHooksDefault {
	return &ListSiteBuildHooksDefault{
		_statusCode: code,
	}
}

/*ListSiteBuildHooksDefault handles this case with default header values.

error
*/
type ListSiteBuildHooksDefault struct {
	_statusCode int

	Payload *models.Error
}

// Code gets the status code for the list site build hooks default response
func (o *ListSiteBuildHooksDefault) Code() int {
	return o._statusCode
}

func (o *ListSiteBuildHooksDefault) Error() string {
	return fmt.Sprintf("[GET /sites/{site_id}/build_hooks][%d] listSiteBuildHooks default  %+v", o._statusCode, o.Payload)
}

func (o *ListSiteBuildHooksDefault) GetPayload() *models.Error {
	return o.Payload
}

func (o *ListSiteBuildHooksDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
