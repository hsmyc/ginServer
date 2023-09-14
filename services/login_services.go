package services

import (
	"context"
	"ginGonic/learn/configs"
	"ginGonic/learn/models"
	"ginGonic/learn/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
)

type Claims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(userId string) (string, error) {
	claims := &Claims{
		UserId: userId,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(configs.EnvSecret()))
}

func VerifyToken(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(configs.EnvSecret()), nil
	})
	if err != nil {
		return claims, err
	}
	if !token.Valid {
		return claims, err
	}
	return claims, nil
}

var validate = validator.New()

func LoginHandler() gin.HandlerFunc {
	return func(c *gin.Context) {

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var user models.User
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "model didn't get error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&user); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "not valid error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		hashedPassword, err := configs.HashPassword(user.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "Pasword cannot be hashed!", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		loggedUser := models.User{
			Name:     user.Name,
			Email:    user.Email,
			Password: hashedPassword,
			Image:    user.Image,
		}

		err = userCollection.FindOne(ctx, bson.M{"email": loggedUser.Email}).Decode(&loggedUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "user can't find error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if err := configs.ComparePassword(loggedUser.Password, user.Password); err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "compare error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		token, err := GenerateToken(loggedUser.ID.Hex())
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "can't create a token error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "You logged in!", Data: map[string]interface{}{"data": token}})
	}
}