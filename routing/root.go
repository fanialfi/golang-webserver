package routing

import (
	"fmt"
	"html/template"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"Name":    "Fani Alfirdaus",
		"message": "memiliki hari yang bagus",
	}

	t, err := template.ParseFiles("./template/template.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	t.Execute(w, data)
}
