package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/service/go/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "CAOS PRESENTA INNOVATION BOX ")
	})

	http.ListenAndServe(":8090", nil)
}
