# ToDo

## ToDo

### container/oneway/list
- make Value private and give accessors:
  This allows for easy and type-safe use via anonymous embedding!

- move up to container/list (analog to standard package)???
  Just: then there would be no chance for concurrent access,
  which would e.g. lock the relevant *list.List
  Or make a 'fake' which 'only' lifts the oneway!

### xxx/stack:
- Drop&Pop panic on underflow implicitly (slice index out of boudary (=1-):
  - TODO: Drop() more tolerant (may even be a great time for shrinking?)
  - TODO: Pop() with private panic

### genny
- make more generic with genny (=> new repo: "container.gen": e.g.:
  - das
  - lsm
  - svp
  - tagmap
  - stack
