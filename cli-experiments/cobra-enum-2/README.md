# cobra-enum-2

Cobra with enum option.

Experiment 2: using [enumflag][]

The enum implementation is in [status.go](cmd/status.go), the main thing
required is the map of statuses to string representation(s), but providing a
method for parsing is not required and there's less duplication.

The usage of [enumflag][] itself is in [list.go](cmd/list.go).

The code is much simpler however it requires the zero-value to be a value that
signifies _unspecified_, i.e. if the option was not specified on the CLI, the
variable will just default to the zero-value of the enum, which is it's first
element. This isn't really anything specific to [enumflag][] just how things
are in Go.

[enumflag]: https://github.com/thediveo/enumflag
