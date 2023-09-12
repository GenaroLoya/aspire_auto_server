package server

import "net/http"

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getCountries(w, r)
		case http.MethodPost:
			postCountrie(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
