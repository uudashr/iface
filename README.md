[![Go Reference](https://pkg.go.dev/badge/github.com/uudashr/iface.svg)](https://pkg.go.dev/github.com/uudashr/iface)

# iface

`iface` is a linter designed to identify the incorrect use of interfaces in Go code, helping developers avoid interface pollution. By detecting unnecessary or poorly implemented interfaces, `iface` ensures your Go code remains clean, efficient, and maintainable.

It consists of several analyzers:
1. `unused`: Identifies interfaces that are not used anywhere in the same package where the interface is defined.
2. `identical`: Identifies interfaces in the same package with identical methods or constraints.
3. `opaque`: Identifies functions that return interfaces, but the actual returned value is always a single concrete implementation.
4. `unexported`: Identifies interfaces that are not exported but are used in exported functions or methods as parameters or return values.

## Usage

To install the linter which has all the checks:
```sh
go install github.com/uudashr/iface/cmd/ifacecheck@latest
```

To install individual linter, use the following command:
```sh
go install github.com/uudashr/iface/unused/cmd/unusediface@latest
go install github.com/uudashr/iface/identical/cmd/identicaliface@latest
go install github.com/uudashr/iface/opaque/cmd/opaqueiface@latest
go install github.com/uudashr/iface/unexported/cmd/unexportediface@latest
```

Run the linter
```sh
ifacecheck ./...
```

or show the help
```sh
ifacecheck help

# or
ifacecheck help <analyzer-name>
```

## Exclusion

### Package exclusion

We encourage to use default behavior and put effot to follow the rules. But, for some reason rules are not applicable. Due to this we can exclude specific package to be scanned by the analyzers. Use `-unused.exclude` flag and currently only `unused` has this feature. See help for more information:

Example usage: 
```sh
ifacecheck -unused.exclude=github.com/example/log ./...
```

### Ignore Directive

Exclusion can be done by using directive in the code by placing the `//iface:ignore` to ignore the code from being scanned by the analyzer. Example:

1. `iface:ignore` to ignore the from all analyzers.
2. `iface:ignore=[analyzer names]` which names is comma separators to exclude from defined names only. Ex: 
    - `iface:ignore=unused` ignore from `unused` analyzer.
    - `iface:ignore=unused,identical` ignore from `unused` and `identical` analyzers.


Note: use exclusion with careful consideration.

## Background

One of Go's powerful features is interfaces. However, sometimes people misuse the interfaces event though the code works but the code polluted with interfaces.

The following quotes inspired the creation of these analyzers:

> "Go interfaces generally belong in the package that uses values of the interface type, not the package that implements those values. The implementing package should return concrete (usually pointer or struct) types: that way, new methods can be added to implementations without requiring extensive refactoring."
>
> [Go Code Review Comments, go.dev](https://go.dev/wiki/CodeReviewComments#interfaces)


> "Don’t export any interfaces until you have to."
>
> [Interface pollution in Go, rakyll.org](https://rakyll.org/interface-pollution/)


> "The use of interfaces when they are not necessary is called interface pollution."
> 
> [Avoid Interface Pollution, ardanlabs.com](https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html)

> "Go interfaces generally belong in the package that consumes values of the interface type, not a package that implements the interface type."
>
> [Go Style Decisions, google.github.io](https://google.github.io/styleguide/go/decisions#interfaces)

