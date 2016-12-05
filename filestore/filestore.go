// Package filestore  provides a filebased timeit.Store
package filestore

import (
	"errors"
	"sync"
	"timeit"
)

type store struct {
	sync.RWMutex
	entries map[string]*timeit.Entry
}

func New() *store {
	return &store{
		entries: make(map[string]*timeit.Entry),
	}
}

func (s *store) GetEntry(id string) (*timeit.Entry, error) {

	s.RLock()
	entry, found := s.entries[id]
	s.RUnlock()

	if !found {
		return nil, errors.New("Entry not found")
	}

	return entry, nil
}

func (s *store) WriteEntry(entry *timeit.Entry) error {
	s.Lock()
	s.entries[entry.Id] = entry
	s.Unlock()
	return nil
}
