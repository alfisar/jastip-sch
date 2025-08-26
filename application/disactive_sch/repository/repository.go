package repository

import (
	"context"
	"fmt"

	"github.com/alfisar/jastip-import/helpers/errorhandler"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type disactiveSCH struct{}

func NewDisactiveSCH() *disactiveSCH {
	return &disactiveSCH{}
}

func (r *disactiveSCH) Update(ctx context.Context, conn *gorm.DB, where clause.Expr, update map[string]any) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
		return
	}

	err = conn.WithContext(ctx).Table("traveler_schedule").Where(where).Updates(update).Error
	if err != nil {
		err = fmt.Errorf("update treveler schedule error : %w", err)
		return
	}
	return
}
