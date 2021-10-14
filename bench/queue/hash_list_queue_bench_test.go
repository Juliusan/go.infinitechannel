package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func BenchmarkDefaultLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityHashListQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, b)
}

func BenchmarkLimitLimitedPriorityHashListQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, b)
}

func BenchmarkLimitLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashListQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, b)
}

func BenchmarkHashLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewHashLimitedPriorityHashListQueue(true) }, b)
}

func BenchmarkPriorityHashLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, b)
}

func BenchmarkLimitHashLimitedPriorityHashListQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, b)
}

func BenchmarkLimitHashLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueAdd10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, b)
}

func BenchmarkLimitedPriorityHashListQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, b)
}

func BenchmarkLimitedPriorityHashListQueueAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, b)
}

//--

func BenchmarkDefaultLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityHashListQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityLimitedPriorityHashListQueue(fun)
	}, b)
}

func BenchmarkLimitLimitedPriorityHashListQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, b)
}

func BenchmarkLimitLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashListQueue(limit) }, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashListQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashListQueue(fun, limit)
	}, b)
}

func BenchmarkHashLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewHashLimitedPriorityHashListQueue(true) }, b)
}

func BenchmarkPriorityHashLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashListQueue(fun, true)
	}, b)
}

func BenchmarkLimitHashLimitedPriorityHashListQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, b)
}

func BenchmarkLimitHashLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueRemove10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashListQueue(limit, true) }, b)
}

func BenchmarkLimitedPriorityHashListQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, b)
}

func BenchmarkLimitedPriorityHashListQueueRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashListQueue(fun, limit, true)
	}, b)
}
