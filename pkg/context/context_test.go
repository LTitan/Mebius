// Package context self context base context.CancelContext
package context

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_mcontext_WaitGoroutine(t *testing.T) {
	mctx := WithWaitGroup(context.Background())
	for i := 0; i < 100; i++ {
		index := i
		mctx.Go(func() {
			t.Logf("fake goroutine %d\n", index)
		})
	}
	mctx.WaitGoroutine()
	assert.NotNil(t, mctx, "context not nil")
}
