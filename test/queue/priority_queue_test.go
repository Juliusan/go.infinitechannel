package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
	"github.com/Juliusan/go.infinitechannel/src/util"
)

func TestDefaultPriorityQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueSimple(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod3)
	resultArray := []int{9, 6, 3, 0, 1, 2, 4, 5, 7, 8}
	result := func(index int) int {
		return resultArray[index]
	}
	elementsToAdd := len(resultArray)
	testQueueSimple(q, elementsToAdd, elementsToAdd, result, t)
}

//--

func TestDefaultPriorityQueueTwice(t *testing.T) {
	testDefaultQueueTwice(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueTwice(t *testing.T) {
	testPriorityQueueTwice(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueTwice(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod3)
	elementsToAddSingle := 50
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int {
		if index <= 16 {
			return 48 - 3*index
		} else if index <= 33 {
			return 99 - 3*index
		} else if index <= 66 {
			if index%2 == 0 {
				return 3*index/2 - 50
			} else {
				return (3*index - 101) / 2
			}
		} else {
			if index%2 == 1 {
				return (3*index - 199) / 2
			} else {
				return 3*index/2 - 100
			}
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, 2*elementsToAddSingle, resultFun, t)
}

//--

func TestDefaultPriorityQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueAddRemove(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod3)
	result := func(index int) int {
		if index%2 == 0 {
			return 3*index/2 + 1
		} else {
			return (3*index + 1) / 2
		}
	}
	elementsToAdd := 100
	testQueueAddRemove(q, elementsToAdd, 50, elementsToAdd, result, t)
}

//--

func TestDefaultPriorityQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueLength(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod3)
	elementsToAdd := 1000
	testQueueLength(q, elementsToAdd, elementsToAdd, t)
}

//--

func TestDefaultPriorityQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueGet(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod2)
	result := func(iteration int, index int) int {
		if index <= iteration/2 {
			return iteration - iteration%2 - 2*index
		} else {
			return -iteration + iteration%2 + 2*index - 1
		}
	}
	testQueueGet(q, 1000, result, t)
}

//--

func TestDefaultPriorityQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueGetNegative(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod2)
	result := func(iteration int, index int) int {
		if index >= -(iteration+iteration%2)/2 {
			return iteration + iteration%2 + 2*index + 1
		} else {
			return -iteration - iteration%2 - 2*index - 2
		}
	}
	testQueueGetNegative(q, 1000, result, t)
}

//--

func TestDefaultPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueGetOutOfRangePanics(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod2)
	testQueueGetOutOfRangePanics(q, t)
}

//--

func TestDefaultPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueuePeekOutOfRangePanics(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod2)
	testQueuePeekOutOfRangePanics(q, t)
}

//--

func TestDefaultPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultPriorityQueue(), t)
}

func TestPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityQueue(fun) }, t)
}

func testPriorityQueueRemoveOutOfRangePanics(makePriorityQueueFun func(func(i interface{}) bool) queue.Queue, t *testing.T) {
	q := makePriorityQueueFun(util.PriorityFunMod2)
	testQueueRemoveOutOfRangePanics(q, t)
}
