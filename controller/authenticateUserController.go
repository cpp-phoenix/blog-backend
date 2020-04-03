package controller

import (
	"fmt"
	"net/http"
)

func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Authenticating User")
}
