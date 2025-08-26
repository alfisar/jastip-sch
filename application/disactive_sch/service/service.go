package service

import (
	"context"
	"fmt"
	"log"
	"scheduler/application/disactive_sch/repository"
	"time"

	"github.com/alfisar/jastip-import/domain"
	"github.com/alfisar/jastip-import/helpers/consts"
	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
)

type disactiveSCHService struct {
	repo repository.DisActiveRepositoryContract
}

func NewDisactiveSCHService(repo repository.DisActiveRepositoryContract) *disactiveSCHService {
	return &disactiveSCHService{
		repo: repo,
	}
}
func (s *disactiveSCHService) DisactiveSch(ctx context.Context, poolData *domain.ConfigSch) (err domain.ErrorData) {
	where := gorm.Expr("period_end < ?", time.Now().Truncate(24*time.Hour))
	updates := map[string]any{
		"status": 0,
	}

	errData := s.repo.Update(ctx, poolData.DBSql[consts.DBCore], where, updates)
	if errData != nil {
		message := fmt.Sprintf("Error update data on func DisactiveSch : %s", errData.Error())
		log.Println(message)

		if errData.Error() == errorhandler.ErrMsgConnEmpty {
			err = errorhandler.ErrInternal(errorhandler.ErrCodeConnection, errData)
		} else {
			errorhandler.ErrUpdateData(fmt.Errorf(message))
		}
	}

	return
}
