package controller

import "monitoring-potensi-energi/repository"

type Controller struct {
	Repo repository.Repository
}

func New(repo repository.Repository) (controller Controller) {
	controller = Controller{
		Repo: repo,
	}
	return
}
