package main

import (
	"log"
	"github.com/timut2/Practice3/pkg/handler"
	"github.com/timut2/Practice3"
	"github.com/spf13/viper"
	"github.com/timu2/Practice3/repository"
	"github.com/timu2/Practice3/service"
)
func main(){
	if err:= initConfig(); err!= nil{
		log.Fatalf("error init config",err)
	}
	repos:= repository.NewRepository()
	services:=service.NewService(repos)
	handlers:= new(services)
	s:= new(server.Server)
	if err:= s.Run(viper.GetString("port"), handlers.InitRoutes()); err!= nil{
		log.Fatalf("error occured ",err)
	}
}


func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadConfig()
}