package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	dbHost = "localhost"
	dbPort = "3306"
	dbUser = "user"
	dbPass = "password"
	dbName = "ascii_art"
)

type image struct {
	SHA256    string `json:"sha256"`
	Size      int    `json:"size"`
	ChunkSize int    `json:"chunk_size"`
	Chunks    []chunk
}

type chunk struct {
	ID   int    `json:"id"`
	Size int    `json:"size"`
	Data string `json:"data"`
}

func main() {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/image", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			registerImage(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	http.HandleFunc("/image/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			uploadChunk(db, w, r)
		case http.MethodGet:
			downloadImage(db, w, r)
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		}
	})

	log.Fatal(http.ListenAndServe(":4444", nil))
}

func registerImage(db *sql.DB, w http.ResponseWriter, r *http.Request) (int64, error) {
	// Begin transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}

	// Insert image metadata into database
	result, err := tx.Exec("INSERT INTO images (sha256, size, chunk_size) VALUES (?, ?, ?)", image.Sha256, image.Size, image.ChunkSize)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Get ID of inserted image
	imageID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	// Insert image chunks into database
	for i, chunk := range image.Chunks {
		_, err = tx.Exec("INSERT INTO chunks (image_id, index, data) VALUES (?, ?, ?)", imageID, i, chunk)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	// Commit transaction
	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return imageID, nil
}

func uploadChunk(db *sql.DB,w http.ResponseWriter, r *http.Request) {
	// Get the parameters from the query string
	query := r.URL.Query()
	sha256 := query.Get("sha256")
	index, err := strconv.Atoi(query.Get("index"))
	if err != nil {
		http.Error(w, "Invalid chunk index", http.StatusBadRequest)
		return
	}
	chunkSize, err := strconv.Atoi(query.Get("chunkSize"))
	if err != nil {
		http.Error(w, "Invalid chunk size", http.StatusBadRequest)
		return
	}

	// Read the chunk data from the request body
	chunkData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read chunk data", http.StatusInternalServerError)
		return
	}

	// Check if the chunk already exists in the database
	rows, err := db.Query("SELECT COUNT(*) FROM chunks WHERE sha256 = ? AND index = ?", sha256, index)
	if err != nil {
		http.Error(w, "Failed to query database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var count int
	if rows.Next() {
		if err := rows.Scan(&count); err != nil {
			http.Error(w, "Failed to scan result", http.StatusInternalServerError)
			return
		}
	}

	if count > 0 {
		// Chunk already exists, do nothing
		return
	}

	// Insert the chunk data into the database
	result, err := db.Exec("INSERT INTO chunks (sha256, index, data) VALUES (?, ?, ?)", sha256, index, chunkData)
	if err != nil {
		http.Error(w, "Failed to insert chunk into database", http.StatusInternalServerError)
		return
	}

	// Check if the image is complete
	if index == chunkSize-1 {
		// Calculate the total size of the image
		rows, err := db.Query("SELECT SUM(LENGTH(data)) FROM chunks WHERE sha256 = ?", sha256)
		if err != nil {
			http.Error(w, "Failed to query database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var size int
		if rows.Next() {
			if err := rows.Scan(&size); err != nil {
				http.Error(w, "Failed to scan result", http.StatusInternalServerError)
				return
			}
		}

		// Update the image size in the database
		_, err = db.Exec("UPDATE images SET size = ? WHERE sha256 = ?", size, sha256)
		if err != nil {
			http.Error(w, "Failed to update image size in database", http.StatusInternalServerError)
			return
		}
	}

	// Return success status
	w.WriteHeader(http.StatusOK)
}
func (s *server) downloadImage(stream pb.ImageService_DownloadImageServer) error {
	// read the image metadata (SHA256, size, and number of chunks)
	image, err := s.readImage(stream)
	if err != nil {
		return err
	}

	// check if the image exists in the database
	exists, err := s.checkImageExists(image.Sha256)
	if err != nil {
		return err
	}
	if !exists {
		return status.Errorf(codes.NotFound, "Image with SHA256 %s not found", image.Sha256)
	}

	// send the image metadata to the client
	if err := stream.Send(&pb.DownloadImageResponse{Image: image}); err != nil {
		return err
	}

	// download the chunks one by one and send them to the client
	for i := uint64(0); i < image.Chunks; i++ {
		chunk, err := s.getChunk(image.Sha256, i)
		if err != nil {
			return err
		}
		if err := stream.Send(&pb.DownloadImageResponse{Chunk: chunk}); err != nil {
			return err
		}
	}

	return nil
}

func (s *server) GetChunks(sha256 string) (uint64, error) {
	row := s.db.QueryRow("SELECT chunk_size FROM images WHERE sha256 = ?", sha256)
	var chunkSize uint64
	if err := row.Scan(&chunkSize); err != nil {
		if err == sql.ErrNoRows {
			return 0, status.Errorf(codes.NotFound, "Image with SHA256 %s not found", sha256)
		}
		return 0, err
	}
	return chunkSize, nil
}

func getChunkSize(size int64) int64 {
    const maxChunkSize int64 = 1024 * 1024 * 2 // 2 MB
    chunkSize := size / 100
    if chunkSize > maxChunkSize {
        return maxChunkSize
    }
    if chunkSize < 1024 {
        return 1024
    }
    return chunkSize
}

