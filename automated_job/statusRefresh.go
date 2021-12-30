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
	var lastUpdate time.Time
	for _, v := range sensorsData {

		if v.TerakhirUpdate != nil {
			lastUpdate = v.TerakhirUpdate.(time.Time)
		} else {
			lastUpdate = v.DitempatkanPada
		}
		lastUpdate = lastUpdate.UTC()

		if time.Now().After(lastUpdate.Add(asr.interval)) {
			err := asr.queries.UpdateStatusSensor(ctx, postgres.UpdateStatusSensorParams{
				Status: false,
				ID:     v.ID,
			})
			if err != nil {
				log.Println(err)
				continue
			}

			log.Printf("sensor %v status changed to offline", v.ID)
		}
		time.Sleep(1 * time.Second)
	}
}
