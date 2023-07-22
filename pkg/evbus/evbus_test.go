package evbus

import (
	"context"
	"github.com/stretchr/testify/assert"
	"sync/atomic"
	"testing"
)

type orderCreatedEvent struct {
	OrderId uint64
}

func (ev orderCreatedEvent) Topic() Topic {
	return "order.created"
}

type orderPaidEvent struct {
	OrderId uint64
	Amount  string
}

func (ev orderPaidEvent) Topic() Topic {
	return "order.paid"
}

type orderCanceledEvent struct {
	OrderId uint64
}

func (ev orderCanceledEvent) Topic() Topic {
	return "order.canceled"
}

func TestEvBus(t *testing.T) {
	bus := New()

	countCreated := 0
	countPaid := 0

	Sub(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		countCreated++
		assert.Equal(t, Topic("order.created"), event.Topic())
		assert.NotEmpty(t, UUID(ctx))
		assert.NotEmpty(t, OccurredAt(ctx))
		return nil
	})
	Sub(bus, orderPaidEvent{}, func(ctx context.Context, event orderPaidEvent) error {
		countPaid++
		assert.Equal(t, Topic("order.paid"), event.Topic())
		assert.Equal(t, "1000", event.Amount)
		assert.NotEmpty(t, UUID(ctx))
		assert.NotEmpty(t, OccurredAt(ctx))
		return nil
	})

	err := bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	assert.NoError(t, err)

	err = bus.Pub(context.Background(), orderPaidEvent{OrderId: 1, Amount: "1000"})
	assert.NoError(t, err)

	err = bus.Pub(context.Background(), orderPaidEvent{OrderId: 1, Amount: "1000"})
	assert.NoError(t, err)

	err = bus.Pub(context.Background(), orderCanceledEvent{OrderId: 1})
	assert.NoError(t, err)

	assert.Equal(t, 1, countCreated)
	assert.Equal(t, 2, countPaid)
}

func TestBroadcast(t *testing.T) {
	bus := New()
	count := 0
	Sub(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		count++
		return nil
	})
	Sub(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		count++
		return nil
	})

	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})

	assert.Equal(t, 4, count)
}

func TestCancel(t *testing.T) {
	bus := New()
	count1 := 0
	count2 := 0

	Sub(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		count1++
		return nil
	})
	cancel := Sub(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		count2++
		return nil
	})

	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	assert.Equal(t, 1, count1)
	assert.Equal(t, 1, count2)

	cancel()

	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	assert.Equal(t, 2, count1)
	assert.Equal(t, 1, count2)

	cancel() // 幂等
}

func TestNonSubscriber(t *testing.T) {
	bus := New()
	assert.NoError(t, bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1}))
}

func TestSubOnce(t *testing.T) {
	bus := New()

	var count int
	SubOnce(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		count++
		return nil
	})
	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCanceledEvent{OrderId: 1})

	assert.Equal(t, 1, count)
}

func TestSubAsync(t *testing.T) {
	bus := New()

	var count int64
	SubAsync(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	SubAsync(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		atomic.AddInt64(&count, 1)
		return nil
	})

	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCanceledEvent{OrderId: 1})

	bus.WaitAsync()

	assert.Equal(t, int64(4), atomic.LoadInt64(&count))
}

func TestSubAsyncOnce(t *testing.T) {
	bus := New()

	var count int64
	SubAsyncOnce(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		atomic.AddInt64(&count, 1)
		return nil
	})
	SubAsyncOnce(bus, orderCreatedEvent{}, func(ctx context.Context, event orderCreatedEvent) error {
		atomic.AddInt64(&count, 1)
		return nil
	})

	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCanceledEvent{OrderId: 1})
	_ = bus.Pub(context.Background(), orderCreatedEvent{OrderId: 1})

	bus.WaitAsync()

	assert.Equal(t, int64(2), atomic.LoadInt64(&count))
}
