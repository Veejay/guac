package images

import (
	"fmt"
	"html/template"
	"net/http"
)

type ip struct {
	Address string
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Path[len("/images/"):]
	fmt.Fprintf(response, "%s", fileName)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(response, ip{Address: request.RemoteAddr})
}

func init() {
	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
