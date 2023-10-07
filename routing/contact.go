package routing

import (
	"fmt"
	"html/template"
	"net/http"
)

func Contact(w http.ResponseWriter, r *http.Request) {
	data := map[string]any{
		"nama":      "fani alfirdaus",
		"telephone": "08123457890",
		"email":     "fanialfirdaus@gmail.com",
	}

	template, err := template.ParseFiles("./template/contact.html")

	if err != nil {
		fmt.Println(err.Error())
		return
	}
	template.Execute(w, data)
}
