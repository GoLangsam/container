*Don't show me Your code - show me Your data!*

---
Any `container` is either
- safe for concurrent processing (short: `ccsafe`)
- or not.

Not being safe for concurrent processing has been *normal* for a looong time.
Just: `normal` is not a good, meaningful name.
Neither is `oldway` - another discarded idea.

Thus: we came up with `oneway` - as there is only **one way** to use it safely: **single** threaded - ***not*** concurrent.

Note: Some kinds may deserve to be implemented both ways, as the `ccsafe` version is usually less performant.

Note: Some kinds are 'safe for concurrent processing' by construction. Notably *immutable* things.

----
## only *ccsafe/* implementation
- `das` - a *Dictionary for Any String*
- `fs` - type safe alternative to direct use of *`path/filepath`*
- `fscache` - provides cached file data for any `*fs.FsFile`
- `svp` - a *String Value Pair* (aka Named Constant)
- `tag` - a *Tag* (= a *String Value Pair* (aka Named Constant))


## only *oneway/* implementation
- `lsb` - a *Lazy String Buffer* -
  a gem hidden in *`path/path.go`* -
  chapeaux to [Rob Pike](https://en.wikipedia.org/wiki/Rob_Pike)
- `sync` - a drop-in replacement for standard `sync` with empty/no-op equivalents (used e.g. in `oneway/lsm`)

## both *ccsafe/* & *oneway/* implementations
- `drum` - a simple beat-counter (with a histogram)
- `lsm` - a *Lazy String Map*
- `stack` - a simple stack for *anything*; **not** typesafe; just: it's pattern is used typesafe elsewhere. You may like to do same.
