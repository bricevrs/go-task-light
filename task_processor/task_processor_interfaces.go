package processor

import (
	"context"
)

type TaskProcessorQuerier interface {
	Execute(ctx context.Context) error
	ExecuteBatch(ctx context.Context) error
}
