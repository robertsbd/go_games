package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type GameFunction func(int64) []int64

type Game struct {
	in  string
	out string
	f   GameFunction
}

func main() {

	add_game := Game{
		in: "Please enter a number: ",
		out: "Adding 10 to your number: ",
		f: func(a int64) []int64 {
			return []int64{a + 10}
		},
	}

	fib_game := Game{
		in: "Enter the length of the sequence: ",
		out: "Here is your fibonacci sequence: ",
		f: func(n int64) []int64 {
			var i int64
			out := make([]int64, n)

			if n > 93 {
				fmt.Println("\nSorry n must be 93 or less")
				out[0] = -99
			} else {
				if n>0 {out[0] = 0}
				if n>1 {out[1] = 1}
				if n>2 {
					for i=2; i<n; i++ {
						out[i] = out[i-2]+out[i-1]
					}
				}
			}
			return out
		},
	}

	fmt.Print("(1) Add 10, (2) Fibonacci, (3) Fun with pointers: ")

	switch getString(){
		case "1":
		gameWrapper(add_game)
		case "2":
		gameWrapper(fib_game)
		case "3":
		funWithPointers()
		default:
		fmt.Print("Invalid response try again\n")

	}
}

func getString() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("An error occured while reading input. Please try again", err)
 		return "Error at line 25"
	}

	return strings.TrimSuffix(input, "\r\n")
}

func gameWrapper(g Game) {

	fmt.Print(g.in)

	input := getString()

	input_conv, err := strconv.ParseInt(input, 10, 64)

	if err != nil {
		fmt.Println("String conversion failed", err)
		return
	}

	fmt.Println(g.out, g.f(input_conv))
}

func funWithPointers() {

	var a,b,c int

	a = 1
	b = 2
	c = 3

	fmt.Println("a = ", a, ", b = ", b, " c = ", c)
	fmt.Println("b = c")

	b = c
	fmt.Println("a = ", a, ", b = ", b, " c = ", c)


	c = 99
	fmt.Println("c = 99")
	fmt.Println("a = ", a, ", b = ", b, " c = ", c)

	
	var d *int = &b
	*d=-99
	fmt.Println("var d *int = &b; *d=-99")
	fmt.Println("a = ", a, ", b = ", b, " c = ", c, "*d = ", *d)
	
}












