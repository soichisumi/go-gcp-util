package spanner

import (
	"context"

	"cloud.google.com/go/spanner"
)

type Config struct {
	Project  string
	Instance string
	Database string
}

func NewClient(cfg Config) (*spanner.Client, error) {
	return spanner.NewClient(context.Background(), cfg.Database)
}
