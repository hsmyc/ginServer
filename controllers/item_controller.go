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

var itemCollection *mongo.Collection = configs.GetCollection(configs.DB, "item")

func CreateItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var item models.Item
		defer cancel()

		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if validationErr := validate.Struct(&item); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}
		if err := item.CType(item.Type); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		newItem := models.Item{
			Name:  item.Name,
			Type:  item.Type,
			Price: item.Price,
			Level: item.Level,
		}

		result, err := itemCollection.InsertOne(ctx, newItem)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.AppResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		itemId := c.Param("itemId")
		var item models.Item
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(itemId)
		err := itemCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&item)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": item}})
	}
}

func EditItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		itemId := c.Param("itemId")
		var item models.Item
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(itemId)

		if err := c.BindJSON(&item); err != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&item); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.AppResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{
			"name":  item.Name,
			"type":  item.Type,
			"price": item.Price,
			"level": item.Level}

		result, err := itemCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var updatedItem models.Item
		if result.MatchedCount == 1 {
			err := itemCollection.FindOneAndUpdate(ctx, bson.M{"_id": objId}, bson.M{"$set": update}).Decode(&updatedItem)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedItem}})

	}
}

func DeleteItem() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		itemId := c.Param("itemId")
		defer cancel()
		objId, _ := primitive.ObjectIDFromHex(itemId)
		result, err := itemCollection.DeleteOne(ctx, bson.M{"_id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
		}
		if result.DeletedCount == 1 {
			c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "item deleted"}})
		} else {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": "error deleting item"}})
		}
	}
}

func GetAllItems() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var items []models.Item
		defer cancel()

		cursor, err := itemCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err}})
			return
		}
		defer cursor.Close(ctx)
		for cursor.Next(ctx) {
			var item models.Item
			cursor.Decode(&item)
			items = append(items, item)
		}
		if err := cursor.Err(); err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err}})
			return
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": items}})
	}
}
