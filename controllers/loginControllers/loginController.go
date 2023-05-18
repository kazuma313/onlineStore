package logincontrollers

import (
	"html/template"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {

		temp, err := template.ParseFiles("views/login/login.html")

		if err != nil {
			panic(err)
		}
		
		temp.Execute(w, nil)

}

func Auth(w http.ResponseWriter, r *http.Request) {
	
}