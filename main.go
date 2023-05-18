package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazuma313/onlineStore/config"
	logincontrollers "github.com/kazuma313/onlineStore/controllers/loginControllers"
	productcategorycontrollers "github.com/kazuma313/onlineStore/controllers/productCategoryControllers"
	registercontrollers "github.com/kazuma313/onlineStore/controllers/registerControllers"
)

func main() {
    router := mux.NewRouter()

	// router.HandleFunc("/", homecontrollers.Index)
	// router.HandleFunc("/productCategory", productcategorycontrollers.Index)
	router.HandleFunc("/", productcategorycontrollers.Index)
	router.HandleFunc("/productCategory/add", func (w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "haven't Done")
	})

	router.HandleFunc("/api/product/{category_id}", productcategorycontrollers.Api)

	router.HandleFunc("/login", logincontrollers.Login)
	router.HandleFunc("/login", logincontrollers.Auth)
	router.HandleFunc("/register", registercontrollers.Register)
	config.ConnectDB()

	fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", router)
}