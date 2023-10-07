package routing

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	data := []map[string]any{
		{
			"nama":   "fani alfirdaus",
			"alamat": "indonesia",
			"umur":   17,
		}, {
			"nama":   "sandhika galih",
			"alamat": "indonesia",
			"umur":   35,
		}, {
			"nama":   "eko kurniawan  khanedy",
			"alamat": "indonesia",
			"umur":   35,
		},
	}

	w.Header().Add("Content-Type", "application/json")

	b, err := json.Marshal(data)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Fprintln(w, string(b))
}
