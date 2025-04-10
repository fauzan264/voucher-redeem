package response

import "github.com/google/uuid"

type UserResponse struct {
	ID				uuid.UUID	`json:"id"`
	Name			string		`json:"name"`
	Email			string		`json:"email"`
	TotalPoints		int			`json:"total_points"`
	IsAdmin			bool		`json:"is_admin"`
	Token			string		`json:"token,omitempty"`
}