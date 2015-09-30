# Jvalin: CLI for querying JSON data by path

Working with JSON on the commandline can be a pain. Jvalin is designed
to make this easy by allowing you to query JSON using a path-style
syntax.

For example, consider this JSON document:

```json
{
"people": [
     {
       "name: "Matt"
     }
  ]
}
```

The outer object contains an array with one entry. Say we want to get
the name of the first person in the `people` list. We do that like this:

```
$ jv /people/0/name example.json
```

## About The Name

Jvalin (pronounced 'javelin') is a shortening of "JSON Value In..."
