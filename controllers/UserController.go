package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/PrasGi/learn-golang/initializers"
	"github.com/PrasGi/learn-golang/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	// get email/pass
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Failed read body",
		})
		return
	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed generating password",
		})
		return
	}

	// Create user
	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(400, gin.H{
			"message": "Failed creating user",
		})

		return
	}

	// Response
	c.JSON(200, gin.H{
		"message": "Success creating user",
		"data":    user,
	})
	return
}

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{
			"message": "Failed read body",
		})
		return
	}

	var user models.User
	initializers.DB.Find(&user, "email = ?", body.Email)

	if user.ID == 0 {
		c.JSON(400, gin.H{
			"message": "Failed login, email or password is wrong : user not found",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed login, email or password is wrong : password wrong",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"nbf": time.Now().Add(time.Hour * 24 * 30).Unix(),
		"exp": time.Now().Add(time.Hour * 24 * 30).Add(time.Hour).Unix(), // Contoh: 31 hari kedepan
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(400, gin.H{
			"message": "Failed login, email or password is wrong : error make token",
		})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(200, gin.H{
		"message": "Success login",
		"user":    user,
	})
}

func Validate(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Success validate",
	})
	return
}
