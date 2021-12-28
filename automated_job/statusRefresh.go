package automated_job

import (
	"context"
	"log"
	"monitoring-potensi-energi/database/postgres"
	"time"
)

type StatusRefresh struct {
	interval time.Duration
	queries  *postgres.Queries
}

func NewStatusRefresh(interval time.Duration, queries *postgres.Queries) (automate StatusRefresh) {
	automate = StatusRefresh{
		interval: interval,
		queries:  queries,
	}
	return
}

func (asr StatusRefresh) GetData(ctx context.Context) ([]postgres.GetAllSensorOnStatusRow, error) {
	return asr.queries.GetAllSensorOnStatus(ctx, true)
}

func (asr StatusRefresh) RunJob(ctx context.Context, sensorsData []postgres.GetAllSensorOnStatusRow) {
	for _, v := range sensorsData {
		lastUpdate := v.DitempatkanPada.(time.Time)
		if v.DibuatPada != nil {
			lastUpdate = v.DibuatPada.(time.Time)
		}

		log.Printf("sensor %v status online", v.Identity)
		lastUpdate = lastUpdate.UTC()

		log.Printf("now %v, last update %v", time.Now(), lastUpdate)
		if time.Now().After(lastUpdate.Add(asr.interval)) {
			err := asr.queries.UpdateStatusSensor(ctx, postgres.UpdateStatusSensorParams{
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
