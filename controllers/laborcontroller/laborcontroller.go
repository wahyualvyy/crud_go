package laborcontroller

import (
	"go-web/models/labormodel"
	"net/http"
	"text/template"
)

func Index(w http.ResponseWriter, r*http.Request){
	labors := labormodel.GetAll()
	data := map[string]any{
		"labors" : labors,
	}

	temp, err := template.ParseFiles("views/admin/lab/index.html")
	if err != nil {
		panic(err)
	}

	temp.Execute(w, data)
}

func Add(w http.ResponseWriter, r*http.Request){
	temp,err := template.ParseFiles("views/admin/lab/add.html")
	if err != nil {
		panic(err)
	}
	temp.Execute(w, nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	
}

func Delete(w http.ResponseWriter, r *http.Request) {
	
}