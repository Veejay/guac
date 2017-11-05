package main

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"

	"golang.org/x/net/context"
)

type ip struct {
	Address string
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	ctx := context.Background()
	blobKey, _ := blobstore.BlobKeyForFile(ctx, "/gs/images-a-gogo.appspot.com/lukaku.jpg")
	url, _ := image.ServingURL(ctx, blobKey, &image.ServingURLOptions{Secure: true, Size: 800, Crop: false})
	fmt.Fprintf(response, "Serving URL: %s", url)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(response, ip{Address: request.RemoteAddr})
}

func main() {
	http.HandleFunc("/images/", imageHandler)
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8080", nil)
}
