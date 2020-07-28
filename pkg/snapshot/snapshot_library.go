package snapshot

import (
	"errors"
	"sync"
)

var ErrSnapshotNotFound = errors.New("snapshot: error could not find snapshot")

type Library interface {
	Add(s *snapshot) (*snapshot, error)
	Update(s *snapshot) (*snapshot, error)
	Delete(s *snapshot) (*snapshot, error)
	Get(id string) (*snapshot, error)
}

func NewLibrary() Library {
	return &library{snapshots: map[string]*snapshot{}}
}

type library struct {
	mu        sync.Mutex
	snapshots map[string]*snapshot
}

func (l *library) Add(s *snapshot) (*snapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.snapshots[s.ID] = s
	return s, nil
}

func (l *library) Update(s *snapshot) (*snapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	_, ok := l.snapshots[s.ID]
	if !ok {
		return nil, ErrSnapshotNotFound
	}
	l.snapshots[s.ID] = s
	return s, nil
}

func (l *library) Delete(s *snapshot) (*snapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	delete(l.snapshots, s.ID)
	return s, nil
}

func (l *library) Get(id string) (*snapshot, error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	ss, ok := l.snapshots[id]
	if !ok {
		return nil, ErrSnapshotNotFound
	}
	return ss, nil
}
