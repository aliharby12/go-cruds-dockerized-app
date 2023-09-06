package models

// Define roles
const (
	AdminRole   = "admin"
	BloggerRole = "blogger"
)

// IsAdmin checks if a user has an admin role.
func (u *User) IsAdmin() bool {
	return u.Role == AdminRole
}

// IsBlogger checks if a user has a blogger role.
func (u *User) IsBlogger() bool {
	return u.Role == BloggerRole
}
