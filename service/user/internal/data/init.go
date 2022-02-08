package data

import (
	"context"
	"gorm.io/gorm"
	"user/internal/biz"

	"github.com/pkg/errors"
)

type DbInitializer struct {
	Ds *gorm.DB
}

func (p *DbInitializer) Name() string {
	return "db_initializer"
}

func (p *DbInitializer) IsNeedInit(ctx context.Context) (bool, error) {
	return true, nil
}

// Initialize AutoMigrate自动建表
func (p *DbInitializer) Initialize(ctx context.Context) error {
	err := p.Ds.AutoMigrate(
		&biz.User{},
	)
	return errors.WithStack(err)
}
