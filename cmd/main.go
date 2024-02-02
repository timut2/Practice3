package main

import (
	"log"
	"github.com/timut2/Practice3/pkg/handler"
	"github.com/timut2/Practice3"
	"github.com/spf13/viper"
	"github.com/timut2/Practice3/pkg/repository"
	"github.com/timut2/Practice3/pkg/service"
)
func main(){
	if err:= initConfig(); err!= nil{
		log.Fatalf("error initialization config",err)
	}
	repos:= repository.NewRepository()
	services:=service.NewService(repos)
	handlers:= handler.(services)
	s:= new(server.Server)
	if err:= s.Run(viper.GetString("port"), handlers.InitRoutes()); err!= nil{
		log.Fatalf("error occured while running  http server",err)
	}
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}