package engine

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item // 增加 itemSaver
	Processor   Processor
}

type Processor func(re Request) (ParseResult, error)

type Scheduler interface {
	ReadyNotifier
	Submit(request Request)
	WorkChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			result, err := e.Processor(<-in)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	e.Scheduler.Run()

	out := make(chan ParseResult)
	for i := 0; i < e.WorkerCount; i++ {
		e.createWorker(e.Scheduler.WorkChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.Url) {
			//log.Printf("Duplicate request: %s", r.Url)
			continue
		}
		e.Scheduler.Submit(r)
	}

	for {
		parseResults := <-out

		// URL dedup
		for _, v := range parseResults.Requests {
			if isDuplicate(v.Url) {
				//log.Printf("Duplicate request: %s", v.Url)
				continue
			}
			e.Scheduler.Submit(v)
		}

		for _, item := range parseResults.Items {
			it := item
			go func() { e.ItemChan <- it }()
		}
	}

}

var visitedUrl = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrl[url] {
		return true
	}
	visitedUrl[url] = true
	return false
}
