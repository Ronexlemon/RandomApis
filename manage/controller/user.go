package controller

import (
	"context"
	"encoding/json"
	"gomanage/models"
	"gomanage/response"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func SignUp() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		user.CreatedAt = time.Now()
		result, err := user.CreateUser(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

// login
func Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.UserResponses{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		result, err := user.Login(ctx, user.Email, user.Password)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}

// profile
func Profile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, _ := primitive.ObjectIDFromHex(id)
		var user models.User
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		result, err := user.Profile(ctx, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)
	}
}
