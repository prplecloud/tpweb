package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"regexp"
)

func main() {

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	type Promo struct {
		NomPromo string
		Filière  string
		Niveau   int
		Nombre   int
	}

	type Etudiant struct {
		Nom  string
		Age  int
		Sexe string
	}

	type Data struct {
		Promo          Promo
		ListeEtudiants []Etudiant
	}

	listeEtudiants := []Etudiant{{"Cyril RODRIGUES", 22, "Homme"}, {"Kheir-Eddine MEDERREG", 22, "Homme"}, {"Alan PHILIPIERT", 26, "Homme"}}

	http.HandleFunc("/promo", func(w http.ResponseWriter, r *http.Request) {
		dataPage := Data{Promo{"Mentor'ac", "Informatique", 5, 3}, listeEtudiants}

		temp.ExecuteTemplate(w, "promo", dataPage)
	})

	type DataForm struct {
		CheckData bool
		Data      string
	}
	http.HandleFunc("/form/treatement", func(w http.ResponseWriter, r *http.Request) {
		var data DataForm
		if r.Method == http.MethodGet {
			data = DataForm{
				CheckData: false,
				Data:      r.URL.Query().Get("display"),
			}
		}

		if r.Method == http.MethodPost {
			data = DataForm{
				CheckData: false,
				Data:      r.FormValue("display"),
			}
		}

		checkValue, _ := regexp.MatchString("^[a-zA-Z-]{1,64}$", data.Data)
		if !checkValue {
			data.CheckData = true
		}

		temp.ExecuteTemplate(w, "display", data)
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	fmt.Println("localhost:6969")
	http.ListenAndServe("localhost:6969", nil)
}
