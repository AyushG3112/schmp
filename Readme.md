**THIS IS STILL A WORK IN PROGRESS**

# schmp


schmp is a schema comparision tool, which allows you to compare schemas of data in JSON, YAML and TOML format.
It is available as both a CLI, or a library which can be imported in your project.

This only compares the types of your data members, and ignores the values.

For objects or dictionaries, it traverses nested objects, as long as the type of that specific key is same in all the data samples provided.

## Installing as CLI

``` console
$ go install github.com/AyushG3112/schmp/cmd/schmp
```


### Usage

```
$ schmp -f path/to/first/file -f path/to/second/file --out-type stdout --mode json
  -f, --file stringArray   Files to compare. Use this flag multiple times, once for each file.
  -m, --mode string        file format (default "json")
  -o, --out-type string    Output format (default "stdout")
```

## Installing as a library

``` console
$ go get -u github.com/AyushG3112/schmp
```

### TODO:

 - ~~Handle diff of nested objects~~
 - ~~Format CLI output~~
 - ~~Allow writing of cli output to file~~
 - Add tests
 - ~~Add Documentation~~
 - Add examples
 