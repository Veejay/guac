package main

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
	"google.golang.org/appengine/log"
)

type ip struct {
	Address string
}

func imageHandler(response http.ResponseWriter, request *http.Request) {
	ctx := request.Context()
	blobKey, err := blobstore.BlobKeyForFile(ctx, "/gs/images-a-gogo.appspot.com/lukaku.jpg")
	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		http.Error(response, "An error occurred. Try again.", http.StatusInternalServerError)
		return
	}
	imageURL, err := image.ServingURL(ctx, blobKey, &image.ServingURLOptions{Secure: true, Size: 800, Crop: false})
	if err != nil {
		log.Errorf(ctx, "could not put into datastore: %v", err)
		http.Error(response, "An error occurred. Try again.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(response, "Serving URL: %s / Blob key is %s", imageURL, blobKey)
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
