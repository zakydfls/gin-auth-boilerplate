package request

type CreateUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"passowrd"`
}
type UpdateUserRequest struct {
	Id       int    `validate:"required" json:"id"`
	Username string `validate:"required,min=2,max=100" json:"username"`
	Email    string `validate:"required,min=2,max=100" json:"email"`
	Password string `validate:"required,min=2,max=100" json:"passowrd"`
}
type LoginUserRequest struct {
	Username string `validate:"required,min=2,max=100" json:"username"`
	Password string `validate:"required,min=2,max=100" json:"passowrd"`
}
