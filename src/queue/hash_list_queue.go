package queue

import "github.com/Juliusan/go.infinitechannel/src/set"

// LimitedPriorityHashListQueue is an improvement of SimpleQueue, which
// can prioritise elements, limit its growth and rejects already included elements.
type LimitedPriorityHashListQueue struct {
	head        *LimitedPriorityListQueueElem
	ptail       *LimitedPriorityListQueueElem
	tail        *LimitedPriorityListQueueElem
	count       int
	priorityFun func(interface{}) bool
	limit       int
	hashTable   set.Set
}

var _ Queue = &LimitedPriorityHashListQueue{}

// New constructs and returns a new Queue. Code duplication needed for benchmarks.
func NewDefaultLimitedPriorityHashListQueue() *LimitedPriorityHashListQueue {
	return NewHashLimitedPriorityHashListQueue(false)
}

func NewPriorityLimitedPriorityHashListQueue(fun func(interface{}) bool) *LimitedPriorityHashListQueue {
	return NewPriorityHashLimitedPriorityHashListQueue(fun, false)
}

func NewLimitLimitedPriorityHashListQueue(limit int) *LimitedPriorityHashListQueue {
	return NewLimitHashLimitedPriorityHashListQueue(limit, false)
}

func NewLimitPriorityLimitedPriorityHashListQueue(fun func(interface{}) bool, limit int) *LimitedPriorityHashListQueue {
	return NewLimitedPriorityHashListQueue(fun, limit, false)
}

func NewHashLimitedPriorityHashListQueue(hashNeeded bool) *LimitedPriorityHashListQueue {
	return NewLimitHashLimitedPriorityHashListQueue(Infinity, hashNeeded)
}

func NewPriorityHashLimitedPriorityHashListQueue(fun func(interface{}) bool, hashNeeded bool) *LimitedPriorityHashListQueue {
	return NewLimitedPriorityHashListQueue(fun, Infinity, hashNeeded)
}

func NewLimitHashLimitedPriorityHashListQueue(limit int, hashNeeded bool) *LimitedPriorityHashListQueue {
	return NewLimitedPriorityHashListQueue(func(interface{}) bool { return false }, limit, hashNeeded)
}

func NewLimitedPriorityHashListQueue(priorityFun func(interface{}) bool, limit int, hashNeeded bool) *LimitedPriorityHashListQueue {
	var hashTable set.Set
	if hashNeeded {
		hashTable = set.NewHashSet()
	} else {
		hashTable = nil
	}
	return &LimitedPriorityHashListQueue{
		priorityFun: priorityFun,
		limit:       limit,
		hashTable:   hashTable,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *LimitedPriorityHashListQueue) Length() int {
	return q.count
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *LimitedPriorityHashListQueue) Add(value interface{}) bool {
	if q.hashTable != nil && q.hashTable.Contains(value) {
		// duplicate element; ignoring
		return false
	}
	priority := q.priorityFun(value)
	if q.head == nil && q.tail == nil {
		elem := &LimitedPriorityListQueueElem{
			value: value,
			next:  nil,
			prev:  nil,
		}
		q.head = elem
		q.tail = elem
		if priority {
			q.ptail = elem
		}
		q.count++
		if q.hashTable != nil {
			q.hashTable.Add(value)
		}
		return true
	} else {
		limitReached := (q.limit != Infinity) && (q.count >= q.limit)
		if limitReached && !priority && (q.ptail == q.tail) {
			//Not possible to add not priority element in queue full of priority elements
			return false
		}
		elem := &LimitedPriorityListQueueElem{value: value}
		if priority {
			elem.next = q.head
			elem.prev = nil
			q.head.prev = elem
			q.head = elem
			if q.ptail == nil {
				q.ptail = elem
			}
		} else {
			elem.next = nil
			elem.prev = q.tail
			q.tail.next = elem
			q.tail = elem
		}
		if limitReached { //then delete one element
			var deleteElem *LimitedPriorityListQueueElem
			if priority {
				//here ptail != nil - at least added [priority] element is present
				if q.ptail == q.tail {
					deleteElem = q.ptail
					q.ptail = q.ptail.prev
				} else {
					deleteElem = q.ptail.next
				}
				deleteElem.prev.next = deleteElem.next
				if deleteElem.next == nil {
					q.tail = deleteElem.prev
				} else {
					deleteElem.next.prev = deleteElem.prev
				}
			} else {
				if q.ptail == nil {
					deleteElem = q.head
				} else {
					deleteElem = q.ptail.next
				}
				if deleteElem.prev == nil {
					q.head = deleteElem.next
				} else {
					deleteElem.prev.next = deleteElem.next
				}
				deleteElem.next.prev = deleteElem.prev
			}
			if q.hashTable != nil {
				q.hashTable.Remove(deleteElem.value)
			}
		} else {
			q.count++
		}
		if q.hashTable != nil {
			q.hashTable.Add(value)
		}
		return true
	}
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *LimitedPriorityHashListQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.head.value
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *LimitedPriorityHashListQueue) Get(i int) interface{} {
	// If indexing backwards, convert to positive index.
	if i < 0 {
		i += q.count
	}
	if i < 0 || i >= q.count {
		panic("queue: Get() called with index out of range")
	}
	elem := q.head
	for j := 0; j < i; j++ {
		elem = elem.next
	}
	// bitwise modulus
	return elem.value
}

// Remove removes and returns the element from the front of the queue. If the
// queue is empty, the call will panic.
func (q *LimitedPriorityHashListQueue) Remove() interface{} {
	if q.count <= 0 {
		panic("queue: Remove() called on empty queue")
	}
	value := q.head.value
	if q.head == q.ptail {
		q.ptail = nil
	}
	if q.head == q.tail {
		q.tail = nil
	}
	if q.head.next != nil {
		q.head.next.prev = nil
	}
	q.head = q.head.next
	q.count--
	if q.hashTable != nil {
		q.hashTable.Remove(value)
	}
	return value
}
