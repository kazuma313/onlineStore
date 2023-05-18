package registercontrollers

import (
	"html/template"
	"net/http"

	"github.com/kazuma313/onlineStore/entities"
	costumermodels "github.com/kazuma313/onlineStore/models/costumerModels"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET"{
		temp, err := template.ParseFiles("views/login/register.html")

		if err != nil {
			panic(err)
		}
		
		temp.Execute(w, nil)

	}else if r.Method == "POST"{
		var Costumer entities.Costumer

		Costumer.Email = r.FormValue("email")
		Costumer.Password = r.FormValue("password")
		Costumer.Username = r.FormValue("username")

		ok := costumermodels.Create(Costumer)

		if !ok {
			temp, _ := template.ParseFiles("views/login/register.html")
			temp.Execute(w, nil)
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}




}