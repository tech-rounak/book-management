package controllers

import (
	"context"
	"github/tech-rounak/book-management/database"
	"github/tech-rounak/book-management/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var bookCollection = database.OpenCollection(database.Client, "book")
var validate = validator.New()

func CreateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		var book models.Book

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err.Error()})
			return
		}

		if validationErr := validate.Struct(&book); validationErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": validationErr.Error()})
			return
		}
		book.ID = primitive.NewObjectID()
		book.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		book.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

		result, err := bookCollection.InsertOne(ctx, book)
		defer cancel()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "Error while creating the document"})
			return
		}
		c.JSON(http.StatusAccepted, gin.H{"success": true, "msg": "Book Created Succesfully", "result": result})
	}
}
func UpdateBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		bid := c.Param("bookId")
		bookId, _ := primitive.ObjectIDFromHex(bid)
		filter := bson.M{"_id": bookId}

		var book models.Book
		var foundBook models.Book
		var updatedObj primitive.D

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": err.Error()})
			return
		}

		err := bookCollection.FindOne(ctx, filter).Decode(&foundBook)
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{"success": false, "msg": "No Book Found with this id"})
			return
		}
		defer cancel()

		book.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC1123))
		updatedObj = append(updatedObj, bson.E{"updatedAt", book.UpdatedAt})
		if book.Name != nil {
			updatedObj = append(updatedObj, bson.E{"name", book.Name})
		}
		if book.Publisher != nil {
			updatedObj = append(updatedObj, bson.E{"publisher", book.Publisher})
		}
		if book.Author != nil {
			updatedObj = append(updatedObj, bson.E{"author", book.Author})
		}
		if book.Price != nil {
			updatedObj = append(updatedObj, bson.E{"price", book.Price})
		}
		if book.ISBN != nil {
			updatedObj = append(updatedObj, bson.E{"isbn", book.ISBN})
		}

		opts := options.Update().SetUpsert(true)

		res, err := bookCollection.UpdateOne(
			ctx,
			filter,
			bson.D{{"$set", updatedObj}},
			opts,
		)
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{"success": false, "msg": "Cannot Update Book"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Books Updated Succesfully", "result": res})
	}
}
func GetBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		bid := c.Param("bookId")
		bookId, _ := primitive.ObjectIDFromHex(bid)
		filter := bson.M{"_id": bookId}

		var book models.Book

		err := bookCollection.FindOne(ctx, filter).Decode(&book)
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{"success": false, "msg": "No Book Found with this id"})
			return
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Books Fetched Succesfully", "result": book})
	}
}
func GetBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		var books []models.Book

		cur, err := bookCollection.Find(ctx, bson.M{})
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{"success": false, "msg": "No Book Found with this id"})
			return
		}
		defer cancel()
		
		for cur.Next(ctx) {
			var book models.Book
			err := cur.Decode(&book)
			if err != nil {
				log.Fatal(err)
			}

			// add item our array
			books = append(books, book)
		}

		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Books Fetched Succesfully", "result": books})
	}
}
func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)

		bid := c.Param("bookId")
		bookId, _ := primitive.ObjectIDFromHex(bid)
		filter := bson.M{"_id": bookId}

		var book models.Book

		err := bookCollection.FindOne(ctx, filter).Decode(&book)
		if err != nil {
			c.JSON(http.StatusAccepted, gin.H{"success": false, "msg": "No Book Found with this id"})
			return
		}
		defer cancel()

		res, err := bookCollection.DeleteOne(ctx, filter)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"success": false, "msg": "Data Couldn't "})
		}
		defer cancel()
		c.JSON(http.StatusOK, gin.H{"success": true, "msg": "Books Deleted Succesfully", "result": res})
	}
}
