**THIS IS STILL A WORK IN PROGRESS**

# schmp


schmp is a schema comparision tool, which allows you to compare schemas of data in JSON, YAML and TOML format.
It is available as both a CLI, or a library which can be imported in your project.

## Installing as CLI

``` console
$ go install github.com/AyushG3112/schmp/cmd/schmp
```


### Usage

```
Usage of schmp:
  -f, --file stringArray   Files to compare. Use this flag multiple times, once for each file.
  -m, --mode string        file format (default "json")
  -o, --out-type string    Output format (default "stdout")
```

Example:

``` console
$ schmp -f path/to/first/file -f path/to/second/file --out-type stdout --mode json
```

## Installing as a library

``` console
$ go get -u github.com/AyushG3112/schmp
```

### TODO:

 - Format CLI output
 - Add tests
 - Add Documentation
 - Bug fixes
 - Add examples
 - Allow writing of cli output to file
 