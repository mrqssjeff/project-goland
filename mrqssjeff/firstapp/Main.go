package main

import (
	"fmt"
	"strconv"
)

var (
	anime string = "One Piece"
	i float32 = 42;
)

/**
declared variables must be used;
lower case variables are scoped to the package;
upper case variables then its globally visible;
to convert to string you need to use the strconv package;
**/
func main(){
	var luke int = 89
	k := 99
	var luka float32 = float32(luke)
	var text string = strconv.Itoa(k)
	fmt.Printf("%v, %T", i ,luka)
	fmt.Println(k)
	fmt.Println(anime)
	fmt.Println(luke)
	fmt.Println(text)

}
