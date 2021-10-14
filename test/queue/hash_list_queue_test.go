package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
	"github.com/Juliusan/go.infinitechannel/src/util"
)

func TestDefaultLimitedPriorityHashListQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitSimple(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashListQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueSimple(t *testing.T) {
	testDefaultQueueSimple(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueSimple(t *testing.T) {
	testPriorityQueueSimple(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueNoLimitSimple(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitSimple(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashListQueueSimple(t *testing.T) {
	testLimitLimitedPriorityQueueSimple(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueNoLimitSimple(t *testing.T) {
	testLimitedPriorityQueueNoLimitSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueSimple(t *testing.T) {
	testLimitedPriorityQueueSimple(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueTwice(t *testing.T) {
	testDefaultQueueTwice(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueTwice(t *testing.T) {
	testPriorityQueueTwice(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitTwice(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashListQueueTwice(t *testing.T) {
	testLimitLimitedPriorityQueueTwice(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitTwice(t *testing.T) {
	testLimitedPriorityQueueNoLimitTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueTwice(t *testing.T) {
	testLimitedPriorityQueueTwice(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueTwice(t *testing.T) {
	testHashLimitedPriorityHashListQueueTwice(func(hashNeeded bool) queue.Queue {
		return queue.NewHashLimitedPriorityHashListQueue(hashNeeded)
	}, t)
}

func testHashLimitedPriorityHashListQueueTwice(makeHashQueueFun func(hashNeeded bool) queue.Queue, t *testing.T) {
	q := makeHashQueueFun(true)
	elementsToAddSingle := 50
	addResultFun := func(index int) bool { return index < elementsToAddSingle }
	resultFun := func(index int) int { return index }
	testQueueTwice(q, elementsToAddSingle, addResultFun, elementsToAddSingle, resultFun, t)
}

func TestPriorityHashLimitedPriorityHashListQueueTwice(t *testing.T) {
	testPriorityHashLimitedPriorityHashListQueueTwice(func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, hashNeeded)
	}, t)
}

func testPriorityHashLimitedPriorityHashListQueueTwice(makeHashQueueFun func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue, t *testing.T) {
	q := makeHashQueueFun(util.PriorityFunMod3, true)
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

func TestLimitHashLimitedPriorityHashListQueueNoLimitTwice(t *testing.T) {
	testHashLimitedPriorityHashListQueueTwice(func(hashNeeded bool) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashListQueue(80, hashNeeded)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueTwice(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	indexDiff := elementsToAddSingle - limit
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int { return index + indexDiff }
	q := queue.NewLimitHashLimitedPriorityHashListQueue(limit, true)
	testQueueTwice(q, elementsToAddSingle, addResultFun, limit, resultFun, t)
}

func TestLimitedPriorityHashListQueueNoLimitTwice(t *testing.T) {
	testPriorityHashLimitedPriorityHashListQueueTwice(func(fun func(i interface{}) bool, hashNeeded bool) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, 80, hashNeeded)
	}, t)
}

func TestLimitedPriorityHashListQueueTwice(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	q := queue.NewLimitedPriorityHashListQueue(util.PriorityFunMod3, limit, true)
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

func TestLimitPriorityLimitedPriorityHashListQueueOverflow(t *testing.T) {
	testLimitedPriorityQueueOverflow(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitedPriorityHashListQueueOverflow(t *testing.T) {
	limit := 30
	elementsToAddSingle := 50
	cutOffLow := util.SimpleHashable(20)
	cutOffHigh := util.SimpleHashable(40)
	q := queue.NewLimitedPriorityHashListQueue(func(i interface{}) bool {
		value := i.(util.SimpleHashable)
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

func TestLimitedPriorityHashListQueueDuplicates(t *testing.T) {
	limit := 80
	elementsToAddFirstIteration := 50
	q := queue.NewLimitedPriorityHashListQueue(util.PriorityFunMod3, limit, true)
	addFun := func(index int) int {
		if index < elementsToAddFirstIteration {
			return 2 * index
		} else {
			return index - elementsToAddFirstIteration
		}
	}
	addResultFun := func(index int) bool {
		return (index < elementsToAddFirstIteration) || ((index-elementsToAddFirstIteration)%2 == 1)
	}
	resultFun := func(index int) int {
		if index <= 16 {
			return 99 - 6*index
		} else if index <= 33 {
			return 198 - 6*index
		} else if index <= 46 {
			if index%2 == 0 {
				return 3*index - 40
			} else {
				return 3*index - 41
			}
		} else {
			if index%2 == 0 {
				return 3*index - 139
			} else {
				return 3*index - 140
			}
		}
	}
	testQueueBasicAddLengthPeekRemove(q, 3*elementsToAddFirstIteration, addFun, addResultFun, limit, resultFun, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testDefaultQueueAddRemove(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testPriorityQueueAddRemove(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueNoLimitAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitAddRemove(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testLimitLimitedPriorityQueueAddRemove(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueNoLimitAddRemove(t *testing.T) {
	testLimitedPriorityQueueNoLimitAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueAddRemove(t *testing.T) {
	testLimitedPriorityQueueAddRemove(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TesDefaultLimitedPriorityHashListQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitLength(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityHashListQueue(limit)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitLimitedPriorityHashListQueue(limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TesHashLimitedPriorityHashListQueueLength(t *testing.T) {
	testDefaultQueueLength(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueLength(t *testing.T) {
	testPriorityQueueLength(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueNoLimitLength(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitLength(func(limit int) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueLength(t *testing.T) {
	testLimitLimitedPriorityQueueLength(func(limit int) queue.Queue {
		return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueNoLimitLength(t *testing.T) {
	testLimitedPriorityQueueNoLimitLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueLength(t *testing.T) {
	testLimitedPriorityQueueLength(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitGet(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashListQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueGet(t *testing.T) {
	testDefaultQueueGet(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueGet(t *testing.T) {
	testPriorityQueueGet(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueNoLimitGet(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGet(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashListQueueGet(t *testing.T) {
	testLimitLimitedPriorityQueueGet(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueNoLimitGet(t *testing.T) {
	testLimitedPriorityQueueNoLimitGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueGet(t *testing.T) {
	testLimitedPriorityQueueGet(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testDefaultQueueGetNegative(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testPriorityQueueGetNegative(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueNoLimitGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueNoLimitGetNegative(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitHashLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testLimitLimitedPriorityQueueGetNegative(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueNoLimitGetNegative(t *testing.T) {
	testLimitedPriorityQueueNoLimitGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

func TestLimitedPriorityHashListQueueGetNegative(t *testing.T) {
	testLimitedPriorityQueueGetNegative(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testQueueGetOutOfRangePanics(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueGetOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueGetOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueGetOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueuePeekOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashtLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testQueuePeekOutOfRangePanics(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueuePeekOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueuePeekOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueuePeekOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}

//--

func TestDefaultLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewDefaultLimitedPriorityHashListQueue(), t)
}

func TestPriorityLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, t)
}

func TestLimitLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueRemoveOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, t)
}

func TestLimitPriorityLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, t)
}

func TestHashLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testQueueRemoveOutOfRangePanics(queue.NewHashLimitedPriorityHashListQueue(true), t)
}

func TestPriorityHashLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, t)
}

func TestLimitHashLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitLimitedPriorityQueueRemoveOutOfRangePanics(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, t)
}

func TestLimitedPriorityHashListQueueRemoveOutOfRangePanics(t *testing.T) {
	testLimitedPriorityQueueRemoveOutOfRangePanics(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, t)
}
