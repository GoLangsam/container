# Package `fs`

Package `fs` represents the constituents of a file system and provides type safe access.

Beyond a couple of convenience functions, it provides a tree of types and related functionalities.

## The Type Tree

- `fsPath` is just a string, and provides many basic functionalities (also from standard packages `"path/filepath"`, `"os"`, `"io/ioutil"`).
	- `fsInfo` extends `fsPath` with an `os.FileInfo` and more functionalities and thus includes some 'reality check'.
		- `FsFold` extends `fsInfo` and represents the name of a folder / directory
		- `FsFile` extends `fsInfo` and represents the name of a file
			- `FsData` extends `FsFile` and represents the name and data of a file (used e.g. in `"FsCache"`)
	- `FsBase` extends `fsPath` and represents a base name (no slashes)

	- `Pattern` extends `fsPath` and represents a pattern (may contain wildcards)

## A Friendly Package

`fs` aims to be a friendly package.

Intentionally it provides a lot of functionalitites and methods.

Sometimes the full sorted lists given by tools such as `godoc` or `go doc`
may be more confusing than helpful.

For better overview there is a tree of interfaces where each node summarises related methods.


## Concurrency safe

As all types are intentionally immutable and thus safe in concurrent use.


## The `TypeS` collections

For each type (incl. non-exported!) there is a corresponding container type suffixed with `...S`,
which is just a typesslice (thus the Uppercase "S": for both: plural, and slice) of typepointers,
e.g. `FsInfoS = []*fsInfo`.

Such collections are provided / accepted by functions, and can easily be passed around as arguments or iterated over.


## Creators and derivators

"New" as a prefix for creator functions is intentionally not used; "Force" is offered instead.

`ForceFile`, ForceFold (and `NotDown` and `Recurse`), `ForceBase`, `ForceData` are intentionally tolerant.
They perform no reality check whatshowever but allow the user to express his intentions.
Thus, they are to be used with good care, if at all.

The `f.As<Type>` methods are a safer way to create/derive related types.

Various corresponding `Try<Type>` methods even give access to eventual 'reality check' failures;
in order to avoid panics, they should be used before any `f.As<Type>` .

Note: `NewS` is an exception - it creates an `FsPathS`, a collections of unqualified `fsPathS`.


## Matching / Patterns

The type `Pattern` extends `fsPath` and represents a pattern, which may contain wildchars such as "`*`" & "`?`".

Patterns are supported in various ways:
- `Match` is a convenience for `filepath.Match` and accepts variadic lists
- `MatchDisk` and friends implicitly interpret the given string / fsPath as such,
  as internally a Glob against the disk is used.
- A couple of *Matches* methods accept patterns and even lists of patterns.

Note: Currently, standard package `"path/filepath"` does not export it's `isMeta` function, which is useful for checking.
Thus, duplication of knowledge would be required - and this is very much disliked.
