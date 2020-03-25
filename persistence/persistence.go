package persistence

import (
	"github.com/LettsDoSomeCoding/ranko/logging"
)

// Persistable will be implemented by any objects intended to be persisted for Ranko
type Persistable interface {
	GetIdentifier() string
	Save() (Persistable, error)
	Update() (Persistable, error)
	SaveOrUpdate() (Persistable, error)
}

// todo - get a real database
var playerDatabase = make([]Persistable, 0, 100)

func save(obj Persistable) (Persistable, error) {
	database := getDatabase(obj)
	_, existing := findExisting(*database, obj)

	if existing != nil {
		err := ObjectAlreadyExistsError{obj}
		logging.GetLogger().Println(err.Error())
		return obj, err
	}

	*database = append(*database, obj)
	return (*database)[len(*database)-1], nil
}

func update(obj Persistable) (Persistable, error) {
	database := *getDatabase(obj)
	index, existing := findExisting(database, obj)

	if existing == nil {
		err := ObjectNotExistsError{obj}
		logging.GetLogger().Println(err.Error())
		return obj, err
	}

	database[index] = obj
	return database[index], nil
}

func saveOrUpdate(obj Persistable) (Persistable, error) {
	database := getDatabase(obj)
	_, existing := findExisting(*database, obj)

	var result Persistable
	var err error

	if existing == nil {
		result, err = save(obj)
	} else {
		result, err = update(obj)
	}

	return result, err
}

func getDatabase(obj Persistable) *[]Persistable {
	switch obj.(type) {
	case Player:
		return &playerDatabase
	default:
		err := UnpersistableObjectError{obj}
		logging.GetLogger().Println(err.Error())
		panic(err)
	}
}

func findExisting(database []Persistable, obj Persistable) (index int, existing Persistable) {
	for index, existing := range database {
		if existing.GetIdentifier() == obj.GetIdentifier() {
			return index, existing
		}
	}

	return 0, nil
}
