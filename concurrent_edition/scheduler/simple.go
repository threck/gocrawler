package scheduler

import "gocrawler/concurrent_edition/engine"

type SimpleScheduler struct {
	Request chan engine.Request
}

func (s *SimpleScheduler) WorkChan() chan engine.Request {
	return s.Request
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {
}

func (s *SimpleScheduler) Submit(r engine.Request) {
	go func() { s.Request <- r }()
	//s.Request <- r
}

func (s *SimpleScheduler) Run() {
	s.Request = make(chan engine.Request)
}
