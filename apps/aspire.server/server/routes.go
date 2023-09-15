package server

import "net/http"

func initRoutes() {
	http.HandleFunc("/steps", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			postAspireSteps(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
