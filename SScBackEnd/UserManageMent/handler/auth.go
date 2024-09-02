// auth.go
package handler

import (
    "net/http"
    "github.com/gin-gonic/gin"
    "github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/comman/userModel"
)

// Auth struct definition
type Auth struct {
     validate Validate
}

// Handler for user registration
func (au *Auth) RegisterUser(c *gin.Context) {
    c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

// Handler for user login
func (au *Auth) LoginUser(c *gin.Context) {
    var validateIns Validate
    var loginDetails userModel.Login
    if err := c.BindJSON(&loginDetails); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    loginResponse, err := validateIns.ValidateLogin(loginDetails.Email, loginDetails.Password, loginDetails.LoginType)
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": loginResponse})
}

// Handler for user logout
func (au *Auth) LogoutUser(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

// Handler for token refresh
func (au *Auth) RefreshToken(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Token refreshed successfully"})
}

// Handler for forgot password
func (au *Auth) ForgotPassword(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Password reset email sent"})
}

// Handler for password reset
func (au *Auth) ResetPassword(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
