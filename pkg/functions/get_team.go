package functions

import (
	"encoding/json"
	_struct "github.com/Andreika228/awesomeProjest1/pkg/struct"
	"net/http"
)

func Persons(w http.ResponseWriter, r *http.Request) {
	var response _struct.Team

	team := Team()

	response.Person = team

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}
