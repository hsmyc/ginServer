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

var classCollection *mongo.Collection = configs.GetCollection(configs.DB, "class")

func CreateClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var class models.Class
		defer cancel()

		if err := c.BindJSON(&class); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&class); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		newClass := models.Class{
			Name:         class.Name,
			Base_Attack:  class.Base_Attack,
			Base_Defense: class.Base_Defense,
			Base_Health:  class.Base_Health,
		}

		result, err := classCollection.InsertOne(ctx, newClass)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AppResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		classId := c.Param("classId")
		var class models.Class
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(classId)
		err := classCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&class)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": class}})
	}
}

func EditClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		classId := c.Param("classId")
		var class models.Class
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(classId)

		if err := c.BindJSON(&class); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&class); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"name":         class.Name,
			"base_attack":  class.Base_Attack,
			"base_defense": class.Base_Defense,
			"base_health":  class.Base_Health,
		}

		result, err := itemCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var updatedClass models.Class
		if result.MatchedCount == 1 {
			err := classCollection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.M{"$set": update}).Decode(&updatedClass)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedClass}})

	}
}

func DeleteClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		classId := c.Param("classId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(classId)
		result, err := classCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetAllClass() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var classes []models.Class
		defer cancel()
		cursor, err := classCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		if err = cursor.All(ctx, &classes); err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": classes}})
	}
}
