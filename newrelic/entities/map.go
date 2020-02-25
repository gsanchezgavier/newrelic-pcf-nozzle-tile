// Copyright 2020 New Relic Corporation. All rights reserved.
// SPDX-License-Identifier: Apache-2.0

package entities

import (
	"sync"

	"github.com/newrelic/newrelic-pcf-nozzle-tile/newrelic/uid"
)

// Collection ...
type Collection map[uid.ID]*Entity

// Map ...
type Map struct {
	collection Collection
	sync       *sync.RWMutex
}

// Drain returns collection of Entities ...
func (m *Map) Drain() (c []*Entity) {
	m.sync.Lock()
	defer m.sync.Unlock()
	for _, v := range m.collection {
		c = append(c, v)
	}
	m.collection = Collection{}
	return c
}

// Has ...
func (m *Map) Has(id uid.ID) (entity *Entity, found bool) {
	m.sync.RLock()
	defer m.sync.RUnlock()
	if entity, found = m.collection[id]; found {
		return entity, true
	}
	return entity, false
}

// Put ...
func (m *Map) Put(entity *Entity) {
	m.sync.Lock()
	m.collection[entity.Signature()] = entity
	m.sync.Unlock()
}
