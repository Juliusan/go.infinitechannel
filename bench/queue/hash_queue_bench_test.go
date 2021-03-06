package queue_bench

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func BenchmarkDefaultLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityHashQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, b)
}

func BenchmarkLimitLimitedPriorityHashQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, b)
}

func BenchmarkLimitLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueAdd10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, b)
}

func BenchmarkHashLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkQueueAdd10k(func() queue.Queue { return queue.NewHashLimitedPriorityHashQueue(true) }, b)
}

func BenchmarkPriorityHashLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkPriorityQueueAdd10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, b)
}

func BenchmarkLimitHashLimitedPriorityHashQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitAdd10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, b)
}

func BenchmarkLimitHashLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueAdd10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, b)
}

func BenchmarkLimitedPriorityHashQueueNoLimitAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, b)
}

func BenchmarkLimitedPriorityHashQueueAdd10k(b *testing.B) {
	benchmarkLimitedPriorityQueueAdd10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, b)
}

//--

func BenchmarkDefaultLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewDefaultLimitedPriorityHashQueue() }, b)
}

func BenchmarkPriorityLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue { return queue.NewPriorityLimitedPriorityHashQueue(fun) }, b)
}

func BenchmarkLimitLimitedPriorityHashQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, b)
}

func BenchmarkLimitLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueRemove10k(func(limit int) queue.Queue { return queue.NewLimitLimitedPriorityHashQueue(limit) }, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, b)
}

func BenchmarkLimitPriorityLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitPriorityLimitedPriorityHashQueue(fun, limit)
	}, b)
}

func BenchmarkHashLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkDefaultQueueRemove10k(func() queue.Queue { return queue.NewHashLimitedPriorityHashQueue(true) }, b)
}

func BenchmarkPriorityHashLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkPriorityQueueRemove10k(func(fun func(i interface{}) bool) queue.Queue {
		return queue.NewPriorityHashLimitedPriorityHashQueue(fun, true)
	}, b)
}

func BenchmarkLimitHashLimitedPriorityHashQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueNoLimitRemove10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, b)
}

func BenchmarkLimitHashLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkLimitLimitedPriorityQueueRemove10k(func(limit int) queue.Queue { return queue.NewLimitHashLimitedPriorityHashQueue(limit, true) }, b)
}

func BenchmarkLimitedPriorityHashQueueNoLimitRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueNoLimitRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, b)
}

func BenchmarkLimitedPriorityHashQueueRemove10k(b *testing.B) {
	benchmarkLimitedPriorityQueueRemove10k(func(fun func(i interface{}) bool, limit int) queue.Queue {
		return queue.NewLimitedPriorityHashQueue(fun, limit, true)
	}, b)
}
