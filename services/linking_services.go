package services

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

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "user")
var charCollection *mongo.Collection = configs.GetCollection(configs.DB, "character")

func LinkUserToCharacter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		var user models.User
		var char models.Character

		defer cancel()

		userId := c.Param("userId")
		charId := c.Param("charId")

		objId, _ := primitive.ObjectIDFromHex(userId)
		objId2, _ := primitive.ObjectIDFromHex(charId)

		err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&user)
		err2 := charCollection.FindOne(ctx, bson.M{"_id": objId2}).Decode(&char)

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error1", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		if err2 != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error2", Data: map[string]interface{}{"data": err2.Error()}})
			return
		}
		update := bson.M{
			"name":     user.Name,
			"location": user.Location,
			"title":    user.Title,
			"char_id":  char.ID,
		}
		result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objId}, bson.M{"$set": update})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error3", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		var updatedUser models.User
		if result.MatchedCount == 1 {
			err := userCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&updatedUser)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.AppResponse{Status: http.StatusInternalServerError, Message: "error4", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.AppResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedUser}})
	}
}
