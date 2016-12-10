package main

import (
	"fmt"
	"timeit"
	"timeit/filestore"
)

func main() {
	e := timeit.NewEntry()

	s := filestore.New()
	s.WriteEntry(e)

	fmt.Println(s.GetEntry(e.Id))

}
