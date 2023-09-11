package controllers

import (
	"context"
	"ginGonic/learn/configs"
	"ginGonic/learn/models"
	"ginGonic/learn/responses"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var charCollection *mongo.Collection = configs.GetCollection(configs.DB, "character")

func CreateChar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var char models.Character
		defer cancel()

		if err := c.BindJSON(&char); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&char); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newChar := models.Character{
			Name:  char.Name,
			Level: char.Level,
		}

		result, err := charCollection.InsertOne(ctx, newChar)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AppResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetChar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		charId := c.Param("charId")
		var char models.Character
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(charId)
		err := charCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&char)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": char}})
	}
}

func EditChar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		charId := c.Param("charId")
		var char models.Character
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(charId)
		err := charCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&char)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if err := c.BindJSON(&char); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&char); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		update := bson.M{
			"name":  char.Name,
			"level": char.Level,
		}

		result, err := charCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedChar models.Character
		if result.MatchedCount == 1 {
			err := charCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedChar)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedChar}})
	}
}

func DeleteChar() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		charId := c.Param("charId")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(charId)
		result, err := charCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "no document deleted"}})
			return
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "character deleted"}})
	}
}

func GetAllChars() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var chars []models.Character
		defer cancel()

		results, err := charCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		defer results.Close(ctx)
		for results.Next(ctx) {
			var char models.Character
			if err := results.Decode(&char); err != nil {
				c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
			chars = append(chars, char)
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": chars}})

	}
}
