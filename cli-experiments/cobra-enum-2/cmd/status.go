package cmd

// FooStatus is an enum of possible statuses of Foo
type FooStatus uint

const (
	StatusActive FooStatus = iota
	StatusDisabled
	StatusExpired
)

var FooStatusIds = map[FooStatus][]string{
	StatusActive:   {"ACTIVE"},
	StatusDisabled: {"DISABLED"},
	StatusExpired:  {"EXPIRED"},
}
