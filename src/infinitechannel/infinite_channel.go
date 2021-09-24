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
	return newInfiniteChannel(queue.NewSimpleQueue())
}

func NewInfiniteDefaultPriorityChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultPriorityQueue())
}

func NewInfinitePriorityChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityQueue(fun))
}

func NewInfiniteDefaultLimitedPriorityChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityQueue())
}

func NewInfinitePriorityLimitedPriorityChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityQueue(fun))
}

func NewInfiniteLimitLimitedPriorityChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityQueue(limit))
}

func NewInfiniteLimitedPriorityChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityQueue(fun, limit))
}

func NewInfiniteDefaultLimitedPriorityListChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityListQueue())
}

func NewInfinitePriorityLimitedPriorityListChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityListQueue(fun))
}

func NewInfiniteLimitLimitedPriorityListChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityListQueue(limit))
}

func NewInfiniteLimitedPriorityListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityListQueue(fun, limit))
}

func NewInfiniteDefaultLimitedPriorityHashChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityHashQueue())
}

func NewInfinitePriorityLimitedPriorityHashChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityHashQueue(fun))
}

func NewInfiniteLimitLimitedPriorityHashChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityHashQueue(limit))
}

func NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit))
}

func NewInfiniteHashLimitedPriorityHashChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewHashLimitedPriorityHashQueue(true))
}

func NewInfinitePriorityHashLimitedPriorityHashChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityHashLimitedPriorityHashQueue(fun, true))
}

func NewInfiniteLimitHashLimitedPriorityHashChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitHashLimitedPriorityHashQueue(limit, true))
}

func NewInfiniteLimitedPriorityHashChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityHashQueue(fun, limit, true))
}

func NewInfiniteDefaultLimitedPriorityHashListChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityHashListQueue())
}

func NewInfinitePriorityLimitedPriorityHashListChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityHashListQueue(fun))
}

func NewInfiniteLimitLimitedPriorityHashListChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityHashListQueue(limit))
}

func NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit))
}

func NewInfiniteHashLimitedPriorityHashListChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewHashLimitedPriorityHashListQueue(true))
}

func NewInfinitePriorityHashLimitedPriorityHashListChannel(fun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true))
}

func NewInfiniteLimitHashLimitedPriorityHashListChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitHashLimitedPriorityHashListQueue(limit, true))
}

func NewInfiniteLimitedPriorityHashListChannel(fun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityHashListQueue(fun, limit, true))
}

func newInfiniteChannel(queue queue.Queue) *InfiniteChannel {
	ch := &InfiniteChannel{
		input:  make(chan interface{}),
		output: make(chan interface{}),
		length: make(chan int),
		buffer: queue,
	}
	go ch.infiniteBuffer()
	return ch
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
