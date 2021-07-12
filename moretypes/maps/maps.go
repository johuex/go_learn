package main

import "fmt"

type Vertex_1 struct {
	Lat, Long float64
}

var m map[string]Vertex_1

func main() {
	m = make(map[string]Vertex_1)
	m["Bell Labs"] = Vertex_1{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}