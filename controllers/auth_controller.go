package controllers

import (
	"time"

	"github.com/Raghunandan-79/auth-service/database"
	"github.com/Raghunandan-79/auth-service/models"
	"github.com/Raghunandan-79/auth-service/utils"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var body struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	c.BindJSON(&body)

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 12)
	if err != nil {
		c.JSON(500, gin.H{"error": "Passowrd hashing failed"})
		return
	}

	user := models.User {
		Name: body.Name,
		Email: body.Email,
		Password: string(hash),
	}

	if err := database.DB.Create(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "user exists"})
		return
	}

	c.JSON(201, gin.H{"message": "registered"})
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string
	}

	c.BindJSON(&body)

	var user models.User
	if err := database.DB.Where("email = ?", body.Email).First(&user).Error; err != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	if bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)) != nil {
		c.JSON(401, gin.H{"error": "invalid credentials"})
		return
	}

	accessToken, _ := utils.GenerateAccessToken(user.ID)
	refreshToken := utils.GenerateRefreshToken()

	rt := models.RefreshToken {
		UserID: user.ID,
		TokenHash: refreshToken,
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
	}

	database.DB.Create(&rt)

	c.SetCookie("refresh_token", refreshToken, 24592000, "/", "", false, true)

	c.JSON(200, gin.H {
		"access_token": accessToken,
	})
}

func Refresh(c *gin.Context) {
	rt, err := c.Cookie("refresh_token")
	if err != nil {
		c.JSON(401, gin.H{"error": "missing refresh token"})
		return
	}

	var token models.RefreshToken
	if err := database.DB.Where("token_hash = ?", rt).First(&token).Error; err != nil {
		c.JSON(401, gin.H{"error": "invalid refresh token"})
		return
	}

	if time.Now().After(token.ExpiresAt) {
		c.JSON(401, gin.H{"error": "refresh token expired"})
		return
	}

	newAccess, _ := utils.GenerateAccessToken(token.UserID)

	c.JSON(200, gin.H{"access_token": newAccess})
}

func Logout(c *gin.Context) {
	rt, _ := c.Cookie("refresh_token")
	database.DB.Where("token_hash = ?", rt).Delete(&models.RefreshToken{})
	c.SetCookie("refresh_token", "", -1, "/", "", false, true)
	c.JSON(200, gin.H{"message": "logged out"})
}

func Me(c *gin.Context) {
	userID := c.GetUint("user_id")
	c.JSON(200, gin.H{"user_id": userID})
}
