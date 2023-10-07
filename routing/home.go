package routing

import (
	"fmt"
	"net/http"
	"text/template"
)

func Home(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"nama":       "ini adalah halaman home",
		"keterangan": "Lorem ipsum dolor sit amet, consectetur adipisicing elit. Facilis eveniet libero et nemo incidunt autem repellendus quos repellat nisi obcaecati, aliquid omnis quis cum saepe voluptatum nesciunt, possimus distinctio laudantium?",
	}

	template, err := template.ParseFiles("./template/home.html")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	template.Execute(w, data)
}
