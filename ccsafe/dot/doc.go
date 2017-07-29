// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// package dot implements a container of named stringmaps.
//
// Yes, it's just a tree of strings and Stringers - free of duplicate names.
//
// And, yes, it's recursive.
// And, yes it's concurrency-safe.
// And, yes it carries stuff as You name it.
// And, yes it can also carry arbitrary Stringers - things which can name themselves.
// And, yes it supports concurrency via generic piping functions - found elsewhere.
//
// Children can be obtained as map[string]string, and as sorted slice []string
//
// Thus, it can be used as an iterated map "map[string]map[string]...map[string]string"
// or as an iterated slice "[]...[]string"
//
// Any of it's elements has a name (=string) and (optionally) a map of named elements,
// the key's of which are strings, and the values are elements,
// which are strings which are the name of an (optional) map of named elements,
// the key's of which are strings, and the values are elements,
// which are strings which are the name of an (optional) map of named elements,
// the key's of which are strings, and the values are elements,
// which are strings which are the name of an (optional) map of named elements,
// the key's of which are strings, and the values are elements,
// ...
//
// Such structure can be useful when information is acquired
// recursivlely such as by parsing some object (directory, template, URL, ...),
// which contains and 'reveals' additional information when being parsed/inspected.
//
// Information, which 'reveals' additional information when being parsed/inspected.
//
// ...
//
// Such can be used e.g. in selfdefining template structures
// Such may be extended to all kind's of types which are suitable as both key and value
// of a map.
//
// This prototype has it's focus on type "string", as this is helpful in dynamic handling
// of template-driven processes.
//
package dot
