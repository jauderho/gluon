package store

import (
	"io"
	"sync"
	"sync/atomic"

	"github.com/ProtonMail/gluon/imap"
)

type syncRef struct {
	lock    sync.RWMutex
	counter int32
}

// WriteControlledStore ensures that a given file on disk can safely be accessed by multiple readers and only
// one writer. Internally we maintain a list of RWLocks per message ID.
type WriteControlledStore struct {
	impl Store

	lock       sync.Mutex
	entryTable map[imap.InternalMessageID]*syncRef
	lockPool   sync.Pool
}

func NewWriteControlledStore(impl Store) *WriteControlledStore {
	return &WriteControlledStore{
		impl:       impl,
		entryTable: make(map[imap.InternalMessageID]*syncRef),
		lockPool: sync.Pool{
			New: func() any {
				return &syncRef{counter: 1}
			},
		},
	}
}

func (w *WriteControlledStore) acquireSyncRef(id imap.InternalMessageID) *syncRef {
	w.lock.Lock()
	defer w.lock.Unlock()

	v, ok := w.entryTable[id]
	if !ok {
		v, ok := w.lockPool.Get().(*syncRef)
		if !ok {
			panic("invalid Type Cast")
		}

		v.counter = 1
		w.entryTable[id] = v

		return v
	}

	atomic.AddInt32(&v.counter, 1)

	return v
}

func (w *WriteControlledStore) releaseSyncRef(id imap.InternalMessageID, ref *syncRef) {
	if atomic.AddInt32(&ref.counter, -1) <= 0 {
		w.lock.Lock()
		defer w.lock.Unlock()

		if atomic.LoadInt32(&ref.counter) <= 0 {
			delete(w.entryTable, id)
			w.lockPool.Put(ref)
		}
	}
}

func (w *WriteControlledStore) Get(messageID imap.InternalMessageID) ([]byte, error) {
	syncRef := w.acquireSyncRef(messageID)
	defer w.releaseSyncRef(messageID, syncRef)

	syncRef.lock.RLock()
	defer syncRef.lock.RUnlock()

	return w.impl.Get(messageID)
}

func (w *WriteControlledStore) Set(messageID imap.InternalMessageID, reader io.Reader) error {
	syncRef := w.acquireSyncRef(messageID)
	defer w.releaseSyncRef(messageID, syncRef)

	syncRef.lock.Lock()
	defer syncRef.lock.Unlock()

	return w.impl.Set(messageID, reader)
}

// SetUnchecked allows the user to bypass lock access. This will only work if you can guarantee that the data being
// set does not previously exit (e.g: New message).
func (w *WriteControlledStore) SetUnchecked(messageID imap.InternalMessageID, reader io.Reader) error {
	return w.impl.Set(messageID, reader)
}

func (w *WriteControlledStore) Delete(messageID ...imap.InternalMessageID) error {
	for _, id := range messageID {
		if err := func() error {
			syncRef := w.acquireSyncRef(id)
			defer w.releaseSyncRef(id, syncRef)

			syncRef.lock.Lock()
			defer syncRef.lock.Unlock()

			return w.impl.Delete(id)
		}(); err != nil {
			return err
		}
	}

	return nil
}

func (w *WriteControlledStore) Close() error {
	return w.impl.Close()
}

func (w *WriteControlledStore) List() ([]imap.InternalMessageID, error) {
	return w.impl.List()
}

type WriteControlledStoreBuilder struct {
	builder Builder
}

func NewWriteControlledStoreBuilder(builder Builder) *WriteControlledStoreBuilder {
	return &WriteControlledStoreBuilder{builder: builder}
}

func (w *WriteControlledStoreBuilder) New(dir, userID string, passphrase []byte) (Store, error) {
	impl, err := w.builder.New(dir, userID, passphrase)
	if err != nil {
		return nil, err
	}

	return NewWriteControlledStore(impl), nil
}

func (w *WriteControlledStoreBuilder) Delete(dir, userID string) error {
	return w.builder.Delete(dir, userID)
}
