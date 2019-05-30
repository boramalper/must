# must
*A delightful way to deal with unrecoverable errors in Go*

`boramalper/must` offers utility functions to deal with *unrecoverable* errors in Go,
in other words, errors that you log only and exit, or panic on. 

```go
Must(err error)
M(e error)

MustVal(val interface{}, err error) interface{}
MV(v interface{}, e error) interface{}

MustValVoid(val interface{}, err error)
MVV(v interface{}, e error)
```

The functions should be self-explanatory; their default behaviour is to `panic()`
when `err != nil`, but this can be customised by `SetHandler` function or by
creating a new `must.Context`. 

## Example
```go
database := must.MV(sql.Open("sqlite3", "...")).(*sql.DB)
defer must.M(database.Close())

must.M(database.Ping())

// Use MustValVoid (or MVV shortly) if you don't care about
// the return value.
must.MVV(database.Exec(`
    PRAGMA foreign_key_check;
    ...
`))
```

## License
ISC License, see [LICENSE](./LICENSE) for details.

