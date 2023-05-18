package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazuma313/onlineStore/config"
	homecontrollers "github.com/kazuma313/onlineStore/controllers/homeControllers"
	logincontrollers "github.com/kazuma313/onlineStore/controllers/loginControllers"
	productcategorycontrollers "github.com/kazuma313/onlineStore/controllers/productCategoryControllers"
	registercontrollers "github.com/kazuma313/onlineStore/controllers/registerControllers"
)

func main() {
    router := mux.NewRouter()

	router.HandleFunc("/", homecontrollers.Index)
	router.HandleFunc("/productCategory", productcategorycontrollers.Index)
	router.HandleFunc("/api/product/{category_id}", productcategorycontrollers.Api)
	router.HandleFunc("/productCategory/add", productcategorycontrollers.Add)
	router.HandleFunc("/productCategory/edit", productcategorycontrollers.Edit)
	router.HandleFunc("/productCategory/delete", productcategorycontrollers.Delete)
	router.HandleFunc("/productCategory/view", productcategorycontrollers.View)

	router.HandleFunc("/login", logincontrollers.Login)
	router.HandleFunc("/login", logincontrollers.Auth)
	router.HandleFunc("/register", registercontrollers.Register)
	config.ConnectDB()

	fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", router)
}