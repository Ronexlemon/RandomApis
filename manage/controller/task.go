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

func TaskCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var task models.Task

		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.UserResponses{Status: http.StatusBadRequest, Message: "incorrect body", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		result, err := task.Create(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "internal server error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

// get all user tasks
func TaskUserTasks() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var task models.Task
		//pass the user id from the middleware -> TODO

		//roundit
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, _ := primitive.ObjectIDFromHex(id)
		result, err := task.UserTasks(ctx, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "internal server error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

// update
func TaskUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var task models.Task

		//pass the user id from the middleware -> TODO

		//roundit
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, _ := primitive.ObjectIDFromHex(id)
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.UserResponses{Status: http.StatusBadRequest, Message: "Incorrect Body", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		result, err := task.Update(ctx, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "Internal server error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

// delete
func TaskDelete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-Type", "application/json")
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()
		var task models.Task

		//pass the user id from the middleware -> TODO

		//roundit
		vars := mux.Vars(r)
		id := vars["id"]
		user_id, _ := primitive.ObjectIDFromHex(id)
		result, err := task.Delete(ctx, user_id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.UserResponses{Status: http.StatusInternalServerError, Message: "Internal server error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		json.NewEncoder(w).Encode(result)

	}
}
