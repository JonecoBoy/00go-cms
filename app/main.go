package main

import (
	"00go-cms/presentation/routes"
	"net/http"
)

func main() {

	r := routes.Routes()

	http.ListenAndServe(":8080", r)
}
