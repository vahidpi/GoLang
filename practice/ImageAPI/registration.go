package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ImageReg struct {
	Sha256       string   `json:"sha256,omitempty"`
	Size int   `json:"size,omitempty"`
	Chunk_size int   `json:"chunk_size,omitempty"`
}

func imageRegistration(w http.ResponseWriter, req *http.Request) {
	var image ImageReg
	
	err := json.NewDecoder(req.Body).Decode(&image)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
		httpResponse(w,http.StatusBadRequest,"Malformed request") // 400    
        return
    }


	imageExist, err := isImageExist(image.Sha256)
	if err != nil {
		fmt.Println("isImageExist error log: ",err)
		return 
	}
	if  !imageExist {
        err:=registerImage(image)
		if err != nil {
			httpResponse(w,http.StatusUnsupportedMediaType,"Unsupported payload format ")  // 415 
		} else {
			httpResponse(w,http.StatusCreated,"Image successfully registered")     
		}
    } else {
		httpResponse(w,http.StatusConflict,"Image already exists")  // 409
	}
}

func isImageExist(Sha256 string) (bool,error) {
	// Query db, if it was exist return true
	fmt.Println("Image not exist")
	return false,nil
}

func registerImage(img ImageReg) error {
	// Register image
	fmt.Println("Image registed ",img.Sha256,img.Size,img.Chunk_size)
	return nil	// if error happend return error
}