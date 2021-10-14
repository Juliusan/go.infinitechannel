package set

type Hashable interface {
	//For requirements of this function see https://docs.oracle.com/javase/8/docs/api/java/lang/Object.html#hashCode--
	//Additional requirement: the returned value must be valid as a map key
	GetHash() interface{}
	//For requirements of this function see https://docs.oracle.com/javase/8/docs/api/java/lang/Object.html#equals-java.lang.Object-
	Equals(elem interface{}) bool
}

type Set interface {
	Size() int
	Add(elem interface{}) bool
	Clear()
	Contains(elem interface{}) bool
	Remove(elem interface{}) bool
}
