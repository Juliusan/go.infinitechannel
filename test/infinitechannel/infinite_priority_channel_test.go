package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultPriorityChannelWriteReadLen(t *testing.T) {
	testInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultPriorityChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityChannelWriteReadLen(t *testing.T) {
	ch := infinitechannel.NewInfinitePriorityChannel(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	result := func(index int) int {
		if index <= 333 {
			return -3*index + 999
		} else {
			if index%2 == 0 {
				return 3*index/2 - 500
			} else {
				return (3*index - 1001) / 2
			}
		}
	}
	testInfiniteChannelWriteReadLen(ch, 1000, result, t)
}

func TestInfiniteDefaultPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultPriorityChannel(), 1000, &result, t)
}

func TestInfinitePriorityChannelConcurrentWriteReadLen(t *testing.T) {
	ch := infinitechannel.NewInfinitePriorityChannel(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	testInfiniteChannelConcurrentWriteReadLen(ch, 1000, nil, t)
}
