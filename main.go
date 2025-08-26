package main

import (
	"context"
	"flag"
	"log"
	"scheduler/router"

	"github.com/alfisar/jastip-import/config"
	"github.com/alfisar/jastip-import/domain"
)

func main() {
	config.InitSch()
	poolData := domain.DataPool.Get().(*domain.ConfigSch)
	jobs := flag.String("job", "", "Specify the job to run")

	flag.Parse()

	if *jobs == "" {
		log.Fatalln("Tidak ada arguments")
	}
	ctx := context.Background()
	if *jobs == "disactive_sch" {
		controller := router.DisactiveSCHInit()
		controller.DisactiveSch(ctx, poolData)
	}

}
