package util

import (
	"github.com/Juliusan/go.infinitechannel/src/set"
)

type SimpleHashable int

var _ set.Hashable = SimpleHashable(0)

func (sh SimpleHashable) GetHash() interface{} {
	return sh
}

func (sh SimpleHashable) Equals(elem interface{}) bool {
	other, ok := elem.(SimpleHashable)
	if !ok {
		return false
	}
	return sh == other
}
