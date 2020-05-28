package forms

// SignupUserCommand defines user form struct
type SignupUserCommand struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginUserCommand defines user login form struct
type LoginUserCommand struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// PasswordResetCommand defines user password reset form struct
type PasswordResetCommand struct {
	Password string `json:"password" binding:"required"`
	Confirm  string `json:"confirm" binding:"required"`
}

// ResendCommand defines resend email payload
type ResendCommand struct {
	Email string `json:"email" binding:"required"`
}
