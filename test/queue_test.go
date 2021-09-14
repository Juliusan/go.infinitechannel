package test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestQueueSimple(t *testing.T) {
	q := infinitechannel.NewQueue()

	for i := 0; i < infinitechannel.MinQueueLen; i++ {
		q.Add(i)
	}
	for i := 0; i < infinitechannel.MinQueueLen; i++ {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		x := q.Remove()
		if x != i {
			t.Error("remove", i, "had value", x)
		}
	}
}

func TestQueueWrapping(t *testing.T) {
	q := infinitechannel.NewQueue()

	for i := 0; i < infinitechannel.MinQueueLen; i++ {
		q.Add(i)
	}
	for i := 0; i < 3; i++ {
		q.Remove()
		q.Add(infinitechannel.MinQueueLen + i)
	}

	for i := 0; i < infinitechannel.MinQueueLen; i++ {
		if q.Peek().(int) != i+3 {
			t.Error("peek", i, "had value", q.Peek())
		}
		q.Remove()
	}
}

func TestQueuePrioritising(t *testing.T) {
	q := infinitechannel.NewPriorityQueue(func(i interface{}) bool {
		return i.(int)%3 == 0
	})

	for i := 0; i < infinitechannel.MinQueueLen; i++ {
		q.Add(i)
	}

	peekAndRemoveFun := func(i int) {
		if q.Peek().(int) != i {
			t.Error("peek", i, "had value", q.Peek())
		}
		x := q.Remove()
		if x != i {
			t.Error("remove", i, "had value", x)
		}
	}
	last3 := infinitechannel.MinQueueLen - 1 - ((infinitechannel.MinQueueLen - 1) % 3)
	for i := last3; i >= 0; i -= 3 {
		peekAndRemoveFun(i)
	}
	for i := 0; i < last3; i += 3 {
		peekAndRemoveFun(i + 1)
		peekAndRemoveFun(i + 2)
	}
	for i := last3 + 1; i < infinitechannel.MinQueueLen; i++ {
		peekAndRemoveFun(i)
	}
}

func TestQueueLength(t *testing.T) {
	q := infinitechannel.NewQueue()

	if q.Length() != 0 {
		t.Error("empty queue length not 0")
	}

	for i := 0; i < 1000; i++ {
		q.Add(i)
		if q.Length() != i+1 {
			t.Error("adding: queue with", i, "elements has length", q.Length())
		}
	}
	for i := 0; i < 1000; i++ {
		q.Remove()
		if q.Length() != 1000-i-1 {
			t.Error("removing: queue with", 1000-i-i, "elements has length", q.Length())
		}
	}
}

func TestQueueGet(t *testing.T) {
	q := infinitechannel.NewQueue()

	for i := 0; i < 1000; i++ {
		q.Add(i)
		for j := 0; j < q.Length(); j++ {
			if q.Get(j).(int) != j {
				t.Errorf("index %d doesn't contain %d", j, j)
			}
		}
	}
}

func TestQueueGetNegative(t *testing.T) {
	q := infinitechannel.NewQueue()

	for i := 0; i < 1000; i++ {
		q.Add(i)
		for j := 1; j <= q.Length(); j++ {
			if q.Get(-j).(int) != q.Length()-j {
				t.Errorf("index %d doesn't contain %d", -j, q.Length()-j)
			}
		}
	}
}

func TestQueueGetOutOfRangePanics(t *testing.T) {
	q := infinitechannel.NewQueue()

	q.Add(1)
	q.Add(2)
	q.Add(3)

	assertPanics(t, "should panic when negative index", func() {
		q.Get(-4)
	})

	assertPanics(t, "should panic when index greater than length", func() {
		q.Get(4)
	})
}

func TestQueuePeekOutOfRangePanics(t *testing.T) {
	q := infinitechannel.NewQueue()

	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	q.Add(1)
	q.Remove()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

func TestQueueRemoveOutOfRangePanics(t *testing.T) {
	q := infinitechannel.NewQueue()

	assertPanics(t, "should panic when removing empty queue", func() {
		q.Remove()
	})

	q.Add(1)
	q.Remove()

	assertPanics(t, "should panic when removing emptied queue", func() {
		q.Remove()
	})
}

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		}
	}()

	f()
}

// General warning: Go's benchmark utility (go test -bench .) increases the number of
// iterations until the benchmarks take a reasonable amount of time to run; memory usage
// is *NOT* considered. On my machine, these benchmarks hit around ~1GB before they've had
// enough, but if you have less than that available and start swapping, then all bets are off.

func BenchmarkQueueSerial(b *testing.B) {
	q := infinitechannel.NewQueue()
	for i := 0; i < b.N; i++ {
		q.Add(nil)
	}
	for i := 0; i < b.N; i++ {
		q.Peek()
		q.Remove()
	}
}

func BenchmarkQueueGet(b *testing.B) {
	q := infinitechannel.NewQueue()
	for i := 0; i < b.N; i++ {
		q.Add(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Get(i)
	}
}

func BenchmarkQueueTickTock(b *testing.B) {
	q := infinitechannel.NewQueue()
	for i := 0; i < b.N; i++ {
		q.Add(nil)
		q.Peek()
		q.Remove()
	}
}
