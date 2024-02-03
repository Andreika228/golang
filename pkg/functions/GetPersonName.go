package functions

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPersonName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	var jsonResponse []byte
	var er error

	for _, item := range Team() {
		if item.SecondName == params["second_name"] {
			jsonResponse, er = json.Marshal(item)

		}

	}
	if er != nil {
		return
	}
	w.Write(jsonResponse)
}
