package models

import "fmt"

var (
	ErrUserNotFound       = fmt.Errorf("User not found")
	ErrInvalidCredentials = fmt.Errorf("Invalid credentials")
	ErrUserExists         = fmt.Errorf("User, with these credentials, is already exists")
	ErrInvalidToken       = fmt.Errorf("Invalid token")
	ErrServer             = fmt.Errorf("Server error")

	ErrComplaintNotFound = fmt.Errorf("Complaint not found")
)
