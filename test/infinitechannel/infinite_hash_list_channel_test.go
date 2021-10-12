package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityHashListChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashListChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, t)
}

func TestInfiniteHashLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteHashLimitedPriorityHashListChannel(func(elem interface{}) interface{} { return elem }), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityHashLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashListChannel(fun, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashListChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashListChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

//--

func TestInfiniteDefaultLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityHashListChannel(), 1000, &result, t)
}

func TestInfinitePriorityLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashListChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashListChannel(limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashListChannel(fun, limit)
	}, t)
}

func TestInfiniteHashLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteHashLimitedPriorityHashListChannel(func(elem interface{}) interface{} { return elem }), 1000, &result, t)
}

func TestInfinitePriorityHashLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashListChannel(fun, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashListChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashListChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashListChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashListChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}
