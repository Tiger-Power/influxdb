package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/go-openapi/errors"
)

/*Permissions permissions

swagger:model Permissions
*/
type Permissions struct {

	/* permissions are scoped to a specific database
	 */
	Database string `json:"database,omitempty"`

	/* permissions
	 */
	Permissions []Permission `json:"permissions,omitempty"`
}

// Validate validates this permissions
func (m *Permissions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Permissions) validatePermissions(formats strfmt.Registry) error {

	if swag.IsZero(m.Permissions) { // not required
		return nil
	}

	return nil
}
