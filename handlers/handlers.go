package handlers

import (
	"log"
	"net/http"
	"os"
	"github.com/PSTPglobal/tweet-go/middlewares"
	"github.com/PSTPglobal/tweet-go/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Manejadores()  {
	router := mux.NewRouter()
	/*Endpoints*/
	router.HandleFunc("/registro", middlewares.ChequeoBD(routes.Registro)).Methods("POST")
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":" + PORT, handler))
}