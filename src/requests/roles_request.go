package requests

type Role struct {
	RoleName string `json:"role_name" validate:"required"`
}
