Stringer
========

A string processing helper, and supported case-sensitive.

Return value same as the built-in `strings`.

If you don't care case-sensitive, please use built-in `strings` lib.

## Install

```go
go get github.com/hiscaler/stringer
```

## Usage

```go
stringer := NewStringer()
stringer.New("hello world!")
stringer.SetCaseSensitive(false)
stringer.HasPrefix("HEllO") // return `true`
stringer.TrimPrefix("HEllO").NewValue() // return ` world!`
```

You can use `NewValue()` method return processed string, and use `OriginalValue()` method return original string if you want. 