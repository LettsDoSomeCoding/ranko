package persistence

import "fmt"

// UnpersistableObjectError is thrown if persistence operations are attempted on an unsupported object
type UnpersistableObjectError struct {
	obj interface{}
}

func (e UnpersistableObjectError) Error() string {
	return fmt.Sprintf("Attempted persistence operations for object %v of type %T which is not persistable", e.obj, e.obj)
}

// ObjectAlreadyExistsError is thrown on an attempt to save an existing Persistable
type ObjectAlreadyExistsError struct {
	obj Persistable
}

func (e ObjectAlreadyExistsError) Error() string {
	return fmt.Sprintf("Attempted to save %T %v which already exists", e.obj, e.obj)
}

// ObjectNotExistsError is thrown on an attempt to access a non-existent Persistable
type ObjectNotExistsError struct {
	obj Persistable
}

func (e ObjectNotExistsError) Error() string {
	return fmt.Sprintf("Attempted to access %T %v which does not exist", e.obj, e.obj)
}
