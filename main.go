package main

import (
	"context"
	"log"
	"monitoring-potensi-energi/automated_job"
	"monitoring-potensi-energi/config"
	"monitoring-potensi-energi/controller"
	"monitoring-potensi-energi/database"
	"monitoring-potensi-energi/database/postgres"
	"monitoring-potensi-energi/repository"
	"monitoring-potensi-energi/server"
	"time"
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

	ctx := context.Background()
	go AutomatedJobs(ctx, db.Queries)

	repo := repository.New(db)
	ctr := controller.New(repo)

	server := server.New(env, ctr)
	server.APIRoute()
	server.Run()
}

func AutomatedJobs(ctx context.Context, db *postgres.Queries) {
	var (
		statusRefreshInterval, _ = time.ParseDuration("35m")
	)

	log.Println("Running automated job")
	statusRefresh := automated_job.NewStatusRefresh(statusRefreshInterval, db)

	for {
		time.Sleep(1 * time.Second)
		sensors, err := statusRefresh.GetData(ctx)
		if err != nil {
			continue
		}

		statusRefresh.RunJob(ctx, sensors)
	}
}
