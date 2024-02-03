package functions

import (
	"fmt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "Now let's see Argentin's Players wich played on World Cup 2022 Done by:Golovinov Andrey")
}
