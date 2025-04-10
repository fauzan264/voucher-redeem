package request

import "github.com/google/uuid"

type GetUser struct {
	ID		uuid.UUID 	`json:"id" validate:"required"`
}