package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityLimitedPriorityChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteLimitLimitedPriorityChannel(1200), 1000, func(index int) int { return index }, t)
}

func TestInfiniteLimitLimitedPriorityChannelWriteReadLen(t *testing.T) {
	limit := 800
	elementsToAdd := 1000
	indexDiff := elementsToAdd - limit
	result := func(index int) int {
		return index + indexDiff
	}
	testInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteLimitLimitedPriorityChannel(limit), elementsToAdd, limit, result, t)
}

func TestInfiniteLimitedPriorityChannelNoLimitWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityChannelWriteReadLen(t *testing.T) {
	limit := 800
	ch := infinitechannel.NewInfiniteLimitedPriorityChannel(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	result := func(index int) int {
		if index <= 333 {
			return -3*index + 999
		} else {
			if index%2 == 0 {
				return 3*index/2 - 200
			} else {
				return (3*index - 401) / 2
			}
		}
	}
	testInfiniteChannelWriteReadLen(ch, 1000, limit, result, t)
}

//--

func TestInfiniteDefaultLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityChannel(), 1000, &result, t)
}

func TestInfinitePriorityLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteLimitLimitedPriorityChannel(1200), 1000, &result, t)
}

func TestInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteLimitLimitedPriorityChannel(800), 1000, 800, nil, t)
}

func TestInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityChannel(fun, 1200)
	}, t)
}

func TestInfiniteLimitedPriorityChannelConcurrentWriteReadLen(t *testing.T) {
	limit := 800
	ch := infinitechannel.NewInfiniteLimitedPriorityChannel(func(i interface{}) bool {
		return i.(int)%3 == 0
	}, limit)
	testInfiniteChannelConcurrentWriteReadLen(ch, 1000, limit, nil, t)
}
