package main
import "fmt"

func main() {

		// declare our fatefull test variable
        x:= 0
        
        // try to change it to the value in the function
        changeNoPointer(x)
        fmt.Println("Without pointer")
        fmt.Println("x =", x)

        // pass refference to the adress of x, with a func
        // that uses pointer in order to change the memory 
        changeYesPointer(&x)
        fmt.Println("With pointer")
        fmt.Println("x =", x)
}


// function to attempt change, out of scope
func changeNoPointer(x int) {

	x = 2
}

// function to change using a pointer, like having a passport
func changeYesPointer(x *int) {

	// has the start and both places, like a journey or a wormhole
	*x = 2
}