package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"strings"
)

var db *sql.DB

func startServer() {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("MYSQL_USER"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
		os.Getenv("MYSQL_PORT"),
		os.Getenv("MYSQL_DATABASE"),
	))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	router := mux.NewRouter()

	// Define your API routes
	router.HandleFunc("/api/v1/auth/login", loginHandler).Methods("POST")
	router.HandleFunc("/api/v1/animals", animalsHandler).Methods("GET")
	router.HandleFunc("/api/v1/animal/{id}", animalDetailsHandler).Methods("GET")
	router.HandleFunc("/api/v1/animal/{id}", deleteAnimalHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/animals", animalsAddHandler).Methods("POST")

	// Use middleware to set CORS headers
	router.Use(corsMiddleware)
	// Add OPTIONS handlers for all routes
	router.PathPrefix("/api/v1/").Methods("OPTIONS").HandlerFunc(optionsHandler)


	log.Fatal(http.ListenAndServe(":8080", router))
}

// General OPTIONS handler
func optionsHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// CORS middleware to handle CORS preflight requests
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace with your frontend URL
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		// Handle OPTIONS request
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Proceed to the next handler
		next.ServeHTTP(w, r)
	})
}

type Animal struct {
	ID    int      `json:"id"`
	Name  string   `json:"name"`
	Tags  []string `json:"tags"`
	Owner string   `json:"owner"`
	Image string   `json:"image"`
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if creds.Username == "admin" && creds.Password == "password" {
		json.NewEncoder(w).Encode(map[string]bool{"success": true})
	} else {
		json.NewEncoder(w).Encode(map[string]bool{"success": false})
	}
}

func animalsHandler(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id, name, owner, tags, image FROM a379.animals")
	if err != nil {
		fmt.Printf("Error querying database: %v\n", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var animals []Animal

	for rows.Next() {
		var animal Animal
		var tags string

		err := rows.Scan(&animal.ID, &animal.Name, &animal.Owner, &tags, &animal.Image)
		if err != nil {
			fmt.Printf("Error scanning row: %v\n", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		// Split tags by commas
		tagStrings := strings.Split(tags, ",")
		animal.Tags = tagStrings
		animals = append(animals, animal)
	}
	json.NewEncoder(w).Encode(animals)
}

func animalsAddHandler(w http.ResponseWriter, r *http.Request) {
	var animal struct {
	    ID    int    `json:"id"`
		Name  string `json:"name"`
		Owner string `json:"owner"`
		Tags  string `json:"tags"`
		Image string `json:"image"`
	}

	// Decode incoming JSON
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
	    fmt.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the new animal into the database
	result, err := db.Exec("INSERT INTO animals (name, owner, tags, image) VALUES (?, ?, ?, ?)", animal.Name, animal.Owner, animal.Tags, animal.Image)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the ID of the newly inserted animal
	id, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the ID of the newly inserted animal
	json.NewEncoder(w).Encode(map[string]int{"id": int(id)})
}

func animalDetailsHandler(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    id := params["id"]

    var animal Animal
    var tagsString string
    err := db.QueryRow("SELECT id, name, tags, owner, image FROM animals WHERE id = ?", id).Scan(
       &animal.ID, &animal.Name, &tagsString, &animal.Owner, &animal.Image,
    )
    if err != nil {
       http.Error(w, err.Error(), http.StatusInternalServerError)
       return
    }

    // Split the tags string into a slice of strings
    animal.Tags = strings.Split(tagsString, ",")

    json.NewEncoder(w).Encode(animal)
}

func deleteAnimalHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	_, err := db.Exec("DELETE FROM animals WHERE id = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"result": "success"})
}

func main() {
	startServer()
}
