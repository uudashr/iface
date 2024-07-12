[![Go Reference](https://pkg.go.dev/badge/github.com/uudashr/iface.svg)](https://pkg.go.dev/github.com/uudashr/iface)

# iface
`iface` is a linter designed to identify the incorrect use of interfaces in Go code, helping developers avoid interface pollution. By detecting unnecessary or poorly implemented interfaces, `iface` ensures your Go code remains clean, efficient, and maintainable.

It consists of several analyzers :
1. `empty`: Finds empty interfaces.
2. `unused`: Finds unused interfaces within the package.
3. `duplicate`: Finds duplicate interfaces within the package.
4. `opaque`: Find the interfaces that is used to abstract a single concrete implementation only.

## Usage

To install the linter which has all the checks:
```sh
go install github.com/uudashr/iface/cmd/ifacecheck@latest
```

To install individual linter, use the following command:
```sh
go install github.com/uudashr/iface/empty/cmd/emptyiface@latest
go install github.com/uudashr/iface/unused/cmd/unusediface@latest
go install github.com/uudashr/iface/duplicate/cmd/duplicateiface@latest
go install github.com/uudashr/iface/opaque/cmd/opaqueiface@latest
```

Run the linter
```sh
ifacecheck ./...
```

or show the help
```sh
ifacecheck help
```

## Related Articles
- [Interface Pollution](https://rakyll.org/interface-pollution/)
- [Avoid Interface Pollution](https://www.ardanlabs.com/blog/2016/10/avoid-interface-pollution.html)

