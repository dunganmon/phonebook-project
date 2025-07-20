package controllers

import (
	"TODO/databases"
	"TODO/dtos"
	"TODO/models"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Register(c *gin.Context) {
	var userReq dtos.UserRegisterRequest
	if err := c.ShouldBind(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newUser := models.CreateUser(userReq.Username, userReq.Password, userReq.FullName)
	if err := databases.DB.Create(&newUser).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "created",
		"data": gin.H{
			"user": newUser,
		},
	})
}

var jwtSecret = []byte("YOUR_SECRET_KEY")

func Login(c *gin.Context) {
	var userReq dtos.UserLoginRequest
	if err := c.ShouldBindJSON(&userReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user models.User
	err := databases.DB.Where("username = ?", userReq.Username).Find(&user).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Lỗi hệ thống",
			"message": err,
		})
		return
	}

	if user.Password != userReq.Password {
		c.JSON(400, gin.H{
			"status":  "failed",
			"message": "Sai mật khẩu",
		})
		return
	}

	var token = jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.ID,
		"username": user.Username,
		"exp":      time.Now().Add(time.Hour * 6).Unix(),
	})
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Lỗi hệ thống"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  tokenString,
		"user": gin.H{
			"id":        user.ID,
			"username":  user.Username,
			"full_name": user.FullName,
		},
	})
}
func GetUserProfile(c *gin.Context) {
	user := c.MustGet("user")
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"user": user,
		},
	})
}
