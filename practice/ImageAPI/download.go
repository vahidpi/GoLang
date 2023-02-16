package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)


func downloadImage(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
    sha, ok := vars["sha"]
	if ok{
		find,err:=searchImage(sha)
		if err != nil { 
			httpResponse(w,http.StatusBadRequest,"Bad request")    
			return 
		} else {
			if find {
				httpResponse(w,http.StatusOK,"Image successfully downloaded")
			}else{          
				httpResponse(w,http.StatusNotFound,"Image not found")
			}
		}
	}
}


func searchImage(sha string) (bool,error) {
	// find image
	fmt.Println("Image Found ",sha)
	return true,nil
}