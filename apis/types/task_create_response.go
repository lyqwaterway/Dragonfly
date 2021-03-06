// Code generated by go-swagger; DO NOT EDIT.

package types

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// TaskCreateResponse response get from task creation request.
// swagger:model TaskCreateResponse
type TaskCreateResponse struct {

	// The length of the file dfget requests to download in bytes.
	//
	FileLength int64 `json:"FileLength,omitempty"`

	// ID of the created task.
	ID string `json:"ID,omitempty"`

	// The size of pieces which is calculated as per the following strategy
	// 1. If file's total size is less than 200MB, then the piece size is 4MB by default.
	// 2. Otherwise, it equals to the smaller value between totalSize/100MB + 2 MB and 15MB.
	//
	PieceSize int32 `json:"PieceSize,omitempty"`
}

// Validate validates this task create response
func (m *TaskCreateResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *TaskCreateResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *TaskCreateResponse) UnmarshalBinary(b []byte) error {
	var res TaskCreateResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
