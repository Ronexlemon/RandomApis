package controller

import (
	"encoding/json"

	"gojobquest/model"
	"gojobquest/response"
	"net/http"
)

func QuestCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}

func Quests() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}

func QuestApplied() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}

func QuestRejected() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}

func QuestApprove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}

func QuestOngoing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest

		err := json.NewDecoder(r.Body).Decode(&quest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			response := response.Response{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}

		//post the data

	}
}
