// Package runtime ...
package runtime

import (
	"context"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
)

type (
	Resources interface {
		Init(context.Context) error
		Watch(context.Context) error
		Stop()
		Release() error
	}

	Application struct {
		MainFunc  func(ctx context.Context) error
		Resources Resources

		err   error
		state int32
	}
)

const (
	appStateInit int32 = iota
	appStateRunning
	appStateHalt
	appStateShutdown
)

func (a *Application) Run() error {
	if a.MainFunc == nil {
		return ErrMainOmitted
	}

	if a.checkState(appStateInit, appStateRunning) {
		if err := a.init(); err != nil {
			a.err = err
		}

		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)
	}

	return nil
}

func (a *Application) init() error {
	return nil
}

func (a *Application) checkState(oldState, newState int32) bool {
	return atomic.CompareAndSwapInt32(&a.state, oldState, newState)
}
