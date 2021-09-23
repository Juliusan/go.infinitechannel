package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func BenchmarkDefaultLimitedPriorityQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, b)
}

func BenchmarkLimitLimitedPriorityQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityQueue(limit) }, b)
}

func benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(makeLimitedQueueFun func(limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedQueueFun(12000)
	}
	benchmarkQueueAdd10k(makeQueueFun, b)
}

func BenchmarkLimitLimitedPriorityQueueAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityQueue(limit) }, b)
}

func benchmarkLimitLimitedPriorityQueueAdd10k(makeLimitedQueueFun func(limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedQueueFun(8000)
	}
	benchmarkQueueAdd10k(makeQueueFun, b)
}

func BenchmarkLimitedPriorityQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityQueue(fun, limit)
	}, b)
}

func benchmarkLimitedPriorityQueueNoLimitAdd10k(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 12000) }, b)
}

func BenchmarkLimitedPriorityQueueAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityQueue(fun, limit)
	}, b)
}

func benchmarkLimitedPriorityQueueAdd10k(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedPriorityQueueFun(func(i interface{}) bool {
			return i.(int)%2 == 0
		}, 8000)
	}
	benchmarkQueueAdd10k(makeQueueFun, b)
}

//--

func BenchmarkDefaultLimitedPriorityQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityQueue(fun) }, b)
}

func BenchmarkLimitLimitedPriorityQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityQueue(limit) }, b)
}

func benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(makeLimitedQueueFun func(limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedQueueFun(12000)
	}
	benchmarkDefaultQueueRemove10k(makeQueueFun, b)
}

func BenchmarkLimitLimitedPriorityQueueRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityQueue(limit) }, b)
}

func benchmarkLimitLimitedPriorityQueueRemove10k(makeLimitedQueueFun func(limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedQueueFun(10000)
	}
	benchmarkQueueRemove10k(makeQueueFun, 12000, b)
}

func BenchmarkLimitedPriorityQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityQueue(fun, limit)
	}, b)
}

func benchmarkLimitedPriorityQueueNoLimitRemove10k(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue { return makeLimitedPriorityQueueFun(fun, 12000) }, b)
}

func BenchmarkLimitedPriorityQueueRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityQueue(fun, limit)
	}, b)
}

func benchmarkLimitedPriorityQueueRemove10k(makeLimitedPriorityQueueFun func(fun func(i interface{}) bool, limit int) queue.Queue, b *testing.B) {
	makeQueueFun := func() queue.Queue {
		return makeLimitedPriorityQueueFun(func(i interface{}) bool {
			return i.(int)%2 == 0
		}, 10000)
	}
	benchmarkQueueRemove10k(makeQueueFun, 12000, b)
}
