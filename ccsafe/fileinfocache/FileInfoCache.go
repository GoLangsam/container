// Copyright 2016 Andreas Pannewitz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package fic represents an os.FileInfo based cache of file data
package fic

import (
	"io/ioutil"
	"os"
	"sync"
)

// FiCache provides cached file data for any os.FileInfo
type FiCache struct {
	dict map[string]FiData
	l    *sync.RWMutex
}

// New returns a fresh FiCache, ready to be populated via Register,
// or via Done (listening to a channel where os.FileInfo'es are supplied),
// and subsequently queried via Lookup or LookupData.
func New() *FiCache {
	return &FiCache{
		dict: make(map[string]FiData),
		l:    new(sync.RWMutex),
	}
}

// Cache launches a inp listener and returns a channel to receive
// a new populated FiCache (like a done channel) upon close of inp.
func Cache(inp <-chan os.FileInfo) <-chan *FiCache {
	cha := make(chan *FiCache)
	fc := New()
	go func(inp <-chan os.FileInfo) {
		defer close(cha)
		<-fc.Done(inp)
		cha <- fc
	}(inp)
	return cha
}

// Done launches a inp listener and returns a done channel
//  Note: while listening, the FiCache is Locked,
//  and any error upon registration is silently ignored
func (fc *FiCache) Done(inp <-chan os.FileInfo) <-chan struct{} {
	cha := make(chan struct{})
	go func(inp <-chan os.FileInfo, done chan<- struct{}, fi *FiCache) {
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
func (fc *FiCache) Register(key os.FileInfo) error {
	fc.l.Lock()
	defer fc.l.Unlock()
	return fc.register(key)
}

// register key with it's data into cache iff not known yet
func (fc *FiCache) register(key os.FileInfo) error {
	finfo, ok := fc.dict[key.Name()]
	if !ok || finfo.fileInfo != key {
		byteS, err := ioutil.ReadFile(key.Name())
		if err == nil {
			fc.dict[key.Name()] = FiData{key, byteS}
		} else {
			delete(fc.dict, key.Name())
		}
		return err
	}
	return nil
}

// lookup returns the FsData object assigned to key (if any) or false
func (fc *FiCache) lookup(key os.FileInfo) (FiData, bool) {
	fdata, ok := fc.dict[key.Name()]
	return fdata, ok
}

// Lookup returns the FsData object assigned to key (if any) or false
func (fc *FiCache) Lookup(key os.FileInfo) (FiData, bool) {
	fc.l.RLock()
	defer fc.l.RUnlock()
	fdata, ok := fc.lookup(key)
	return fdata, ok
}

// LookupData returns the data assigned to the FsData of key (or an empty string)
func (fc *FiCache) LookupData(key os.FileInfo) string {
	fc.l.RLock()
	defer fc.l.RUnlock()
	fdata, ok := fc.lookup(key)
	if !ok {
		return ""
	}
	return string(fdata.byteS)
}
