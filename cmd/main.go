package main

import (
	"log"
	"github.com/timut2/project2/pkg/handler"
)
func main(){
	handlers:= new(handler.Handler)

	s:= new(server.Server)
	if err:= s.Run("8080", handlers.InitRoutes()); err!= nil{
		log.Fatalf("error occured ",err)
	}
}