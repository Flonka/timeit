package main

import "encoding/json"
import "fmt"
import "timeit"

func main() {
	fmt.Println("hello world")
	e := timeit.NewEntry()
	b, _ := json.Marshal(e)

	fmt.Println(e.End.IsZero())
	fmt.Println(e.Start.IsZero())
	fmt.Println(string(b))

}
