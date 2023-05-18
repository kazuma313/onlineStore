package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazuma313/onlineStore/config"
	homecontrollers "github.com/kazuma313/onlineStore/controllers/homeControllers"
	productcategorycontrollers "github.com/kazuma313/onlineStore/controllers/productCategoryControllers"
)

// func handler(route string, massege string){
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/"+route, func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprint(w, massege)
// 	})
// http.HandleFunc("/"+route, func(w http.ResponseWriter, r *http.Request) {
// 	w.Write([]byte(massege))
// })
// }

func main() {
    router := mux.NewRouter()

    // Register a handler function for the "/hello/{name}" route
    // router.HandleFunc("/hello/{name}", func(w http.ResponseWriter, r *http.Request) {
    //     // Get the name parameter from the route URL
    //     vars := mux.Vars(r)
    //     name := vars["name"]

    //     fmt.Fprintf(w, "Hello, %s!", name)
    // })

	// handlerIndex := func (w http.ResponseWriter, r *http.Request)  {
	// w.Write([]byte("Your program worked!!!"))	
	// }

	// http.HandleFunc("/", handlerIndex)
	// http.HandleFunc("/index", handlerIndex)

	// handler("chart", "shopping chart")
	// handler("payment", "your payment")
	// handler("login", "please login")
	// handler("registration", "Please register")

	router.HandleFunc("/", homecontrollers.Index)
	router.HandleFunc("/productCategory", productcategorycontrollers.Index)
	router.HandleFunc("/api/product/{category_id}", productcategorycontrollers.Api)
	router.HandleFunc("/productCategory/add", productcategorycontrollers.Add)
	router.HandleFunc("/productCategory/edit", productcategorycontrollers.Edit)
	router.HandleFunc("/productCategory/delete", productcategorycontrollers.Delete)
	router.HandleFunc("/productCategory/view", productcategorycontrollers.View)
	config.ConnectDB()

	fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", router)
}