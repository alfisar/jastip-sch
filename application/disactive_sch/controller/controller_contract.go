package controller

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type DisactiveSCHContorllerContract interface {
	DisactiveSch(ctx context.Context, poolData *domain.ConfigSch)
}
