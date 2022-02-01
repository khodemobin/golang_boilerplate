package domain

import "context"

type Sample struct {
	ID     *int    `db:"id"`
	Params *string `db:"params" json:"p,omitempty"`
}

type SampleService interface {
	Sample(ctx context.Context) (string, error)
}

type SampleRepository interface {
	Sample(ctx context.Context) (*int64, error)
}
