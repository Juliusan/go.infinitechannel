package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func TestDefaultLimitedPriorityListQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitSimple(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitLimitedPriorityListQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueTwice(t *testing.T) {
	testDefaultQueueTwice(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueTwice(t *testing.T) {
	testPriorityQueueTwice(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitTwice(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitLimitedPriorityListQueueTwice(t *testing.T) {
	testLimitLimitedPriorityQueueTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueNoLimitTwice(t *testing.T) {
	testLimitedPriorityQueueNoLimitTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueTwice(t *testing.T) {
	testLimitedPriorityQueueTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestLimitedPriorityListQueueOverflow(t *testing.T) {
	testLimitedPriorityQueueOverflow(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitLimitedPriorityListQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TesDefaultLimitedPriorityListQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitLength(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityListQueue(limit)
	}, t)
}

func TestLimitLimitedPriorityListQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityListQueue(limit)
	}, t)
}

func TestLimitedPriorityListQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitGet(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitLimitedPriorityListQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitLimitedPriorityListQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityListQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueuePeekOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

//--

func TestDefaultLimitedPriorityListQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueRemoveOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func TestLimitedPriorityListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}
