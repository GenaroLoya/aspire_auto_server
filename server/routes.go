package server

import "net/http"

func initRoutes() {
	// finalIndex := middleFirsLayer(middleReqLogger(func(w http.ResponseWriter, r *http.Request) {
	// 	switch r.Method {
	// 	case http.MethodGet:
	// 		index(w, r)
	// 	default:
	// 		w.WriteHeader(http.StatusMethodNotAllowed)
	// 	}
	// }))

	// http.HandleFunc("/", finalIndex)

	http.HandleFunc("/countries", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAspireSteps(w, r)
		// case http.MethodPost:
		// 	postCountrie(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
}
