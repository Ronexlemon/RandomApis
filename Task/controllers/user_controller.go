package controllers

import (
	"context"
	"encoding/json"
	"gomanagetask/configs"
	"gomanagetask/models"
	"gomanagetask/responses"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var userCollection *mongo.Collection = configs.GetCollection(configs.DB, "users")
var validate = validator.New()

func CreateUser() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var user models.User
		defer cancel()
		//validate the request body -> unmarshial
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		//use the validator library to validate required fields
		if validationErr := validate.Struct(&user); validationErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			response := responses.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		newUser := models.User{
			Id:       primitive.NewObjectID(),
			Name:     user.Name,
			Location: user.Location,
			Title:    user.Title,
		}
		result, err := userCollection.InsertOne(ctx, newUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponses{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusCreated)
		response := responses.UserResponses{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}}
		json.NewEncoder(w).Encode(response)
	}
}

//get all user events

func GetAUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		vars := mux.Vars(r)
		id := vars["id"]
		var user models.User
		//get collection
		objid, _ := primitive.ObjectIDFromHex(id)
		err := userCollection.FindOne(ctx, bson.M{"id": objid}).Decode(&user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response := responses.UserResponses{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		w.WriteHeader(http.StatusOK)
		response := responses.UserResponses{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": user}}
		json.NewEncoder(w).Encode(response)

	}
}

//update user field
func Update() http.HandlerFunc{
	return func(w http.ResponseWriter,r *http.Request){
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		vars:= mux.Vars(r)
		var user models.User
		id:=vars["id"]
		objId,_:= primitive.ObjectIDFromHex(id)
		//validate the request body
		if err:= json.NewDecoder(r.Body).Decode(&user);err !=nil{
			w.WriteHeader(http.StatusBadRequest)
			response:= responses.UserResponses{Status:http.StatusBadRequest,Message: "error",Data: map[string]interface{}{"data":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//use the validator library to validate required fields
		if validatorerr := validate.Struct(&user);validatorerr !=nil{
			w.WriteHeader(http.StatusBadRequest)
			response:= responses.UserResponses{Status:http.StatusBadRequest,Message: "error",Data: map[string]interface{}{"data":validatorerr.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		update:=bson.M{"name":user.Name,"location":user.Location,"title":user.Title}
		result,err:= userCollection.UpdateOne(ctx,bson.M{"id":objId},bson.M{"$set":update})
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			response:= responses.UserResponses{Status:http.StatusInternalServerError,Message: "error",Data: map[string]interface{}{"data":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return



	}
	//get updated user details
	var updatedUser models.User
	if result.MatchedCount ==1{
		err:= userCollection.FindOne(ctx,bson.M{"id":objId}).Decode(&updatedUser)
		if err!=nil{
			w.WriteHeader(http.StatusInternalServerError)
			response:= responses.UserResponses{Status:http.StatusInternalServerError,Message: "error",Data: map[string]interface{}{"data":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
	}
	w.WriteHeader(http.StatusOK)
	response:= responses.UserResponses{Status:http.StatusOK,Message: "success",Data: map[string]interface{}{"data":updatedUser}}
	json.NewEncoder(w).Encode(response)
}
}}