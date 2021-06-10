package reflectio

// Setter allows for custom elements to be set with a custom method
type Setter interface {
	SetValueAsString(string) error
}
