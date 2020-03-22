This is a verbatim copy from "golang.org/x/build/internal/lru".

The original cannot be imported directly due to being internal.

Which is a shame! 'cause it's simple & efficient & safe => good!

## Other implementations
as seen in the wild:

### "github.com\bluele\gcache"
- dedicated cache package
- lots of options & features
	- expiry
	- event handlers: loader- and callback functions
- provides
	- LRU - Least Recently Used
	- LFU - Least Frequently Used
	- ARC - Adaptive Replacement Cache

- re ARC see also: http://en.wikipedia.org/wiki/Adaptive_replacement_cache

### "github.com/youtube/vitess/go/cache/lru_cache.go"
- More complex; uses string as Key and a type which can reports its Size() int.
- Aims to limit such Size (and not #-of-items).
- Uses timestamps for entries.
- Has methods such as Peek (without changing LRU) and SetIfAbsent,
  which may be usefull additions to this lru here.

### "github.com/tideland/golib/cache/cache.go"
- by Frank Mueller / Tideland / Oldenburg / Germany
- big
- clever
- imports other components of "github.com/tideland/golib"

### "github.com/syndtr/goleveldb/leveldb/cache/lru.go
- part of a larger cacher
- implements its own doubly linked list
- uses unsafe pointers

### "github.com/syncthing/syncthing/lib/ignore/cache.go
- more simple & dedicated map-based cache

### "github.com\hashicorp\golang-lru\ARC.go"
- ARC is an important strategy! IBM-Patent! 
- ARCCache is a thread-safe fixed size Adaptive Replacement Cache (ARC).

### "github.com\hashicorp\golang-lru\simplelru\lru.go"
- Not concurrency-safe.
- Optional EvictCallback function
- Code looks less go-ish - a little more verbose than need.
