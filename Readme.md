# schmp
[![GoDoc](https://godoc.org/github.com/ayushg3112/schmp?status.svg)](https://pkg.go.dev/github.com/ayushg3112/schmp)
[![Master branch sanity](https://github.com/AyushG3112/schmp/actions/workflows/master-merge-checks.yml/badge.svg)](https://github.com/AyushG3112/schmp/actions/workflows/master-merge-checks.yml)

schmp is a schema comparision tool, which allows you to compare schemas of data in JSON, YAML and TOML format.
It is available as both a CLI, or a library which can be imported in your project.

This only compares the types of your data members, and ignores the values.

For objects or dictionaries, it traverses nested objects, as long as the type of that specific key is same in all the data samples provided.

Arrays are compared by the schema of their elements rather than their length: every element is folded onto a synthetic `<path>[]` key. So `[1, 2, 3]` and `[1]` match, while `["a"]` and `[1]` differ at `tags[]`. Objects inside arrays are traversed too (for example `users[].id`), and a mixed array such as `[1, "a"]` reports its element type as the set `int | string`.

**NOTE:** This package is close to stable and is awaiting feedback before a `v1.0.0` release.

# Why? 

Most projects use config files which store configuration settings which are different for different environments.
As new features are added, need for more global configuration settings arise, which then need to be updated across all environments on release.

The motivation to develop `schmp` arose as a requirement in our release process, to autoverify whether our config structure matches, in schema, not values, to our development or example config file.


## Installing as CLI

If you are not using Go:
 - Download the release from the [Releases](https://github.com/AyushG3112/schmp/releases) page and download the latest release archive for your OS and architecture.
 - Extract the archive, and move the extracted file to your `PATH` (optional, use complete path to file if skipping this step)


If you are using Go:

``` console
$ go install github.com/ayushg3112/schmp/cmd/schmp@latest
```


### Usage

```
$ schmp --help
Usage of schmp:
  -f, --file stringArray      Files to compare. Use this flag multiple times, once for each file.
  -m, --mode string           input file format.  Allowed values: json, yaml, toml (default "json")
      --out-file --out-type   Output file for the json --out-type. If omitted, JSON is written to stdout.
  -o, --out-type string       Output format. Allowed values: stdout, json (default "stdout")
```

schmp exits with status `7` when schema differences are found and `0` when there are none, which makes it convenient as a release/CI gate.

**Examples**:

```
# human-readable diff to stdout
$ schmp -f path/to/first/file.json -f path/to/second/file.json --out-type stdout --mode json

# JSON diff to stdout
$ schmp -f a.yaml -f b.yaml --mode yaml --out-type json

# JSON diff to a file
$ schmp -f a.toml -f b.toml --mode toml --out-type json --out-file diff.json
```


## Installing as a library

``` console
$ go get -u github.com/ayushg3112/schmp
```

**Usage Example**:

``` go
import (
  "github.com/ayushg3112/schmp"
  "io"
  "strings"
  "fmt"
)

func main() {
  reader1 := strings.NewReader(`{ "a": 1, "b": "2", "c": { "d": {"e" : 3}, "f": 4}}`)
  reader2 := strings.NewReader(`{ "b": "5", "c": { "d": "6", "f": 7}, "g": null}`)
  result, err := schmp.Compare(schmp.ProcessingOptions{
    Mode: "json",
    Sources: []io.Reader{
      reader1,
      reader2,
    },
  })
  if err != nil {
    panic(err)
  }
  fmt.Printf("%+v", result)
}
```