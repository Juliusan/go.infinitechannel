package queue

// LimitedPriorityHashQueue is an improvement of SimpleQueue, which
// can prioritise elements, limit its growth and rejects already included elements.
type LimitedPriorityHashQueue struct {
	head        *LimitedPriorityQueueElem
	ptail       *LimitedPriorityQueueElem
	tail        *LimitedPriorityQueueElem
	count       int
	priorityFun func(interface{}) bool
	limit       int
	hash        *map[interface{}]bool
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

func NewLimitedPriorityHashQueue(fun func(interface{}) bool, limit int, hashNeeded bool) *LimitedPriorityHashQueue {
	var hash *map[interface{}]bool
	if hashNeeded {
		hashT := make(map[interface{}]bool)
		hash = &hashT
	} else {
		hash = nil
	}
	return &LimitedPriorityHashQueue{
		priorityFun: fun,
		limit:       limit,
		hash:        hash,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *LimitedPriorityHashQueue) Length() int {
	return q.count
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *LimitedPriorityHashQueue) Add(value interface{}) bool {
	if q.hash != nil {
		_, exists := (*(q.hash))[value]
		if exists {
			// duplicate element; ignoring
			return false
		}
	}
	priority := q.priorityFun(value)
	if q.head == nil && q.tail == nil {
		elem := &LimitedPriorityQueueElem{
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
		if q.hash != nil {
			(*(q.hash))[value] = true
		}
		return true
	} else {
		limitReached := (q.limit != Infinity) && (q.count >= q.limit)
		if limitReached && !priority && (q.ptail == q.tail) {
			//Not possible to add not priority element in queue full of priority elements
			return false
		}
		elem := &LimitedPriorityQueueElem{value: value}
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
			var deleteElem *LimitedPriorityQueueElem
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
			if q.hash != nil {
				delete(*(q.hash), deleteElem.value)
			}
		} else {
			q.count++
		}
		if q.hash != nil {
			(*(q.hash))[value] = true
		}
		return true
	}
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *LimitedPriorityHashQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.head.value
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
	elem := q.head
	for j := 0; j < i; j++ {
		elem = elem.next
	}
	// bitwise modulus
	return elem.value
}

// Remove removes and returns the element from the front of the queue. If the
// queue is empty, the call will panic.
func (q *LimitedPriorityHashQueue) Remove() interface{} {
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
	if q.hash != nil {
		delete(*(q.hash), value)
	}
	return value
}
