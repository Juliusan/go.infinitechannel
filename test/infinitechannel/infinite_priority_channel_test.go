package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultPriorityChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultPriorityChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(fun)
	}, t)
}

func testInfinitePriorityChannelWriteReadLen(makeInfinitePriorityChannelFun func(func(i interface{}) bool) *infinitechannel.InfiniteChannel, t *testing.T) {
	ch := makeInfinitePriorityChannelFun(func(i interface{}) bool {
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
	testDefaultInfiniteChannelWriteReadLen(ch, 1000, result, t)
}

//--

func TestInfiniteDefaultPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultPriorityChannel(), 1000, &result, t)
}

func TestInfinitePriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityChannel(fun)
	}, t)
}

func testInfinitePriorityChannelConcurrentWriteReadLen(makeInfinitePriorityChannelFun func(func(i interface{}) bool) *infinitechannel.InfiniteChannel, t *testing.T) {
	ch := makeInfinitePriorityChannelFun(func(i interface{}) bool {
		return i.(int)%3 == 0
	})
	testDefaultInfiniteChannelConcurrentWriteReadLen(ch, 1000, nil, t)
}
