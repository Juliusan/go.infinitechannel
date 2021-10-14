package queue

import "github.com/Juliusan/go.infinitechannel/src/set"

// LimitedPriorityHashQueue is an improvement of SimpleQueue, which
// can prioritise elements, limit its growth and rejects already included elements.
type LimitedPriorityHashQueue struct {
	buf         []interface{}
	head        int
	pend        int
	tail        int
	count       int
	priorityFun func(interface{}) bool
	limit       int
	hashTable   set.Set
}

var _ Queue = &LimitedPriorityHashQueue{}

// New constructs and returns a new Queue. Code duplication needed for benchmarks.
func NewDefaultLimitedPriorityHashQueue() *LimitedPriorityHashQueue {
	return NewHashLimitedPriorityHashQueue(false)
}

func NewPriorityLimitedPriorityHashQueue(fun func(interface{}) bool) *LimitedPriorityHashQueue {
	return NewPriorityHashLimitedPriorityHashQueue(fun, false)
}

func NewLimitLimitedPriorityHashQueue(limit int) *LimitedPriorityHashQueue {
	return NewLimitHashLimitedPriorityHashQueue(limit, false)
}

func NewLimitPriorityLimitedPriorityHashQueue(fun func(interface{}) bool, limit int) *LimitedPriorityHashQueue {
	return NewLimitedPriorityHashQueue(fun, limit, false)
}

func NewHashLimitedPriorityHashQueue(hashNeeded bool) *LimitedPriorityHashQueue {
	return NewLimitHashLimitedPriorityHashQueue(Infinity, hashNeeded)
}

func NewPriorityHashLimitedPriorityHashQueue(fun func(interface{}) bool, hashNeeded bool) *LimitedPriorityHashQueue {
	return NewLimitedPriorityHashQueue(fun, Infinity, hashNeeded)
}

func NewLimitHashLimitedPriorityHashQueue(limit int, hashNeeded bool) *LimitedPriorityHashQueue {
	return NewLimitedPriorityHashQueue(func(interface{}) bool { return false }, limit, hashNeeded)
}

func NewLimitedPriorityHashQueue(priorityFun func(interface{}) bool, limit int, hashNeeded bool) *LimitedPriorityHashQueue {
	var initBufSize int
	if (limit != Infinity) && (limit < minQueueLen) {
		initBufSize = limit
	} else {
		initBufSize = minQueueLen
	}
	var hashTable set.Set
	if hashNeeded {
		hashTable = set.NewHashSet()
	} else {
		hashTable = nil
	}
	return &LimitedPriorityHashQueue{
		head:        0,
		pend:        -1,
		tail:        0,
		count:       0,
		buf:         make([]interface{}, initBufSize),
		priorityFun: priorityFun,
		limit:       limit,
		hashTable:   hashTable,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *LimitedPriorityHashQueue) Length() int {
	return q.count
}

func (q *LimitedPriorityHashQueue) getIndex(rawIndex int) int {
	index := rawIndex % len(q.buf)
	if index < 0 {
		return index + len(q.buf)
	}
	return index
}

// resizes the queue to fit exactly twice its current contents
// this can result in shrinking if the queue is less than half-full
func (q *LimitedPriorityHashQueue) resize() {
	newSize := q.count << 1
	if newSize < minQueueLen {
		newSize = minQueueLen
	}
	if (q.limit != Infinity) && (newSize > q.limit) {
		newSize = q.limit
	}
	newBuf := make([]interface{}, newSize)

	if q.tail > q.head {
		copy(newBuf, q.buf[q.head:q.tail])
	} else {
		n := copy(newBuf, q.buf[q.head:])
		copy(newBuf[n:], q.buf[:q.tail])
	}

	if q.pend >= 0 {
		q.pend = q.getIndex(q.pend - q.head)
	}
	q.head = 0
	q.tail = q.count
	q.buf = newBuf
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *LimitedPriorityHashQueue) Add(elem interface{}) bool {
	if q.hashTable != nil && q.hashTable.Contains(elem) {
		// duplicate element; ignoring
		return false
	}
	limitReached := false
	if q.count == len(q.buf) {
		if (q.limit != Infinity) && (q.count >= q.limit) {
			limitReached = true
		} else {
			q.resize()
		}
	}
	priority := q.priorityFun(elem)
	if limitReached && !priority && (q.pend == q.getIndex(q.tail-1)) {
		//Not possible to add not priority element in queue full of priority elements
		return false
	}
	if limitReached {
		var deleteElem interface{}
		if q.pend < 0 {
			deleteElem = q.buf[q.head]
			q.head = q.getIndex(q.head + 1)
		} else {
			ptail := q.getIndex(q.pend + 1)
			if ptail == q.tail {
				deleteElem = q.buf[q.pend]
				q.tail = q.getIndex(q.tail - 1)
				q.pend = q.getIndex(q.pend - 1)
			} else {
				deleteElem = q.buf[ptail]
				if ptail > q.head {
					copy(q.buf[q.head+1:ptail+1], q.buf[q.head:ptail])
				} else {
					oldHead := q.buf[q.head]
					if ptail > 0 {
						copy(q.buf[1:ptail+1], q.buf[:ptail])
					}
					lastIndex := len(q.buf) - 1
					q.buf[0] = q.buf[lastIndex]
					if q.head < lastIndex {
						copy(q.buf[q.head+1:], q.buf[q.head:lastIndex])
					}
					q.buf[q.getIndex(q.head+1)] = oldHead
				}
				q.pend = q.getIndex(q.pend + 1)
				q.head = q.getIndex(q.head + 1)
			}
		}
		if q.hashTable != nil {
			q.hashTable.Remove(deleteElem)
		}
	}
	if priority {
		q.head = q.getIndex(q.head - 1)
		q.buf[q.head] = elem
		if q.pend < 0 {
			q.pend = q.head
		}
	} else {
		q.buf[q.tail] = elem
		// bitwise modulus
		q.tail = q.getIndex(q.tail + 1)
	}
	if !limitReached {
		q.count++
	}
	if q.hashTable != nil {
		q.hashTable.Add(elem)
	}
	return true
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *LimitedPriorityHashQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.buf[q.head]
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *LimitedPriorityHashQueue) Get(i int) interface{} {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		panic("queue: Get() called with index out of range")
	}
	// bitwise modulus
	return q.buf[q.getIndex(q.head+i)]
}

// Remove removes and returns the element from the front of the queue. If the
// queue is empty, the call will panic.
func (q *LimitedPriorityHashQueue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}
	ret := q.buf[q.head]
	q.buf[q.head] = nil
	if q.head == q.pend {
		q.pend = -1
	}
	q.head = q.getIndex(q.head + 1)
	q.count--
	// Resize down if buffer 1/4 full.
	if (len(q.buf) > minQueueLen) && ((q.count << 2) <= len(q.buf)) {
		q.resize()
	}
	if q.hashTable != nil {
		q.hashTable.Remove(ret)
	}
	return ret
}
