package cmd

import "flag"

// CompatValue provides compatibility code to allow using
// github.com/creachadair/goflags/enumflag
// with pflag/cobra.
// It can't be used as is because of
// https://github.com/spf13/pflag/issues/228
// so `pflag.Value` interface declares a `Type()` method, while `flag.Value`
// doesn't, and so `enumflag.Value` doesn't implement it.
// This wraps a `flag.Value`, forwarding `String()` and `Set()` methods to it,
// and implements a `Type()` method
type CompatValue struct {
	wrapped flag.Value
}

func (cv CompatValue) String() string {
	return cv.wrapped.String()
}

func (cv CompatValue) Set(s string) error {
	return cv.wrapped.Set(s)
}

func (cv CompatValue) Type() string {
	return ""
}
