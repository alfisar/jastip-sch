package controller

import (
	"context"
	"fmt"
	"scheduler/application/disactive_sch/service"

	"github.com/alfisar/jastip-import/domain"
)

type disactiveSCHController struct {
	serv service.DisactiveSCHServiceContract
}

func NewDisactiveSCHController(serv service.DisactiveSCHServiceContract) *disactiveSCHController {
	return &disactiveSCHController{
		serv: serv,
	}
}

func (c *disactiveSCHController) DisactiveSch(ctx context.Context, poolData *domain.ConfigSch) {
	err := c.serv.DisactiveSch(ctx, poolData)
	if err.Code != 0 {
		errors := err.Errors.(error)
		fmt.Println(err.Message + " " + errors.Error())
	}

	fmt.Println("Successfully running DisactiveSch")
}
