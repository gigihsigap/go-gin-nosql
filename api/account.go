package api

import (
	"context"
	"fmt"
	"net/http"
	"time"

	database "backend/db"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/go-playground/validator"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create a validator object
var validate = validator.New()

// Connect to to the database and open an "account" collection
var accountsCollection *mongo.Collection = database.OpenCollection(database.Client, "account")

type Account struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Username   string             `json:"username" validate:"required"`
	LineID     string             `json:"line_id" validate:"required" `
	Email      string             `json:"email" validate:"required"`
	Created_at time.Time          `json:"created_at"`
	Updated_at time.Time          `json:"updated_at"`
}

func (server *Server) createAccount(ctx *gin.Context) {

	// This is used to determine how long the API call should last
	var timeCtx, cancel = context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Declare a variable of type Account
	var accounts Account

	// Bind the object that comes in with the declared varaible. thrrow an error if one occurs
	if err := ctx.BindJSON(&accounts); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Use the validation packge to verify that all items coming in meet the requirements of the struct
	validationErr := validate.Struct(accounts)
	if validationErr != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": validationErr.Error()})
		return
	}

	// Create time stamps
	accounts.Created_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	accounts.Updated_at, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))

	// Generate new ID for the object to be created
	accounts.ID = primitive.NewObjectID()

	// Insert the newly created object into mongodb
	result, insertErr := accountsCollection.InsertOne(timeCtx, accounts)
	if insertErr != nil {
		msg := fmt.Sprintf("Account item was not created")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		return
	}

	// Return the id of the created object to the frontend
	ctx.JSON(http.StatusOK, result)
}

// func (server *Server) getAccount(ctx *gin.Context) {

// }

// func (server *Server) listAccounts(ctx *gin.Context) {

// }
