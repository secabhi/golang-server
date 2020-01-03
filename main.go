package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func home(w http.ResponseWriter, t *http.Request) {
	fmt.Fprintf(w, "home")
}

func getArticles(w http.ResponseWriter, t *http.Request) {
	articles := Articles{
		Article{Title: "Abhi", Desc: "This is test", Content: "this is content"},
	}
	fmt.Println("Endpoint hit - All Article")
	json.NewEncoder(w).Encode(articles)
}

func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", home)
	router.HandleFunc("/articles", getArticles)
	router.HandleFunc("/signupuser", SignUpUser).Methods("POST")
	router.HandleFunc("/loginuser", LoginUser).Methods("POST")
	router.HandleFunc("/adduser", AddUser).Methods("POST")
	router.HandleFunc("/users", AllUser).Methods("GET")
	router.HandleFunc("/removeuser", RemoveUser).Methods("DELETE")
	router.HandleFunc("/updateuser", UpdateUser).Methods("PUT")
	//router.Use(mux.CORSMethodMiddleware(router))
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {
	dbConnection()
	handleRequest()
	fmt.Println("Server started . . .")
}

// mongodb+srv://secabhi:<password>@secabhi-007-uxqsb.mongodb.net/test?retryWrites=true&w=majority
