package controllers

import (
	"app/db"
	"app/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type Claims struct {
	UserId int    `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

func Register(context *gin.Context) {
	var req models.User

	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Unable to bind JSON"})
		return
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	_, err = db.DB.Exec("INSERT INTO users(username, email, password, role) VALUES($1, $2, $3, 'user')", req.Username, req.Email, hashPass)

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to insert user"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(context *gin.Context) {
	var req models.User

	err := context.BindJSON(&req)

	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"err": "Failed to bing json"})
		return
	}

	var user models.User

	err = db.DB.QueryRow("SELECT id, username, email, role, password FROM users WHERE email = $1", req.Email).Scan(&user.Id, &user.Username, &user.Email, &user.Role, &user.Password)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Faile to select user"})
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
		return
	}

	expiration := time.Now().Add(5 * time.Hour)

	claims := &Claims{
		UserId: user.Id,
		Role:   user.Role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiration),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret_key"))

	if err != nil {
		context.IndentedJSON(http.StatusInternalServerError, gin.H{"error": "Unable to generate token"})
		return
	}

	context.IndentedJSON(http.StatusOK, gin.H{"message": "Login successful", "user_id": user.Id, "role": user.Role, "token": tokenString})
}
