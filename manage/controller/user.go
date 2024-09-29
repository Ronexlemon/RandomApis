package controller

import (
	"context"
	"encoding/json"
	"gomanage/models"
	"net/http"
	"time"
)

func SignUp()http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Set("content-Type","application/json")
		ctx,cancel := context.WithTimeout(context.Background(),time.Second *10)
		defer cancel()
		var user models.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
			}
		result,err:= user.CreateUser(ctx)	
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
			}
			json.NewEncoder(w).Encode(result)

	}
}