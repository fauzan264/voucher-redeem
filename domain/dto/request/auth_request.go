package request

type LoginRequest struct {
	Email			string		`json:"email" validate:"required"`
	Password		string		`json:"password" validate:"required"`
}

type RegisterRequest struct {
	Name			string		`json:"name" validate:"required"`
	Email			string		`json:"email" validate:"required"`
	Password		string		`json:"password" validate:"required"`
	IsAdmin			bool		`json:"is_admin"`
}