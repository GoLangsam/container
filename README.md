# *Don't show me Your code - show me Your data!*

[![Software License](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat-square)](LICENSE.md)
[![Go Report Card](https://goreportcard.com/badge/github.com/GoLangsam/container)](https://goreportcard.com/report/github.com/GoLangsam/container)
[![Build Status](https://travis-ci.org/GoLangsam/container.svg?branch=master)](https://travis-ci.org/GoLangsam/container)
[![GoDoc](https://godoc.org/github.com/GoLangsam/container?status.svg)](https://godoc.org/github.com/GoLangsam/container)

Any `container` is either
- safe for concurrent processing (short: `ccsafe`)
- or not.

Note: Not being safe for concurrent processing has been *normal* for a looong time.
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
- [`dot`](https://github.com/GoLangsam/container/blob/master/ccsafe/dot) -
  a tree of named anythings
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/dot?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/dot)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/dot)
- [`dotpath`](https://github.com/GoLangsam/container/blob/master/ccsafe/dotpath) -
  a parser for user-provided strings to be understood as a (list of) path
  with special awareness of triple dots (to recourse) and of trailing double dots (to inspect)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/dotpath?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/dotpath)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/dotpath)
- [`fs`](https://github.com/GoLangsam/container/blob/master/ccsafe/fs) -
  a type safe alternative to direct use of [`path/filepath`](https://godoc.org/path/filepath)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/fs?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/fs)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/fs)
- [`fscache`](https://github.com/GoLangsam/container/blob/master/ccsafe/fscache) -
  a cache for file data for any
  [`*fs.FsFile`](https://godoc.org/github.com/GoLangsam/container/ccsafe/fs#FsFile)
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
  a drop-in enhancement for a [`tag`]((https://godoc.org/github.com/GoLangsam/container/ccsafe/tag) ):
  adding introspective methods from [`do/ami`](https://godoc.org/github.com/GoLangsam/do/ami)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag/ami?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/tag/ami)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/tag/ami)
- [`tagmap`](https://github.com/GoLangsam/container/blob/master/ccsafe/tagmap) -
  a *Tag Map* (= a dictionary of *String Value Pair* s (as [`map`](https://golang.org/ref/spec#Map_types))
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/ccsafe/tagmap?status.svg)](https://godoc.org/github.com/GoLangsam/container/ccsafe/tagmap)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/ccsafe/tagmap)
  
## only *oneway/* implementation
- [`lsb`](https://github.com/GoLangsam/container/blob/master/oneway/lsb) - 
  a *Lazy String Buffer* -
  a gem hidden in [`path/path.go`](https://golang.org/src/path/path.go) -
  chapeaux to [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike)
  |
  [![GoDoc](https://godoc.org/github.com/GoLangsam/container/oneway/lsb?status.svg)](https://godoc.org/github.com/GoLangsam/container/oneway/lsb)
  |
  [lint](http://go-lint.appspot.com/github.com/GoLangsam/container/oneway/lsb)
- [`sync`](https://github.com/GoLangsam/container/blob/master/oneway/sync) - 
  a drop-in replacement for [standard `sync`](https://godoc.org/path/filepath) with empty/no-op equivalents (used e.g. in `oneway/lsm`)
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
## Remarks

Note: Some kinds may deserve to be implemented both ways, as the `ccsafe` version is usually less performant.

Note: Some kinds are *safe for concurrent processing* by construction. Notably *immutable* things.

## Support on Beerpay
Hey dude! Help me out for a couple of :beers:!

[![Beerpay](https://beerpay.io/GoLangsam/container/badge.svg?style=beer-square)](https://beerpay.io/GoLangsam/container)  [![Beerpay](https://beerpay.io/GoLangsam/container/make-wish.svg?style=flat-square)](https://beerpay.io/GoLangsam/container?focus=wish)