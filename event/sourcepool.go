package event

import (
	"container/list"
	"log"
	"sync"

	chrononpb "github.com/vflopes/chronon/gen"
)

type SourcePool[R Reducer] struct {
	size      int
	evictList *list.List
	index     map[string]*list.Element

	eventStore    chrononpb.EventStoreClient
	snapshotStore chrononpb.SnapshotStoreClient

	readPool sync.Pool

	lock sync.RWMutex
}

func NewSourcePool[R Reducer](size int, eventStore chrononpb.EventStoreClient, snapshotStore chrononpb.SnapshotStoreClient, zeroVersionReducer R) *SourcePool[R] {

	if size <= 0 {
		log.Fatalln("SourcePool size must have a positive and non-zero size")
	}

	return &SourcePool[R]{
		size:      size,
		evictList: list.New(),
		index:     make(map[string]*list.Element),

		eventStore:    eventStore,
		snapshotStore: snapshotStore,

		readPool: sync.Pool{
			New: func() any {
				return NewSource("", zeroVersionReducer.Clone().(R), eventStore, snapshotStore)
			},
		},
	}

}

func (sp *SourcePool[R]) Take(key string) *Source[R] {

	source := sp.readPool.Get().(*Source[R])

	source.key = key
	source.reducer.Reset()

	return source

}

func (sp *SourcePool[R]) Return(source *Source[R]) {

	sp.readPool.Put(source)

}

func (sp *SourcePool[R]) Add(key string) (*Source[R], bool) {

	sp.lock.Lock()

	defer sp.lock.Unlock()

	return sp.add(key)

}

func (sp *SourcePool[R]) add(key string) (*Source[R], bool) {

	if elem, ok := sp.index[key]; ok {

		sp.evictList.MoveToFront(elem)

		return elem.Value.(*Source[R]), false

	}

	source := sp.Take(key)

	elem := sp.evictList.PushFront(source)

	sp.index[key] = elem

	evict := sp.evictList.Len() > sp.size

	if evict {

		sp.removeOldest()

	}

	return source, evict

}

func (sp *SourcePool[R]) RemoveOldest() (key string, source *Source[R], ok bool) {

	sp.lock.Lock()

	defer sp.lock.Unlock()

	if elem := sp.removeOldest(); elem != nil {

		source := elem.Value.(*Source[R])

		return source.key, source, true
	}

	return "", nil, false

}

func (sp *SourcePool[R]) removeOldest() *list.Element {

	elem := sp.evictList.Back()

	if elem != nil {

		sp.removeElement(elem)

	}

	return elem

}

func (sp *SourcePool[R]) removeElement(elem *list.Element) {

	sp.evictList.Remove(elem)

	source := elem.Value.(*Source[R])

	delete(sp.index, source.key)

}

func (sp *SourcePool[R]) Purge() {

	sp.lock.Lock()

	for k := range sp.index {

		delete(sp.index, k)

	}

	sp.evictList.Init()

	sp.lock.Unlock()

}

func (sp *SourcePool[R]) Keys() []string {

	sp.lock.RLock()

	keys := make([]string, len(sp.index))

	i := 0

	for elem := sp.evictList.Back(); elem != nil; elem = elem.Prev() {

		keys[i] = elem.Value.(*Source[R]).key

		i++

	}

	sp.lock.RUnlock()

	return keys

}

func (sp *SourcePool[R]) Len() int {

	sp.lock.RLock()

	length := sp.evictList.Len()

	sp.lock.RUnlock()

	return length

}

func (sp *SourcePool[R]) Resize(size int) int {

	sp.lock.Lock()

	diff := sp.Len() - size

	if diff < 0 {
		diff = 0
	}

	for i := 0; i < diff; i++ {
		sp.removeOldest()
	}

	sp.size = size

	sp.lock.Unlock()

	return diff

}

func (sp *SourcePool[R]) Contains(key string) bool {

	sp.lock.RLock()

	_, ok := sp.index[key]

	sp.lock.RUnlock()

	return ok

}

func (sp *SourcePool[R]) Peek(key string) (*Source[R], bool) {

	sp.lock.RLock()

	defer sp.lock.RUnlock()

	if elem, ok := sp.index[key]; ok {

		return elem.Value.(*Source[R]), true

	}

	return nil, false
}

func (sp *SourcePool[R]) Get(key string) (*Source[R], bool) {

	sp.lock.Lock()

	defer sp.lock.Unlock()

	return sp.get(key)

}

func (sp *SourcePool[R]) get(key string) (*Source[R], bool) {

	if elem, ok := sp.index[key]; ok {

		sp.evictList.MoveToFront(elem)

		source := elem.Value.(*Source[R])

		if source == nil {

			return nil, false

		}

		return source, true
	}

	return nil, false

}

func (sp *SourcePool[R]) GetOldest() (string, *Source[R], bool) {

	sp.lock.RLock()

	defer sp.lock.RUnlock()

	elem := sp.evictList.Back()

	if elem != nil {

		source := elem.Value.(*Source[R])

		return source.key, source, true

	}
	return "", nil, false
}
