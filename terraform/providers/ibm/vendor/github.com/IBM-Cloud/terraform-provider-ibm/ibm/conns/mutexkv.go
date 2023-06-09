// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package conns

import (
	"log"
	"sync"
)

// MutexKV is a simple key/value store for arbitrary mutexes. It can be used to
// serialize changes across arbitrary collaborators that share knowledge of the
// keys they must serialize on.
//
// Deprecated: This will be removed in v2 without replacement. If you need
// its functionality, you can copy it or reference the v1 package.
//
// The initial use case is to let aws_security_group_rule resources serialize
// their access to individual security groups based on SG ID.

// This is a global MutexKV for use within this plugin.
var IbmMutexKV = NewMutexKV()

type MutexKV struct {
	lock  sync.Mutex
	store map[string]*sync.Mutex
}

// Lock the mutex for the given key. Caller is responsible for calling Unlock
// for the same key
//
// Deprecated: This will be removed in v2 without replacement. If you need
// its functionality, you can copy it or reference the v1 package.
func (m *MutexKV) Lock(key string) {
	log.Printf("[DEBUG] Locking %q", key)
	m.get(key).Lock()
	log.Printf("[DEBUG] Locked %q", key)
}

// Unlock the mutex for the given key. Caller must have called Lock for the same key first
//
// Deprecated: This will be removed in v2 without replacement. If you need
// its functionality, you can copy it or reference the v1 package.
func (m *MutexKV) Unlock(key string) {
	log.Printf("[DEBUG] Unlocking %q", key)
	m.get(key).Unlock()
	log.Printf("[DEBUG] Unlocked %q", key)
}

// Returns a mutex for the given key, no guarantee of its lock status
func (m *MutexKV) get(key string) *sync.Mutex {
	m.lock.Lock()
	defer m.lock.Unlock()
	mutex, ok := m.store[key]
	if !ok {
		mutex = &sync.Mutex{}
		m.store[key] = mutex
	}
	return mutex
}

// NewMutexKV Returns a properly initalized MutexKV
//
// Deprecated: This will be removed in v2 without replacement. If you need
// its functionality, you can copy it or reference the v1 package.
func NewMutexKV() *MutexKV {
	return &MutexKV{
		store: make(map[string]*sync.Mutex),
	}
}
