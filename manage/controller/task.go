package controller

import (
	"context"
	"encoding/json"
	"gomanage/models"
	"gomanage/response"
	"net/http"
	"time"
)


func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
		defer cancel()
		var task models.Task

		err:= json.NewDecoder(r.Body).Decode(&task)
		if err !=nil{
			http.Error(w, err.Error(), http.StatusBadRequest)
			response:= response.UserResponses{Status: http.StatusBadRequest,Message: "incorrect body",Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		result,err := task.Create(ctx)
		if err !=nil{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response:= response.UserResponses{Status: http.StatusInternalServerError,Message: "internal server error",Data:map[string]interface{}{"data":err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

		

	}
}