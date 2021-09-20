package queue_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/queue"
)

func testQueueBasicAddLengthPeekRemove(q queue.Queue, elementsToAdd int, add func(index int) int, addResult func(index int) bool, elementsToRemove int, result func(index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		value := add(i)
		expected := addResult(i)
		obtained := q.Add(value)
		if obtained != expected {
			t.Errorf("add result of element %d value %d expected %v obtained %v", i, value, expected, obtained)
		}

	}
	obtained := q.Length()
	if obtained != elementsToRemove {
		t.Errorf("expected full queue length %d, obtained %d", elementsToAdd, obtained)
	}
	for i := 0; i < elementsToRemove; i++ {
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

//--

func testDefaultQueueSimple(q queue.Queue, t *testing.T) {
	elementsToAdd := 10
	testQueueSimple(q, elementsToAdd, elementsToAdd, func(index int) int { return index }, t)
}

func testQueueSimple(q queue.Queue, elementsToAdd int, elementsToRemove int, result func(index int) int, t *testing.T) {
	testQueueBasicAddLengthPeekRemove(q, elementsToAdd, func(index int) int { return index }, func(index int) bool { return true }, elementsToRemove, result, t)
}

//--

func testDefaultQueueTwice(q queue.Queue, t *testing.T) {
	elementsToAddSingle := 50
	addResultFun := func(index int) bool { return true }
	resultFun := func(index int) int { return index % elementsToAddSingle }
	testQueueTwice(q, elementsToAddSingle, addResultFun, 2*elementsToAddSingle, resultFun, t)
}

func testQueueTwice(q queue.Queue, elementsToAddSingle int, addResult func(index int) bool, elementsToRemove int, result func(index int) int, t *testing.T) {
	addFun := func(index int) int {
		return index % elementsToAddSingle
	}
	testQueueBasicAddLengthPeekRemove(q, 2*elementsToAddSingle, addFun, addResult, elementsToRemove, result, t)
}

//--

func testDefaultQueueAddRemove(q queue.Queue, t *testing.T) {
	elementsToAdd := 100
	elementsToRemoveAdd := 50
	testQueueAddRemove(q, elementsToAdd, elementsToRemoveAdd, elementsToAdd, func(index int) int { return index + elementsToRemoveAdd }, t)
}

func testQueueAddRemove(q queue.Queue, elementsToAdd int, elementsToRemoveAdd int, elementsToRemove int, result func(index int) int, t *testing.T) {
	for i := 0; i < elementsToAdd; i++ {
		if !q.Add(i) {
			t.Errorf("failed to add element %d", i)
		}
	}
	for i := 0; i < elementsToRemoveAdd; i++ {
		q.Remove()
		add := elementsToAdd + i
		if !q.Add(add) {
			t.Errorf("failed to add element %d", add)
		}
	}
	obtained := q.Length()
	if obtained != elementsToRemove {
		t.Errorf("expected full queue length %d, obtained %d", elementsToAdd, obtained)
	}

	for i := 0; i < elementsToRemove; i++ {
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

//--

func testDefaultQueueLength(q queue.Queue, t *testing.T) {
	elementsToAdd := 1000
	testQueueLength(q, elementsToAdd, elementsToAdd, t)
}

func testQueueLength(q queue.Queue, elementsToRemoveAdd int, elementsToRemove int, t *testing.T) {
	obtained := q.Length()
	if obtained != 0 {
		t.Errorf("empty queue length is %d", obtained)
	}

	for i := 0; i < elementsToRemoveAdd; i++ {
		if !q.Add(i) {
			t.Errorf("failed to add element %d", i)
		}
		var expected int
		if i >= elementsToRemove {
			expected = elementsToRemove
		} else {
			expected = i + 1
		}
		obtained := q.Length()
		if obtained != expected {
			t.Errorf("adding: expected queue length %d, obtained %d", expected, obtained)
		}
	}
	for i := 0; i < elementsToRemove; i++ {
		q.Remove()
		expected := elementsToRemove - i - 1
		obtained := q.Length()
		if obtained != expected {
			t.Errorf("removing: expected queue length %d, obtained %d", expected, obtained)
		}
	}
}

//--

func testDefaultQueueGet(q queue.Queue, t *testing.T) {
	testQueueGet(q, 1000, func(iteration int, index int) int { return index }, t)
}

func testQueueGet(q queue.Queue, elementsToAdd int, result func(iteration int, index int) int, t *testing.T) {
	if testing.Short() {
		t.Skip("skipping Get test in short mode")
	}
	for i := 0; i < elementsToAdd; i++ {
		if !q.Add(i) {
			t.Errorf("failed to add element %d", i)
		}
		for j := 0; j < q.Length(); j++ {
			expected := result(i, j)
			obtained := q.Get(j).(int)
			if obtained != expected {
				t.Errorf("iteration %d index %d contains %d instead of %d", i, j, obtained, expected)
			}
		}
	}
}

//--

func testDefaultQueueGetNegative(q queue.Queue, t *testing.T) {
	testQueueGetNegative(q, 1000, func(iteration int, index int) int { return iteration + index + 1 }, t)
}

func testQueueGetNegative(q queue.Queue, elementsToAdd int, result func(iteration int, index int) int, t *testing.T) {
	if testing.Short() {
		t.Skip("skipping GetNegative test in short mode")
	}
	for i := 0; i < elementsToAdd; i++ {
		if !q.Add(i) {
			t.Errorf("failed to add element %d", i)
		}
		for j := -1; j >= -q.Length(); j-- {
			expected := result(i, j)
			obtained := q.Get(j).(int)
			if obtained != expected {
				t.Errorf("iteration %d index %d contains %d instead of %d", i, j, obtained, expected)
			}
		}
	}
}

//--

func testQueueGetOutOfRangePanics(q queue.Queue, t *testing.T) {
	for i := 0; i < 3; i++ {
		if !q.Add(i) {
			t.Errorf("failed to add element %d", i)
		}
	}

	assertPanics(t, "should panic when negative index", func() {
		q.Get(-4)
	})

	assertPanics(t, "should panic when index greater than length", func() {
		q.Get(4)
	})
}

//--

func testQueuePeekOutOfRangePanics(q queue.Queue, t *testing.T) {
	assertPanics(t, "should panic when peeking empty queue", func() {
		q.Peek()
	})

	if !q.Add(0) {
		t.Errorf("failed to add element 0")
	}
	q.Remove()

	assertPanics(t, "should panic when peeking emptied queue", func() {
		q.Peek()
	})
}

//--

func testQueueRemoveOutOfRangePanics(q queue.Queue, t *testing.T) {
	assertPanics(t, "should panic when removing empty queue", func() {
		q.Remove()
	})

	if !q.Add(0) {
		t.Errorf("failed to add element 0")
	}
	q.Remove()

	assertPanics(t, "should panic when removing emptied queue", func() {
		q.Remove()
	})
}

//--

func assertPanics(t *testing.T, name string, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("%s: didn't panic as expected", name)
		}
	}()

	f()
}
