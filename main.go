package images

import (
	"fmt"
	"net/http"
)

func handler(response http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Path[len("/images/"):]
	fmt.Fprintf(response, "%s", fileName)
}

func init() {
	http.HandleFunc("/images/", handler)
	http.ListenAndServe(":8080", nil)
}
