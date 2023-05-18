package productcategorycontrollers

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	categorymodels "github.com/kazuma313/onlineStore/models/categoryModels"
)

func Index(w http.ResponseWriter, r *http.Request) {
	categories := categorymodels.GetData()
	data := map[string]any{
		"categories": categories,
	}

	temp, err := template.ParseFiles("views/productCategory/index.html")

	if err != nil {
		panic(err)	
	}

	temp.Execute(w, data)
}


func Api(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    strid := vars["category_id"]
	intid, _ := strconv.Atoi(strid)
	
	product := categorymodels.GetDataProductByCategory(intid) 
	json.NewEncoder(w).Encode(product)
}


func Add(w http.ResponseWriter, r *http.Request) {

}

func Delete(w http.ResponseWriter, r *http.Request) {

}

func Edit(w http.ResponseWriter, r *http.Request) {

}

func View(w http.ResponseWriter, r *http.Request) {

}