package service

import (
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestWeCanAsyncServiceCallWithSeveralActionsConcurrently(t *testing.T) {
	action1 := &action{}
	action2 := &action{}
	action3 := &action{}

	s := NewCyclicAsyncService(
		&Config{LoopSleepTime: time.Millisecond * 2},
		[]AsyncAction{action1, action2, action3},
	)

	s.Start()

	time.Sleep(time.Millisecond)

	assert.True(t, action1.isDone())
	assert.True(t, action2.isDone())
	assert.True(t, action3.isDone())
	s.Stop()
}

type action struct {
	done bool
}

func (a *action) Run(wg *sync.WaitGroup) {
	defer wg.Done()

	a.done = true
}

func (a *action) isDone() bool {
	return a.done
}
