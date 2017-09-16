// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fscache represents a file data cache based on an FsInfo
package fscache

import (
	"sync"

	"github.com/GoLangsam/container/ccsafe/fs"
)

// FsCache provides cached file data for any *fs.FsFile
type FsCache struct {
	dict map[string]*fs.FsData
	l    *sync.RWMutex
}

// New returns a fresh FsCache, ready to be populated via Register,
// or via Done (listening to a channel where FsPath'es are supplied),
// and subsequently queried via Lookup or LookupData.
func New() *FsCache {
	return &FsCache{
		dict: make(map[string]*fs.FsData),
		l:    new(sync.RWMutex),
	}
}

// Cache launches a inp listener and returns a channel to receive
// a new populated FsCache (like a done channel) upon close of inp.
func Cache(inp <-chan *fs.FsFile) <-chan *FsCache {
	cha := make(chan *FsCache)
	fc := New()
	go func(inp <-chan *fs.FsFile) {
		defer close(cha)
		<-fc.Done(inp)
		cha <- fc
	}(inp)
	return cha
}

// Done launches a inp listener and returns a done channel
//  Note: while listening, the FsCache is Locked,
//  and any error upon registration is silently ignored
func (fi *FsCache) Done(inp <-chan *fs.FsFile) <-chan struct{} {
	cha := make(chan struct{})
	go func(inp <-chan *fs.FsFile, done chan<- struct{}, fi *FsCache) {
		defer close(cha)
		fi.l.Lock()
		defer fi.l.Unlock()
		for op := range inp {
			_ = fi.register(op) // skip error
		}
		done <- struct{}{}
	}(inp, cha, fi)
	return cha
}

// Register a key in the FsCache; related data is refreshed,
// if underlying os.FileInfo has changed.
func (fi *FsCache) Register(key *fs.FsFile) error {
	fi.l.Lock()
	defer fi.l.Unlock()
	return fi.register(key)
}

func (fi *FsCache) register(key *fs.FsFile) error {
	finfo, ok := fi.dict[key.String()]
	if !ok || !finfo.InfoEquals(key) {
		fi.dict[key.String()] = key.AsData()
	}

	return nil
}

// lookup returns the FsData object assigned to key (if any) or false
func (fi *FsCache) lookup(key *fs.FsFile) (*fs.FsData, bool) {
	fdata, ok := fi.dict[key.String()]
	return fdata, ok
}

// Lookup returns the FsData object assigned to key (if any) or false
func (fi *FsCache) Lookup(key *fs.FsFile) (*fs.FsData, bool) {
	fi.l.RLock()
	defer fi.l.RUnlock()
	fdata, ok := fi.lookup(key)
	return fdata, ok
}

// LookupData returns the data assigned to the FsData of key (or an empty string)
func (fi *FsCache) LookupData(key *fs.FsFile) string {
	fi.l.RLock()
	defer fi.l.RUnlock()
	fdata, ok := fi.lookup(key)
	if !ok {
		return ""
	}
	return fdata.Data()
}
