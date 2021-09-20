package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func TestDefaultLimitedPriorityHashQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitSimple(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitSimple(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitSimple(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueueTwice(t *testing.T) {
	testDefaultQueueTwice(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueTwice(t *testing.T) {
	testPriorityQueueTwice(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitTwice(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashQueueTwice(t *testing.T) {
	testLimitLimitedPriorityQueueTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitTwice(t *testing.T) {
	testLimitedPriorityQueueNoLimitTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueTwice(t *testing.T) {
	testLimitedPriorityQueueTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueTwice(t *testing.T) {
	testHashLimitedPriorityHashQueueTwice(func(hashNeeded bool) queue.Queue {
		return queue.NewHashLimitedPriorityHashQueue(hashNeeded)
	}, t)
}

func testHashLimitedPriorityHashQueueTwice(makeHashQueueFun func(hashNeeded bool) queue.Queue, t *testing.T) {
	q := makeHashQueueFun(true)
	elementsToAddSingle := 50
	addResultFun := func(index int) bool { return index < elementsToAddSingle }
	resultFun := func(index int) int { return index }
	testQueueTwice(q, elementsToAddSingle, addResultFun, elementsToAddSingle, resultFun, t)
}

func TestPriorityHashLimitedPriorityHashQueueTwice(t *testing.T) {
	testPriorityHashLimitedPriorityHashQueueTwice(func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, hashNeeded)
	}, t)
}

func testPriorityHashLimitedPriorityHashQueueTwice(makeHashQueueFun func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue, t *testing.T) {
	q := makeHashQueueFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, true)
	elementsToAddSingle := 50
	addResultFun := func(index int) bool { return index < elementsToAddSingle }
	resultFun := func(index int) int {
		if index <= 16 {
			return 48 - 3*index
		} else {
			if index%2 == 1 {
				return (3*index - 49) / 2
			} else {
				return 3*index/2 - 25
			}
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, elementsToAddSingle, resultFun, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitTwice(t *testing.T) {
	testHashLimitedPriorityHashQueueTwice(func(hashNeeded bool) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashQueue(80, hashNeeded)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueTwice(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	indexDiff := elementsToAddSingle - limit
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int { return index + indexDiff }
	q := queue.NewLimitHashLimitedPriorityHashQueue(limit, true)
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

func TestLimitedPriorityHashQueueNoLimitTwice(t *testing.T) {
	testPriorityHashLimitedPriorityHashQueueTwice(func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, 80, hashNeeded)
	}, t)
}

func TestLimitedPriorityHashQueueTwice(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	q := queue.NewLimitedPriorityHashQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit, true)
	addResultFun := func(index int) bool { return (index < elementsToAddSingle) || ((index-elementsToAddSingle)%3 != 0) }
	resultFun := func(index int) int {
		if index <= 16 {
			return 48 - 3*index
		} else {
			if index%2 == 1 {
				return (3*index + 11) / 2
			} else {
				return 3*index/2 + 5
			}
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

//--

func TestLimitPriorityLimitedPriorityHashQueueOverflow(t *testing.T) {
	testLimitedPriorityQueueOverflow(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityHashQueueOverflow(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	cutOffLow := 20
	cutOffHigh := 40
	q := queue.NewLimitedPriorityHashQueue(func(i interface{}) bool {
		value := i.(int)
		return value < cutOffLow || cutOffHigh <= value
	}, limit, true)
	addResultFun := func(index int) bool {
		return index < elementsToAddSingle
	}
	resultFun := func(index int) int {
		if index < 10 {
			return 49 - index
		} else {
			return 29 - index
		}
	}
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

//--

func TestDefaultLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitAddRemove(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TesDefaultLimitedPriorityHashQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitLength(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityHashQueue(limit)
	}, t)
}

func TestLimitLimitedPriorityHashQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityHashQueue(limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TesHashLimitedPriorityHashQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitLength(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitLength(func(limit int) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashQueue(limit, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashQueue(limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitGet(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitGet(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGet(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityQueueHashGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityQueueHashGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueuePeekOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashtLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueuePeekOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultLimitedPriorityHashQueue(), t)
}

func TestPriorityLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, t)
}

func TestLimitLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueRemoveOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewHashLimitedPriorityHashQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueRemoveOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, t)
}

func TestLimitedPriorityHashQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, t)
}
