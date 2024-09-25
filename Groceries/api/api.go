package api

import (
	"encoding/json"
	"gogroceries/models"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

var groceries = []models.Grocery{
	{Name: "Mango",Price: 100,Quantity: 20,Unit: "kg"},
	{Name: "Apple",Price: 50,Quantity: 30,Unit: "kg"},
	{Name: "Banana",Price: 20,Quantity: 40,Unit: "kg"},
	
}
func AllGroceries(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusFound)
	json.NewEncoder(w).Encode(groceries)

}

func SingleGroceries(w http.ResponseWriter, r * http.Request){
	w.Header().Set("Content-Type", "application/json")
	vars:= mux.Vars(r)
	name:= vars["name"]

	for _,grocery:= range groceries{
		if grocery.Name == name{
			json.NewEncoder(w).Encode(models.Grocery{
				Name: grocery.Name,
				Price: grocery.Price,
				Quantity: grocery.Quantity,
				Unit: grocery.Unit,
			})
		}
		
	}

}

//grocery to buy
func GroceriesToBuy(w http.ResponseWriter, r * http.Request){
	reqBody,err := ioutil.ReadAll(r.Body)
	if err != nil{
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//unmarshal
	var grocery models.Grocery
	json.Unmarshal(reqBody,&grocery)
	groceries = append(groceries, grocery)

	json.NewEncoder(w).Encode(groceries)
}

//delete

func DeleteGrocery(w http.ResponseWriter, r * http.Request){

	vars:= mux.Vars(r)
	name:= vars["name"]
	for index,grocery:= range groceries{
		if grocery.Name == name{
			groceries = append(groceries[:index],groceries[index+1:]...)}

	}
}

//update grocery

func UpdateGrocery(w http.ResponseWriter, r * http.Request){
	vars:=mux.Vars(r)
	name:= vars["name"]
	for index,grocery:= range groceries{
		if grocery.Name == name{
			var updateGrocery models.Grocery
			json.NewDecoder(r.Body).Decode(&updateGrocery)
			groceries[index].Name = updateGrocery.Name
			groceries[index].Price = updateGrocery.Price
			groceries[index].Quantity= updateGrocery.Quantity
			groceries[index].Unit = updateGrocery.Unit
			json.NewEncoder(w).Encode(updateGrocery)
			return
		}
	}
}