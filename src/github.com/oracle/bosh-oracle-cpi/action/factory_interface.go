package action

// Factory interface for creating an Action.
type Factory interface {
	// Create an action for the given request method
	Create(method string) (Action, error)
}
