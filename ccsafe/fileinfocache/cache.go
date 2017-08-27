// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fic represents an os.FileInfo based cache of file data
package fic

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/golangsam/container/ccsafe/lsm"
)

// FiCache provides cached file data for any os.FileInfo
type FiCache struct {
	dict *lsm.LazyStringerMap
	l    *sync.RWMutex
}

// New returns a fresh FiCache, ready to be populated via Register,
// or via Done (listening to a channel where os.FileInfo'es are supplied),
// and subsequently queried via Lookup or LookupData.
func New() *FiCache {
	return &FiCache{
		dict: lsm.New(),
		l:    new(sync.RWMutex),
	}
}

// Cache launches a inp listener and returns a channel to receive
// a new populated FiCache (like a done channel) upon close of inp.
func Cache(inp <-chan string) <-chan *FiCache {
	cha := make(chan *FiCache)
	fc := New()
	go func(inp <-chan string) {
		defer close(cha)
		<-fc.Done(inp)
		cha <- fc
	}(inp)
	return cha
}

// Done launches a inp listener and returns a done channel
//  Note: while listening, the FiCache is Locked,
//  and any error upon registration is silently ignored
func (fc *FiCache) Done(inp <-chan string) <-chan struct{} {
	cha := make(chan struct{})
	go func(inp <-chan string, done chan<- struct{}, fi *FiCache) {
		defer close(cha)
		fi.l.Lock()
		defer fi.l.Unlock()
		for op := range inp {
			_ = fi.register(op) // skip error
		}
		done <- struct{}{}
	}(inp, cha, fc)
	return cha
}

// Register a key in the FiCache; related data is refreshed,
// iff underlying os.FileInfo has changed.
func (fc *FiCache) Register(key string) error {
	fc.l.Lock()
	defer fc.l.Unlock()
	return fc.register(key)
}

// register key with it's data into cache iff not known yet
func (fc *FiCache) register(key string) (err error) {
	if have, err := os.Stat(key); err == nil {
		item, ok := fc.fetch(key)
		if !ok || item.fileInfo != have {
			if byteS, err := ioutil.ReadFile(key); err == nil {
				fc.dict.Assign(key, Item{have, byteS})
			} else {
				fc.dict.Delete(key)
			}
		}
	}
	return err
}

// lookup returns the FsData object assigned to key (if any) or false
func (fc *FiCache) lookup(key string) (Item, bool) {
	fdata, ok := fc.fetch(key)
	return fdata, ok
}

// Lookup returns the FsData object assigned to key (if any) or false
func (fc *FiCache) Lookup(key string) (Item, bool) {
	fc.l.RLock()
	defer fc.l.RUnlock()
	fdata, ok := fc.lookup(key)
	return fdata, ok
}

// LookupData returns the data assigned to the FsData of key (or an empty string)
func (fc *FiCache) LookupData(key string) string {
	fc.l.RLock()
	defer fc.l.RUnlock()
	fdata, ok := fc.lookup(key)
	if !ok {
		return ""
	}
	return string(fdata.byteS)
}

func (fc *FiCache) fetch(key string) (item Item, ok bool) {
	this, ok := fc.dict.Fetch(key)
	if ok {
		switch item := this.(type) {
		case Item:
			return item, ok
		}
	}
	return item, false
}
