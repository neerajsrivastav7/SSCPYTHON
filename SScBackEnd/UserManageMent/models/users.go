package models

// User represents a user in the system with Email, Password, and OTP fields.
type User struct {
    Email    string `json:"email" validate:"required"`  
    Password string `json:"password,omitempty"`         
    OTP      string `json:"otp,omitempty"`              
}
