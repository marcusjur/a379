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

	router.HandleFunc("/api/v1/auth/login", loginHandler).Methods("POST")
	router.HandleFunc("/api/v1/animals", animalsHandler).Methods("GET")
	router.HandleFunc("/api/v1/animal/{id}", animalDetailsHandler).Methods("GET")
	router.HandleFunc("/api/v1/animal/{id}", deleteAnimalHandler).Methods("DELETE")
	router.HandleFunc("/api/v1/animals", animalsAddHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
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
	// New code for POST method
	var animal Animal
	err := json.NewDecoder(r.Body).Decode(&animal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// TODO:Perform necessary validation on the animal object

	// Insert the new animal into the database
	result, err := db.Exec("INSERT INTO animals (name, owner, image) VALUES (?, ?, ?)", animal.Name, animal.Owner, animal.Image)
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
	err := db.QueryRow("SELECT id, name, owner, image FROM animals WHERE id = ?", id).Scan(
		&animal.ID, &animal.Name, &animal.Owner, &animal.Image,
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	animal.Tags = []string{"dog", "pet"} // Example tags

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
