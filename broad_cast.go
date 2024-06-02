package lssec_go

import "context"

type BroadCast struct {
	broadcast  chan interface{}
	Register   chan chan interface{}
	Unregister chan chan interface{}
	child      map[chan interface{}]bool
}

func NewBroadCast() *BroadCast {
	return &BroadCast{
		broadcast:  make(chan interface{}, 10),
		Register:   make(chan chan interface{}, 2),
		Unregister: make(chan chan interface{}, 1),
		child:      map[chan interface{}]bool{},
	}
}

// Start event engine start function need goroutine
func (e *BroadCast) Start(ctx context.Context) {
Done:
	for {
		select {
		case <-ctx.Done():
			break Done
		case evt := <-e.broadcast:
			for c := range e.child {
				c <- evt
			}
		case cli := <-e.Register:

			e.child[cli] = true
		case cli := <-e.Unregister:
			delete(e.child, cli)
		}
	}
}

func (e *BroadCast) BroadCast(evt interface{}) {
	e.broadcast <- evt
}
