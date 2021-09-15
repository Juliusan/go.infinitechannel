package test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func TestDefaultPriorityQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueSimple(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	resultArray := []int{9, 6, 3, 0, 1, 2, 4, 5, 7, 8}
	result := func(index int) int {
		return resultArray[index]
	}
	testQueueSimple(q, len(resultArray), result, t)
}

func TestDefaultPriorityQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueAddRemove(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	result := func(index int) int {
		if index%2 == 0 {
			return 3*index/2 + 1
		} else {
			return (3*index + 1) / 2
		}
	}
	testQueueAddRemove(q, 100, 50, result, t)
}

func TestDefaultPriorityQueueLength(t *testing.T) {
	testQueueLength(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueLength(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	testQueueLength(q, t)
}

func TestDefaultPriorityQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGet(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	result := func(iteration int, index int) int {
		if index <= iteration/2 {
			return iteration - iteration%2 - 2*index
		} else {
			return -iteration + iteration%2 + 2*index - 1
		}
	}
	testQueueGet(q, 1000, result, t)
}

func TestDefaultPriorityQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGetNegative(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	result := func(iteration int, index int) int {
		if index >= -(iteration+iteration%2)/2 {
			return iteration + iteration%2 + 2*index + 1
		} else {
			return -iteration - iteration%2 - 2*index - 2
		}
	}
	testQueueGetNegative(q, 1000, result, t)
}

func TestDefaultPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGetOutOfRangePanics(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	testQueueGetOutOfRangePanics(q, t)
}

func TestDefaultPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	testQueuePeekOutOfRangePanics(q, t)
}

func TestDefaultPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	q := queue.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%2 == 0
	})
	testQueueRemoveOutOfRangePanics(q, t)
}
