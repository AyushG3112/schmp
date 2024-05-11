# schmp
[![GoDoc](https://godoc.org/github.com/ayushg3112/schmp?status.svg)](https://pkg.go.dev/github.com/ayushg3112/schmp)
[![Build Status](https://travis-ci.org/AyushG3112/schmp.svg?branch=master)](https://travis-ci.org/AyushG3112/schmp)

schmp is a schema comparision tool, which allows you to compare schemas of data in JSON, YAML and TOML format.
It is available as both a CLI, or a library which can be imported in your project.

This only compares the types of your data members, and ignores the values.

For objects or dictionaries, it traverses nested objects, as long as the type of that specific key is same in all the data samples provided.

**NOTE:** This package is currently considered stable and no further changes are intended. I am only waiting on some feedback before a `v1.0.0` release.

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
      --out-file --out-type   Output file. Only used if --out-type is not stdout
  -o, --out-type string       Output format. Allowed values: stdout, json (default "stdout")
```

**Example**:

```
$ schmp -f path/to/first/file.json -f path/to/second/file.json --out-type stdout --mode json
```


## Installing as a library

``` console
$ go get -u github.com/ayushg3112/schmp
```

**Usage Example**:

``` go
import (
  "github.com/ayushg3112/schmp"
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