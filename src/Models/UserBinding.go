package Models

type UserBinding struct {
	Email    string `json:"Email" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
