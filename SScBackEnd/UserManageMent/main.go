package main

import (
    "log"
    "github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/userManageMent/handler"
    "github.com/neerajsrivastav7/SSCPYTHON/SScBackEnd/userManageMent/utils"
    "github.com/gin-gonic/gin"
)

type UserManageMentRest struct {
	auth   *handler.Auth
	notify  *handler.Notification
	admin   *handler.Admin
	user    *handler.User
}
func main() {
    // Load configuration
	var usermngrest UserManageMentRest
    config, err := utils.LoadConfig("config/userConfig.json")
    if err != nil {
        log.Fatalf("Error loading configuration: %v", err)
    }

    router := gin.Default()

    // Authentication routes
    router.POST("/api/auth/register", usermngrest.auth.RegisterUser)
    router.POST("/api/auth/login", usermngrest.auth.LoginUser)
    router.POST("/api/auth/logout", usermngrest.auth.LogoutUser)
    router.POST("/api/auth/refresh-token", usermngrest.auth.RefreshToken)
    router.POST("/api/auth/forgot-password", usermngrest.auth.ForgotPassword)
    router.POST("/api/auth/reset-password",usermngrest.auth.ResetPassword)

    // User profile routes
    router.GET("/api/users/:userId", usermngrest.user.GetUser)
    router.PUT("/api/users/:userId", usermngrest.user.UpdateUser)
    router.DELETE("/api/users/:userId", usermngrest.user.DeleteUser)

    // User listing and search
    router.GET("/api/users", usermngrest.user.ListUsers)
    router.GET("/api/users/search", usermngrest.user.SearchUsers)

    // Admin-specific routes
    router.GET("/api/admin/dashboard",usermngrest.admin.AdminDashboard)
    router.PUT("/api/admin/users/:userId/status", usermngrest.admin.ChangeUserStatus)

    // Notification routes
    router.GET("/api/users/:userId/notifications", usermngrest.notify.GetUserNotifications)
    router.POST("/api/users/:userId/notifications", usermngrest.notify.SendUserNotification)

    // Start the server using configuration settings
    address := config.Server.IP + ":" + config.Server.Port
    router.Run(address)
}
