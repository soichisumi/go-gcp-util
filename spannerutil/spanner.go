package spannerutil

import (
	"context"
	"fmt"

	"cloud.google.com/go/spanner"
)

type Config struct {
	Project  string
	Instance string
	Database string
}

func NewClient(cfg Config) (*spanner.Client, error) {
	return spanner.NewClient(context.Background(), fmt.Sprintf("projects/%s/instances/%s/databases/%s", cfg.Project, cfg.Instance, cfg.Database))
}
