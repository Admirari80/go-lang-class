package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m1 = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68, -74.39,
	},
	"Google": Vertex{
		37.42, -122.08,
	},
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		Lat:  40.68,
		Long: -74.39,
	}
	fmt.Println(m["Bell Labs"])
	_, ok := m["ashish"]
	fmt.Println(ok)
	fmt.Println(m1)
}
