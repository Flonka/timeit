package timeit

import (
	"time"
	"timeit/uuid"
)

type entry struct {
	Start time.Time
	End   time.Time `json:"end",omitifempty`
	Id    string    `json:"id"`
}

func NewEntry() *entry {
	return &entry{
		Id:    uuid.NewHexUuid(16),
		Start: time.Now().UTC(),
	}
}
