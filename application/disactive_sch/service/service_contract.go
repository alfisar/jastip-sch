package service

import (
	"context"

	"github.com/alfisar/jastip-import/domain"
)

type DisactiveSCHServiceContract interface {
	DisactiveSch(ctx context.Context, poolData *domain.ConfigSch) (err domain.ErrorData)
}
