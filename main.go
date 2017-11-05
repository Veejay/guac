package images

import (
	"fmt"
	"html/template"
	"net/http"
)

type ip struct {
	address string
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Path[len("/images/"):]
	fmt.Fprintf(response, "%s", fileName)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(response, ip{address: request.Header.Get("X-Forwarded-For")})
}

func init() {
	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
