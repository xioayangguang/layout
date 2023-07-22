// Package evbus inspired by https://github.com/asaskevich/EventBus.git
// and make some changes to suitable for our project.
package evbus

import (
	"context"
	"layout/pkg/helper/uuid"
	"log"
	"reflect"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type Topic string

type Event interface {
	Topic() Topic
}

type Handler interface {
	Handle(ctx context.Context, event Event) error
}

type HandleFunc[T Event] func(ctx context.Context, event T) error

func (h HandleFunc[T]) Handle(ctx context.Context, event Event) error {
	e := event.(T)
	return h(ctx, e)
}

type CancelFunc func()

func Sub[T Event](bus *Bus, event T, h HandleFunc[T]) CancelFunc {
	//log.Printf("[evbus] subscribe topic %s, handler: %v\n", event.Topic(), getFuncName(h))
	return bus.sub(event.Topic(), h)
}

func SubOnce[T Event](bus *Bus, event T, h HandleFunc[T]) CancelFunc {
	//log.Printf("[evbus] subscribe topic %s once, handler: %v\n", event.Topic(), getFuncName(h))
	return bus.subOnce(event.Topic(), h)
}

func SubAsync[T Event](bus *Bus, event T, h HandleFunc[T]) CancelFunc {
	log.Printf("[evbus] subscribe topic %s async, handler: %v\n", event.Topic(), getFuncName(h))
	return bus.subAsync(event.Topic(), h)
}
func SubAsyncOnce[T Event](bus *Bus, event T, h HandleFunc[T]) CancelFunc {
	//log.Printf("[evbus] subscribe topic %s async once, handler: %v\n", event.Topic(), getFuncName(h))
	return bus.subAsyncOnce(event.Topic(), h)
}

func getFuncName(i interface{}) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

type eventUUID struct{}

var eventUUIDKey = eventUUID{}

// UUID returns the UUID of the event
func UUID(ctx context.Context) string {
	return ctx.Value(eventUUIDKey).(string)
}

type eventOccurredAt struct{}

var eventOccurredKey = eventOccurredAt{}

// OccurredAt returns the time when the event was published
func OccurredAt(ctx context.Context) time.Time {
	return ctx.Value(eventOccurredKey).(time.Time)
}

type Bus struct {
	handlers map[Topic][]*eventHandler
	lock     sync.RWMutex
	wg       sync.WaitGroup
}

var eventIdCounter int64

func nextHandlerId() int64 {
	return atomic.AddInt64(&eventIdCounter, 1)
}

type eventHandler struct {
	id           int64
	callback     Handler
	once         bool
	async        bool
	triggerCount int64
	sync.Mutex
}

func New() *Bus {
	b := &Bus{handlers: make(map[Topic][]*eventHandler)}
	return b
}

func (bus *Bus) addHandler(topic Topic, handler *eventHandler) CancelFunc {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	handler.id = nextHandlerId()
	bus.handlers[topic] = append(bus.handlers[topic], handler)

	return func() {
		bus.removeHandler(topic, handler)
	}
}

func (bus *Bus) sub(topic Topic, handler Handler) CancelFunc {
	return bus.addHandler(topic, &eventHandler{
		callback: handler,
	})
}

func (bus *Bus) subAsync(topic Topic, handler Handler) CancelFunc {
	return bus.addHandler(topic, &eventHandler{
		callback: handler,
		async:    true,
	})
}

func (bus *Bus) subOnce(topic Topic, handler Handler) CancelFunc {
	return bus.addHandler(topic, &eventHandler{
		callback: handler,
		once:     true,
	})
}

func (bus *Bus) subAsyncOnce(topic Topic, handler Handler) CancelFunc {
	return bus.addHandler(topic, &eventHandler{
		callback: handler,
		async:    true,
		once:     true,
	})
}

func (bus *Bus) Pub(ctx context.Context, event Event) error {
	bus.lock.RLock()

	var deadHandlers []*eventHandler
	topic := event.Topic()

	id := uuid.GenUUID()
	occurredAt := time.Now()

	ctx = context.WithValue(ctx, eventUUIDKey, id)
	ctx = context.WithValue(ctx, eventOccurredKey, occurredAt)

	log.Printf("event published, topic: %s, uuid: %v\n", event.Topic(), id)

	for _, h := range bus.handlers[topic] {
		if h.once {
			if atomic.LoadInt64(&h.triggerCount) > 0 {
				continue
			}
			// 记录将要移除的 handler
			deadHandlers = append(deadHandlers, h)
		}

		// 增加触发次数, 保证设置为一次性的 handler 只触发一次
		atomic.AddInt64(&h.triggerCount, 1)

		if h.async {
			bus.wg.Add(1)
			go func(h *eventHandler) {
				defer bus.wg.Done()
				_ = h.callback.Handle(ctx, event) // 异步调用忽略错误
			}(h)
		} else {
			if err := h.callback.Handle(ctx, event); err != nil {
				return err
			}
		}
	}

	bus.lock.RUnlock()

	for _, h := range deadHandlers {
		bus.removeHandler(topic, h)
	}

	return nil
}

func (bus *Bus) removeHandler(topic Topic, handler *eventHandler) {
	bus.lock.Lock()
	defer bus.lock.Unlock()

	if _, ok := bus.handlers[topic]; !ok {
		return
	}

	for i, h := range bus.handlers[topic] {
		if h.id == handler.id {
			bus.handlers[topic] = append(bus.handlers[topic][:i], bus.handlers[topic][i+1:]...)
			break
		}
	}
}

// WaitAsync waits for all async callbacks to complete
func (bus *Bus) WaitAsync() {
	bus.wg.Wait()
}
