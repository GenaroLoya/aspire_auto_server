package server

import (
	"fmt"
	"net/http"

	"github.com/fatih/color"
)

func New(addr string) *http.Server {
	initRoutes()

	fmt.Println(color.GreenString("Server running on: "), color.CyanString(addr+"\n"))

	return &http.Server{
		Addr: addr,
	}
}
