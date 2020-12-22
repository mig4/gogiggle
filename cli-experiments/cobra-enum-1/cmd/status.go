package cmd

import (
	"fmt"
	"strings"
)

// FooStatus is an enum of possible statuses of Foo
type FooStatus uint

const (
	StatusActive FooStatus = iota
	StatusDisabled
	StatusExpired
)

// strings are duplicated but in go maps don't have a defined order of keys

var fooStatusMap = map[string]FooStatus{
	"ACTIVE":   StatusActive,
	"DISABLED": StatusDisabled,
	"EXPIRED":  StatusExpired,
}

var fooStatuses = []string{"ACTIVE", "DISABLED", "EXPIRED"}

func Parse(s string) (FooStatus, error) {
	s = strings.ToUpper(s)
	for _, stat := range fooStatuses {
		if s == stat {
			return fooStatusMap[s], nil
		}
	}
	return StatusActive, fmt.Errorf("Not a valid FooStatus: %s", s)
}

func (s FooStatus) String() string {
	return fooStatuses[s]
}
