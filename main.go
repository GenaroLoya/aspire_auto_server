package main

import (
	. "aspire_auto_server/server"
)

func main() {

	srv := New(":8080")

	err := srv.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
