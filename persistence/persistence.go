package persistence

import (
	"sync"

	"github.com/LettsDoSomeCoding/ranko/logging"
)

// Persistable will be implemented by any objects intended to be persisted for Ranko
type Persistable interface {
	getIdentifier() (id int, exists bool)
	setIdentifier(id int)
	Save() (Persistable, error)
}

// todo = get a real database
type databaseTable struct {
	nextID   int
	contents map[int]Persistable
	mux      sync.Mutex
}

func (t *databaseTable) save(obj Persistable) (Persistable, error) {
	if _, exists := obj.getIdentifier(); exists {
		return t.update(obj)
	}
	return t.insert(obj)
}

func (t *databaseTable) insert(obj Persistable) (Persistable, error) {
	t.mux.Lock()
	defer t.mux.Unlock()

	id := t.nextID
	t.nextID++
	obj.setIdentifier(id)
	t.contents[id] = obj
	return t.contents[id], nil
}

func (t *databaseTable) update(obj Persistable) (Persistable, error) {
	id, _ := obj.getIdentifier()

	t.mux.Lock()
	defer t.mux.Unlock()

	if _, present := t.contents[id]; !present {
		err := ObjectNotExistsError{obj}
		logging.GetLogger().Println(err.Error())
		return obj, err
	}

	t.contents[id] = obj
	return t.contents[id], nil
}
