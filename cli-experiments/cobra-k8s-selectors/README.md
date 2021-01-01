# cobra-k8s-selectors

Cobra with K8s field and label selectors for filtering a list of objects.

Using [k8s.io/apimachinery][apimachinery]

The difference between a field selector and a label selector is that a field
selector is a simple set of `field == value` matchers, while a label selector
allows for more complex expressions like `field > num`, `!field` or
`field in (val1, val2)`.

The app architecture is hexagonal (ports and adapters), good references for
that are:

- [Hexagonal architecture - Wikipedia][wp-hex-arch]
- [Implementing Ports and Adapters by Yuri Khomyakov][yk-impl-pa]

Project layout is inspired by [standard go layout][gostd-layout].

[apimachinery]: https://github.com/kubernetes/apimachinery
[gostd-layout]: https://github.com/golang-standards/project-layout
[wp-hex-arch]: https://en.wikipedia.org/wiki/Hexagonal_architecture_(software)
[yk-impl-pa]: https://yuriktech.com/2020/02/01/Implementing-Ports-and-Adapters/
