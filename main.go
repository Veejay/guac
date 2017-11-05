package images

import (
	"fmt"
	"html/template"
	"net/http"

	"cloud.google.com/go/storage"
	"golang.org/x/net/context"
)

type ip struct {
	Address string
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	fileName := request.URL.Path[len("/images/"):]
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	bucketName := "images-a-gogo.appspot.com"
	bucket := client.Bucket(bucketName)
	fmt.Fprintf(response, "<h4>%s</h4>", bucket.Attrs(ctx).Name)
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
