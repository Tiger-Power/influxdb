package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/influxdata/mrfusion/models"
)

/*DeleteSourcesIDRolesRoleIDNoContent Role has been removed

swagger:response deleteSourcesIdRolesRoleIdNoContent
*/
type DeleteSourcesIDRolesRoleIDNoContent struct {
}

// NewDeleteSourcesIDRolesRoleIDNoContent creates DeleteSourcesIDRolesRoleIDNoContent with default headers values
func NewDeleteSourcesIDRolesRoleIDNoContent() *DeleteSourcesIDRolesRoleIDNoContent {
	return &DeleteSourcesIDRolesRoleIDNoContent{}
}

// WriteResponse to the client
func (o *DeleteSourcesIDRolesRoleIDNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(204)
}

/*DeleteSourcesIDRolesRoleIDNotFound Unknown role id

swagger:response deleteSourcesIdRolesRoleIdNotFound
*/
type DeleteSourcesIDRolesRoleIDNotFound struct {

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDRolesRoleIDNotFound creates DeleteSourcesIDRolesRoleIDNotFound with default headers values
func NewDeleteSourcesIDRolesRoleIDNotFound() *DeleteSourcesIDRolesRoleIDNotFound {
	return &DeleteSourcesIDRolesRoleIDNotFound{}
}

// WithPayload adds the payload to the delete sources Id roles role Id not found response
func (o *DeleteSourcesIDRolesRoleIDNotFound) WithPayload(payload *models.Error) *DeleteSourcesIDRolesRoleIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources Id roles role Id not found response
func (o *DeleteSourcesIDRolesRoleIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDRolesRoleIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*DeleteSourcesIDRolesRoleIDDefault Unexpected internal service error

swagger:response deleteSourcesIdRolesRoleIdDefault
*/
type DeleteSourcesIDRolesRoleIDDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewDeleteSourcesIDRolesRoleIDDefault creates DeleteSourcesIDRolesRoleIDDefault with default headers values
func NewDeleteSourcesIDRolesRoleIDDefault(code int) *DeleteSourcesIDRolesRoleIDDefault {
	if code <= 0 {
		code = 500
	}

	return &DeleteSourcesIDRolesRoleIDDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the delete sources ID roles role ID default response
func (o *DeleteSourcesIDRolesRoleIDDefault) WithStatusCode(code int) *DeleteSourcesIDRolesRoleIDDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the delete sources ID roles role ID default response
func (o *DeleteSourcesIDRolesRoleIDDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the delete sources ID roles role ID default response
func (o *DeleteSourcesIDRolesRoleIDDefault) WithPayload(payload *models.Error) *DeleteSourcesIDRolesRoleIDDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the delete sources ID roles role ID default response
func (o *DeleteSourcesIDRolesRoleIDDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DeleteSourcesIDRolesRoleIDDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
