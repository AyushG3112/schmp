# schmp


schmp is a schema comparision tool, which allows you to compare schemas of data in JSON, YAML and TOML format.
It is available as both a CLI, or a library which can be imported in your project.

This only compares the types of your data members, and ignores the values.

For objects or dictionaries, it traverses nested objects, as long as the type of that specific key is same in all the data samples provided.


# Why? 

Most projects use config files which store configuration settings which are different for different environments.
As new features are added, need for more global configuration settings arise, which then need to be updated across all environments on release.

The motivation to develop `schmp` arose as a requirement in our release process, to autoverify whether our config structure matches, in schema, not values, to our development or example config file.


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
 - ~~Add tests~~
 - ~~Add Documentation~~
 - Add examples
 