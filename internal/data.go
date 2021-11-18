// Package internal defines the types used to create data and their corresponding attributes
package internal

// Data is an activity that needs to be completed within a period of time
type Data struct {
	Key   string
	Value string
}

// Validate ...
func (u Data) Validate() error {
	if u.Key == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Key is required")
	}
	if u.Value == "" {
		return NewErrorf(ErrorCodeInvalidArgument, "Value must be stored")
	}
	return nil
}
