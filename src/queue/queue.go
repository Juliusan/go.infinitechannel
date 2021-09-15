package queue

// minQueueLen is smallest capacity that queue may have.
// Must be power of 2 for bitwise modulus: x % n == x & (n - 1).
const minQueueLen = 16

type Queue interface {
	Length() int
	Add(elem interface{})
	Peek() interface{}
	Get(i int) interface{}
	Remove() interface{}
}
