package main

type scenarioFn func()

type Runner struct {
	scenarios []scenarioFn
	para      int
	done      chan struct{}
}

func NewRunner(para int, funcs []scenarioFn) *Runner {
	return &Runner{
		scenarios: funcs,
		done:      make(chan struct{}, para),
	}
}

func (r *Runner) start(ch chan *Runner) {
	para := cap(r.done)
	for i := 0; i < para; i++ {
		ch <- r
	}
	for range r.done {
		ch <- r
	}
}
func (r *Runner) run() {
	for _, fn := range r.scenarios {
		fn()
	}
	r.done <- struct{}{}
}
