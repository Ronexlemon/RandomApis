package api

import (
	"encoding/json"
	"gomyapp/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetWord(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type","application/json")
	vars:=mux.Vars(r)
	searchTerm,ok:= vars["search"]
	if !ok{
		http.Error(w,http.StatusText(http.StatusBadRequest),http.StatusBadRequest)
		return}
		responseData:= models.ResponseData{SearchTerm: searchTerm}

		if err:= json.NewEncoder(w).Encode(responseData); err !=nil{
			http.Error(w,err.Error(),http.StatusInternalServerError)
			return
		}


}