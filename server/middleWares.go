package server

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func middleReqLogger(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(color.CyanString("Request:"), color.HiBlueString(r.Method), color.HiBlueString(r.URL.Path))
		next.ServeHTTP(w, r)
		fmt.Println(color.CyanString("Response:"), color.HiBlueString(r.Method), color.HiBlueString(r.URL.Path))
	})
}

func middleFirsLayer(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println(color.CyanString("Request1:"), color.HiBlueString(r.Method), color.HiBlueString(r.URL.Path))
		next.ServeHTTP(w, r)
		fmt.Println(color.CyanString("Response1:"), color.HiBlueString(r.Method), color.HiBlueString(r.URL.Path))
	})
}
