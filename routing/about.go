package routing

import (
	"fmt"
	"net/http"
	"text/template"
)

func About(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"nama":       "Ini adalah halaman About",
		"keterangan": "selamat belajar golang web",
	}

	template, err := template.ParseFiles("./template/about.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	template.Execute(w, data)
}
