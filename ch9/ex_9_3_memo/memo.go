package memo

// Func is the type of the function to memoize.
type Func func(key string, done <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{}
}

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	done     <-chan struct{}
	response chan<- result
}

type Memo struct {
	requests, cancels chan request
}

// New returns a memoization of f. Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{make(chan request), make(chan request)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, done <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	req := request{key, done, response}
	memo.requests <- req
	res := <-response
	select {
	case <-done:
		memo.cancels <- req
	default:
	}
	return res.value, res.err
}

func (memo *Memo) Close() { close(memo.requests) }

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
Loop:
	for {
	Cancel:
		for {
			select {
			case req := <-memo.cancels:
				delete(cache, req.key)
			default:
				break Cancel
			}
		}
		select {
		case req := <-memo.cancels:
			delete(cache, req.key)
			continue Loop
		case req := <-memo.requests:
			e := cache[req.key]
			if e == nil {
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.done)
			}
			go e.deliver(req.response)
		}
	}
}

func (e *entry) call(f Func, key string, done <-chan struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.ready)
}

func (e *entry) deliver(response chan<- result) {
	<-e.ready
	response <- e.res
}
