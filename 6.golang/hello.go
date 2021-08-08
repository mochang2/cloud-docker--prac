// like Java
package main

// include <stdio.h>
import "fmt"

func my_func() {
	fmt.Println("This is my_func")
}

// int main(int argc, char **argv)
func main() {
	var i int
	var j int = 2
	k := 3
	// var str = "한글"

	fmt.Println("Hello, Go")
	fmt.Print("i =", i)
	fmt.Println("j =", j)
	fmt.Println("k =", k)

	for m:=1; m<=5; m++{
		println("m =", m)
	}

	var scores [3]int = [3]int{333, 444, 555}
	// scores := [...]int{333, 444, 555} // dynamic array length
	for i:=0; i<len(scores); i++{
		fmt.Println(scores[i])
	}
}

// gcc -o hello hello.c
// go run hello.go (without build, just run)