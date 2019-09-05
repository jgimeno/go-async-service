package service

import (
	"sync"
	"time"
)

//NewCyclicService construct a service that runs an action, encapsulated in ServiceCycle in a loop
//until calling Stop()
func NewCyclicAsyncService(
	config *Config,
	actions []AsyncAction,
) StoppableService {
	return &cyclicAsyncService{
		config:  config,
		actions: actions,
		wg:      &sync.WaitGroup{},
	}
}

type cyclicAsyncService struct {
	config *Config
	wg     *sync.WaitGroup

	actions []AsyncAction

	pendingStop bool
}

func (c *cyclicAsyncService) Start() error {
	c.pendingStop = false

	go c.startCycle()

	return nil
}

func (c *cyclicAsyncService) startCycle() {
	for {
		c.wg.Add(len(c.actions))

		for _, action := range c.actions {
			go action.Run(c.wg)
		}

		c.wg.Wait()
		if c.pendingStop {
			return
		}

		time.Sleep(c.config.LoopSleepTime)
	}
}

func (c *cyclicAsyncService) Stop() error {
	c.pendingStop = true
	return nil
}

