package timeit

type Store interface {
	WriteEntry(*Entry) error
	GetEntry(string) (*Entry, error)
}
