package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"simple-image-gallery/services"

	"github.com/gorilla/mux"
)

type ImageRequest struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Text   string `json:"text"`
}

func GenerateImage(w http.ResponseWriter, r *http.Request) {
	var req ImageRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	imageURL, err := services.GenerateAndUploadImage(req.Width, req.Height, req.Text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{
		"url": imageURL,
	})
}

func GetImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["id"]

	contents, contentType, err := services.GetImageContents(imageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Set appropriate headers
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("Cache-Control", "public,max-age=3600")

	w.Write(contents)
}

func StreamImage(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	imageId := vars["id"]

	reader, contentType, err := services.GetImageStream(imageId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	defer reader.Close()

	// Set headers
	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Content-Disposition", "inline")
	w.Header().Set("Cache-Control", "public,max-age=3600")

	// Stream the contents
	_, err = io.Copy(w, reader)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
