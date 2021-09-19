package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func TestDefaultLimitedPriorityQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueNoLimitSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewLimitLimitedPriorityQueue(15), t)
}

func TestLimitLimitedPriorityQueueSimple(t *testing.T) {
	limit := 8
	elementsToAdd := 10
	indexDiff := elementsToAdd - limit
	q := queue.NewLimitLimitedPriorityQueue(limit)
	result := func(index int) int {
		return index + indexDiff
	}
	testQueueSimple(q, elementsToAdd, limit, result, t)
}

func TestLimitedPriorityQueueNoLimitSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return queue.NewLimitedPriorityQueue(fun, 15) }, t)
}

func TestLimitedPriorityQueueSimple(t *testing.T) {
	resultArray := []int{9, 6, 3, 0, 4, 5, 7, 8}
	limit := len(resultArray)
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	result := func(index int) int {
		return resultArray[index]
	}
	testQueueSimple(q, 10, limit, result, t)
}

//--

func TestDefaultLimitedPriorityQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueNoLimitAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewLimitLimitedPriorityQueue(150), t)
}

func TestLimitLimitedPriorityQueueAddRemove(t *testing.T) {
	limit := 80
	elementsToAdd := 100
	elementsToRemoveAdd := 50
	indexDiff := elementsToAdd - limit + elementsToRemoveAdd
	q := queue.NewLimitLimitedPriorityQueue(limit)
	result := func(index int) int {
		return index + indexDiff
	}
	testQueueAddRemove(q, elementsToAdd, elementsToRemoveAdd, limit, result, t)
}

func TestLimitedPriorityQueueNoLimitAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return queue.NewLimitedPriorityQueue(fun, 150) }, t)
}

func TestLimitedPriorityQueueAddRemove(t *testing.T) {
	limit := 80
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	result := func(index int) int {
		if index%2 == 0 {
			return 3*index/2 + 31
		} else {
			return (3*index + 61) / 2
		}
	}
	testQueueAddRemove(q, 100, 50, limit, result, t)
}

//--

func TesDefaultLimitedPriorityQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueNoLimitLength(t *testing.T) {
	testDefaultQueueLength(queue.NewLimitLimitedPriorityQueue(1500), t)
}

func TestLimitLimitedPriorityQueueLength(t *testing.T) {
	limit := 800
	q := queue.NewLimitLimitedPriorityQueue(limit)
	testQueueLength(q, 1000, limit, t)
}

func TestLimitedPriorityQueueNoLimitLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return queue.NewLimitedPriorityQueue(fun, 1500) }, t)
}

func TestLimitedPriorityQueueLength(t *testing.T) {
	limit := 800
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	testQueueLength(q, 1000, limit, t)
}

//--

func TestDefaultLimitedPriorityQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueNoLimitGet(t *testing.T) {
	testDefaultQueueGet(queue.NewLimitLimitedPriorityQueue(1500), t)
}

func TestLimitLimitedPriorityQueueGet(t *testing.T) {
	limit := 800
	q := queue.NewLimitLimitedPriorityQueue(limit)
	result := func(iteration int, index int) int {
		if iteration < limit {
			return index
		} else {
			return index + iteration - limit + 1
		}
	}
	testQueueGet(q, 1000, result, t)
}

func TestLimitedPriorityQueueNoLimitGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return queue.NewLimitedPriorityQueue(fun, 1500) }, t)
}

func TestLimitedPriorityQueueGet(t *testing.T) {
	limit := 800
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, limit)
	result := func(iteration int, index int) int {
		if index <= iteration/2 {
			return iteration - iteration%2 - 2*index
		} else {
			if iteration < limit {
				return -iteration + iteration%2 + 2*index - 1
			} else {
				return iteration + iteration%2 + 2*index - 2*limit + 1
			}
		}
	}
	testQueueGet(q, 1000, result, t)
}

//--

func TestDefaultLimitedPriorityQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueNoLimitGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewLimitLimitedPriorityQueue(1500), t)
}

func TestLimitLimitedPriorityQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewLimitLimitedPriorityQueue(800), t)
}

func TestLimitedPriorityQueueNoLimitGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewLimitedPriorityQueue(fun, 1500) }, t)
}

func TestLimitedPriorityQueueGetNegative(t *testing.T) {
	limit := 800
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, limit)
	result := func(iteration int, index int) int {
		if iteration < limit {
			if index >= -(iteration+iteration%2)/2 {
				return iteration + iteration%2 + 2*index + 1
			} else {
				return -iteration - iteration%2 - 2*index - 2
			}
		} else {
			if index <= (iteration-iteration%2)/2-limit {
				return iteration - iteration%2 - 2*index - 2*limit
			} else {
				return iteration + iteration%2 + 2*index + 1
			}
		}
	}
	testQueueGetNegative(q, 1000, result, t)
}

//--

func TestDefaultLimitedPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewLimitLimitedPriorityQueue(800), t)
}

func TestLimitedPriorityQueueGetOutOfRangePanics(t *testing.T) {
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueueGetOutOfRangePanics(q, t)
}

//--

func TestDefaultLimitedPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewLimitLimitedPriorityQueue(800), t)
}

func TestLimitedPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueuePeekOutOfRangePanics(q, t)
}

//--

func TestDefaultLimitedPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultLimitedPriorityQueue(), t)
}

func TestPriorityLimitedPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewLimitLimitedPriorityQueue(800), t)
}

func TestLimitedPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	q := queue.NewLimitedPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueueRemoveOutOfRangePanics(q, t)
}
