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

func testLimitLimitedPriorityQueueNoLimitSimple(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueSimple(makeLimitedQueueFun(15), t)
}

func TestLimitLimitedPriorityListQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueSimple(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	limit := 8
	elementsToAdd := 10
	indexDiff := elementsToAdd - limit
	q := makeLimitedQueueFun(limit)
	result := func(index int) int {
		return index + indexDiff
	}
	testQueueSimple(q, elementsToAdd, limit, result, t)
}

func TestLimitedPriorityListQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitSimple(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 15) }, t)
}

func TestLimitedPriorityListQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueSimple(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	resultArray := []int{9, 6, 3, 0, 4, 5, 7, 8}
	limit := len(resultArray)
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	result := func(index int) int {
		return resultArray[index]
	}
	testQueueSimple(q, 10, limit, result, t)
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

func testLimitLimitedPriorityQueueNoLimitTwice(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueTwice(makeLimitedQueueFun(150), t)
}

func TestLimitLimitedPriorityListQueueTwice(t *testing.T) {
	testLimitLimitedPriorityQueueTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueTwice(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	limit := 80
	elementsToAddSingle := 50
	indexDiff := 2*elementsToAddSingle - limit
	q := makeLimitedQueueFun(limit)
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int {
		return (index + indexDiff) % elementsToAddSingle
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

func TestLimitedPriorityListQueueNoLimitTwice(t *testing.T) {
	testLimitedPriorityQueueNoLimitTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitTwice(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueTwice(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 150) }, t)
}

func TestLimitedPriorityListQueueTwice(t *testing.T) {
	testLimitedPriorityQueueTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueTwice(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 80
	elementsToAddSingle := 50
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int {
		if index <= 16 {
			return 48 - 3*index
		} else if index <= 33 {
			return 99 - 3*index
		} else if index <= 46 {
			if index%2 == 0 {
				return 3*index/2 - 20
			} else {
				return (3*index - 41) / 2
			}
		} else {
			if index%2 == 1 {
				return (3*index - 139) / 2
			} else {
				return 3*index/2 - 70
			}
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

//--

func TestLimitedPriorityListQueueOverflow(t *testing.T) {
	testLimitedPriorityQueueOverflow(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueOverflow(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	cutOff := elementsToAddSingle / 2
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int) < cutOff
	}, limit)
	addResultFun := func(index int) bool {
		return index < elementsToAddSingle+cutOff
	}
	resultFun := func(index int) int {
		if index < 25 {
			return 24 - index
		} else {
			return 49 - index
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
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

func testLimitLimitedPriorityQueueNoLimitAddRemove(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueAddRemove(makeLimitedQueueFun(150), t)
}

func TestLimitLimitedPriorityListQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueAddRemove(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	limit := 80
	elementsToAdd := 100
	elementsToRemoveAdd := 50
	indexDiff := elementsToAdd - limit + elementsToRemoveAdd
	q := makeLimitedQueueFun(limit)
	result := func(index int) int {
		return index + indexDiff
	}
	testQueueAddRemove(q, elementsToAdd, elementsToRemoveAdd, limit, result, t)
}

func TestLimitedPriorityListQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitAddRemove(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 150) }, t)
}

func TestLimitedPriorityListQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueAddRemove(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 80
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
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

func testLimitLimitedPriorityQueueNoLimitLength(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueLength(makeLimitedQueueFun(1500), t)
}

func TestLimitLimitedPriorityListQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityListQueue(limit)
	}, t)
}

func testLimitLimitedPriorityQueueLength(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	limit := 800
	q := makeLimitedQueueFun(limit)
	testQueueLength(q, 1000, limit, t)
}

func TestLimitedPriorityListQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitLength(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 1500) }, t)
}

func TestLimitedPriorityListQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueLength(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 800
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	testQueueLength(q, 1000, limit, t)
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

func testLimitLimitedPriorityQueueNoLimitGet(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueGet(makeLimitedQueueFun(1500), t)
}

func TestLimitLimitedPriorityListQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueGet(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	limit := 800
	q := makeLimitedQueueFun(limit)
	result := func(iteration int, index int) int {
		if iteration < limit {
			return index
		} else {
			return index + iteration - limit + 1
		}
	}
	testQueueGet(q, 1000, result, t)
}

func TestLimitedPriorityListQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitGet(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 1500) }, t)
}

func TestLimitedPriorityListQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueGet(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 800
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
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

func TestDefaultLimitedPriorityListQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueNoLimitGetNegative(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueGetNegative(makeLimitedQueueFun(1500), t)
}

func TestLimitLimitedPriorityListQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueGetNegative(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testDefaultQueueGetNegative(makeLimitedQueueFun(800), t)
}

func TestLimitedPriorityListQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueNoLimitGetNegative(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 1500) }, t)
}

func TestLimitedPriorityListQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueGetNegative(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	limit := 800
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
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

func TestDefaultLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultLimitedPriorityListQueue(), t)
}

func TestPriorityLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityListQueue(fun) }, t)
}

func TestLimitLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityListQueue(limit) }, t)
}

func testLimitLimitedPriorityQueueGetOutOfRangePanics(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testQueueGetOutOfRangePanics(makeLimitedQueueFun(800), t)
}

func TestLimitedPriorityListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueGetOutOfRangePanics(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueueGetOutOfRangePanics(q, t)
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

func testLimitLimitedPriorityQueuePeekOutOfRangePanics(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testQueuePeekOutOfRangePanics(makeLimitedQueueFun(800), t)
}

func TestLimitedPriorityListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueuePeekOutOfRangePanics(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueuePeekOutOfRangePanics(q, t)
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

func testLimitLimitedPriorityQueueRemoveOutOfRangePanics(makeLimitedQueueFun func(limit int) queue.Queue, t *testing.T) {
	testQueueRemoveOutOfRangePanics(makeLimitedQueueFun(800), t)
}

func TestLimitedPriorityListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityListQueue(fun, limit)
	}, t)
}

func testLimitedPriorityQueueRemoveOutOfRangePanics(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, t *testing.T) {
	q := makeLimitedPriorityQueueFun(func(i interface{}) bool {
		return i.(int)%2 == 0
	}, 800)
	testQueueRemoveOutOfRangePanics(q, t)
}
