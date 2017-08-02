package main

import "fmt"
import "time"
import "math/rand"

func random(min, max int) int {
    rand.Seed(time.Now().Unix())
    return rand.Intn(max - min) + min
}

func roll(number,sides,bonus int) int {
	damage := 0
	
	for i := 0; i < number; i++ {
		damage += random(1, sides)
		fmt.Println("Hello")
	}
	return damage + bonus
}

func main() {
	fmt.Println(roll(1,20,0))
}