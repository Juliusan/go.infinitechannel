package set

type Set interface {
	Size() int
	Add(elem interface{}) bool
	Clear()
	Contains(elem interface{}) bool
	Remove(elem interface{}) bool
}
