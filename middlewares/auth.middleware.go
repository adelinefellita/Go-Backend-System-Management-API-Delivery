package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// RoleBasedAuth checks if the user has one of the allowed roles
func RoleBasedAuth(allowedRoles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user role from the header
		userRole := c.GetHeader("Role")
		if userRole == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role header is missing"})
			c.Abort()
			return
		}

		// Validate the role
		for _, role := range allowedRoles {
			if strings.EqualFold(userRole, role) {
				c.Next()
				return
			}
		}

		// If the role is not allowed
		c.JSON(http.StatusForbidden, gin.H{"error": "Access denied"})
		c.Abort()
	}
}

// RoleBasedAccess checks if the user has a specific role
func RoleBasedAccess(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the user role from the header
		userRole := c.GetHeader("Role")
		if userRole == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Role header is missing"})
			c.Abort()
			return
		}

		// Check if the role matches
		if !strings.EqualFold(userRole, requiredRole) {
			c.JSON(http.StatusForbidden, gin.H{"error": "Forbidden: You do not have the required role"})
			c.Abort()
			return
		}

		c.Next()
	}
}
