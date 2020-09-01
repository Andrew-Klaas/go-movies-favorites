package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Favorite struct {
	UserName string
	Movies   []Movie
}
type FavoriteRecord struct {
	UserName string `json:"UserName"`
	Title    string `json:"Title"`
}
type Movie struct {
	MovieID  string `json:"MovieID"`
	Title    string `json:"Title"`
	Director string `json:"Director"`
	Price    string `json:"Price"`
}

//CustomerFavorites is a Map of Usernames to their favorite movies
var CustomerFavorites = map[string][]string{}

func main() {
	http.HandleFunc("/getFavorite", GetFavorites)
	http.HandleFunc("/addtoFavorite", AddToFavorite)
	http.HandleFunc("/Favorites", MethodFavorites)
	http.HandleFunc("/deletefromFavorite", DeleteFromFavorite)
	http.ListenAndServe(":8081", nil)
}

//GetFavorites returns Favorite for User in request
func GetFavorites(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received GET Request for user Favorites\n")
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}
	un := r.FormValue("username")
	if un == "" {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)

	}

	//requested customer Favorite
	rcc := CustomerFavorites[r.FormValue("username")]

	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(rcc)
	if err != nil {
		log.Println(err)
	}

	/*
			slcD := []string{"apple", "peach", "pear"}
		    slcB, _ := json.Marshal(slcD)
			fmt.Println(string(slcB))
	*/
}

//AddToFavorite adds a movie to a users Favorite
func AddToFavorite(w http.ResponseWriter, r *http.Request) {
	log.Printf("Received POST Request to add item to user's Favorite\n")
	fmt.Printf("r.Body: %v\n", r.Body)
	if r.Method != http.MethodPost {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
	}

	cr := FavoriteRecord{}
	err := json.NewDecoder(r.Body).Decode(&cr)
	if err != nil {
		fmt.Printf("err %v\n", err)
	}

	u := cr.UserName
	t := cr.Title

	fmt.Printf("username: %v, title %v\n", u, t)
	if _, ok := CustomerFavorites[u]; !ok {
		movieTitles := make([]string, 0)
		movieTitles = append(movieTitles, t)
		CustomerFavorites[u] = movieTitles
	} else {
		CustomerFavorites[u] = append(CustomerFavorites[u], t)
	}
	fmt.Printf("CustomerFavorites[%v]: %v\n", u, CustomerFavorites[u])

	// Write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated) // 201
}

//DeleteFromFavorite ...
func DeleteFromFavorite(w http.ResponseWriter, r *http.Request) {

}

//MethodFavorites is general purpose for the demo
//TODO
func MethodFavorites(w http.ResponseWriter, r *http.Request) {

}
