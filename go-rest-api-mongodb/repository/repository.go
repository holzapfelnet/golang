package repository

// Repository interface represents all methods provided by the ToDoItem Repo.
type Repository interface {
	// Adds a new ToDo item to the collection
	AddItem(p ToDoItem) error
	// Gets all items stored within the collection
	GetAllItems() ([]*ToDoItem, error)
}
