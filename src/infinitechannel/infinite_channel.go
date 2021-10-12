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

func NewInfinitePriorityChannel(priorityFun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityQueue(priorityFun))
}

func NewInfiniteDefaultLimitedPriorityChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityQueue())
}

func NewInfinitePriorityLimitedPriorityChannel(priorityFun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityQueue(priorityFun))
}

func NewInfiniteLimitLimitedPriorityChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityQueue(limit))
}

func NewInfiniteLimitedPriorityChannel(priorityFun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityQueue(priorityFun, limit))
}

func NewInfiniteDefaultLimitedPriorityListChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityListQueue())
}

func NewInfinitePriorityLimitedPriorityListChannel(priorityFun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityListQueue(priorityFun))
}

func NewInfiniteLimitLimitedPriorityListChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityListQueue(limit))
}

func NewInfiniteLimitedPriorityListChannel(priorityFun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityListQueue(priorityFun, limit))
}

func NewInfiniteDefaultLimitedPriorityHashChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityHashQueue())
}

func NewInfinitePriorityLimitedPriorityHashChannel(priorityFun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityHashQueue(priorityFun))
}

func NewInfiniteLimitLimitedPriorityHashChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityHashQueue(limit))
}

func NewInfiniteLimitPriorityLimitedPriorityHashChannel(priorityFun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitPriorityLimitedPriorityHashQueue(priorityFun, limit))
}

func NewInfiniteHashLimitedPriorityHashChannel(hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewHashLimitedPriorityHashQueue(&hashFun))
}

func NewInfinitePriorityHashLimitedPriorityHashChannel(priorityFun func(interface{}) bool, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityHashLimitedPriorityHashQueue(priorityFun, &hashFun))
}

func NewInfiniteLimitHashLimitedPriorityHashChannel(limit int, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitHashLimitedPriorityHashQueue(limit, &hashFun))
}

func NewInfiniteLimitedPriorityHashChannel(priorityFun func(interface{}) bool, limit int, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityHashQueue(priorityFun, limit, &hashFun))
}

func NewInfiniteDefaultLimitedPriorityHashListChannel() *InfiniteChannel {
	return newInfiniteChannel(queue.NewDefaultLimitedPriorityHashListQueue())
}

func NewInfinitePriorityLimitedPriorityHashListChannel(priorityFun func(interface{}) bool) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityLimitedPriorityHashListQueue(priorityFun))
}

func NewInfiniteLimitLimitedPriorityHashListChannel(limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitLimitedPriorityHashListQueue(limit))
}

func NewInfiniteLimitPriorityLimitedPriorityHashListChannel(priorityFun func(interface{}) bool, limit int) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitPriorityLimitedPriorityHashListQueue(priorityFun, limit))
}

func NewInfiniteHashLimitedPriorityHashListChannel(hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewHashLimitedPriorityHashListQueue(&hashFun))
}

func NewInfinitePriorityHashLimitedPriorityHashListChannel(priorityFun func(interface{}) bool, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewPriorityHashLimitedPriorityHashListQueue(priorityFun, &hashFun))
}

func NewInfiniteLimitHashLimitedPriorityHashListChannel(limit int, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitHashLimitedPriorityHashListQueue(limit, &hashFun))
}

func NewInfiniteLimitedPriorityHashListChannel(priorityFun func(interface{}) bool, limit int, hashFun func(interface{}) interface{}) *InfiniteChannel {
	return newInfiniteChannel(queue.NewLimitedPriorityHashListQueue(priorityFun, limit, &hashFun))
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
