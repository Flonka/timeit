package timeit

import (
	"time"
	"timeit/uuid"
)

type Entry struct {
	Start time.Time
	End   time.Time
	Id    string
}

func NewEntry() *Entry {
	return &Entry{
		Id:    uuid.NewHexUuid(16),
		Start: time.Now().UTC(),
	}
}
