# error-handling-1

This is just an experiment to see how error wrapping and handling works in
Go. E.g. case like some lower level component throws an error, and in a caller
we want to attach some details/context, e.g. we're doing something in a loop
and want to attach info on which item was being processed when the error
ocurred.

Here component in `pkg/comp1/` is the loop and it calls function from
`pkg/comp2/` which returns an error.
