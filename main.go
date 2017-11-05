package images

import (
	"fmt"
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/blobstore"
	"google.golang.org/appengine/image"
)

type ip struct {
	Address string
}

func imageHandler(response http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	blobKey, err := blobstore.BlobKeyForFile(ctx, "/gs/images-a-gogo.appspot.com/lukaku.jpg")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	imageURL, err := image.ServingURL(ctx, blobKey, nil)
	if err != nil {
		http.Error(response, "An error occurred. Try again.", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(response, "Serving URL: %s", imageURL)
}

func rootHandler(response http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("index.html")
	t.Execute(response, ip{Address: request.RemoteAddr})
}

func init() {
	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/images/", imageHandler)
}
