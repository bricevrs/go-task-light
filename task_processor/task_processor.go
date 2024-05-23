package processor

import (
	"context"
)

/*

The goal of this file is to define what every task processor should have in common.


*/

type TaskProcessor interface {
	ExecTask(ctx context.Context, arg ...interface{}) error
}
