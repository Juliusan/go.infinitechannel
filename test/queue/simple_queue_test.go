package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func TestSimpleQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueLength(t *testing.T) {
	testQueueLength(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewSimpleQueue(), t)
}

func TestSimpleQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewSimpleQueue(), t)
}

func TestSimpleQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewSimpleQueue(), t)
}
