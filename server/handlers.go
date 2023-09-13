package server

import (
	"fmt"
	"net/http"
)

type Country struct {
	Name     string
	Language string
}

func getAspireSteps(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}
