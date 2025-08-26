package repository

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type DisActiveRepositoryContract interface {
	Update(ctx context.Context, conn *gorm.DB, where clause.Expr, update map[string]any) (err error)
}
