# cobra-enum-1

Cobra with enum option.

Experiment 1: using [goflags][]/enumflag

The enum implementation is in [status.go](cmd/status.go), the main thing
required is the `Parse(string)` method for transforming from a string used in
CLI into the enum.

The main problem with the implementation here is the duplication of keys
between the map of string to status and the list of statuses as strings.

Also the `Parse(string)` method has to return a real `FooStatus` even if it
returns an error in which case it should be ignored. It works but it doesn't
seem great.

Although those aren't anything specific to [goflags][]/enumflag.

As for [goflags][]/enumflag, there is a bit of compatibility code required to
make it work with [cobra][], mostly due to [spf13/pflag#228][pflag-228]. And
the usage of [goflags][]/enumflag itself is in [list.go](cmd/list.go).

It basically sets a [flag.Value][flagv] object to a string given on command
line, then it's up to you to turn that into an actual enum. It does manage
throwing error in case an invalid value was given. What's nice about it is it
can generate a help message with all allowed values.

[cobra]: https://github.com/spf13/cobra
[flagv]: https://golang.org/pkg/flag/#Value
[goflags]: https://github.com/creachadair/goflags
[pflag-228]: https://github.com/spf13/pflag/issues/228
