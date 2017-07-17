*Don't show me Your code - show me Your data!*

---
Any `container` is either
- safe for concurrent processing (short: `ccsafe`)
- or not.

Not being safe for concurrent processing has been *normal* for a looong time.
Just: `normal` is not a good, meaningful name.
Neither is `oldway` - another discarded idea.

Thus: we came up with `oneway` - as there is only **one way** to use it safely: **single** threaded - ***not*** concurrent.

---
## only *ccsafe/* implementation
- [`das`](https://github.com/GoLangsam/container/blob/master/ccsafe/das) -
  a *Dictionary for Any String*
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/das?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/das) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/das)
  |
- [`fs`](https://github.com/GoLangsam/container/blob/master/ccsafe/fs) -
  type safe alternative to direct use of *`path/filepath`*
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/fs?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/fs) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/fs)
- [`fscache`](https://github.com/GoLangsam/container/blob/master/ccsafe/fscache) -
  provides cached file data for any `*fs.FsFile`
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/fscache?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/fscache) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/fscache)
- [`svp`](https://github.com/GoLangsam/container/blob/master/ccsafe/svp) -
  a *String Value Pair* (aka Named Constant)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/svp?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/svp) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/svp)
- [`tag`](https://github.com/GoLangsam/container/blob/master/ccsafe/tag) -
  a *Tag* (= a *String Value Pair* (aka Named Constant))
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/tag)
- [`tag/ami`](https://github.com/GoLangsam/container/blob/master/ccsafe/tag/ami) -
  drop-in enhancement for a `tag`: adding introspective methods from `do/ami`
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag/ami?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag/ami) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/tag/ami)
- [`tagmap`](https://github.com/GoLangsam/container/blob/master/ccsafe/tagmap) -
  a *Tag Map* (= a dictionary of *String Value Pair* s (as `map`)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/tagmap?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/tagmap) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/tagmap)
  
  


## only *oneway/* implementation
- [`lsb`](https://github.com/GoLangsam/container/blob/master/oneway/lsb) - 
  a *Lazy String Buffer* -
  a gem hidden in *`path/path.go`* -
  chapeaux to [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/lsb?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/lsb)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/lsb)
- [`sync`](https://github.com/GoLangsam/container/blob/master/oneway/sync) - 
  a drop-in replacement for standard `sync` with empty/no-op equivalents (used e.g. in `oneway/lsm`)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/sync?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/sync) 
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/sync)

## both *ccsafe/* & *oneway/* implementations
- `drum` - 
  a simple beat-counter (with a histogram)
	- [`ccsafe/drum`](https://github.com/GoLangsam/container/blob/master/ccsafe/drum)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/drum?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/drum)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/drum)
	- [`oneway/drum`](https://github.com/GoLangsam/container/blob/master/oneway/drum)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/drum?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/drum)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/drum)
- `lsm` - 
  a *Lazy String Map*
	- [`ccsafe/lsm`](https://github.com/GoLangsam/container/blob/master/ccsafe/lsm)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/lsm?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/lsm)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/lsm)
	- [`oneway/lsm`](https://github.com/GoLangsam/container/blob/master/oneway/lsm)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/lsm?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/lsm)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/lsm)
- `stack` - 
  a simple stack for *anything*; **not** typesafe; just: it's pattern is used typesafe elsewhere. You may like to do same.
	- [`ccsafe/stack`](https://github.com/GoLangsam/container/blob/master/ccsafe/stack)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/stack?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/stack)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/stack)
	- [`oneway/stack`](https://github.com/GoLangsam/container/blob/master/oneway/stack)
	  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/stack?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/stack)
	  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/stack)

---

Note: Some kinds may deserve to be implemented both ways, as the `ccsafe` version is usually less performant.

Note: Some kinds are 'safe for concurrent processing' by construction. Notably *immutable* things.
