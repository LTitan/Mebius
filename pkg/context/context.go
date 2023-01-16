// Package context self context base context.CancelContext
package context

import (
	"context"
	"sync"
	"time"
)

type (
	_waitGroupKey struct{}
)

type mcontext struct {
	Ctx context.Context
}

// MContext .
type MContext interface {
	WithTimeout(timeout time.Duration) (MContext, context.CancelFunc)
	WithCancel() (MContext, context.CancelFunc)

	// goroutine
	GO(exec func())
	WaitGoroutine()

	// basic
	Done() <-chan struct{}
	Deadline() (time.Time, bool)
	Value(key interface{}) interface{}
	Err() error
}

func warp(ctx context.Context) MContext {
	return &mcontext{
		Ctx: ctx,
	}
}

func (m *mcontext) WithTimeout(timeout time.Duration) (MContext, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(m, timeout)
	return warp(ctx), cancel
}

func (m *mcontext) WithCancel() (MContext, context.CancelFunc) {
	ctx, cancel := context.WithCancel(m)
	return warp(ctx), cancel
}

// WithValue .
func WithValue(parent context.Context, key, value interface{}) MContext {
	return warp(context.WithValue(parent, _waitGroupKey{}, &sync.WaitGroup{}))
}

// WithWaitGroup .
func WithWaitGroup(ctx context.Context) MContext {
	ctx = context.WithValue(ctx, _waitGroupKey{}, &sync.WaitGroup{})
	return warp(ctx)
}

func (m *mcontext) Done() <-chan struct{} {
	return m.Ctx.Done()
}

func (m *mcontext) Deadline() (time.Time, bool) {
	return m.Ctx.Deadline()
}

func (m *mcontext) Value(key interface{}) interface{} {
	return m.Ctx.Value(key)
}

func (m *mcontext) Err() error {
	return m.Ctx.Err()
}

func (m *mcontext) GO(exec func()) {
	wg := m.Value(_waitGroupKey{}).(*sync.WaitGroup)
	go func() {
		wg.Add(1)
		defer wg.Done()
		// defer func() {
		// 	// auto recover
		// 	if err := recover(); err != nil {
		// 		klog.Errorln(err)
		// 	}
		// }()
		exec()
	}()
}

func (m *mcontext) WaitGoroutine() {
	wg := m.Value(_waitGroupKey{}).(*sync.WaitGroup)
	wg.Wait()
}
