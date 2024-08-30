package handlers

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Admin struct {

} 
func (ad *Admin)AdminDashboard(c *gin.Context) {
    // Implement logic to get admin dashboard stats
    c.JSON(http.StatusOK, gin.H{"stats": "Admin dashboard stats"})
}

func (ad *Admin)ChangeUserStatus(c *gin.Context) {
    userId := c.Param("userId")
    // Implement logic to change user status
    c.JSON(http.StatusOK, gin.H{"message": "User status changed successfully", "userId": userId})
}
