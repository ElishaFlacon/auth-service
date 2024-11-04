package service

type AuthService interface {
	CheckAuth() (string, error)
	Register() (string, error)
	Login() (string, error)
	Logout() (string, error)
	// Refresh() (string, error)
	// TwoFactorAuth() (string, error)
	// ResetPassword() (string, error)
	// ChangePassword() (string, error)
}
