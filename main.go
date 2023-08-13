package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// No utilizaremos una base de datos. Si no que utilizaremos structs y slices
type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
}

var movies []Movie

func main() {
	r := mux.NewRouter()

	// Añadiendo peliculas manualmente para corroborar que funciona el request
	movies = append(movies, Movie{
		ID:    "1",
		ISBN:  "123456",
		Title: "Movie One",
		Director: &Director{
			Name:    "John",
			Surname: "Salchichon",
		},
	})

	movies = append(movies, Movie{
		ID:    "2",
		ISBN:  "12340101",
		Title: "Movie Two",
		Director: &Director{
			Name:    "Juan",
			Surname: "Carlos",
		},
	})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getMovies(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	/*
		El operador _ = se utiliza para ignorar el valor de retorno del método .Decode().
		En este caso, se está descartando el valor de retorno que puede ser un error (error) si ocurre
		algún problema durante la decodificación del JSON.
		Al utilizar _ =, se evita la necesidad de manejar explícitamente el error en este punto.
	*/
	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, req *http.Request) {
	// Set header
	w.Header().Set("Content-Type", "application/json")

	// Obtengo parametros del request
	params := mux.Vars(req)

	// Itero el slice de peliculas
	for index, item := range movies {
		// Si el id del request es igual de la pelicula almacenada entonces elimina la pelicula y agrega una nueva pelicula
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie

			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}
