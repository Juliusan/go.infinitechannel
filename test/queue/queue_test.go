package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func testDefaultQueueSimple(q queue.Queue, t *testing.T) {
	testQueueSimple(q, 10, func(index int) int { return index }, t)
}

func testQueueSimple(q queue.Queue, elementsToAdd int, result func(index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		q.Add(i)
	}
	obtained := q.Length()
	if obtained != elementsToAdd {
		t.Errorf("expected full queue length %d, obtained %d", elementsToAdd, obtained)
	}
	for i := 0; i < elementsToAdd; i++ {
		expected := result(i)
		obtained = q.Peek().(int)
		if obtained != expected {
			t.Errorf("peek %d obtained %d instead of %d", i, obtained, expected)
		}
		obtained = q.Remove().(int)
		if obtained != expected {
			t.Errorf("remove %d obtained %d instead of %d", i, obtained, expected)
		}
	}
	obtained = q.Length()
	if obtained != 0 {
		t.Errorf("expected empty queue length 0, obtained %d", obtained)
	}
}

func testDefaultQueueAddRemove(q queue.Queue, t *testing.T) {
	elementsToRemoveAdd := 50
	testQueueAddRemove(q, 100, elementsToRemoveAdd, func(index int) int { return index + elementsToRemoveAdd }, t)
}

func testQueueAddRemove(q queue.Queue, elementsToAdd int, elementsToRemoveAdd int, result func(index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		q.Add(i)
	}
	for i := 0; i < elementsToRemoveAdd; i++ {
		q.Remove()
		q.Add(elementsToAdd + i)
	}
	obtained := q.Length()
	if obtained != elementsToAdd {
		t.Errorf("expected full queue length %d, obtained %d", elementsToAdd, obtained)
	}

	for i := 0; i < elementsToAdd; i++ {
		expected := result(i)
		obtained = q.Peek().(int)
		if obtained != expected {
			t.Errorf("peek %d obtained %d instead of %d", i, obtained, expected)
		}
		obtained = q.Remove().(int)
		if obtained != expected {
			t.Errorf("remove %d obtained %d instead of %d", i, obtained, expected)
		}
	}
	obtained = q.Length()
	if obtained != 0 {
		t.Errorf("expected empty queue length 0, obtained %d", obtained)
	}
}

func testQueueLength(q queue.Queue, t *testing.T) {
	obtained := q.Length()
	if obtained != 0 {
		t.Errorf("empty queue length is %d", obtained)
	}

	for i := 0; i < 1000; i++ {
		q.Add(i)
		expected := i + 1
		obtained := q.Length()
		if obtained != expected {
			t.Errorf("adding: expected queue length %d, obtained %d", expected, obtained)
		}
	}
	for i := 0; i < 1000; i++ {
		q.Remove()
		expected := 1000 - i - 1
		obtained := q.Length()
		if obtained != expected {
			t.Errorf("removing: expected queue length %d, obtained %d", expected, obtained)
		}
	}
}

func testDefaultQueueGet(q queue.Queue, t *testing.T) {
	testQueueGet(q, 1000, func(iteration int, index int) int { return index }, t)
}

func testQueueGet(q queue.Queue, elementsToAdd int, result func(iteration int, index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		q.Add(i)
		for j := 0; j < q.Length(); j++ {
			expected := result(i, j)
			obtained := q.Get(j).(int)
			if obtained != expected {
				t.Errorf("iteration %d index %d contains %d instead of %d", i, j, obtained, expected)
			}
		}
	}
}

func testDefaultQueueGetNegative(q queue.Queue, t *testing.T) {
	testQueueGetNegative(q, 1000, func(iteration int, index int) int { return iteration + index + 1 }, t)
}

func testQueueGetNegative(q queue.Queue, elementsToAdd int, result func(iteration int, index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		q.Add(i)
		for j := -1; j >= -q.Length(); j-- {
			expected := result(i, j)
			obtained := q.Get(j).(int)
			if obtained != expected {
				t.Errorf("iteration %d index %d contains %d instead of %d", i, j, obtained, expected)
			}
		}
	}
}

func testQueueGetOutOfRangePanics(q queue.Queue, t *testing.T) {
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

func testQueuePeekOutOfRangePanics(q queue.Queue, t *testing.T) {
	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	q.Add(1)
	q.Remove()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

func testQueueRemoveOutOfRangePanics(q queue.Queue, t *testing.T) {
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
