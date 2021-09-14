package infinitechannel_test

import (
	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
	"testing"
)

// SimpleInChannel is an interface representing a writeable channel that does not necessarily
// implement the Buffer interface.
type SimpleInChannel interface {
	In() chan<- interface{} // The writeable end of the channel.
	Close()                 // Closes the channel. It is an error to write to In() after calling Close().
}

// InChannel is an interface representing a writeable channel with a buffer.
type InChannel interface {
	SimpleInChannel
	Buffer
}

// SimpleOutChannel is an interface representing a readable channel that does not necessarily
// implement the Buffer interface.
type SimpleOutChannel interface {
	Out() <-chan interface{} // The readable end of the channel.
}

// OutChannel is an interface representing a readable channel implementing the Buffer interface.
type OutChannel interface {
	SimpleOutChannel
	Buffer
}

// Buffer is an interface for any channel that provides access to query the state of its buffer.
// Even unbuffered channels can implement this interface by simply returning 0 from Len() and None from Cap().
type Buffer interface {
	Len() int // The number of elements currently buffered.
	Cap() int // The maximum number of elements that can be buffered.
}

func testChannel(t *testing.T, name string, ch infinitechannel.InfiniteChannel) {
	go func() {
		for i := 0; i < 1000; i++ {
			ch.In() <- i
		}
		ch.Close()
	}()
	for i := 0; i < 1000; i++ {
		val := <-ch.Out()
		if i != val.(int) {
			t.Fatal(name, "expected", i, "but got", val.(int))
		}
	}
}

func testChannelPair(t *testing.T, name string, in InChannel, out OutChannel) {
	go func() {
		for i := 0; i < 1000; i++ {
			in.In() <- i
		}
		in.Close()
	}()
	for i := 0; i < 1000; i++ {
		val := <-out.Out()
		if i != val.(int) {
			t.Fatal("pair", name, "expected", i, "but got", val.(int))
		}
	}
}

func testChannelConcurrentAccessors(t *testing.T, name string, ch infinitechannel.InfiniteChannel) {
	// no asserts here, this is just for the race detector's benefit
	go ch.Len()
	go ch.Cap()

	go func() {
		ch.In() <- nil
	}()

	go func() {
		<-ch.Out()
	}()
}
