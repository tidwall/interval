# `interval`

[![GoDoc](https://godoc.org/github.com/tidwall/interval?status.svg)](https://godoc.org/github.com/tidwall/interval)

Repeatedly call a function with a fixed time delay between each call. This is
an alternative to the `time.Ticker`.

## Usage

There are only two functions `Set()` and `Clear()`. The Set function begins
repeating the function until Clear is called.

### Functions

```
func Set(fn func(t time.Time), delay time.Duration) Interval
func (iv Interval) Clear()
```

### Example

Let's say that you have a custom file type which needs to perform a
background operation every hour.

This example starts the interval function when the object opens and clears
the interval when the object is closed.

```go

type SpecialFile struct {
    iv Interval
}

func Open(path string) (*SpecialFile, error) {
    f := new(SpecialFile)
    f.iv = interval.Start(func(t time.Time){
        db.bgOperation()
    }, time.Hour)
    return f, nil
}

func Close() error{
    f.iv.Clear()
    return nil
}

func (f *SpecialFile) bgOperation() {
    // this operation runs in the backgound
}
```

## Contact

Josh Baker [@tidwall](http://twitter.com/tidwall)

## License

`interval` source code is available under the MIT [License](/LICENSE).