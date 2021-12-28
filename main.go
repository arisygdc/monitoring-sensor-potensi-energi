package main

import (
	"context"
	"errors"
	"log"
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
	go AutomatedstatusRefresh(ctx, db)

	repo := repository.New(db)
	ctr := controller.New(repo)

	server := server.New(env, ctr)
	server.APIRoute()
	server.Run()
}

func AutomatedstatusRefresh(ctx context.Context, db database.DB) {
	var (
		refreshAt, _ = time.ParseDuration("35m")
		ErrNoRows    = errors.New("sql: no rows in result set")
	)
	log.Println("Running automated job")
	for {
		time.Sleep(1 * time.Minute)
		sensors, err := db.Queries.GetAllSensorOnStatus(ctx, true)
		if err != nil {
			if err != ErrNoRows {
				log.Println(err)
			}
			continue
		}

		for _, v := range sensors {
			lastUpdate := v.DitempatkanPada.(time.Time)
			if v.DibuatPada != nil {
				lastUpdate = v.DibuatPada.(time.Time)
			}
			log.Printf("sensor %v status online", v.Identity)
			lastUpdate = lastUpdate.UTC()

			log.Printf("now %v, last update %v", time.Now(), lastUpdate)
			if time.Now().After(lastUpdate.Add(refreshAt)) {
				err = db.Queries.UpdateStatusSensor(ctx, postgres.UpdateStatusSensorParams{
					Status: false,
					ID:     v.InfID,
				})
				if err != nil {
					log.Println(err)
					continue
				}

				log.Printf("sensor %v status changed to offline", v.Identity)
			}
			time.Sleep(1 * time.Second)
		}
	}
}
