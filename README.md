# calc-go

A flexible command-line calculator written in Go. Supports arithmetic, scientific functions, constants, and custom grouping with (), [], {}.

## Features
- Basic arithmetic: `+`, `-`, `*`, `/`, `^`
- Grouping: `()`, `[]`, `{}` (can be mixed)
- Scientific functions: `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `ln`, `log`, `sqrt`, `abs`, `ceil`, `floor`, `round`, `exp`, `max`, `min`, `gcd`, `lcm`, `factorial`, and more
- Constants: `pi`, `e`
- Multi-argument functions: e.g. `max(2, 5)`, `lcm(3, 4)`
- Clean output: prints integers without a decimal, trims trailing zeros for floats

## How to Use

Build the binary:

```
make build
```

Run interactively:

```
./bin/calc
```

Or pass an expression directly:

```
./bin/calc "sin(pi/2) + max(2, 5)"
```

## Example Expressions

```
1 + 2 * 3
sin(pi/2)
lcm(3, 4)
max(2, 5)
(2 + 3) * [4 - 1]
log(100)
2.5 + 3.1
```

## Supported Functions & Constants

- Arithmetic: `+`, `-`, `*`, `/`, `^`
- Grouping: `()`, `[]`, `{}`
- Functions: `sin`, `cos`, `tan`, `asin`, `acos`, `atan`, `ln`, `log`, `sqrt`, `abs`, `ceil`, `floor`, `round`, `exp`, `max`, `min`, `gcd`, `lcm`, `factorial`, `mod`
- Constants: `pi`, `e`

## Installation
You can install `calc-go` using `go install`:

```
go install github.com/amjadjibon/calc-go/cmd/calc@latest
```

Or download a pre-built binary from the [releases page](https://github.com/amjadjibon/calc-go/releases)

## Releasing

This project uses [GoReleaser](https://goreleaser.com/) for multi-platform releases. See `.goreleaser.yaml` for details.

## License

MIT License. See [LICENSE](LICENSE).
