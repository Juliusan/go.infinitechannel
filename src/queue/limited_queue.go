package queue

type LimitedPriorityQueueElem struct {
	value interface{}
	next  *LimitedPriorityQueueElem
	prev  *LimitedPriorityQueueElem
}

// LimittedPriorityQueuePriorityQueue is an improvement of SimpleQueue, which
// can prioritise elements and limit its growth. When the limit is reached, oldest
// elements are discarded.
type LimitedPriorityQueue struct {
	head        *LimitedPriorityQueueElem
	ptail       *LimitedPriorityQueueElem
	tail        *LimitedPriorityQueueElem
	count       int
	priorityFun func(interface{}) bool
	limit       int
}

var _ Queue = &LimitedPriorityQueue{}

const Infinity = 0

// New constructs and returns a new Queue. Code duplication needed for benchmarks.
func NewDefaultLimitedPriorityQueue() *LimitedPriorityQueue {
	return NewLimitLimitedPriorityQueue(Infinity)
}

func NewPriorityLimitedPriorityQueue(fun func(interface{}) bool) *LimitedPriorityQueue {
	return NewLimitedPriorityQueue(fun, Infinity)
}

func NewLimitLimitedPriorityQueue(limit int) *LimitedPriorityQueue {
	return NewLimitedPriorityQueue(func(interface{}) bool { return false }, limit)
}

func NewLimitedPriorityQueue(fun func(interface{}) bool, limit int) *LimitedPriorityQueue {
	return &LimitedPriorityQueue{
		priorityFun: fun,
		limit:       limit,
	}
}

// Length returns the number of elements currently stored in the queue.
func (q *LimitedPriorityQueue) Length() int {
	return q.count
}

// Add puts an element to the start of end of the queue, depending
// on the result of priorityFun.
func (q *LimitedPriorityQueue) Add(value interface{}) bool {
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
			var delete *LimitedPriorityQueueElem
			if priority {
				//here ptail != nil - at least added [priority] element is present
				if q.ptail == q.tail {
					delete = q.ptail
					q.ptail = q.ptail.prev
				} else {
					delete = q.ptail.next
				}
				delete.prev.next = delete.next
				if delete.next == nil {
					q.tail = delete.prev
				} else {
					delete.next.prev = delete.prev
				}
			} else {
				if q.ptail == nil {
					delete = q.head
				} else {
					delete = q.ptail.next
				}
				if delete.prev == nil {
					q.head = delete.next
				} else {
					delete.prev.next = delete.next
				}
				delete.next.prev = delete.prev
			}
		} else {
			q.count++
		}
		return true
	}
}

// Peek returns the element at the head of the queue. This call panics
// if the queue is empty.
func (q *LimitedPriorityQueue) Peek() interface{} {
	if q.count <= 0 {
		panic("queue: Peek() called on empty queue")
	}
	return q.head.value
}

// Get returns the element at index i in the queue. If the index is
// invalid, the call will panic. This method accepts both positive and
// negative index values. Index 0 refers to the first element, and
// index -1 refers to the last.
func (q *LimitedPriorityQueue) Get(i int) interface{} {
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
func (q *LimitedPriorityQueue) Remove() interface{} {
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
	return value
}
