package movie

import "fmt"

func init() {
	fmt.Println("init movie")
}

func Review(name string, ratting float64) {
	fmt.Printf("I reviewed %s and it's ratting is %f\n", name, ratting)
}
