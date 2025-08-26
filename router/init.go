package router

import (
	"scheduler/application/disactive_sch/controller"
	"scheduler/application/disactive_sch/repository"
	"scheduler/application/disactive_sch/service"
)

func DisactiveSCHInit() controller.DisactiveSCHContorllerContract {
	repo := repository.NewDisactiveSCH()
	serv := service.NewDisactiveSCHService(repo)
	controll := controller.NewDisactiveSCHController(serv)
	return controll
}
