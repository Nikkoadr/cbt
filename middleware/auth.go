package middleware

import (
	"net/http"
	"strings"
	"time"

	"cbt/models"
	"cbt/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// LoginPayload defines the structure for the login request
type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// LoginHandler handles the user login, validates credentials, and returns a JWT.
func LoginHandler(jwtSecret string, db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var payload LoginPayload
		if err := c.ShouldBindJSON(&payload); err != nil {
			utils.SendError(c, http.StatusBadRequest, "Invalid request payload")
			return
		}

		var user models.User
		if err := db.Where("username = ?", payload.Username).First(&user).Error; err != nil {
			utils.SendError(c, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(payload.Password)); err != nil {
			utils.SendError(c, http.StatusUnauthorized, "Invalid username or password")
			return
		}

		// Create the JWT claims, including the user role
		claims := jwt.MapClaims{
			"id":       user.ID,
			"username": user.Username,
			"role":     user.Role, // 'siswa' or 'admin'
			"exp":      time.Now().Add(time.Hour * 72).Unix(),
		}

		// Create token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		t, err := token.SignedString([]byte(jwtSecret))
		if err != nil {
			utils.SendError(c, http.StatusInternalServerError, "Could not generate token")
			return
		}

		utils.SendSuccess(c, "Login successful", gin.H{"token": t})
	}
}

// genericRoleMiddleware provides a template for role-based authentication.
func genericRoleMiddleware(jwtSecret, requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.SendError(c, http.StatusUnauthorized, "Missing Authorization header")
			c.Abort()
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Don't forget to validate the alg is what you expect:
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, http.ErrAbortHandler
			}
			return []byte(jwtSecret), nil
		})

		if err != nil {
			utils.SendError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			role, ok := claims["role"].(string)
			if !ok || role != requiredRole {
				utils.SendError(c, http.StatusForbidden, "You are not authorized to access this resource")
				c.Abort()
				return
			}
			// Set user context
			c.Set("user_id", claims["id"])
			c.Set("username", claims["username"])
			// Continue to the next handler
			c.Next()
		} else {
			utils.SendError(c, http.StatusUnauthorized, "Invalid token")
			c.Abort()
			return
		}
	}
}

// SiswaMiddleware checks if the user has the 'siswa' role.
func SiswaMiddleware(jwtSecret string) gin.HandlerFunc {
	return genericRoleMiddleware(jwtSecret, "siswa")
}

// AdminMiddleware checks if the user has the 'admin' role.
func AdminMiddleware(jwtSecret string) gin.HandlerFunc {
	return genericRoleMiddleware(jwtSecret, "admin")
}
