package watcher

import (
	"reflect"

	"github.com/ProtonMail/gluon/queue"
)

type Watcher[T any] struct {
	types   map[reflect.Type]struct{}
	eventCh *queue.QueuedChannel[T]
}

func New[T any](ofType ...T) *Watcher[T] {
	types := make(map[reflect.Type]struct{}, len(ofType))

	for _, t := range ofType {
		types[reflect.TypeOf(t)] = struct{}{}
	}

	return &Watcher[T]{
		types:   types,
		eventCh: queue.NewQueuedChannel[T](1, 1),
	}
}

func (w *Watcher[T]) IsWatching(event T) bool {
	if len(w.types) == 0 {
		return true
	}

	_, ok := w.types[reflect.TypeOf(event)]

	return ok
}

func (w *Watcher[T]) GetChannel() <-chan T {
	return w.eventCh.GetChannel()
}

func (w *Watcher[T]) Send(event T) bool {
	return w.eventCh.Enqueue(event)
}

func (w *Watcher[T]) Close() {
	w.eventCh.CloseAndDiscardQueued()
}
