package main

import (
	"log"
	"github.com/PSTPglobal/tweet-go/handlers"
	"github.com/PSTPglobal/tweet-go/bd"
)

func main()  {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}
	handlers.Manejadores()
}