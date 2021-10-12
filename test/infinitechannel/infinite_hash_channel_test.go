package infinitechannel_test

import (
	"testing"

	"github.com/Juliusan/go.infinitechannel/src/infinitechannel"
)

func TestInfiniteDefaultLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityHashChannel(), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, t)
}

func TestInfiniteHashLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testDefaultInfiniteChannelWriteReadLen(infinitechannel.NewInfiniteHashLimitedPriorityHashChannel(func(elem interface{}) interface{} { return elem }), 1000, func(index int) int { return index }, t)
}

func TestInfinitePriorityHashLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashChannel(fun, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashChannelNoLimitWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashChannelWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

//--

func TestInfiniteDefaultLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteDefaultLimitedPriorityHashChannel(), 1000, &result, t)
}

func TestInfinitePriorityLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityLimitedPriorityHashChannel(fun)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, t)
}

func TestInfiniteLimitLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitLimitedPriorityHashChannel(limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, t)
}

func TestInfiniteLimitPriorityLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitPriorityLimitedPriorityHashChannel(fun, limit)
	}, t)
}

func TestInfiniteHashLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	result := func(index int) int { return index }
	testDefaultInfiniteChannelConcurrentWriteReadLen(infinitechannel.NewInfiniteHashLimitedPriorityHashChannel(func(elem interface{}) interface{} { return elem }), 1000, &result, t)
}

func TestInfinitePriorityHashLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfinitePriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfinitePriorityHashLimitedPriorityHashChannel(fun, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitHashLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitLimitedPriorityChannelConcurrentWriteReadLen(func(limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitHashLimitedPriorityHashChannel(limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashChannelNoLimitConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelNoLimitConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}

func TestInfiniteLimitedPriorityHashChannelConcurrentWriteReadLen(t *testing.T) {
	testInfiniteLimitedPriorityChannelConcurrentWriteReadLen(func(fun func(i interface{}) bool, limit int) *infinitechannel.InfiniteChannel {
		return infinitechannel.NewInfiniteLimitedPriorityHashChannel(fun, limit, func(elem interface{}) interface{} { return elem })
	}, t)
}
