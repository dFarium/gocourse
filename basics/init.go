package main

// Se ejecutan en orden las funciones init antes de main
func init () {
	println("Init function called")
}

func init () {
	println("Init 2 function called")
}

func init () {
	println("Init 3 function called")
}

func main () {
	println("Main function called")
}
