package main 
import "fmt"


type Entry struct {
	age int
	quote string
}

var dict = map[string]Entry {
	"Thadryan": {29, "t"},
	"Scott":    {43, "s"},
}

var mass = map[string]int {
	"C" : 10, "H" : 5, "L": 3,
}

func main() {
	fmt.Println(dict)
	fmt.Println(dict)
	fmt.Println(mass)

	
	count := 0

	for value := range mass {
		count += mass[value]

	fmt.Println(count)
	}
}

/*
type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println("Bell Labs", m["Bell Labs"])
}
*/