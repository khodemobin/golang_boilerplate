package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pio/provider/internal/cache"
	"github.com/khodemobin/pio/provider/internal/domain"
)

type sampleRepo struct {
	db    *sqlx.DB
	cache cache.Cache
}

func NewSampleRepo(db *sqlx.DB, cache cache.Cache) domain.SampleRepository {
	return &sampleRepo{
		db:    db,
		cache: cache,
	}
}

func (r *sampleRepo) Sample(ctx context.Context) (*int64, error) {
	return nil, nil
}
