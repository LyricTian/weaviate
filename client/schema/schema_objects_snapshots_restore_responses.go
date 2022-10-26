//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package schema

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// SchemaObjectsSnapshotsRestoreReader is a Reader for the SchemaObjectsSnapshotsRestore structure.
type SchemaObjectsSnapshotsRestoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SchemaObjectsSnapshotsRestoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSchemaObjectsSnapshotsRestoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewSchemaObjectsSnapshotsRestoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewSchemaObjectsSnapshotsRestoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewSchemaObjectsSnapshotsRestoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewSchemaObjectsSnapshotsRestoreUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewSchemaObjectsSnapshotsRestoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewSchemaObjectsSnapshotsRestoreOK creates a SchemaObjectsSnapshotsRestoreOK with default headers values
func NewSchemaObjectsSnapshotsRestoreOK() *SchemaObjectsSnapshotsRestoreOK {
	return &SchemaObjectsSnapshotsRestoreOK{}
}

/*
SchemaObjectsSnapshotsRestoreOK handles this case with default header values.

Snapshot restoring process successfully started.
*/
type SchemaObjectsSnapshotsRestoreOK struct {
	Payload *models.SnapshotRestoreMeta
}

func (o *SchemaObjectsSnapshotsRestoreOK) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreOK  %+v", 200, o.Payload)
}

func (o *SchemaObjectsSnapshotsRestoreOK) GetPayload() *models.SnapshotRestoreMeta {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsRestoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnapshotRestoreMeta)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsRestoreUnauthorized creates a SchemaObjectsSnapshotsRestoreUnauthorized with default headers values
func NewSchemaObjectsSnapshotsRestoreUnauthorized() *SchemaObjectsSnapshotsRestoreUnauthorized {
	return &SchemaObjectsSnapshotsRestoreUnauthorized{}
}

/*
SchemaObjectsSnapshotsRestoreUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type SchemaObjectsSnapshotsRestoreUnauthorized struct {
}

func (o *SchemaObjectsSnapshotsRestoreUnauthorized) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreUnauthorized ", 401)
}

func (o *SchemaObjectsSnapshotsRestoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSchemaObjectsSnapshotsRestoreForbidden creates a SchemaObjectsSnapshotsRestoreForbidden with default headers values
func NewSchemaObjectsSnapshotsRestoreForbidden() *SchemaObjectsSnapshotsRestoreForbidden {
	return &SchemaObjectsSnapshotsRestoreForbidden{}
}

/*
SchemaObjectsSnapshotsRestoreForbidden handles this case with default header values.

Forbidden
*/
type SchemaObjectsSnapshotsRestoreForbidden struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsRestoreForbidden) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreForbidden  %+v", 403, o.Payload)
}

func (o *SchemaObjectsSnapshotsRestoreForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsRestoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsRestoreNotFound creates a SchemaObjectsSnapshotsRestoreNotFound with default headers values
func NewSchemaObjectsSnapshotsRestoreNotFound() *SchemaObjectsSnapshotsRestoreNotFound {
	return &SchemaObjectsSnapshotsRestoreNotFound{}
}

/*
SchemaObjectsSnapshotsRestoreNotFound handles this case with default header values.

Not Found - Snapshot does not exist
*/
type SchemaObjectsSnapshotsRestoreNotFound struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsRestoreNotFound) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreNotFound  %+v", 404, o.Payload)
}

func (o *SchemaObjectsSnapshotsRestoreNotFound) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsRestoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsRestoreUnprocessableEntity creates a SchemaObjectsSnapshotsRestoreUnprocessableEntity with default headers values
func NewSchemaObjectsSnapshotsRestoreUnprocessableEntity() *SchemaObjectsSnapshotsRestoreUnprocessableEntity {
	return &SchemaObjectsSnapshotsRestoreUnprocessableEntity{}
}

/*
SchemaObjectsSnapshotsRestoreUnprocessableEntity handles this case with default header values.

Invalid restore snapshot attempt.
*/
type SchemaObjectsSnapshotsRestoreUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsRestoreUnprocessableEntity) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *SchemaObjectsSnapshotsRestoreUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsRestoreUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSchemaObjectsSnapshotsRestoreInternalServerError creates a SchemaObjectsSnapshotsRestoreInternalServerError with default headers values
func NewSchemaObjectsSnapshotsRestoreInternalServerError() *SchemaObjectsSnapshotsRestoreInternalServerError {
	return &SchemaObjectsSnapshotsRestoreInternalServerError{}
}

/*
SchemaObjectsSnapshotsRestoreInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type SchemaObjectsSnapshotsRestoreInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *SchemaObjectsSnapshotsRestoreInternalServerError) Error() string {
	return fmt.Sprintf("[POST /schema/{className}/snapshots/{storageName}/{id}/restore][%d] schemaObjectsSnapshotsRestoreInternalServerError  %+v", 500, o.Payload)
}

func (o *SchemaObjectsSnapshotsRestoreInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SchemaObjectsSnapshotsRestoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}