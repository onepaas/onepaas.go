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

// CreateProjectReader is a Reader for the CreateProject structure.
type CreateProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 422:
		result := NewCreateProjectUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateProjectCreated creates a CreateProjectCreated with default headers values
func NewCreateProjectCreated() *CreateProjectCreated {
	return &CreateProjectCreated{}
}

/* CreateProjectCreated describes a response with status code 201, with default header values.

Created
*/
type CreateProjectCreated struct {
	Payload *models.Project
}

func (o *CreateProjectCreated) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createProjectCreated  %+v", 201, o.Payload)
}
func (o *CreateProjectCreated) GetPayload() *models.Project {
	return o.Payload
}

func (o *CreateProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Project)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectUnprocessableEntity creates a CreateProjectUnprocessableEntity with default headers values
func NewCreateProjectUnprocessableEntity() *CreateProjectUnprocessableEntity {
	return &CreateProjectUnprocessableEntity{}
}

/* CreateProjectUnprocessableEntity describes a response with status code 422, with default header values.

Unprocessable Entity
*/
type CreateProjectUnprocessableEntity struct {
	Payload *models.Problem
}

func (o *CreateProjectUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createProjectUnprocessableEntity  %+v", 422, o.Payload)
}
func (o *CreateProjectUnprocessableEntity) GetPayload() *models.Problem {
	return o.Payload
}

func (o *CreateProjectUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectInternalServerError creates a CreateProjectInternalServerError with default headers values
func NewCreateProjectInternalServerError() *CreateProjectInternalServerError {
	return &CreateProjectInternalServerError{}
}

/* CreateProjectInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type CreateProjectInternalServerError struct {
	Payload *models.Problem
}

func (o *CreateProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /v1/projects][%d] createProjectInternalServerError  %+v", 500, o.Payload)
}
func (o *CreateProjectInternalServerError) GetPayload() *models.Problem {
	return o.Payload
}

func (o *CreateProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Problem)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}