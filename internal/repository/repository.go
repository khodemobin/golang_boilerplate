package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/khodemobin/pio/provider/internal/cache"
	"github.com/khodemobin/pio/provider/internal/domain"
)

type Repository struct {
	Sample domain.SampleRepository
}

func NewRepository(db *sqlx.DB, cache cache.Cache) *Repository {
	s := NewSampleRepo(db, cache)
	return &Repository{
		Sample: s,
	}
}
