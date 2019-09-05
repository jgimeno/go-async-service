Sometimes we want to have a service that runs several actions in a concurrent mode
and we solve this problem every time the same way.

With go async we can run the CyclicAsyncService and inject several actions that
will be executed in a cyclic fashion in an interval configured by the Config.

## Cyclic service

A cyclic service runs the actions you set in the array of actions and it is executed
until you call the Stop() method.

```
    // Every action must satisfy the service.AsyncAction interface.
    type action1 struct {
    } 

    func (a *action) Run(wg *sync.WaitGroup) {
        defer wg.Done() // The defer is MANDATORY

        .... // Do what you want
    }
    ... we can define more actions

	service := NewCyclicAsyncService(
		&Config{LoopSleepTime: time.Millisecond * 2}, // The time it sleeps to execute next iteration. (Can be zero, meaning that when the actions finish it starts again without sleeping)
		[]AsyncAction{action1, action2, action3}, // Here it goes all the actions
	)

    service.Start() // Check that start is a non-blocking function
    
    ... // we can do other things, or create a channel that listens
        // to signals and once received we can call the Stop method.

    <-exitChan
    service.Stop()
```

On every cycle, all the actions are executed asynchronously. Once all the actions finish
it starts again. It is ideal for a listener that performs several actions
on every iteration.