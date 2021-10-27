// Code generated by go-swagger; DO NOT EDIT.

package project

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	models "github.com/onepaas/onepaas/pkg/api/v1"
)

// ReplaceProjectReader is a Reader for the ReplaceProject structure.
type ReplaceProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReplaceProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReplaceProjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 404:
		result := NewReplaceProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewReplaceProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReplaceProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewReplaceProjectOK creates a ReplaceProjectOK with default headers values
func NewReplaceProjectOK() *ReplaceProjectOK {
	return &ReplaceProjectOK{}
}

/* ReplaceProjectOK describes a response with status code 200, with default header values.

OK
*/
type ReplaceProjectOK struct {
	Payload *models.Project
}

func (o *ReplaceProjectOK) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/{id}][%d] replaceProjectOK  %+v", 200, o.Payload)
}
func (o *ReplaceProjectOK) GetPayload() *models.Project {
	return o.Payload
}

func (o *ReplaceProjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceProjectNotFound creates a ReplaceProjectNotFound with default headers values
func NewReplaceProjectNotFound() *ReplaceProjectNotFound {
	return &ReplaceProjectNotFound{}
}

/* ReplaceProjectNotFound describes a response with status code 404, with default header values.

Not Found
*/
type ReplaceProjectNotFound struct {
	Payload *models.Problem
}

func (o *ReplaceProjectNotFound) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/{id}][%d] replaceProjectNotFound  %+v", 404, o.Payload)
}
func (o *ReplaceProjectNotFound) GetPayload() *models.Problem {
	return o.Payload
}

func (o *ReplaceProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceProjectUnprocessableEntity creates a ReplaceProjectUnprocessableEntity with default headers values
func NewReplaceProjectUnprocessableEntity() *ReplaceProjectUnprocessableEntity {
	return &ReplaceProjectUnprocessableEntity{}
}

/* ReplaceProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type ReplaceProjectUnprocessableEntity struct {
	Payload *models.Problem
}

func (o *ReplaceProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/{id}][%d] replaceProjectUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *ReplaceProjectUnprocessableEntity) GetPayload() *models.Problem {
	return o.Payload
}

func (o *ReplaceProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReplaceProjectInternalServerError creates a ReplaceProjectInternalServerError with default headers values
func NewReplaceProjectInternalServerError() *ReplaceProjectInternalServerError {
	return &ReplaceProjectInternalServerError{}
}

/* ReplaceProjectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type ReplaceProjectInternalServerError struct {
	Payload *models.Problem
}

func (o *ReplaceProjectInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /v1/projects/{id}][%d] replaceProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *ReplaceProjectInternalServerError) GetPayload() *models.Problem {
	return o.Payload
}

func (o *ReplaceProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
