package controller

import (
	"context"
	"encoding/json"
	"time"

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
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := quest.Quests(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)


	}
}

func QuestPending() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := quest.Pendinguests(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

func QuestRejected() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := quest.RejectedQuests(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

func QuestApprove() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := quest.ApprovedQuests(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}

func QuestOngoing() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		var quest model.Quest
		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		result, err := quest.ActiveQuests(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			response := response.Response{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}}
			json.NewEncoder(w).Encode(response)
			return
		}
		json.NewEncoder(w).Encode(result)

	}
}
