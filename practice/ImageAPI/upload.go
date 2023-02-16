package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Chunk struct {
	Id       int   `json:"id,omitempty"`
	Size int   `json:"size,omitempty"`
	Data string   `json:"data,omitempty"`
}


func uploadImage(w http.ResponseWriter, r *http.Request)  {
	var chunk Chunk

	err := json.NewDecoder(r.Body).Decode(&chunk)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		httpResponse(w,http.StatusBadRequest,"Malformed request") // 400
        return
    }

	vars := mux.Vars(r)
    sha, ok := vars["sha"]
	if ok{
		if findImage(sha){
			if isChunkExist(chunk.Id){
					httpResponse(w,http.StatusConflict,"Chunk already exists") // 409
				} else {
					if createChunk(chunk) !=nil {
						httpResponse(w,http.StatusCreated,"Chunk successfully uploaded") // 201
					}
				}
		}
	} else {
		httpResponse(w,http.StatusNotFound,"Image not found") // 404
	}
}

func isChunkExist(id int) bool {
	// Check, if chunk was exist return true
	fmt.Println("Chunk is not exist")
	return false
}

func createChunk(chunk Chunk) error {
	fmt.Println("Chunk created")
	return nil
}

func findImage(sha string) bool {
	// if not find retrun false
	fmt.Println("Image found")
	return true
}