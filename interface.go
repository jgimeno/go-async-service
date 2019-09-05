package service

import "sync"

type Service interface {
	Start() error
}

type StoppableService interface {
	Service
	Stop() error
}

// Action represents the action that a cyclicService executes on every loop.
type AsyncAction interface {
	Run(wg *sync.WaitGroup)
}
