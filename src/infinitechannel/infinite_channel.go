package infinitechannel

import (
	"github.com/Juliusan/go.infinitechannel/src/queue"
)

// InfiniteChannel provides an infinite buffer between the input and the output.
type InfiniteChannel struct {
	input  chan interface{}
	output chan interface{}
	length chan int
	buffer queue.Queue
}

func NewInfiniteChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewSimpleQueue()}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteDefaultPriorityChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewDefaultPriorityQueue()}
	ch.initInfiniteChannel()
	return ch
}

func NewInfinitePriorityChannel(fun func(interface{}) bool) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewPriorityQueue(fun)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteDefaultLimitedPriorityChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewDefaultLimitedPriorityQueue()}
	ch.initInfiniteChannel()
	return ch
}

func NewInfinitePriorityLimitedPriorityChannel(fun func(interface{}) bool) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewPriorityLimitedPriorityQueue(fun)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitLimitedPriorityChannel(limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitLimitedPriorityQueue(limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitedPriorityChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitedPriorityQueue(fun, limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteDefaultLimitedPriorityListChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewDefaultLimitedPriorityListQueue()}
	ch.initInfiniteChannel()
	return ch
}

func NewInfinitePriorityLimitedPriorityListChannel(fun func(interface{}) bool) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewPriorityLimitedPriorityListQueue(fun)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitLimitedPriorityListChannel(limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitLimitedPriorityListQueue(limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitedPriorityListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitedPriorityListQueue(fun, limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteDefaultLimitedPriorityHashListChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewDefaultLimitedPriorityHashListQueue()}
	ch.initInfiniteChannel()
	return ch
}

func NewInfinitePriorityLimitedPriorityHashListChannel(fun func(interface{}) bool) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewPriorityLimitedPriorityHashListQueue(fun)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitLimitedPriorityHashListChannel(limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitLimitedPriorityHashListQueue(limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteHashLimitedPriorityHashListChannel() *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewHashLimitedPriorityHashListQueue(true)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfinitePriorityHashLimitedPriorityHashListChannel(fun func(interface{}) bool) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitHashLimitedPriorityHashListChannel(limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitHashLimitedPriorityHashListQueue(limit, true)}
	ch.initInfiniteChannel()
	return ch
}

func NewInfiniteLimitedPriorityHashListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	ch := &InfiniteChannel{buffer: queue.NewLimitedPriorityHashListQueue(fun, limit, true)}
	ch.initInfiniteChannel()
	return ch
}

func (ch *InfiniteChannel) initInfiniteChannel() {
	ch.input = make(chan interface{})
	ch.output = make(chan interface{})
	ch.length = make(chan int)
	go ch.infiniteBuffer()
}

func (ch *InfiniteChannel) In() chan<- interface{} {
	return ch.input
}

func (ch *InfiniteChannel) Out() <-chan interface{} {
	return ch.output
}

func (ch *InfiniteChannel) Len() int {
	return <-ch.length
}

func (ch *InfiniteChannel) Close() {
	close(ch.input)
}

func (ch *InfiniteChannel) infiniteBuffer() {
	var input, output chan interface{}
	var next interface{}
	input = ch.input

	for input != nil || output != nil {
		select {
		case elem, open := <-input:
			if open {
				ch.buffer.Add(elem)
			} else {
				input = nil
			}
		case output <- next:
			ch.buffer.Remove()
		case ch.length <- ch.buffer.Length():
		}

		if ch.buffer.Length() > 0 {
			output = ch.output
			next = ch.buffer.Peek()
		} else {
			output = nil
			next = nil
		}
	}

	close(ch.output)
	close(ch.length)
}
