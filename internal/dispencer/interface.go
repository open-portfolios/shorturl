package dispencer

import (
	"context"
)

type Interface interface {
	Dispence(ctx context.Context) (uint64, error)
}
