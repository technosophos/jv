# Jvalin: CLI for querying JSON data by path

Working with JSON on the commandline can be a pain. Jvalin is designed
to make this easy by allowing you to query JSON using a path-style
syntax.

For example, consider this JSON document:

```json
{
  "people": [
    {"name": "Matt"},
    {"name": "Skippy", "age": 123}
  ]
}
```

The outer object contains an array with one entry. Say we want to get
the name of the first person in the `people` list. We do that like this:

```
$ jv /people/0/name example.json
Matt
```

When possible, a query will return an individual value (as above), but
it can also return partial JSON:

```
$ jv /people/1 example.json
{"age":123,"name":"Skippy"}
```

## Building

JV uses Go 1.5+. Make sure you have that, then you can simply `go get
github.com/technosophos/jv`.

To build from source, you can use `go install` or:

```
go build -o jv jv.go
```

## About The Name

Jvalin (pronounced 'javelin') is a shortening of "JSON Value In..."
