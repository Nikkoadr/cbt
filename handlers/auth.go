package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"cbt/config"
	"cbt/models"
	"cbt/utils"
)

// AuthHandler holds the database connection
type AuthHandler struct {
	DB  *gorm.DB
	Cfg config.Config
}

// Login handles user login
func (h *AuthHandler) Login(c *gin.Context) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&creds); err != nil {
		utils.SendError(c, http.StatusBadRequest, "Invalid request body")
		return
	}

	var user models.User
	if err := h.DB.Where("username = ?", creds.Username).First(&user).Error; err != nil {
		utils.SendError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	if !utils.CheckPasswordHash(creds.Password, user.Password) {
		utils.SendError(c, http.StatusUnauthorized, "Invalid credentials")
		return
	}

	token, err := utils.GenerateJWT(user, h.Cfg)
	if err != nil {
		utils.SendError(c, http.StatusInternalServerError, "Failed to generate token")
		return
	}

	c.SetCookie("token", token, h.Cfg.JWTExpire*3600, "/", "", false, true)

	utils.SendSuccess(c, "Login successful", gin.H{"token": token})
}

// Logout handles user logout
func (h *AuthHandler) Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	utils.SendSuccess(c, "Logout successful", nil)
}

// Me returns the currently logged in user
func (h *AuthHandler) Me(c *gin.Context) {
	userID, _ := c.Get("user_id")

	var user models.User
	if err := h.DB.First(&user, userID).Error; err != nil {
		utils.SendError(c, http.StatusNotFound, "User not found")
		return
	}

	utils.SendSuccess(c, "User details", user)
}
