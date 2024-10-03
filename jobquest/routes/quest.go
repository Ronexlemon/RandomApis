package routes

import (
	"gojobquest/controller"

	"github.com/gorilla/mux"
)

func Quest(router *mux.Router) {
	router.HandleFunc("/quest", controller.QuestCreate()).Methods("POST")
	router.HandleFunc("/quests", controller.Quests()).Methods("GET")
	router.HandleFunc("/quests/approve", controller.QuestApprove()).Methods("GET")
	router.HandleFunc("/quests/rejected", controller.QuestRejected()).Methods("GET")
	router.HandleFunc("/quests/active", controller.QuestOngoing()).Methods("GET")
	router.HandleFunc("/quests/pending", controller.QuestPending()).Methods("GET")

}
