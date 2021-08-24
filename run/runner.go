package run

import (
	"context"
)

type Interface interface {
	Run(ctx context.Context) error
}
