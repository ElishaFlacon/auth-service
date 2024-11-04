package repository

type AuthRepository interface {
	Register() (string, error)
	Login() (string, error)
	Logout() (string, error)
	CheckAuth() (string, error)
	// Refresh() (string, error)
	// TwoFactorAuth() (string, error)
	// ResetPassword() (string, error)
	// ChangePassword() (string, error)
}
