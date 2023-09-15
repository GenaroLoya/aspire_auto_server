package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type bodyAspire struct {
	Pos   int         `json:"pos"`
	Table []EnumState `json:"table"`
	Dir   EnumDir     `json:"dir"`
}

type bodyAspireRes struct {
	Scenes []Scene `json:"scene"`
}

func postAspireSteps(w http.ResponseWriter, r *http.Request) {
	// Decodifica el JSON de la solicitud en un objeto de tipo bodyAspire.
	var body bodyAspire
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&body); err != nil {
		http.Error(w, "Error al decodificar JSON", http.StatusBadRequest)
		return
	}

	// Ahora, puedes trabajar con 'body' como un objeto de tipo bodyAspire.
	fmt.Println("Pos:", body.Pos)
	fmt.Println("Table:", body.Table)
	fmt.Println("Dir:", body.Dir)

	res, _ := aspireLive(body.Table, body.Pos, body.Dir)

	// Componer la respuesta utilizando la estructura bodyAspireRes.
	response := bodyAspireRes{
		Scenes: res,
	}

	fmt.Println("Response:", response)

	// Codificar la respuesta en formato JSON.
	responseJSON, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Error al codificar JSON de respuesta", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK) // 200 OK
	w.Write(responseJSON)
}
