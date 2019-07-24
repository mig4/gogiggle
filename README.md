# gogiggle
Go Experiments

## TODO

- [ ] figure out what's the proper way to structure packages; according to
  https://github.com/golang/go/wiki/PackagePublishing a `src/` dir isn't
  something common and it's a good idea to have subdirs under repo for packages
  including one for main named after the executable it should produce,
  unfortunately that breaks `go build` because for some reason it wants to put
  the executable in the root of the repo which then conflicts with the package
  directory that already exists there
