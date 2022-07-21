package main

import(
	"fmt"
)

var onePiece string = "We are!"
var s int = 7
var m float32 = 11.5

func main(){
	//%s for strings
	//%d for integers
	//%f for float
	//%q for quoted string
	//%v can replace any type
	//%T display the type
	fmt.Printf("One Piece! %s\n", onePiece, )
	fmt.Printf("There are %d Shichibukais!\n", s)
	fmt.Printf("There are %.1f Mugiwaras!\n", m)
	fmt.Printf("The best opening is %q!\n", onePiece)
	fmt.Printf("%q is of the type %T", onePiece, onePiece)
}