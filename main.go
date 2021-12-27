package main

import (
	"log"
	"monitoring-potensi-energi/config"
	"monitoring-potensi-energi/controller"
	"monitoring-potensi-energi/database"
	"monitoring-potensi-energi/repository"
	"monitoring-potensi-energi/server"
)

func main() {
	env, err := config.NewEnv(".")
	if err != nil {
		log.Fatalln(err)
	}

	db, err := database.NewPostgres(env)
	if err != nil {
		log.Fatalln(err)
	}

	repo := repository.New(db)
	ctr := controller.New(repo)

	server := server.New(env, ctr)
	server.APIRoute()
	server.Run()
}
